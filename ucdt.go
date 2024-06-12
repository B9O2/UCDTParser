package ucdt

import (
	"bytes"
	"context"
	"fmt"

	"github.com/B9O2/Multitasking"
	"github.com/BurntSushi/toml"
	"github.com/google/cel-go/common/types"
)

type Environment struct {
	funcs         map[string]any
	memberMethods map[*types.Type]map[string]any
}

func (e *Environment) Funcs() map[string]any {
	return e.funcs
}

func (e *Environment) MemberMethods() map[*types.Type]map[string]any {
	return e.memberMethods
}


func (e *Environment) PatchFuncs(funcs map[string]any) {
	if funcs == nil {
		return
	}

	for name, f := range funcs {
		e.funcs[name] = f
	}
}

func (e *Environment) PatchMemberFuncs(memberMethods map[*types.Type]map[string]any) {
	if memberMethods == nil {
		return
	}

	for t, funcs := range memberMethods {
		if _, ok := e.memberMethods[t]; !ok {
			e.memberMethods[t] = map[string]any{}
		}
		for name, f := range funcs {
			e.memberMethods[t][name] = f
		}
	}
}

func NewEnviroment(funcs map[string]any, memberMethods map[*types.Type]map[string]any) *Environment {
	env := &Environment{
		funcs:         map[string]any{},
		memberMethods: map[*types.Type]map[string]any{},
	}
	env.PatchFuncs(funcs)
	env.PatchMemberFuncs(memberMethods)
	return env
}

type TagOption struct {
	Traits TraitMap   `toml:"traits"` //[name]TraitOption
	Info   InfoMap    `toml:"info"`
	Expr   Expression `toml:"expression"`
}

func (t *TagOption) Match(env *Environment, sds ...SourceData) MatchResult {
	mr := NewMatchResult()

	//Trait
	scores, detail := t.Traits.Match(sds)
	for _, s := range scores[true] {
		mr.score += s
	}
	mr.scoreDetail = scores
	mr.detail = append(mr.detail, detail...)
	//Expression
	if len(t.Expr) > 0 {
		e, err := NewEvaluate(env.funcs, env.memberMethods)
		if err != nil {
			mr.err = err
			return mr
		}

		args := GenArgs(e, mr.score, sds)
		r, detail := t.Expr.Eval(e, args)
		mr.detail = append(mr.detail, detail...)
		if b, ok := r.(bool); ok {
			mr.expression = b
		} else {
			mr.detail = append(mr.detail, fmt.Sprintf("expression must return a bool,but '%T'", r))
			mr.expression = false
		}

		if mr.expression {
			mr.score = 1
		} else {
			mr.score = 0
		}
	} else {
		mr.detail = append(mr.detail, "expression not enabled")
	}

	//Info

	mr.info, detail = t.Info.Extract(env, mr.score, sds)
	mr.detail = append(mr.detail, detail...)

	return mr
}

type Tags map[string]TagOption

func (t Tags) Match(env *Environment, threads uint, sds ...SourceData) (MatchResults, error) {
	type task struct {
		tag string
		opt TagOption
	}
	type result struct {
		tag string
		mr  MatchResult
	}

	mrs := NewMatchResults()

	mt := Multitasking.NewMultitasking("match", nil)
	mt.Register(
		func(dc Multitasking.DistributeController) {
			for tag, opt := range t {
				dc.AddTask(task{
					tag: tag,
					opt: opt,
				})
			}
		},
		func(ec Multitasking.ExecuteController, a any) any {
			task := a.(task)
			return result{
				tag: task.tag,
				mr:  task.opt.Match(env, sds...),
			}
		},
	)

	mt.SetResultMiddlewares(Multitasking.NewBaseMiddleware(func(ec Multitasking.ExecuteController, i interface{}) (interface{}, error) {
		res := i.(result)
		mrs.Add(res.tag, res.mr)
		return nil, nil
	}))

	_, err := mt.Run(context.Background(), threads)
	if err != nil {
		return mrs, err
	}
	return mrs, nil
}

type UCDT struct {
	Tags Tags `toml:"tags"`
}

func ParseUCDT(data []byte) (*UCDT, error) {
	v := &UCDT{}
	_, err := toml.NewDecoder(bytes.NewReader(data)).Decode(v)
	return v, err
}
