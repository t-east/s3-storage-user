package entities

type Content struct {
	Content     []byte   `json:"content"`
	MetaData    [][]byte `json:"meta_data"`
	HashedData  [][]byte `json:"hashed_data"`
	ContentName string   `json:"name"`
	SplitCount  int      `json:"split_count"`
}

type ContentInput struct {
	Content     []byte `json:"content"`
	ContentName string `json:"name"`
	Owner       string `json:"owner"`
	Param       *Param `json:"param"`
	Key         *Key   `json:"key"`
}

type Key struct {
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}

type Param struct {
	Pairing string `json:"paring"`
	G       []byte `json:"g"`
	U       []byte `json:"u"`
}
