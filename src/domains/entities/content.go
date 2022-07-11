package entities

type Point struct {
	X int `json:"x,string"`
	Y int `json:"y,string"`
}

type Content struct {
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"meta_data"`
}

type ContentIn struct {
	Content Point  `json:"content"`
	PrivKey []byte `json:"priv_key"`
	Address string `json:"address"`
}

type ContentInBlockChain struct {
	HashedData [][]byte `json:"hashed_data"`
	ContentId  string   `json:"content_id"`
	Owner      string   `json:"owner"`
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

type Proof struct {
	Myu       []byte `json:"myu"`
	Gamma     []byte `json:"gamma"`
	ContentId string `json:"content_id"`
}

type Proofs struct {
	DataList []Proof `json:"proofs"`
	Total    int     `json:"total"`
}

type Chals struct {
	DataList []Chal `json:"data"`
	Total    int    `json:"total"`
}

type Chal struct {
	ContentId string `json:"art_id"`
	C         int    `json:"ck"`
	K1        []byte `json:"k1"`
	K2        []byte `json:"k2"`
}

type AuditLog struct {
	Chal      *Chal  `json:"chal"`
	Proof     *Proof `json:"proof"`
	Result    bool   `json:"result"`
	ContentID string `json:"content_id"`
}

type Log struct {
	AuditLog   []*AuditLog            `json:"audit_log"`
	ContentLog []*ContentInBlockChain `json:"content_log"`
}

type ContentForMock struct {
	ID       string   `json:"id"`
	Address  string   `json:"address"`
	Content  Point    `json:"content"`
	MetaData [][]byte `json:"metadata"`
	HashData [][]byte `json:"hashdata"`
}