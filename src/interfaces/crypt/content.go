package crypt

import (
	"pairing_test/src/user/domains/entities"
	contracts "pairing_test/src/user/interfaces/contracts"
	"pairing_test/src/user/usecases/port"
)

// type UserCrypt interface {
// 	MakeMetaData(content entities.ContentInput) (entities.Content, error)
// }

type contentCrypt struct {
	Param contracts.Param
}

func NewUserCrypt(param contracts.Param) port.ContentCrypt {
	return &contentCrypt{
		Param: param,
	}
}

func (cc *contentCrypt) MakeMetaData(contentInput *entities.ContentInput) (*entities.Content, error) {
	result := &entities.Content{}
	return result, nil
}
