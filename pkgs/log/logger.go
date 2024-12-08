package log

import (
	"log/slog"
	"os"
	"sync"
)

var JsonLogger *slog.Logger
var once sync.Once

func init() {
	once.Do(func() {
		logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
		JsonLogger = slog.New(logHandler)
	})
}
