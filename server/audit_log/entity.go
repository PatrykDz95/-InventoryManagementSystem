package audit_log

import (
	"gorsk/server/user"
	"time"
)

type AuditLog struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	User      user.User `gorm:"foreignKey:UserID"`
	Action    string    `gorm:"not null"`
	TableName string
	RecordID  uint
	Changes   string    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
