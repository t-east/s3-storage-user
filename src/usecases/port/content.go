package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	MetaGen(content *entities.ContentIn) (*entities.Content, error)
	GetLog() (*entities.Content, error)
	SetKey(key, ethPrivKey string)  error
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentIn) (*entities.Content, error)
	KeyGen() (*entities.Key, error)
}

type ContentContract interface {
	GetContractLog() ([]*entities.Content, error)
	SetKey(key, address string) error
}
