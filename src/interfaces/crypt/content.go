package crypt

import (
	"user/src/core"
	"user/src/domains/entities"
	"user/src/usecases/port"

	"github.com/Nik-U/pbc"
)

// type UserCrypt interface {
// 	MakeMetaData(content entities.ContentInput) (entities.Content, error)
// }

type contentCrypt struct {
	Param entities.Param
}

func NewUserCrypt(param entities.Param) port.ContentCrypt {
	return &contentCrypt{
		Param: param,
	}
}

func (cc *contentCrypt) MakeMetaData(uc *entities.ContentInput) (*entities.Content, error) {
	pairing, err := pbc.NewPairingFromString(uc.Param.Pairing)
	if err != nil {
		return nil, err
	}
	u := pairing.NewG1().SetBytes([]byte(uc.Param.U))
	splitCount := 3
	splitedFile, err := core.SplitSlice(uc.Content, splitCount)
	if err != nil {
		return nil, err
	}
	privKey := pairing.NewZr().SetBytes([]byte(uc.Key.PrivKey))

	// メタデータの作成
	var metaData [][]byte
	var hashData [][]byte
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
		hashData = append(hashData, mm)
	}

	return &entities.Content{
		Content:     []byte{},
		MetaData:    metaData,
		HashedData:  hashData,
		ContentName: "",
		SplitCount:  splitCount,
		Owner:       "",
		Id:          "",
		UserId:      "",
	}, nil
}
