package bootstrap

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"service/pkg/logger"
	"time"
)

var startedAt time.Time

func Execute() {
	var command = new(cobra.Command)

	startedAt = time.Now()

	command.AddCommand(
		restCommand,
	)

	if err := command.Execute(); err != nil {
		logger.Error("Bootstrap", zap.Any("error", err))
		os.Exit(1)
	}
}
