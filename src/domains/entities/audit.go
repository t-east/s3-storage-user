package entities

type ContentLog struct {
	HashedData [][]byte `json:"hashed_data"`
	ContentId  string   `json:"content_id"`
	Owner      string   `json:"owner"`
}

type AuditLog struct {
	Challenge *Challenge `json:"chal"`
	Proof     *Proof     `json:"proof"`
	Result    bool       `json:"result"`
	ContentID string     `json:"content_id"`
}

type Log struct {
	AuditLog   *AuditLog   `json:"audit_log"`
	ContentLog *ContentLog `json:"content_log"`
}
