package entities

type ArtLog struct {
	HashedData [][]byte `json:"hashed_data"`
	Owner      string   `json:"owner"`
}

type Content struct {
	Content     []byte   `json:"content"`
	MetaData    [][]byte `json:"meta_data"`
	HashedData  [][]byte `json:"hashed_data"`
	ContentName string   `json:"name"`
	SplitCount  int      `json:"split_count"`
	Owner       string   `json:"owner"`
	Id          string   `json:"id"`
	UserId      string   `json:"user_id"`
}

type ContentInput struct {
	Content     []byte `json:"content"`
	ContentName string `json:"name"`
	Owner       string `json:"owner"`
}

type ReceiptFromSP struct {
	ContentName string `json:"name"`
	Owner       string `json:"owner"`
	ArtId       string `json:"art_id"`
}

type ReceiptFromBC struct {
	ContentName string   `json:"name"`
	Owner       string   `json:"owner"`
	ArtId       string   `json:"art_id"`
	HashedData  [][]byte `json:"hashed_data"`
}

func NewContent() *Content {
	return &Content{}
}
