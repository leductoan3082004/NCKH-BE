package appCommon

import "os"

type DbType int

const (
	MainDBName = "nckh"

	DBMain      = "mongodb"
	PluginRedis = "redis"
	PluginJWT   = "jwt"
	PluginAWS   = "aws"
)

const (
	CurrentUser = "user"
)

var (
	S3Domain = os.Getenv("AWS_DOMAIN")
	S3Path   = "images"
)

const (
	ExpiryAccessToken = 60 * 60 * 12 * 365
)

type TokenPayload struct {
	UId       int64  `json:"user_id"`
	URole     int    `json:"role"`
	Type      string `json:"type"`
	SessionId string `json:"session_id"`
}

func (p TokenPayload) UserId() int64 {
	return p.UId
}

func (p TokenPayload) Role() int {
	return p.URole
}

func (p TokenPayload) UType() string {
	return p.Type
}

func (p TokenPayload) USessionId() string {
	return p.SessionId
}
