package gotercore

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot Be Deleted")

// ErrTokenInvalidOrExpired
var ErrTokenInvalidOrExpired = errors.New("token invalid or expired")
