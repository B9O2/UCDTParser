package ucdt_parser

import (
	"os"
	"testing"
)

func TestUCDT(t *testing.T) {
	data, err := os.ReadFile("test.toml")
	if err != nil {
		panic(err)
	}
	u, err := ParseUCDT(data)
	if err != nil {
		panic(err)
	}

	s := NewSourceData("aa", map[string][]byte{
		"test": []byte("Hello World!"),
	})

	mrs := u.Match(nil, s)
	mrs.Draw(0)

}
