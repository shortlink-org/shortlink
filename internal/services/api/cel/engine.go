package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// evaluateRule ...
func evaluateRule(program *cel.Program, inputs map[string]interface{}) (ref.Val, error) {
	// activation, err := cel.NewEnv()
	// if err != nil {
	//	return nil, err
	//}
	out, _, err := eval(*program, inputs)
	return out, err
}

// eval ...
func eval(prg cel.Program,
	vars any,
) (out ref.Val, det *cel.EvalDetails, err error) {

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

// compile will parse and check an expression `expr` against a given
// environment `env` and determine whether the resulting type of the expression
// matches the `exprType` provided as input.
func compile(env *cel.Env, expr string, celType *cel.Type) *cel.Ast {
	ast, iss := env.Compile(expr)
	if iss.Err() != nil {
		glog.Exit(iss.Err())
	}
	if !reflect.DeepEqual(ast.OutputType(), celType) {
		glog.Exitf(
			"Got %v, wanted %v result type", ast.OutputType(), celType)
	}

	return ast
}
