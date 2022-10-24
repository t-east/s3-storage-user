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
	for i := 0; i < len(inList); i++ {
		outList = append(outList, &entities.IndexLog{
			HashedData: inList[i].Hash,
			IndexId:    inList[i].AuditLogId,
			Owner:      inList[i].Owner.String(),
		})
	}
	return outList
}
