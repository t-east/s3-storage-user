package crypt

import (
	"user/src/core"
	"user/src/domains/entities"
	"user/src/interfaces/crypt/schema"
	"user/src/usecases/port"
)

// type UserCrypt interface {
// 	MakeMetaData(content entities.ContentCreateMetaDataput) (entities.Content, error)
// }

type contentCrypt struct {
	Param *entities.Param
}

func NewContentCrypt(param *entities.Param) port.CryptPort {
	return &contentCrypt{
		Param: param,
	}
}

func (cc *contentCrypt) MakeMetaData(uc *entities.ContentCreateMetaData) (*entities.MetaData, error) {
	//* Gen Param
	param, err := schema.GeneratePairingParam(cc.Param)
	if err != nil {
		return nil, err
	}
	key, err := param.GeneratePairingKey(uc.PrivKey)
	if err != nil {
		return nil, err
	}

	//* Slice Content
	splitCount := 3
	splitFile, err := schema.SplitContent(uc.Content, splitCount)
	if err != nil {
		return nil, err
	}

	//* Gen MetaData
	var metaData [][]byte
	metaToHash := ""
	for i := 0; i < len(splitFile); i++ {
		m := param.Pairing.NewG1().SetFromHash(splitFile[i])

		mm := core.GetBinaryBySHA256(m.X().String())
		M := param.Pairing.NewG1().SetBytes(mm)

		um := param.Pairing.NewG1().PowBig(param.U, m.X())
		temp := param.Pairing.NewG1().Mul(um, M)
		meta := param.Pairing.NewG1().MulZn(temp, key.PrivKey)

		metaData = append(metaData, meta.Bytes())
		metaToHash = metaToHash + meta.String()
	}

	return (*entities.MetaData)(&metaData), nil
}

func (cc *contentCrypt) KeyGen() (*entities.Key, error) {
	//* Gen Param
	param, err := schema.GeneratePairingParam(cc.Param)
	if err != nil {
		return nil, err
	}
	privKey := param.Pairing.NewZr().Rand()
	pubKey := param.Pairing.NewG1().MulZn(param.G, privKey)
	return &entities.Key{
		PubKey:  pubKey.Bytes(),
		PrivKey: privKey.Bytes(),
	}, nil
}
