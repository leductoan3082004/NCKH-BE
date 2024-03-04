package sessionmanage

type CredentialsData struct {
	SessionId string `json:"ssid"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
