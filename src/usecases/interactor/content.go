package interactor

import (
	"log"
	entities "user/src/domains/entities"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	Crypt    port.CryptPort
	Contract port.ContractPort
}

func NewContentInputPort(cryptHandler port.CryptPort, contractHandler port.ContractPort) port.ContentInputPort {
	return &ContentHandler{
		Crypt:    cryptHandler,
		Contract: contractHandler,
	}
}

func (c *ContentHandler) MetaGen(contentCreateMetaData *entities.ContentCreateMetaData) (*entities.MetaData, error) {
	//* メタデータ作成
	metaData, err := c.Crypt.MakeMetaData(contentCreateMetaData)
	if err != nil {
		return nil, err
	}
	return metaData,  nil
}

func (c *ContentHandler) SetKey(pubKey []byte) error {
	err := c.Contract.SetPubKey(pubKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContentHandler) ListLog() (*entities.Log, error) {
	contentIDs, _ := c.Contract.ListContentIDs()
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
