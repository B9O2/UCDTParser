package builtin

import (
	"bytes"

	"github.com/B9O2/UCDTParser/ucdt_expr"
)

func Contains(all *ucdt_expr.SourceDataset, position string, substr any) bool {
	var sub []byte
	switch v := substr.(type) {
	case string:
		sub = []byte(v)
	case []byte:
		sub = v
	default:
		return false
	}

	if len(position) > 0 {
		for _, sd := range all.Dataset {
			if v, ok := sd.Data[position]; ok {
				if bytes.Contains(v, sub) {
					return true
				}
			}
		}
	} else {
		for _, sd := range all.Dataset {
			for _, v := range sd.Data {
				if bytes.Contains(v, sub) {
					return true
				}
			}
		}
	}
	return false
}
