package contracts

import (
	"user/src/domains/entities"
	"user/src/drivers/ethereum"
	"user/src/usecases/port"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContentContract struct{}

func NewContentContracts() port.ContractPort {
	return &ContentContract{}
}

func (cc *ContentContract) ListIndexLog() ([]*entities.IndexLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	list, err := conn.ListIndexLog(&bind.CallOpts{})
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

func (cc *ContentContract) FindAuditLogByIndexID(indexID string) (*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	a, err := conn.FindAuditLogByIndexID(&bind.CallOpts{}, indexID)
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
	return conn.ListIndexID(&bind.CallOpts{})
}

func (cc *ContentContract) ListAuditLog(ids []string) ([]*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectContractNetWork()
	al, err := conn.ListAuditLog(&bind.CallOpts{}, ids)
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
