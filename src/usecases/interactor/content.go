package interactor

import (
	entities "user/src/domains/entities"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	Crypt    port.ContentCrypt
	Contract port.ContentContract
}

func NewContentInputPort(cryptHandler port.ContentCrypt, contractHandler port.ContentContract) port.ContentInputPort {
	return &ContentHandler{
		Crypt:    cryptHandler,
		Contract: contractHandler,
	}
}

func (c *ContentHandler) MetaGen(contentIn *entities.ContentIn) (*entities.Content, error) {
	//* メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentIn)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *ContentHandler) SetKey(pubKey, ethPrivKey string) error {
	err := c.Contract.SetKey(pubKey, ethPrivKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContentHandler) GetLog() (*entities.Content, error) {
	return &entities.Content{}, nil
}
