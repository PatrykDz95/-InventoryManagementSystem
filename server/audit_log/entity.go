package audit_log

import (
	"time"
)

type AuditLog struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Action    string `gorm:"not null"`
	TableName string
	RecordID  uint
	Changes   string    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
