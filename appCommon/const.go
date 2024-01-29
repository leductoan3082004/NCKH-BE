package appCommon

type DbType int

const (
	MainDBName = "nckh"

	DBMain      = "mongodb"
	PluginRedis = "redis"
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
