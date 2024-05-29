package ucdt

import "fmt"

type SourceData struct {
	source   string
	contents map[string][]byte
	err      error
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

func NewSourceData(source string, contents map[string][]byte, err error) SourceData {
	return SourceData{
		source:   source,
		contents: contents,
		err:      err,
	}
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
				fmt.Println("  [Expression Hit]")
			} else {
				fmt.Println("  [Expression Not Hit]")
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

			fmt.Println(" Fetch Info:")
			for k, v := range mr.info {
				fmt.Println("  "+k, ":", string(v))
			}
		}
		return true
	})
}
