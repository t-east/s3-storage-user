package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	MetaGen(content *entities.ContentCreateMetaData) (*entities.MetaData, error)
	ListLog() ([]*entities.Log, error)
	SetKey(key []byte) error
}

type CryptPort interface {
	MakeMetaData(contentCreateMetaData *entities.ContentCreateMetaData) (*entities.MetaData, error)
	KeyGen() (*entities.Key, error)
}

type ContractPort interface {
	ListContractLog() ([]*entities.ContentLog, error)
	ListContentIDs() ([]string, error)
	GetAuditLog(string) (*entities.AuditLog, error)
	ListAuditLog([]string) ([]*entities.AuditLog, error)
	SetPubKey(key []byte) error
}
