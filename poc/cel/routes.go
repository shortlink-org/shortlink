package main

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"encoding/json/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func setupRoutes(r *chi.Mux, compiledRules map[string]*cel.Program) chi.Router {
	r.Post("/evaluate", func(w http.ResponseWriter, r *http.Request) {
		var input EvaluateInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		nowTimestamp := timestamppb.New(time.Unix(input.Now, 0))
		results := make(map[string]any)

		for ruleName, compiledRule := range compiledRules {
			result, errEvalRule := evaluateRule(*compiledRule, map[string]any{
				"claims":            map[string]any{"exp": input.Claims.Exp, "aud": input.Claims.Aud},
				"now":               nowTimestamp,
				"expected_audience": expectedValue,
			})
			if errEvalRule != nil {
				http.Error(w, fmt.Sprintf("Error evaluating rule '%s': %v", ruleName, errEvalRule), http.StatusInternalServerError)

				return
			}
			results[ruleName] = result.Value()
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			http.Error(w, "Error marshaling results", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonResults) //nolint:errcheck // ignore error
	})

	return r
}

// the report prints out the result of evaluation in human-friendly terms.
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
		ids := make([]int, len(stateIDs))
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
