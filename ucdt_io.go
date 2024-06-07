package ucdt

import (
	"fmt"
	"strings"
)

type SourceData map[string][]byte

func (sd SourceData) Range(positions []string, f func(string, []byte) bool) {
	if len(positions) <= 0 {
		if raw, ok := sd["raw"]; ok {
			f("raw", raw)
		} else {
			for k, v := range sd {
				if !f(k, v) {
					break
				}
			}
		}
	} else {
		for _, k := range positions {
			if _, ok := sd[k]; ok {
				if !f(k, sd[k]) {
					break
				}
			}
		}
	}
}

func (sd SourceData) ToString(ignoreParts ...string) string {
	builder := strings.Builder{}
	headLine := fmt.Sprintf("============= Ignore: %s ==============\n", ignoreParts)
	builder.WriteString(headLine)
	positions := []string{}
	for k := range sd {
		positions = append(positions, k)
	}
	sd.Range(positions, func(s string, b []byte) bool {
		contains := false
		for _, part := range ignoreParts {
			if strings.Contains(part, s) {
				contains = true
				break
			}
		}
		if !contains {
			builder.WriteString(fmt.Sprintf("\n[%s]:%s\n", s, string(b)))
		}
		return true
	})
	builder.WriteString(strings.Repeat("=", len(headLine)))

	return builder.String()
}

func (sd SourceData) String() string {
	return sd.ToString()
}

type MatchResult struct {
	info        map[string][]byte
	scoreDetail map[bool]map[string]float32
	detail      []string
	score       float32
	expression  bool
	err         error
}

func NewMatchResult() MatchResult {
	return MatchResult{
		info: map[string][]byte{},
		scoreDetail: map[bool]map[string]float32{
			true:  {},
			false: {},
		},
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

func (mrs MatchResults) Dump(suitability float32) {
	mrs.Range(func(name string, mr MatchResult) bool {
		if mr.score >= suitability {
			title := fmt.Sprintf("%s %.1f%%", name, mr.score*100)
			fmt.Println(title)
			if mr.expression {
				fmt.Println("  \\_ Expression Hit")
			} else {
				fmt.Println("  \\_[x] Expression Hit")
			}
			for name, i := range mr.scoreDetail[true] {
				fmt.Println("  \\_", name, i)
			}
			for name, i := range mr.scoreDetail[false] {
				fmt.Println("  \\_[x]", name, i)
			}
			if mr.err != nil {
				fmt.Printf(" [*]%s\n", mr.err)
			}
			for _, i := range mr.detail {
				fmt.Println(" [!]", i)
			}

			if len(mr.info) > 0 {
				fmt.Println(" Fetch Info:")
				for k, v := range mr.info {
					fmt.Println("  "+k, ":", string(v))
				}
			}

			fmt.Println()
		}
		return true
	})
}

func NewMatchResults() MatchResults {
	return make(MatchResults)
}
