package interactor

import (
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

func (c *ContentHandler) ListLog() ([]*entities.Log, error) {
	contentIDs, _ := c.Contract.ListIndexID()
	il, err := c.Contract.ListIndexLog()
	if err != nil {
		return nil, err
	}
	al, err := c.Contract.ListAuditLog(contentIDs)
	if err != nil {
		return nil, err
	}

	var logs []*entities.Log
	for i := 0; i < len(il) ; i++ {
		logs = append(logs, &entities.Log{
			AuditLog:   al[i],
			IndexLog: il[i],
		})
	}
	return logs, nil
}
