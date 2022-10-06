package entities

type AuditLog struct {
	Challenge *Challenge `json:"chal"`
	Proof     *Proof     `json:"proof"`
	Result    bool       `json:"result"`
	ContentID string     `json:"content_id"`
}
