package ucdt_parser

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/B9O2/canvas/containers"
	"github.com/B9O2/evaluate"
	"github.com/BurntSushi/toml"
)

type Trait struct {
	In        []string `toml:"in"`
	Source    []string `toml:"source"`
	Start     int      `toml:"start"`
	End       int      `toml:"end"`
	Contains  []string `toml:"contains"`
	Regexps   []string `toml:"regexps"`
	Weight    float32  `toml:"weight"`
	MinWeight float32  `toml:"min_weight"`
	MaxWeight float32  `toml:"max_weight"`
}

func (t Trait) Match(data []byte) bool {
	if t.End == 0 {
		t.End = len(data)
	}
	data = InterceptData(data, t.Start, t.End)

	contains := false
	for _, str := range t.Contains {
		if bytes.Contains(data, []byte(str)) {
			contains = true
			break
		}
	}

	for _, str := range t.Contains {
		if bytes.Contains(data, []byte(str)) {
			contains = true
			break
		}
	}

	regexps := false

	for _, str := range t.Regexps {
		re := regexp.MustCompile(str)
		if re.Match(data) {
			regexps = true
			break
		}
	}

	return contains || regexps
}

type Info struct {
	Text     string   `toml:"text"`
	In       []string `toml:"in"`
	Source   []string `toml:"source"`
	Decoding []string `toml:"decoding"`
	Fetch    string   `toml:"fetch"`
	Regexps  []string `toml:"regexps"`
	Expr     string   `toml:"expression"`
}

func (info *Info) Extract(data []byte) ([]byte, error) {
	return []byte("INFO EXTRACT"), nil
}

type TagOption struct {
	Source []string         `toml:"source"`
	Traits map[string]Trait `toml:"traits"` //[name]TraitOption
	Info   map[string]Info  `toml:"info"`
	Expr   string           `toml:"expression"`
}

func (t *TagOption) Match(sds map[string]SourceData, e *evaluate.Evaluate) MatchResult {
	mr := NewMatchResult(0, map[string][]byte{}, []string{}, nil)
	weightRanges := map[string]containers.Range{}
	traitHits := map[string]bool{}

	source := t.Source
	if len(source) <= 0 {
		for src := range sds {
			source = append(source, src)
		}
	}

	//Trait
	for name, trait := range t.Traits {
		traitSource := source
		if len(trait.Source) > 0 {
			traitSource = trait.Source
		}

		var weightRange containers.Range
		hit := false
		for _, src := range traitSource {
			if sd, ok := sds[src]; ok {
				sd.Range(trait.In, func(_ string, data []byte) bool {
					if trait.Match(data) {
						hit = true
						return false
					}
					return true
				})
			}
		}

		traitHits[name] = hit
		weight := uint(trait.Weight * 100)
		minWeight := uint(trait.MinWeight * 100)
		maxWeight := uint(trait.MaxWeight * 100)
		if hit {
			if trait.Weight != 0 {
				weightRange = containers.NewRange(weight, weight)
			} else {
				weightRange = containers.NewRange(minWeight, maxWeight)
			}
		} else {
			if trait.Weight == 0 {
				weightRange = containers.NewRange(minWeight, maxWeight)
			} else {
				continue
			}
		}

		weightRanges[name] = weightRange
	}

	var score uint
	results := SpaceAllocate(weightRanges, 100)
	for name, hit := range traitHits {
		if hit {
			score += results[name]
		}
	}
	mr.score = float32(score) / 100

	//Expression
	if e != nil {
		args := map[string]any{
			"score": score,
		}

		r, err := e.Eval(t.Expr, args)
		if err != nil {
			return mr
		}
		if b, ok := r.(bool); ok {
			if b {
				mr.score = 1
			} else {
				mr.score = 0
			}
		} else {
			mr.err = errors.New("expression must return a bool,but '" + reflect.TypeOf(r).String() + "'")
		}
	} else {
		mr.detail = append(mr.detail, "[Warning] Expression Disabled")
	}

	//Info
	for name, info := range t.Info {
		infoSource := source
		if len(info.Source) > 0 {
			infoSource = info.Source
		}
		for _, src := range infoSource {
			if sd, ok := sds[src]; ok {
				sd.Range(info.In, func(_ string, data []byte) bool {
					data, err := info.Extract(data)
					if err != nil {
						mr.detail = append(mr.detail, "[Info Extract]"+name+":"+err.Error())
					} else {
						mr.info[name] = data
					}
					return true
				})
			}
		}
	}
	return mr
}

type UCDT struct {
	Tags map[string]TagOption `toml:"tags"`
}

func (u *UCDT) Patch(patchUCDT UCDT) {

}

func (u *UCDT) String() string {
	return fmt.Sprint(*u)
}

func (u *UCDT) Match(e *evaluate.Evaluate, sds ...SourceData) MatchResults {
	sdMap := map[string]SourceData{}
	for _, sd := range sds {
		sdMap[sd.source] = sd
	}

	mrs := make(MatchResults)
	for tag, opt := range u.Tags {
		mrs.Add(tag, opt.Match(sdMap, e))
	}
	return mrs
}

func ParseUCDT(data []byte) (*UCDT, error) {
	v := &UCDT{}
	_, err := toml.NewDecoder(bytes.NewReader(data)).Decode(v)
	return v, err
}
