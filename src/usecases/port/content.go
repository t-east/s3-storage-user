package port

import (
	entities "user/src/domains/entities"
)

type ContentInputPort interface {
	MetaGen(content *entities.ContentIn) (*entities.Content, error)
	GetKey() (*entities.Key, error)
}

type ContentOutputPort interface {
	Render(*entities.Content, int)
	RenderKey(*entities.Key, int)
	RenderError(error)
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentIn) (*entities.Content, error)
	KeyGen() (*entities.Key, error)
}

type EthContract interface {
	GetContractLog() ([]*entities.Content, error)
}
