package interactor

import (
	entities "user/src/domains/entities"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	OutputPort port.ContentOutputPort
	Crypt      port.ContentCrypt
}

func NewContentInputPort(cryptHandler port.ContentCrypt) port.ContentInputPort {
	return &ContentHandler{
		Crypt:      cryptHandler,
	}
}

func (c *ContentHandler) Upload(contentIn *entities.ContentIn) (*entities.Content, error) {
	//* メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentIn)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentHandler) GetKey() (*entities.Key, error) {
	key, err := c.Crypt.KeyGen()
	if err != nil {
		return nil, err
	}
	return key, nil
}
