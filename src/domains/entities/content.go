package entities

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Content struct {
	Content    Point   `json:"content"`
	MetaData   [][]byte `json:"meta_data"`
}

type ContentIn struct {
	Content Point `json:"content"`
	PrivKey string `json:"priv_key"`
	Address string `json:"address"`
}

type Key struct {
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}

type Param struct {
	Pairing string `json:"pairing"`
	G       string `json:"g"`
	U       string `json:"u"`
}
