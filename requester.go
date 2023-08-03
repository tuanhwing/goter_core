package gotercore

import "context"

const (
	KeyRequester string = "requester"
)

type Requester interface {
	GetUserId() string
	GetTokenId() string
	GetRole() int
}

type requesterData struct {
	UserId string `json:"user_id"`
	Tid    string `json:"tid"`
	Role   int    `json:"role"`
}

func NewRequester(userId, tid string, role int) *requesterData {
	return &requesterData{
		UserId: userId,
		Tid:    tid,
		Role:   role,
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

func GetRequester(ctx context.Context) Requester {
	if requester, ok := ctx.Value(KeyRequester).(Requester); ok {
		return requester
	}

	return &requesterData{}
}
