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
	Status   string      `json:"status"`
}

func LogListToRes(logs []*entities.Log) []*LogListRes {
	var res []*LogListRes
	for _, l := range logs {
		res = append(res, &LogListRes{
			AuditLog: AuditLogRes{
				Challenge: ChallengeRes{
					C:  l.AuditLog.Challenge.C,
					K1: l.AuditLog.Challenge.K1,
					K2: l.AuditLog.Challenge.K2,
				},
				Proof: ProofRes{
					Myu:   l.AuditLog.Proof.Myu,
					Gamma: l.AuditLog.Proof.Gamma,
				},
				Result: l.AuditLog.Result,
			},
			IndexLog: IndexLogRes{
				HashedData: l.IndexLog.HashedData,
				IndexId:    l.IndexLog.IndexId,
				Owner:      l.IndexLog.Owner,
				Provider:   l.IndexLog.Provider,
				AuditLogId: l.IndexLog.AuditLogId,
			},
			Status: string(l.Status),
		})
	}
	return res
}
