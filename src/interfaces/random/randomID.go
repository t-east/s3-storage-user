package random

import (
	"crypto/rand"
	"errors"
	"user/src/usecases/port"
)

type contentRandom struct{}

func NewRandomID() port.RandomIDPort {
	return &contentRandom{}
}

func (c contentRandom) MakeRandomStr() (string, error) {
	const letters = "abcdefghijk"

	// 乱数を生成
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
