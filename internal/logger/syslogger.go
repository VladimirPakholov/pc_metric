package logger

import (
	"log/slog"
	"os"
)

var SysLogger *slog.Logger

func InitSysLogger() {
	SysLogger = slog.New(
		slog.NewTextHandler(os.Stdout, nil),
	)
}
