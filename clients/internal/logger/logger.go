package logger

import (
	"go_grpc/clients/internal/config"
	"log/slog"
	"os"
	"runtime/debug"
)

// Logger using slog
func LoadLogger(conf *config.Config) (*slog.Logger, error) {
	
	buildInfo, _ := debug.ReadBuildInfo()
	
	opts := &slog.HandlerOptions{
		AddSource: true,
        Level: slog.LevelDebug,
    }

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	child := logger.With(
        slog.Group("program_info",
            slog.Int("pid", os.Getpid()),
            slog.String("go_version", buildInfo.GoVersion),
        ),
    )

	slog.SetDefault(child)

	child.Info("Logger initialized")

	return child, nil
}