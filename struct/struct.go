package structs

import "encoding/json"

type Base struct{}

func (b *Base) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Base) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}
