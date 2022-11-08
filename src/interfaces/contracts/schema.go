package contracts

import (
	"user/src/domains/entities"
	"user/src/drivers/ethereum"
)

type IndexLogSchema struct {
	Logs []ethereum.ContractIndexLog
}

func (ccl *IndexLogSchema) BindSchema() []*entities.IndexLog {
	var inList = ccl.Logs
	var outList []*entities.IndexLog
	for _, l := range inList {
		outList = append(outList, &entities.IndexLog{
			HashedData: l.Hash,
			IndexId:    l.AuditLogId,
			Owner:      l.Owner.String(),
			Provider:   l.Provider.String(),
		})
	}
	return outList
}
