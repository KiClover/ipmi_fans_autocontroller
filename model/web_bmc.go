package model

type UserSession struct {
	OK               int    `json:"ok"`
	Privilege        int    `json:"privilege"`
	UserID           int    `json:"user_id"`
	ExtendedPriv     int    `json:"extendedpriv"`
	RACSessionID     int    `json:"racsession_id"`
	RemoteAddr       string `json:"remote_addr"`
	ServerName       string `json:"server_name"`
	ServerAddr       string `json:"server_addr"`
	HTTPSEnabled     int    `json:"HTTPSEnabled"`
	CSRFToken        string `json:"CSRFToken"`
	Channel          int    `json:"channel"`
	PasswordStatus   int    `json:"passwordStatus"`
	LoginLastSeconds int    `json:"login_last_seconds"`
}
