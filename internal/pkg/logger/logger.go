package logger

import (
	"b2b-api/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var NewLogger = fx.Provide(newLogger)

type ILogger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

func newLogger(d dependencies) ILogger {

	level := zapcore.DebugLevel

	core := zapcore.NewCore(getEncoder(), getWriter(), level)

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer log.Sync() // todo УБРАТЬ ЭТО

	return &logger{lg: log.Sugar(), CFG: d.CFG}
}

type dependencies struct {
	fx.In
	CFG config.IConfig
}

type logger struct {
	lg  *zap.SugaredLogger
	CFG config.IConfig
}

func (l *logger) Debug(format string, v ...interface{}) {
	if l.CFG.GetBool("api.logger.debug") {
		l.lg.Debugf(format, v...)
	}
}

func (l *logger) Info(format string, v ...interface{}) {
	if l.CFG.GetBool("api.logger.info") {
		l.lg.Infof(format, v...)
	}
}

func (l *logger) Warning(format string, v ...interface{}) {
	if l.CFG.GetBool("api.logger.warning") {
		l.lg.Warnf(format, v...)
	}
}

func (l *logger) Error(format string, v ...interface{}) {
	if l.CFG.GetBool("api.logger.error") {
		l.lg.Errorf(format, v...)
	}
}

//getEncoder returns Encoder
func getEncoder() zapcore.Encoder {

	var encoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

//getWriter returns WriteSyncer
func getWriter() zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log", //location of log file
		MaxSize:    200,              //maximum size of log file in MBs, before it is rotated
		MaxBackups: 10,               //maximum no. of old files to retain
		MaxAge:     30,               //maximum number of days it will retain old files
		Compress:   false,            //whether to compress/archive old files
	}

	return zapcore.AddSync(lumberJackLogger)
}
