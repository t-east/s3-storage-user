package controllers

import "user/src/domains/entities"

type ContentReq struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type MetaDataReq struct {
	Content ContentReq `json:"content"`
	PrivKey string     `json:"priv_key"`
	Address string     `json:"address"`
}

type MetaDataRes struct {
	Data [][]byte `json:"meta"`
}

func MetaDataReqToEntity(req *MetaDataReq) *entities.ContentCreateMetaData {
	return &entities.ContentCreateMetaData{
		Content: entities.Point{
			X: req.Content.X,
			Y: req.Content.Y,
		},
		PrivKey: []byte(req.PrivKey),
		Address: req.Address,
	}
}

func MetaDataEntityToRes(m *entities.MetaData) *MetaDataRes {
	return &MetaDataRes{
		Data: *m,
	}
}

type SetKeyReq struct {
	PubKey []byte `json:"pub_key"`
}

type InitIndexLogRes struct {
	ID string `json:"id"`
}

type ChallengeRes struct {
	C  int    `json:"c"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

type ProofRes struct {
	Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
}

type AuditLogRes struct {
	Challenge ChallengeRes `json:"challenge"`
	Proof     ProofRes     `json:"proof"`
	Result    bool         `json:"result"`
}

type IndexLogRes struct {
	HashedData [][]byte `json:"hash"`
	IndexId    string   `json:"index_id"`
	Owner      string   `json:"owner"`
	Provider   string   `json:"provider"`
	AuditLogId string   `json:"audit_id"`
}
type LogListRes struct {
	AuditLog AuditLogRes `json:"audit_log"`
	IndexLog IndexLogRes `json:"index_log"`
}

func LogListToRes(logs []*entities.Log) []*LogListRes {
	var res []*LogListRes
	for i := 0; i < len(logs); i++ {
		res = append(res, &LogListRes{
			AuditLog: AuditLogRes{
				Challenge: ChallengeRes{
					C:  logs[i].AuditLog.Challenge.C,
					K1: logs[i].AuditLog.Challenge.K1,
					K2: logs[i].AuditLog.Challenge.K2,
				},
				Proof: ProofRes{
					Myu:   logs[i].AuditLog.Proof.Myu,
					Gamma: logs[i].AuditLog.Proof.Gamma,
				},
				Result: logs[i].AuditLog.Result,
			},
			IndexLog: IndexLogRes{
				HashedData: logs[i].IndexLog.HashedData,
				IndexId:    logs[i].IndexLog.IndexId,
				Owner:      logs[i].IndexLog.Owner,
				Provider:   logs[i].IndexLog.Provider,
				AuditLogId: logs[i].IndexLog.AuditLogId,
			},
		})
	}
	return res
}
