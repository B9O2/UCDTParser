package builtin

import (
	"bytes"

	"github.com/B9O2/UCDTParser/expression"
)

func Contains(all *expression.SourceDataList, position string, substr any) bool {
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
		for _, sd := range all.Sds {
			if v, ok := sd.Data[position]; ok {
				if bytes.Contains(v, sub) {
					return true
				}
			}
		}
	} else {
		for _, sd := range all.Sds {
			for _, v := range sd.Data {
				if bytes.Contains(v, sub) {
					return true
				}
			}
		}
	}
	return false
}
