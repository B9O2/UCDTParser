package ucdt_parser

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type Trait struct {
	decoding []string `toml:"decoding"`
	start    int      `toml:"start"`
	end      int      `toml:"end"`
	contains []string `toml:"contains"`
	regexps  []string `toml:"regexps"`
	expr     string   `toml:"expression"`
}

type Info struct {
	text      string   `toml:"text"`
	positions []string `toml:"positions"`
	decoding  []string `toml:"decoding"`
	extract   string   `toml:"extract"`
	regexps   []string `toml:"regexps"`
	expr      string   `toml:"expression"`
}

type TagOption struct {
	source []string                    `toml:"source"`
	format string                      `toml:"format"`
	trait  map[string]map[string]Trait `toml:"trait"` //[position][name]TraitOption
	info   map[string]Info             `toml:"info"`
}

type UCDT struct {
	tags map[string]TagOption `toml:"tags"`
}

func (u *UCDT) Patch(patchUCDT UCDT) {

}

func (u *UCDT) Match(sd SourceData) MatchResult {

}

func ParseUCDT(data []byte) (*UCDT, error) {
	v := &UCDT{}
	_, err := toml.NewDecoder(bytes.NewReader(data)).Decode(v)
	return v, err
}

type MatchResult map[string]map[string]string //[tag][name][data]

type SourceData struct {
	source  string
	content []byte
}

func NewSourceData(source string, content []byte) SourceData {
	return SourceData{
		source:  source,
		content: content,
	}
}
