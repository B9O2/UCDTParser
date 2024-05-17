package ucdt

import (
	"os"
	"testing"
)

func TestUCDT(t *testing.T) {
	data, err := os.ReadFile("test.toml")
	if err != nil {
		panic(err)
	}
	u, err := ParseUCDT(data, nil)
	if err != nil {
		panic(err)
	}

	s := NewSourceData("aa", map[string][]byte{
		"test": []byte("Hello World!"),
	})
	mrs := u.Match(s)
	mrs.Dump(0)

}
