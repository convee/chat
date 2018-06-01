package logger_test

import (
	"chat/logger"
	"testing"
)

func Test_Log(t *testing.T) {
	logger := logger.New("test", "../")
	logger.Debug("debug test")
	logger.Info("debug test")
	logger.Notice("debug test")
	logger.Warn("debug test")
	logger.Error("debug test")
	t.Log()
}
