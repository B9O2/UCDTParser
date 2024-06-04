package ucdt

import (
	"bytes"
	"fmt"

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

func NewEnviroment(funcs map[string]any, memberMethods map[*types.Type]map[string]any) *Environment {
	if funcs == nil {
		funcs = map[string]any{}
	}
	if memberMethods == nil {
		memberMethods = map[*types.Type]map[string]any{}
	}
	env := &Environment{
		funcs, memberMethods,
	}
	return env
}

type TagOption struct {
	Source []string   `toml:"source"`
	Traits TraitMap   `toml:"traits"` //[name]TraitOption
	Info   InfoMap    `toml:"info"`
	Expr   Expression `toml:"expression"`
}

func (t *TagOption) Match(env *Environment, allSds map[string]SourceData) MatchResult {
	mr := NewMatchResult()

	sds := map[string]SourceData{}
	if len(t.Source) > 0 {
		for _, src := range t.Source {
			if data, ok := allSds[src]; ok {
				sds[src] = data
			} else {
				mr.detail = append(mr.detail, fmt.Sprintf("source '%s' has no data", src))
			}
		}
	} else {
		sds = allSds
		mr.detail = append(mr.detail, "using all source data")
	}

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

		args, err := GenArgs(e, mr.score, sds)
		if err != nil {
			mr.detail = append(detail, fmt.Sprintf("[Expression Eval] error:%s", err))
		}

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

func (t Tags) Match(env *Environment, sds ...SourceData) MatchResults {
	sdMap := map[string]SourceData{}
	for _, sd := range sds {
		sdMap[sd.source] = sd
	}

	mrs := NewMatchResults()
	for tag, opt := range t {
		mrs.Add(tag, opt.Match(env, sdMap))
	}
	return mrs
}

type UCDT struct {
	Tags Tags `toml:"tags"`
}

func ParseUCDT(data []byte) (*UCDT, error) {
	v := &UCDT{}
	_, err := toml.NewDecoder(bytes.NewReader(data)).Decode(v)
	return v, err
}
