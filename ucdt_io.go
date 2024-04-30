package ucdt_parser

import "fmt"

type SourceData struct {
	source   string
	contents map[string][]byte
}

func (sd SourceData) IsValid(limits []string) bool {
	for _, limit := range limits {
		if limit == sd.source {
			return true
		}
	}
	return false
}

func (sd SourceData) Range(positions []string, f func(string, []byte) bool) {
	if len(positions) <= 0 {
		for k, v := range sd.contents {
			if !f(k, v) {
				break
			}
		}
	} else {
		for _, k := range positions {
			if _, ok := sd.contents[k]; ok {
				if !f(k, sd.contents[k]) {
					break
				}
			}
		}
	}
}

func NewSourceData(source string, contents map[string][]byte) SourceData {
	return SourceData{
		source:   source,
		contents: contents,
	}
}

type MatchResult struct {
	info   map[string][]byte
	detail []string
	score  float32
	err    error
}

func NewMatchResult(score float32, info map[string][]byte, detail []string, err error) MatchResult {
	return MatchResult{
		info:   info,
		detail: detail,
		score:  score,
		err:    err,
	}
}

type MatchResults map[string]MatchResult //[tag][name][data]

func (mrs MatchResults) Add(tag string, m MatchResult) {
	mrs[tag] = m
}

func (mrs MatchResults) Range(f func(string, MatchResult) bool) {
	for name, mr := range mrs {
		if !f(name, mr) {
			break
		}
	}
}

func (mrs MatchResults) Draw(suitability float32) {
	mrs.Range(func(name string, mr MatchResult) bool {
		if mr.score >= suitability {
			fmt.Println(name, fmt.Sprint(mr.score*100)+"%")
		}
		return true
	})
}
