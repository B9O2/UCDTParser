package ucdt

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/cel-go/common/types"
)

func DUMP() bool {
	fmt.Println("Global Func OK")
	return true
}

func TestUCDT(t *testing.T) {
	data, err := os.ReadFile("test.toml")
	if err != nil {
		panic(err)
	}
	u, err := ParseUCDT(data)
	if err != nil {
		panic(err)
	}

	s := SourceData(map[string][]byte{
		"test": []byte("Hello World!"),
	})

	mrs, err := u.Tags.Match(NewEnviroment(map[string]any{
		"DUMP": DUMP,
	}, map[*types.Type]map[string]any{
		types.BytesType: {
			"help": func(self []byte) bool {
				fmt.Println("Method OK")
				return true
			},
		},
	}), 20, s)
	if err != nil {
		fmt.Println(err)
	}
	mrs.Dump(0)

}
