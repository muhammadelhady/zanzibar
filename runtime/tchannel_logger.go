// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zanzibar

import (
	"fmt"

	"github.com/uber/tchannel-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewTChannelLogger creates a TChannel logger given a zap logger
func NewTChannelLogger(logger *zap.Logger) tchannel.Logger {
	return TChanLogger{logger: logger}
}

// TChanLogger warps a zap logger to be used for TChannel internal logging
type TChanLogger struct {
	logger *zap.Logger
}

// Enabled returns whether the given level is enabled.
func (l TChanLogger) Enabled(tlevel tchannel.LogLevel) bool {
	var zlevel zapcore.Level

	/*
		Log level mapping:
			  TChannel	Zap
		all	  0		n/a
		debug	  1		-1
		info	  2		0
		warn	  3		1
		error	  4		2
		dpanic	  n/a		3
		panic	  n/a		4
		fatal	  5		5
	*/
	switch tlevel {
	case tchannel.LogLevelAll:
		// zap does not have a log all level, zap minimum log level is debug
		return false
	case tchannel.LogLevelFatal:
		zlevel = zapcore.Level(tlevel)
	default:
		zlevel = zapcore.Level(tlevel - 2)
	}

	return l.logger.Check(zlevel, "") != nil
}

// Fatal logs a message, then exits with os.Exit(1).
func (l TChanLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

// Error logs a message at error priority.
func (l TChanLogger) Error(msg string) {
	l.logger.Error(msg)
}

// Warn logs a message at warning priority.
func (l TChanLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

// Infof logs a message at info priority.
func (l TChanLogger) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}

// Info logs a message at info priority.
func (l TChanLogger) Info(msg string) {
	l.logger.Info(msg)
}

// Debugf logs a message at debug priority.
func (l TChanLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(msg, args...))
}

// Debug logs a message at debug priority.
func (l TChanLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

// Fields returns the fields that this logger contains.
func (l TChanLogger) Fields() tchannel.LogFields {
	// zap logger does not expose the fields
	// zap.With writes the field to underlying buffer
	// fortunately TChannel-go does not call this method except in tests
	return nil
}

// WithFields returns a logger with the current logger's fields and fields.
func (l TChanLogger) WithFields(fields ...tchannel.LogField) tchannel.Logger {
	zfields := []zapcore.Field{}
	for _, tf := range fields {
		zf := zap.Any(tf.Key, tf.Value)
		zfields = append(zfields, zf)
	}
	return TChanLogger{logger: l.logger.With(zfields...)}
}