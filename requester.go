package gotercore

import "context"

const (
	KeyRequester string = "requester"
)

type Requester interface {
	GetUserId() string
	GetTokenId() string
	GetRole() int
	GetLanguageCode() string
	GetDeviceInfo() DeviceInfoHeaderRequest
	GetToken() string
}

type requesterData struct {
	UserId       string                  `json:"user_id"`
	Tid          string                  `json:"tid"`
	Role         int                     `json:"role"`
	LanguageCode string                  `json:"language_code"`
	DeviceInfo   DeviceInfoHeaderRequest `json:"device_info"`
	Token        string                  `json:"token"`
}

func NewRequester(userId, tid, languageCode, token string, role int, deviceInfo DeviceInfoHeaderRequest) *requesterData {
	return &requesterData{
		UserId:       userId,
		Tid:          tid,
		Role:         role,
		LanguageCode: languageCode,
		DeviceInfo:   deviceInfo,
		Token:        token,
	}
}

func (r *requesterData) GetUserId() string {
	return r.UserId
}

func (r *requesterData) GetTokenId() string {
	return r.Tid
}

func (r *requesterData) GetRole() int {
	return r.Role
}

func (r *requesterData) GetLanguageCode() string {
	return r.LanguageCode
}

func (r *requesterData) GetDeviceInfo() DeviceInfoHeaderRequest {
	return r.DeviceInfo
}

func (r *requesterData) GetToken() string {
	return r.Token
}

func GetRequester(ctx context.Context) Requester {
	if requester, ok := ctx.Value(KeyRequester).(Requester); ok {
		return requester
	}

	return &requesterData{}
}

func IsLanguageSupported(lang string) bool {
	supportsLanguages := []string{"vi", "en"}
	for _, v := range supportsLanguages {
		if v == lang {
			return true
		}
	}

	return false
}
