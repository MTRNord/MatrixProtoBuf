package jsonTypes

type HSVersions struct {
	Versions []string `json:"versions"`
}

type ReqLogin struct {
	Type                     string `json:"type"`
	Password                 string `json:"password,omitempty"`
	Medium                   string `json:"medium,omitempty"`
	User                     string `json:"user,omitempty"`
	Address                  string `json:"address,omitempty"`
	Token                    string `json:"token,omitempty"`
	DeviceID                 string `json:"device_id,omitempty"`
	InitialDeviceDisplayName string `json:"initial_device_display_name,omitempty"`
}

type RespLogin struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	HomeServer   string `json:"home_server"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// RespError is the standard JSON error response from Homeservers. It also implements the Golang "error" interface.
// See http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type RespError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

// Error returns the errcode and error message.
func (e RespError) Error() string {
	return e.ErrCode + ": " + e.Err
}
