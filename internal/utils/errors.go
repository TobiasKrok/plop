package utils

import "errors"

var (
	ErrFileNotFound = errors.New("content not found")
	ErrFileTooLarge = errors.New("content too large")
	ErrEmptyContent = errors.New("empty content")
	ErrUnauthorized = errors.New("unauthorized")
)
