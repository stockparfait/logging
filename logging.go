// Copyright 2022 Stock Parfait

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logging impelments custom logging.
//
// It is made loosely compatible with go.chromium.org/luci/common/logging, so
// the latter can be swapped in if needed.
package logging

import (
	"context"
	"flag"

	"github.com/stockparfait/errors"
)

// Level of logging, primarily for user-visible output.
type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
)

// DefaultLevel is the default log level emitted by standard loggers.
const DefaultLevel = Info

var l Level
var _ flag.Value = &l // ensure Level implements flag.Value

// Set implements flag.Value.
func (l *Level) Set(v string) error {
	switch v {
	case "debug":
		*l = Debug
	case "info":
		*l = Info
	case "warning":
		*l = Warning
	case "error":
		*l = Error
	default:
		return errors.Reason("unknown log level value")
	}
	return nil
}

// String implements flag.Value.
func (l Level) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}

// Logger implementations are injected into context and used by the top-level
// methods like Infof() etc. It is loosely compatible with
// go.chromium.org/luci/common/logging.Logger, in a way that a LUCI's Logger
// implementation also implements this Logger.
type Logger interface {
	// Debugf formats its arguments according to the format, analogous to
	// fmt.Printf and records the text as a log message at Debug level.
	Debugf(format string, args ...interface{})
	// Infof is like Debugf, but logs at Info level.
	Infof(format string, args ...interface{})
	// Warningf is like Debugf, but logs at Warning level.
	Warningf(format string, args ...interface{})
	// Errorf is like Debugf, but logs at Error level.
	Errorf(format string, args ...interface{})
}

// Debugf emits debug-level log formatted as fmt.Printf(msg, args...).
func Debugf(ctx context.Context, msg string, args ...interface{}) {
	Get(ctx).Debugf(msg, args...)
}

// Infof emits info-level log formatted as fmt.Printf(msg, args...).
func Infof(ctx context.Context, msg string, args ...interface{}) {
	Get(ctx).Infof(msg, args...)
}

// Warningf emits warning-level log formatted as fmt.Printf(msg, args...).
func Warningf(ctx context.Context, msg string, args ...interface{}) {
	Get(ctx).Warningf(msg, args...)
}

// Errorf emits error-level log formatted as fmt.Printf(msg, args...).
func Errorf(ctx context.Context, msg string, args ...interface{}) {
	Get(ctx).Errorf(msg, args...)
}

// Null is a logger that silently ignores all messages.
var Null Logger = nullLogger{}

type nullLogger struct{}

func (nullLogger) Debugf(string, ...interface{})   {}
func (nullLogger) Infof(string, ...interface{})    {}
func (nullLogger) Warningf(string, ...interface{}) {}
func (nullLogger) Errorf(string, ...interface{})   {}

type contextKey int

const (
	loggerContextKey contextKey = iota
)

// Use injects the Logger into the context.
func Use(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// Get a Logger instance from the context. If no Logger is installed, return
// Null Logger.
func Get(ctx context.Context) Logger {
	l, ok := ctx.Value(loggerContextKey).(Logger)
	if !ok || l == nil {
		return Null
	}
	return l
}
