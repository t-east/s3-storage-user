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

func (cc *ContentContract) ListContractLog() ([]*entities.ContentLog, error) {
	conn, _ := ethereum.ConnectContentNetWork()
	list, err := conn.ListContentLog(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	logSchema := ContentLogSchema{
		Logs: list,
	}
	return logSchema.BindSchema(), nil
}

func (cc *ContentContract) SetPubKey(pubKey []byte) error {
	conn, client := ethereum.ConnectPubkeyNetWork()
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

func (cc *ContentContract) GetAuditLog(id string) (*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectAuditNetWork()
	a, err := conn.GetAuditLog(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}
	return &entities.AuditLog{
		Challenge: &entities.Challenge{
			ContentId: id,
			C:         int(a.Chal),
			K1:        a.K1,
			K2:        a.K2,
		},
		Proof: &entities.Proof{
			Myu:       a.Myu,
			Gamma:     a.Gamma,
			ContentId: id,
		},
		Result:    a.Result,
		ContentID: id,
	}, nil
}

func (cc *ContentContract) ListContentIDs() ([]string, error) {
	conn, _ := ethereum.ConnectContentNetWork()
	return conn.ListContentID(&bind.CallOpts{})
}

func (cc *ContentContract) ListAuditLog(ids []string) ([]*entities.AuditLog, error) {
	conn, _ := ethereum.ConnectAuditNetWork()
	al, err := conn.ListAuditLog(&bind.CallOpts{}, ids)
	if err != nil {
		return nil, err
	}
	var logs []*entities.AuditLog
	for i := 0; i < len(al); i++ {
		logs = append(logs, &entities.AuditLog{
			Challenge: &entities.Challenge{
				ContentId: ids[i],
				C:         int(al[i].Chal),
				K1:        al[i].K1,
				K2:        al[i].K2,
			},
			Proof: &entities.Proof{
				Myu:       al[i].Myu,
				Gamma:     al[i].Gamma,
				ContentId: ids[i],
			},
			Result:    al[i].Result,
			ContentID: ids[i],
		})
	}
	return logs, nil
}
