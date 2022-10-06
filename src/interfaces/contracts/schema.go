package contracts

import (
	"user/src/domains/entities"
	eth "user/src/drivers/ethereum/content"
)

type ContentLogSchema struct {
	Logs []eth.ContentContentLog
}

func (ccl *ContentLogSchema) BindSchema() []*entities.ContentLog {
	var inList = ccl.Logs
	var outList []*entities.ContentLog
	for i := 0; i < len(inList); i++ {
		outList = append(outList, &entities.ContentLog{
			HashedData: inList[i].Hash,
			ContentId:  inList[i].LogId,
		})
	}
	return outList
}
