package entities

type Proof struct {
	Myu   []byte
	Gamma []byte
}

type Proofs struct {
	DataList []Proof
	Total    int
}
