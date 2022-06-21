package entities

type SampleData struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
}

type Content struct {
	Content    SampleData   `json:"content"`
	MetaData   [][]byte `json:"meta_data"`
}

type ContentIn struct {
	Content SampleData `json:"content"`
	PrivKey string `json:"priv_key"`
	Address string `json:"address"`
}

type Key struct {
	PubKey  []byte `json:"pub_key"`
	PrivKey []byte `json:"priv_key"`
}

type Param struct {
	Pairing string `json:"pairing"`
	G       []byte `json:"g"`
	U       []byte `json:"u"`
}
