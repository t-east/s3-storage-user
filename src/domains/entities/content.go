package entities

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MetaData [][]byte
type Content struct {
	Content  Point    `json:"content"`
	MetaData MetaData `json:"meta_data"`
}

type ContentCreateMetaData struct {
	Content Point  `json:"content"`
	PrivKey []byte `json:"priv_key"`
	Address string `json:"address"`
}

