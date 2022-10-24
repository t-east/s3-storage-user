package entities

type Proof struct {
	Myu       []byte `json:"myu"`
	Gamma     []byte `json:"gamma"`
}

type Proofs struct {
	DataList []Proof `json:"proofs"`
	Total    int     `json:"total"`
}