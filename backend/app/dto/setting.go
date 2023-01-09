package dto

import "time"

type SettingInfo struct {
	UserName      string `json:"userName"`
	Email         string `json:"email"`
	SystemVersion string `json:"systemVersion"`

	SessionTimeout string `json:"sessionTimeout"`
	LocalTime      string `json:"localTime"`

	PanelName string `json:"panelName"`
	Theme     string `json:"theme"`
	Language  string `json:"language"`

	ServerPort             string `json:"serverPort"`
	SecurityEntrance       string `json:"securityEntrance"`
	ExpirationDays         string `json:"expirationDays"`
	ExpirationTime         string `json:"expirationTime"`
	ComplexityVerification string `json:"complexityVerification"`
	MFAStatus              string `json:"mfaStatus"`
	MFASecret              string `json:"mfaSecret"`

	MonitorStatus    string `json:"monitorStatus"`
	MonitorStoreDays string `json:"monitorStoreDays"`

	MessageType string `json:"messageType"`
	EmailVars   string `json:"emailVars"`
	WeChatVars  string `json:"weChatVars"`
	DingVars    string `json:"dingVars"`
}

type SettingUpdate struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value"`
}

type PasswordUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type SnapshotCreate struct {
	From        string `json:"from" validate:"required,oneof=OSS S3 SFTP MINIO"`
	Description string `json:"description"`
}
type SnapshotRecover struct {
	IsNew      bool `json:"isNew"`
	ReDownload bool `json:"reDownload"`
	ID         uint `json:"id" validate:"required"`
}
type SnapshotInfo struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	From        string    `json:"from"`
	Status      string    `json:"status"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"createdAt"`
	Version     string    `json:"version"`

	InterruptStep    string `json:"interruptStep"`
	RecoverStatus    string `json:"recoverStatus"`
	RecoverMessage   string `json:"recoverMessage"`
	LastRecoveredAt  string `json:"lastRecoveredAt"`
	RollbackStatus   string `json:"rollbackStatus"`
	RollbackMessage  string `json:"rollbackMessage"`
	LastRollbackedAt string `json:"lastRollbackedAt"`
}

type UpgradeInfo struct {
	NewVersion  string `json:"newVersion"`
	Tag         string `json:"tag"`
	ReleaseNote string `json:"releaseNote"`
	CreatedAt   string `json:"createdAt"`
	PublishedAt string `json:"publishedAt"`
}
