package entities

type IndexLog struct {
	HashedData [][]byte
	IndexId    string
	Owner      string
	Provider   string
	AuditLogId string
}

type AuditLog struct {
	Challenge *Challenge
	Proof     *Proof
	Result    bool
}

type Log struct {
	AuditLog *AuditLog
	IndexLog *IndexLog
}
