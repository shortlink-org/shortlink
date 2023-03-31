package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Claims struct {
	Exp int64  `json:"exp"`
	Aud string `json:"aud"`
	// Add more fields as needed
}

type EvaluateInput struct {
	Rule   string `json:"rule"`
	Claims Claims `json:"claims"`
	Now    int64  `json:"now"`
}

func main() {
	rulesPath := "./internal/services/api/cel/rules"
	rules, err := loadRules(rulesPath)
	if err != nil {
		fmt.Printf("Error loading rules: %v\n", err)
		os.Exit(1)
	}

	compiledRules, err := compileRules(rules)
	if err != nil {
		fmt.Printf("Error compiling rules: %v\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", setupRoutes(r, compiledRules))
	http.ListenAndServe(":8080", r)
}

func setupRoutes(r *chi.Mux, compiledRules map[string]*cel.Program) chi.Router {
	r.Post("/evaluate", func(w http.ResponseWriter, r *http.Request) {
		var input EvaluateInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		nowTimestamp := timestamppb.New(time.Unix(input.Now, 0))
		results := make(map[string]interface{})

		for ruleName, compiledRule := range compiledRules {
			result, err := evaluateRule(compiledRule, map[string]interface{}{
				"claims": map[string]interface{}{"exp": input.Claims.Exp, "aud": input.Claims.Aud},
				"now":    nowTimestamp,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("Error evaluating rule '%s': %v", ruleName, err), http.StatusInternalServerError)
				return
			}
			results[ruleName] = result.Value()
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			http.Error(w, "Error marshalling results", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResults)
	})

	return r
}

func loadRules(path string) (map[string]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	rules := make(map[string]string)
	for _, file := range files {
		if !file.IsDir() {
			filename := filepath.Join(path, file.Name())
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				return nil, err
			}
			rules[file.Name()] = string(content)
		}
	}
	return rules, nil
}

func compileRules(rules map[string]string) (map[string]*cel.Program, error) {
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar(
				"claims",
				decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("now", decls.Timestamp),
			decls.NewVar("aud", decls.String),
		),
	)
	if err != nil {
		return nil, err
	}

	compiledRules := make(map[string]*cel.Program)
	for name, rule := range rules {
		ast, iss := env.Compile(rule)
		if iss.Err() != nil {
			return nil, iss.Err()
		}
		program, err := env.Program(ast)
		if err != nil {
			return nil, err
		}
		compiledRules[name] = &program
	}
	return compiledRules, nil
}

func evaluateRule(program *cel.Program, inputs map[string]interface{}) (ref.Val, error) {
	//activation, err := cel.NewEnv()
	//if err != nil {
	//	return nil, err
	//}
	out, _, err := eval(*program, inputs)
	return out, err
}

func eval(prg cel.Program,
	vars any) (out ref.Val, det *cel.EvalDetails, err error) {
	varMap, isMap := vars.(map[string]any)
	fmt.Println("------ input ------")
	if !isMap {
		fmt.Printf("(%T)\n", vars)
	} else {
		for k, v := range varMap {
			switch val := v.(type) {
			case proto.Message:
				bytes, err := prototext.Marshal(val)
				if err != nil {
					glog.Exitf("failed to marshal proto to text: %v", val)
				}
				fmt.Printf("%s = %s", k, string(bytes))
			case map[string]any:
				b, _ := json.MarshalIndent(v, "", "  ")
				fmt.Printf("%s = %v\n", k, string(b))
			case uint64:
				fmt.Printf("%s = %vu\n", k, v)
			default:
				fmt.Printf("%s = %v\n", k, v)
			}
		}
	}
	fmt.Println()
	out, det, err = prg.Eval(vars)
	report(out, det, err)
	fmt.Println()
	return
}

// report prints out the result of evaluation in human-friendly terms.
func report(result ref.Val, details *cel.EvalDetails, err error) {
	fmt.Println("------ result ------")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("value: %v (%T)\n", result, result)
	}
	if details != nil {
		fmt.Printf("\n------ eval states ------\n")
		state := details.State()
		stateIDs := state.IDs()
		ids := make([]int, len(stateIDs), len(stateIDs))
		for i, id := range stateIDs {
			ids[i] = int(id)
		}
		sort.Ints(ids)
		for _, id := range ids {
			v, found := state.Value(int64(id))
			if !found {
				continue
			}
			fmt.Printf("%d: %v (%T)\n", id, v, v)
		}
	}
}
