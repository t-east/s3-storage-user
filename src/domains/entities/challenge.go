package entities

type Challenge struct {
	C  int
	K1 []byte
	K2 []byte
}

type ChallengeList struct {
	DataList []Challenge
	Total    int
}
