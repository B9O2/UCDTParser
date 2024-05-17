package ucdt

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/B9O2/canvas/containers"
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

func (t Trait) Match(data []byte) (bool, error) {
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
		re, err := regexp.Compile(str)
		if err != nil {
			return false, err
		}
		if re.Match(data) {
			regexps = true
			break
		}
	}

	return contains || regexps, nil
}

type TraitMap map[string]Trait

func (tm TraitMap) Match(sds map[string]SourceData) (map[bool]map[string]float32, []string) {
	var detail []string
	weightRanges := map[string]containers.Range{}
	traitHits := map[string]bool{}
	result := map[bool]map[string]float32{
		true:  {},
		false: {},
	}

	for name, trait := range tm {
		var weightRange containers.Range
		hit := false
		matcher := func(_ string, data []byte) bool {
			if ok, err := trait.Match(data); err != nil {
				detail = append(detail, fmt.Sprintf("[Trait Match]%s error:%s", name, err))
				return true
			} else if ok {
				hit = true
				return false
			} else {
				return true
			}
		}

		if len(trait.Source) > 0 {
			for _, src := range trait.Source {
				if sd, ok := sds[src]; ok {
					sd.Range(trait.In, matcher)
				}
			}
		} else {
			for _, sd := range sds {
				sd.Range(trait.In, matcher)
			}
		}
		traitHits[name] = hit
		weight := uint(trait.Weight * 100)
		minWeight := uint(trait.MinWeight * 100)
		maxWeight := uint(trait.MaxWeight * 100)
		if trait.Weight != 0 {
			weightRange = containers.NewRange(weight, weight)
		} else {
			weightRange = containers.NewRange(minWeight, maxWeight)
		}
		weightRanges[name] = weightRange
	}
	results := SpaceAllocate(weightRanges, 100)
	for name, hit := range traitHits {
		s := float32(results[name]) / 100
		result[hit][name] = s
	}
	return result, detail
}
