package interactor

import (
	"log"
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

func (c *ContentHandler) MetaGen(contentIn *entities.ContentIn) (*entities.Content, []byte, []byte, error) {
	//* メタデータ作成
	content, privKey, pubKey, err := c.Crypt.MakeMetaData(contentIn)
	if err != nil {
		return nil, nil, nil, err
	}
	return content, privKey, pubKey, nil
}

func (c *ContentHandler) SetKey(pubKey []byte) error {
	err := c.Contract.SetKey(pubKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContentHandler) GetLog() (*entities.Log, error) {
	contentIDs, _:= c.Contract.ListContentIDs()
	log.Print(contentIDs)
	// cl, err := c.Contract.ListContractLog()
	// if err != nil {
	// 	return nil, err
	// }
	// al, err := c.Contract.ListAuditLog(contentIDs)
	// if err != nil {
	// 	return nil, err
	// }

	return &entities.Log{
		// AuditLog:   al,
		// ContentLog: cl,
	}, nil
}
