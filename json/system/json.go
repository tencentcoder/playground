package system

import (
	"encoding/json"
	"github.com/intel-go/fastjson"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	JSON     = 0
	FastJSON = 1
	FFJSON   = 2
)

type OneLevel struct {
	Info string `json:"info"`
}

type TwoInfo struct {
	Address string `json:"address"`
}

type TwoLevel struct {
	Info TwoInfo `json:"info"`
}

type ThreeAddress struct {
	City string `json:"city"`
}

type ThreeInfo struct {
	Address ThreeAddress `json:"address"`
}

type ThreeLevel struct {
	Info ThreeInfo `json:"info"`
}

func Parse(data []byte, v interface{}, mode int) error {
	switch mode {
	case FastJSON:
		return fastjson.Unmarshal(data, v)
	case JSON:
		return json.Unmarshal(data, v)
	case FFJSON:
		return ffjson.Unmarshal(data, v)
	default:
		return nil
	}
}

func UnParse(v interface{}, mode int) ([]byte, error) {
	switch mode {
	case FastJSON:
		return fastjson.Marshal(v)
	case JSON:
		return json.Marshal(v)
	case FFJSON:
		return ffjson.Marshal(v)
	default:
		return nil, nil
	}
}
