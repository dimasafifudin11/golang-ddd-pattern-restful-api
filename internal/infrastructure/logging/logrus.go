package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logEntry represents a log message to be processed asynchronously.
type logEntry struct {
	level  logrus.Level
	fields logrus.Fields
	msg    string
}

// asyncLogHook is a Logrus hook that sends log entries to a channel.
type asyncLogHook struct {
	logChannel chan logEntry
}

func (h *asyncLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *asyncLogHook) Fire(entry *logrus.Entry) error {
	h.logChannel <- logEntry{
		level:  entry.Level,
		fields: entry.Data,
		msg:    entry.Message,
	}
	return nil
}

// NewAsyncLogger creates a new Logrus logger that writes to a file asynchronously.
func NewAsyncLogger(filePath string) *logrus.Logger {
	logChannel := make(chan logEntry, 1000) // Buffered channel

	// This goroutine will process log entries from the channel
	go func() {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Errorf("Failed to open log file: %v", err)
			return
		}
		defer file.Close()

		fileLogger := logrus.New()
		fileLogger.SetOutput(file)
		fileLogger.SetFormatter(&logrus.JSONFormatter{})

		for entry := range logChannel {
			fileLogger.WithFields(entry.fields).Log(entry.level, entry.msg)
		}
	}()

	// This is the main logger instance used by the application
	appLogger := logrus.New()
	appLogger.SetOutput(os.Stdout) // Also log to stdout for real-time view
	appLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	appLogger.AddHook(&asyncLogHook{logChannel: logChannel})

	return appLogger
}
