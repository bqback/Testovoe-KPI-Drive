package apperrors

import "errors"

// Привычка определять внутренние ошибки в отдельном модуле

var (
	ErrEnvNotFound        = errors.New("unable to load .env file")
	ErrSaveURLMissing     = errors.New("save fact URL missing from .env file")
	ErrGetURLMissing      = errors.New("get facts URL missing from .env file")
	ErrBearerTokenMissing = errors.New("bearer token missing from .env file")
	ErrAuthUserIDMissing  = errors.New("auth user id missing from .env file")
)

var (
	ErrInvalidLoggingLevel      = errors.New("invalid logging level")
	ErrLoggerMissingFromContext = errors.New("logger missing from context")
)
