package contracts

import (
	"user/src/domains/entities"
	"user/src/drivers/ethereum"
	"user/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContentContract struct{}

func NewContentContracts() port.ContentContract {
	return &ContentContract{}
}

func (cc *ContentContract) GetContractLog() ([]*entities.Content, error) {
	conn, _ := ethereum.ConnectContentNetWork()
	list, err := conn.ListContentLog(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	var logs []*entities.Content
	for i:=0;i<len(list);i++{
		logs = append(logs, &entities.Content{})
	}
	return logs, nil
}

func (cc *ContentContract) SetKey(key, privKey string) error {
	conn, client := ethereum.ConnectPubkeyNetWork()
	auth, err := ethereum.AuthUser(client, privKey)
	if err != nil {
		return err
	}
	_, err = conn.SetPubkey(auth, key)
	if err != nil {
		return err
	}
	return nil
}