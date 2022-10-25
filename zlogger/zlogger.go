package zlogger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

const LOG_FILE string = `../log/app.log`
const CALLER_SKIP int = 1

func init() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCallerSkip(CALLER_SKIP))
	sugarLogger = logger.Sugar()

	defer sugarLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   LOG_FILE,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(message string) {
	sugarLogger.Debug(message)
}

func Info(message string) {
	sugarLogger.Info(message)
}

func Fatal(message string) {
	sugarLogger.Fatal(message)
}

func Error(message string) {
	sugarLogger.Error(message)
}
