package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/ooqls/go-log"
	"github.com/ooqls/go-log/api/v1/gen"
)

var _ gen.ServerInterface = &LoggingServerImpl{}

type LoggingServerImpl struct{}

func (l *LoggingServerImpl) GetLoggingConfig(c *gin.Context) {
	cfg := gen.LoggingConfig{
		Format: log.GetFormat(),
		Level:  log.GetLevel(),
	}

	// Implement the logic to get the logging configuration
	// For example, you might want to read from a config file or environment variable
	c.JSON(200, cfg)
}

func (l *LoggingServerImpl) SetLoggingConfig(c *gin.Context, cfg gen.LoggingConfig) {
	// Implement the logic to set the logging configuration
	// For example, you might want to write to a config file or environment variable
	log.SetFormat(cfg.Format)
	log.SetLogLevel(cfg.Level)

	c.JSON(200, gin.H{
		"message": "Logging configuration updated successfully",
	})
}

func (l *LoggingServerImpl) GetLoggingLevel(c *gin.Context) {
	// Implement the logic to get the current log level
	// For example, you might want to read from a config file or environment variable
	level := log.GetLevel()
	c.JSON(200, gin.H{
		"level": level,
	})
}

func (l *LoggingServerImpl) UpdateLoggingLevel(c *gin.Context, level gen.LevelEnum) {
	// Implement the logic to set the log level
	// For example, you might want to write to a config file or environment variable
	log.SetLogLevel(level)

	c.JSON(200, level)
}

func (l *LoggingServerImpl) GetLoggingFormat(c *gin.Context) {
	// Implement the logic to get the current log format
	// For example, you might want to read from a config file or environment variable
	format := log.GetFormat()
	c.JSON(200, gin.H{
		"format": format,
	})
}

func (l *LoggingServerImpl) UpdateLoggingFormat(c *gin.Context, f gen.FormatEnum) {
	log.SetFormat(gen.FormatEnum(f))
	c.JSON(200, gin.H{
		"message": "Logging format updated successfully",
	})

}
