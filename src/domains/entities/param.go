package entities

type Param struct {
	Pairing string `json:"x,string"`
	G []byte `json:"g"`
	U []byte `json:"u"`
}
