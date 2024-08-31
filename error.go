package gotercore

import "errors"

const (
	ExceededDailyActionsLimitErrorCode = -999
)

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot Be Deleted")

// ErrTokenInvalidOrExpired
var ErrTokenInvalidOrExpired = errors.New("token invalid or expired")

// ErrExceededDailyActionsLimit
var ErrExceededDailyActionsLimit = errors.New("exceeded daily actions limit")
