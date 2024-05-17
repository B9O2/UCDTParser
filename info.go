package ucdt

import (
	"fmt"
	"regexp"

	"github.com/google/cel-go/common/types"
)

type Info struct {
	Text    string     `toml:"text"`
	In      []string   `toml:"in"`
	Source  []string   `toml:"source"`
	Start   int        `toml:"start"`
	End     int        `toml:"end"`
	Regexps []string   `toml:"regexps"`
	Expr    Expression `toml:"expression"`
}

func (info *Info) Extract(data []byte) ([]byte, error) {
	var result []byte
	if info.End == 0 {
		info.End = len(data)
	}
	data = InterceptData(data, info.Start, info.End)

	for _, str := range info.Regexps {
		re, err := regexp.Compile(str)
		if err != nil {
			return []byte{}, err
		}
		result = re.Find(data)
		if result != nil {
			break
		}
	}
	return result, nil
}

type InfoMap map[string]Info

func (im InfoMap) Extract(score float32, sds map[string]SourceData) (map[string][]byte, []string) {
	var detail []string
	result := map[string][]byte{}

	e, err := NewEvaluate(map[string]any{}, make(map[*types.Type]map[string]any))
	if err != nil {
		detail = append(detail, fmt.Sprintf("[Info Extract] error:%s", err))
	}

	args, err := GenArgs(e, score, sds)
	if err != nil {
		detail = append(detail, fmt.Sprintf("[Info Extract] error:%s", err))
	}
	for name, info := range im {
		if len(info.Expr) > 0 {
			r, d := info.Expr.Eval(e, args)
			detail = append(detail, d...)
			switch r := r.(type) {
			case nil:
				result[name] = []byte(info.Text)
			case []byte:
				result[name] = r
			case string:
				result[name] = []byte(r)
			default:
				result[name] = []byte(fmt.Sprint(r))
			}
			continue
		}

		extractor := func(_ string, data []byte) bool {
			data, err := info.Extract(data)
			if err != nil {
				detail = append(detail, fmt.Sprintf("[Info Extract]%s error: %s", name, err))
			} else {
				result[name] = data
			}
			return true
		}
		if len(info.Source) > 0 {
			for _, src := range info.Source {
				if sd, ok := sds[src]; ok {
					sd.Range(info.In, extractor)
				}
			}
		} else {
			for _, sd := range sds {
				sd.Range(info.In, extractor)
			}
		}
	}
	return result, detail
}
