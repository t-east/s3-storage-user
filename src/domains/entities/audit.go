package entities

type IndexLog struct {
	HashedData [][]byte `json:"hashed_data"`
	IndexId  string   `json:"index_id"`
	Owner      string   `json:"owner"`
}

type AuditLog struct {
	Challenge *Challenge `json:"challenge"`
	Proof     *Proof     `json:"proof"`
	Result    bool       `json:"result"`
}

type Log struct {
	AuditLog   *AuditLog   `json:"audit_log"`
	IndexLog *IndexLog `json:"index_log"`
}
