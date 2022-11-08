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
	Status   LogStatus
}

type LogStatus string

var (
	Uploading LogStatus = "Uploading"
	Waiting   LogStatus = "Waiting"
	Verifying LogStatus = "Verifying"
	NG        LogStatus = "NG"
	OK        LogStatus = "OK"
)

func (l *Log) SetStatus() {
	if l.IndexLog.Provider == "" {
		l.Status = Uploading
	} else if l.IndexLog.AuditLogId == "" {
		l.Status = Waiting
	} else if l.AuditLog.Proof == nil || l.AuditLog.Challenge == nil {
		l.Status = Verifying
	} else if !l.AuditLog.Result {
		l.Status = NG
	} else {
		l.Status = OK
	}
}
