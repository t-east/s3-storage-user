package entities

type Point struct {
	X int
	Y int
}

type MetaData [][]byte
type Content struct {
	Content  Point
	MetaData MetaData
}

type ContentCreateMetaData struct {
	Content Point
	PrivKey []byte
	Address string
}
