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
}

type requesterData struct {
	UserId       string `json:"user_id"`
	Tid          string `json:"tid"`
	Role         int    `json:"role"`
	LanguageCode string `json:"language_code"`
}

func NewRequester(userId, tid, languageCode string, role int) *requesterData {
	return &requesterData{
		UserId:       userId,
		Tid:          tid,
		Role:         role,
		LanguageCode: languageCode,
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
