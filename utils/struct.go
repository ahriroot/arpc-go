package utils

import "encoding/json"

type Base struct{}

func (b *Base) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Base) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

type Package struct {
	Language string `json:"language"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

type Procedures struct {
	Name     string `json:"name"`
	Index    int    `json:"index"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Param struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Index int    `json:"index"`
}

type ArpcMeta struct {
	Version    string             `json:"version"`
	Unique     string             `json:"unique"`
	Package    []Package          `json:"package"`
	Procedures []Procedures       `json:"procedures"`
	Param      map[string][]Param `json:"param"`
}
