package ucdt

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/B9O2/Multitasking"
	"github.com/BurntSushi/toml"
	"github.com/google/cel-go/common/types"
)

type Environment struct {
	funcs         map[string]any
	memberMethods map[*types.Type]map[string]any
	lock          *sync.Mutex
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

	e.lock.Lock()
	for name, f := range funcs {
		e.funcs[name] = f
	}
	e.lock.Unlock()
}

func (e *Environment) PatchMemberFuncs(memberMethods map[*types.Type]map[string]any) {
	if memberMethods == nil {
		return
	}
	e.lock.Lock()
	for t, funcs := range memberMethods {
		if _, ok := e.memberMethods[t]; !ok {
			e.memberMethods[t] = map[string]any{}
		}
		for name, f := range funcs {
			e.memberMethods[t][name] = f
		}
	}
	e.lock.Unlock()
}

func NewEnviroment(funcs map[string]any, memberMethods map[*types.Type]map[string]any) *Environment {
	env := &Environment{
		funcs:         map[string]any{},
		memberMethods: map[*types.Type]map[string]any{},
		lock:          &sync.Mutex{},
	}
	env.PatchFuncs(funcs)
	env.PatchMemberFuncs(memberMethods)
	return env
}

type TagOption struct {
	Source []string   `toml:"source"`
	Traits TraitMap   `toml:"traits"` //[name]TraitOption
	Info   InfoMap    `toml:"info"`
	Expr   Expression `toml:"expression"`
}

func (t *TagOption) Match(env *Environment, sds ...*SourceData) MatchResult {
	mr := NewMatchResult()
	sdset := map[string]*SourceData{}
	limited := len(t.Source) > 0
	for _, sd := range sds {
		if limited {
			for _, src := range t.Source {
				if src == sd.source {
					sdset[sd.source] = sd
					break
				}
			}
		} else {
			sdset[sd.source] = sd
		}
	}

	//Trait
	scores, detail := t.Traits.Match(sdset)
	for _, s := range scores[true] {
		mr.Score += s
	}
	mr.ScoreDetail = scores
	mr.Detail = append(mr.Detail, detail...)
	//Expression
	if len(t.Expr) > 0 {
		e, err := NewEvaluate(env.funcs, env.memberMethods)
		if err != nil {
			mr.Err = err
			return mr
		}

		args := GenArgs(e, mr.Score, sdset)
		r, detail := t.Expr.Eval(e, args)
		mr.Detail = append(mr.Detail, detail...)
		if b, ok := r.(bool); ok {
			mr.Expression = b
		} else {
			mr.Detail = append(mr.Detail, fmt.Sprintf("expression must return a bool,but '%T'", r))
			mr.Expression = false
		}

		if mr.Expression {
			mr.Score = 1
		} else {
			mr.Score = 0
		}
	} else {
		mr.Detail = append(mr.Detail, "expression not enabled")
	}

	//Info

	mr.Info, detail = t.Info.Extract(env, mr.Score, sdset)
	mr.Detail = append(mr.Detail, detail...)

	return mr
}

type Tags map[string]TagOption

func (t Tags) Match(env *Environment, threads uint, sds ...*SourceData) (MatchResults, error) {
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
