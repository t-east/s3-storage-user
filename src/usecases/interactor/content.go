package interactor

import (
	entities "user/src/domains/entities"
	port "user/src/usecases/port"
)

type ContentHandler struct {
	Crypt    port.CryptPort
	Contract port.ContractPort
	Random   port.RandomIDPort
}

func NewContentInputPort(cryptHandler port.CryptPort, contractHandler port.ContractPort, randomHandler port.RandomIDPort) port.ContentInputPort {
	return &ContentHandler{
		Crypt:    cryptHandler,
		Contract: contractHandler,
		Random:   randomHandler,
	}
}

func (c *ContentHandler) MetaGen(contentCreateMetaData *entities.ContentCreateMetaData) (*entities.MetaData, error) {
	//* メタデータ作成
	metaData, err := c.Crypt.MakeMetaData(contentCreateMetaData)
	if err != nil {
		return nil, err
	}
	return metaData, nil
}

func (c *ContentHandler) SetKey(pubKey []byte) error {
	err := c.Contract.SetPubKey(pubKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContentHandler) ListLog() ([]*entities.Log, error) {
	indexIds, err := c.Contract.ListIndexID()
	if err != nil {
		return nil, err
	}

	var logs []*entities.Log
	for _, l := range indexIds {
		indexLog, err := c.Contract.FindIndexLogByID(l)
		if err != nil {
			return nil, err
		}
		var auditLog *entities.AuditLog
		if indexLog.AuditLogId != "" {
			auditLog, err = c.Contract.FindAuditLogByID(indexLog.AuditLogId)
			if err != nil {
				return nil, err
			}
		}
		log := &entities.Log{
			AuditLog: auditLog,
			IndexLog: indexLog,
		}
		log.SetStatus()
		logs = append(logs, log)
	}
	return logs, nil
}

func (c *ContentHandler) InitIndexLog() (string, error) {
	id, err := c.Random.MakeRandomStr()
	if err != nil {
		return "", err
	}
	err = c.Contract.InitIndexLog(id)
	if err != nil {
		return "", err
	}
	return id, nil
}
