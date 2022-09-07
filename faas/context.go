// Package faas provides a function as a service framework.
package faas

import (
	"context"
)

// Context 函数执行上文接口
type Context interface {
	Logger() Logger
	Storage() Storage
	Metadata() Metadata
	Context() context.Context
}

// Logger 函数执行日志接口
type Logger interface {
	Warnf(message string, args ...any)
	Logf(message string, args ...any)
	Debugf(message string, args ...any)
	Errorf(message string, args ...any)
}

// Storage 函数执行存储接口
type Storage interface {
	Get(key string) (any, error)
	Set(key string, value any, expireTime int64) error
	Del(key string) error
}

// Metadata 函数执行元数据接口
type Metadata interface {
	Namespace() string
	ID() string
	Name() string
	Version() int
	Attribute(key string) (any, error)
}
