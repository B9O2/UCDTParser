package ucdt

import (
	"fmt"
	"strings"
)

type SourceData struct {
	data   map[string][]byte
	source string
}

func (sd SourceData) Range(positions []string, f func(string, []byte) bool) {
	if len(positions) <= 0 {
		if raw, ok := sd.data["raw"]; ok {
			f("raw", raw)
		} else {
			for k, v := range sd.data {
				if !f(k, v) {
					break
				}
			}
		}
	} else {
		for _, k := range positions {
			if _, ok := sd.data[k]; ok {
				if !f(k, sd.data[k]) {
					break
				}
			}
		}
	}
}

func (sd SourceData) ToString(ignoreParts ...string) string {
	builder := strings.Builder{}
	source := sd.source
	if len(source) <= 0 {
		source = "No Source"
	}
	headLine := fmt.Sprintf("============= Source: %s == Ignore: %s ==============\n", source, ignoreParts)
	builder.WriteString(headLine)
	positions := []string{}
	for k := range sd.data {
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

func NewSourceData(source string, data map[string][]byte) *SourceData {
	sd := &SourceData{
		data:   data,
		source: source,
	}
	if sd.data == nil {
		sd.data = map[string][]byte{}
	}
	return sd
}

func NewNoSourceData(dataset ...map[string][]byte) []*SourceData {
	sds := []*SourceData{}
	for _, data := range dataset {
		sds = append(sds, NewSourceData("", data))
	}
	return sds
}

type MatchResult struct {
	Info        map[string][]byte
	ScoreDetail map[bool]map[string]float32
	Detail      []string
	Score       float32
	Expression  bool
	Err         error
}

func NewMatchResult() MatchResult {
	return MatchResult{
		Info: map[string][]byte{},
		ScoreDetail: map[bool]map[string]float32{
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
		if mr.Score >= suitability {
			title := fmt.Sprintf("%s %.1f%%", name, mr.Score*100)
			fmt.Println(title)
			if mr.Expression {
				fmt.Println("  \\_ Expression Hit")
			} else {
				fmt.Println("  \\_[x] Expression Hit")
			}
			for name, i := range mr.ScoreDetail[true] {
				fmt.Println("  \\_", name, i)
			}
			for name, i := range mr.ScoreDetail[false] {
				fmt.Println("  \\_[x]", name, i)
			}
			if mr.Err != nil {
				fmt.Printf(" [*]%s\n", mr.Err)
			}
			for _, i := range mr.Detail {
				fmt.Println(" [!]", i)
			}

			if len(mr.Info) > 0 {
				fmt.Println(" Fetch Info:")
				for k, v := range mr.Info {
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
