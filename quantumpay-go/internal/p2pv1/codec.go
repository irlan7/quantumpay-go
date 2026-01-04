package p2pv1

import "encoding/json"

func EncodeHeader(h HeaderMsg) ([]byte, error) {
	return json.Marshal(h)
}

func DecodeHeader(b []byte) (HeaderMsg, error) {
	var h HeaderMsg
	err := json.Unmarshal(b, &h)
	return h, err
}
