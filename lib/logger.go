package lib

import (
	"os"

	"github.com/kz/discordrus"
	"github.com/sirupsen/logrus"
)

// LoggerSetting : setting logger
func LoggerSetting() {
	if os.Getenv("LOGGER_WEBHOOK_URL") != "" {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.TraceLevel)

		logrus.AddHook(discordrus.NewHook(
			// Use environment variable for security reasons
			// Set minimum level to DebugLevel to receive all log entries
			os.Getenv("LOGGER_WEBHOOK_URL"),
			// Set minimum level to DebugLevel to receive all log entries
			logrus.TraceLevel,
			&discordrus.Opts{
				Username:           "편지봇 로거",
				DisableTimestamp:   false, // Setting this to true will disable timestamps from appearing in the footer
				TimestampLocale:    nil,   // The timestamp uses this locale; if it is unset, it will use time.Local
				EnableCustomColors: true,  // If set to true, the below CustomLevelColors will apply
				CustomLevelColors: &discordrus.LevelColors{
					Trace: 3092790,
					Debug: 10170623,
					Info:  3581519,
					Warn:  14327864,
					Error: 13631488,
					Panic: 13631488,
					Fatal: 13631488,
				},
				DisableInlineFields: false, // If set to true, fields will not appear in columns ("inline")
			},
		))
	}
}
