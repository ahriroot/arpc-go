package utils

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
	Package    []Package          `json:"package"`
	Procedures []Procedures       `json:"procedures"`
	Param      map[string][]Param `json:"param"`
}
