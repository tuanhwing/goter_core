package gotercore

import (
	"context"
)

const (
	KeyHeaderInfo string = "header_info"
)

type HeaderInfo interface {
	GetLanguageCode() string
}

type headerInfoData struct {
	LanguageCode string `json:"language_code"`
}

func NewHeaderInfo(languageCode string) *headerInfoData {
	return &headerInfoData{
		LanguageCode: languageCode,
	}
}

func (h *headerInfoData) GetLanguageCode() string {
	return h.LanguageCode
}

func GetHeaderInfo(ctx context.Context) HeaderInfo {
	if header, ok := ctx.Value(KeyHeaderInfo).(HeaderInfo); ok {
		return header
	}
	return &headerInfoData{
		LanguageCode: "vi",
	}
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
