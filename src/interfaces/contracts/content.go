package contracts

import (
	"user/src/domains/entities"
	"user/src/drivers/ethereum"
	"user/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ContentContract struct{}

func NewContentContracts() port.ContractPort {
	return &ContentContract{}
}

func (cc *ContentContract) ListIndexLog() ([]*entities.IndexLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	list, err := conn.ListIndexLog(&bind.CallOpts{From: common.HexToAddress(ethereum.EthAddress)})
	if err != nil {
		return nil, err
	}
	logSchema := IndexLogSchema{
		Logs: list,
	}
	return logSchema.BindSchema(), nil
}

func (cc *ContentContract) SetPubKey(pubKey []byte) error {
	conn, client := ethereum.ConnectContractNetWork()
	auth, err := ethereum.AuthUser(client, ethereum.EthPrivKey)
	if err != nil {
		return err
	}
	_, err = conn.SetPubkey(auth, pubKey)
	if err != nil {
		return err
	}
	return nil
}

func (cc *ContentContract) InitIndexLog(indexId string, hash [][]byte) error {
	conn, client := ethereum.ConnectContractNetWork()
	auth, err := ethereum.AuthUser(client, ethereum.EthPrivKey)
	if err != nil {
		return err
	}
	_, err = conn.InitIndexLog(auth, indexId, hash)
	if err != nil {
		return err
	}
	return nil
}

func (cc *ContentContract) FindAuditLogByIndexID(indexID string) (*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	add := ethereum.GetUserAddress(ethereum.EthPrivKey)
	a, err := conn.FindAuditLogByIndexID(&bind.CallOpts{From: add}, indexID)
	if err != nil {
		return nil, err
	}
	return &entities.AuditLog{
		Challenge: &entities.Challenge{
			C:  int(a.Chal),
			K1: a.K1,
			K2: a.K2,
		},
		Proof: &entities.Proof{
			Myu:   a.Myu,
			Gamma: a.Gamma,
		},
		Result: a.Result,
	}, nil
}

func (cc *ContentContract) ListIndexID() ([]string, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	add := ethereum.GetUserAddress(ethereum.EthPrivKey)
	return conn.ListIndexID(&bind.CallOpts{From: add})
}

func (cc *ContentContract) ListAuditLog(ids []string) ([]*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	add := ethereum.GetUserAddress(ethereum.EthPrivKey)
	al, err := conn.ListAuditLog(&bind.CallOpts{From: add}, ids)
	if err != nil {
		return nil, err
	}
	var logs []*entities.AuditLog
	for i := 0; i < len(al); i++ {
		logs = append(logs, &entities.AuditLog{
			Challenge: &entities.Challenge{
				C:  int(al[i].Chal),
				K1: al[i].K1,
				K2: al[i].K2,
			},
			Proof: &entities.Proof{
				Myu:   al[i].Myu,
				Gamma: al[i].Gamma,
			},
			Result: al[i].Result,
		})
	}
	return logs, nil
}
