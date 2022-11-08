package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	MetaGen(content *entities.ContentCreateMetaData) (*entities.MetaData, error)
	ListLog() ([]*entities.Log, error)
	InitIndexLog() (string, error)
	SetKey(key []byte) error
}

type CryptPort interface {
	MakeMetaData(contentCreateMetaData *entities.ContentCreateMetaData) (*entities.MetaData, error)
	KeyGen() (*entities.Key, error)
}

type ContractPort interface {
	ListIndexLog() ([]*entities.IndexLog, error)
	ListIndexID() ([]string, error)
	InitIndexLog(indexID string) error
	FindAuditLogByIndexID(indexID string) (*entities.AuditLog, error)
	ListAuditLog(indexIDs []string) ([]*entities.AuditLog, error)
	SetPubKey(key []byte) error
	FindIndexLogByID(indexID string) (*entities.IndexLog, error)
	FindAuditLogByID(auditID string) (*entities.AuditLog, error)
}

type RandomIDPort interface {
	MakeRandomStr() (string, error)
}
