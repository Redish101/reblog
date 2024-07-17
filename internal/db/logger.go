package db

import (
	"context"
	"reblog/internal/log"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger.Interface
}

func (gl *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	gl.Interface.LogMode(level)
	return gl
}

func (gl *GormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	log.Infof(format, args...)
}

func (gl *GormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func (gl *GormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		gl.Warn(ctx, "[DB] [%v] [rows: %d] %s %v", elapsed, rows, sql, err)
	}
}
