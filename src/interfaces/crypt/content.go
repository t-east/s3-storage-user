package crypt

import (
	"bytes"
	"encoding/gob"
	"log"
	"user/src/core"
	"user/src/domains/entities"
	"user/src/usecases/port"

	"github.com/Nik-U/pbc"
)

// type UserCrypt interface {
// 	MakeMetaData(content entities.ContentInput) (entities.Content, error)
// }

type contentCrypt struct {
	Param *entities.Param
}

func NewContentCrypt(param *entities.Param) port.ContentCrypt {
	return &contentCrypt{
		Param: param,
	}
}

func (cc *contentCrypt) MakeMetaData(uc *entities.ContentIn) (*entities.Content, error) {
	pairing, err := pbc.NewPairingFromString(cc.Param.Pairing)
	if err != nil {
		return nil, err
	}
	var cb bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&cb) // Will write to network.
	err = enc.Encode(uc.Content)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	contentByte := cb.Bytes()
	u := pairing.NewG1().SetBytes([]byte(cc.Param.U))
	splitCount := 3
	splitedFile, err := core.SplitSlice(contentByte, splitCount)
	if err != nil {
		return nil, err
	}
	privKey := pairing.NewZr().SetBytes([]byte(uc.PrivKey))

	// メタデータの作成
	var metaData [][]byte
	metaToHash := ""
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])

		mm := core.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetBytes(mm)

		um := pairing.NewG1().PowBig(u, m.X())
		temp := pairing.NewG1().Mul(um, M)
		meta := pairing.NewG1().MulZn(temp, privKey)

		metaData = append(metaData, meta.Bytes())
		metaToHash = metaToHash + meta.String()
	}

	return &entities.Content{
		Content:    uc.Content,
		MetaData:   metaData,
	}, nil
}

func (cc *contentCrypt) KeyGen() (*entities.Key, error) {
	// pairing, err := pbc.NewPairingFromString(cc.Param.Pairing)
	// if err != nil {
	// 	return nil, err
	// }
	// g := pairing.NewG1().SetBytes(cc.Param.G)
	// privKey := pairing.NewZr().Rand()
	// pubKey := pairing.NewG1().MulZn(g, privKey)
	// return &entities.Key{
	// 	PubKey:  pubKey.Bytes(),
	// 	PrivKey: privKey.Bytes(),
	// }, nil
	return &entities.Key{}, nil
}
