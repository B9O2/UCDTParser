package ucdt

import (
	"fmt"

	"github.com/B9O2/UCDTParser/builtin"
	"github.com/B9O2/UCDTParser/expression"
	"github.com/B9O2/evaluate"
	"github.com/google/cel-go/common/types"
)

type Expression string

func (expr Expression) Eval(e *evaluate.Evaluate, args map[string]any) (any, []string) {
	var detail []string

	r, err := e.Eval(string(expr), args)
	if err != nil {
		detail = append(detail, fmt.Sprintf("[Expression Eval] error: %s", err))
		return nil, detail
	}
	return r, detail
}

func NewEvaluate(funcs map[string]any, memberMethods map[*types.Type]map[string]any) (*evaluate.Evaluate, error) {
	e, err := evaluate.NewEvaluate("ucdt")
	if err != nil {
		return nil, err
	}

	for name, f := range funcs {
		err = e.NewFunction(name, f)
		if err != nil {
			return nil, err
		}
	}

	for t, methods := range memberMethods {
		for name, method := range methods {
			err = e.NewMemberFunction(t, name, method)
			if err != nil {
				return nil, err
			}
		}
	}

	return e, nil
}

func GenArgs(e *evaluate.Evaluate, score float32, sds []*SourceData) map[string]any {
	e.DeclareVariable("score", 0)
	e.NewClass("all", (*expression.SourceDataList)(nil), map[string]any{
		"contains": builtin.Contains,
	})

	sdl := &expression.SourceDataList{}
	for _, sd := range sds {
		sdl.Sds = append(sdl.Sds, &expression.SourceData{
			Data:   sd.data,
			Source: sd.source,
		})
	}

	args := map[string]any{
		"score": score,
		"all":   sdl,
	}
	return args
}

var StandardEnv = NewEnviroment(
	map[string]any{},
	map[*types.Type]map[string]any{},
)
