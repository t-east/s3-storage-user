package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	MetaGen(content *entities.ContentIn) (*entities.Content, []byte, []byte, error)
	GetLog() (*entities.Log, error)
	SetKey(key []byte) error
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentIn) (*entities.Content, []byte, []byte, error)
	KeyGen() (*entities.Key, error)
}

type ContentContract interface {
	ListContractLog() ([]*entities.ContentInBlockChain, error)
	ListContentIDs() ([]string, error)
	GetAuditLog(string) (*entities.AuditLog, error)
	ListAuditLog([]string) ([]*entities.AuditLog, error)
	SetKey(key []byte) error
}
