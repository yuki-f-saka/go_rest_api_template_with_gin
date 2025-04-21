package log

import (
	"bytes"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var defaultLogLevel = "info"

// Config ロガーの設定
type Config struct {
	Env        string
	LogLevel   string
	AppName    string
	AppVersion string
}

var defaultZapLoggerConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(ConvertToZapLevel(defaultLogLevel)),
	Development: false,
	// DisableCaller:     false,
	// DisableStacktrace: false,
	Sampling: &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
		// Hook: func(zapcore.Entry, zapcore.SamplingDecision) {
		// 	panic("TODO")
		// },
	},
	Encoding: "json",
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		NameKey:    "logger",
		CallerKey:  "caller",
		// FunctionKey:    "",
		StacktraceKey: "stacktrace",
		// SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// EncodeName:     nil,
		// NewReflectedEncoder: func(io.Writer) zapcore.ReflectedEncoder {
		// 	panic("TODO")
		// },
		// ConsoleSeparator: "",
	},
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
	// InitialFields:     map[string]interface{}{},
}

// NewLogger 通常のロガー
func NewLogger(cfg *Config) (*zap.Logger, error) {
	defaultZapLoggerConfig.Level = zap.NewAtomicLevelAt(ConvertToZapLevel(cfg.LogLevel))
	l, err := defaultZapLoggerConfig.Build()
	if err != nil {
		return nil, err
	}

	return l.With(
		zap.String("env", cfg.Env),
		zap.String("app_name", cfg.AppName),
		zap.String("app_version", cfg.AppVersion),
	), nil

}

// NewTestLogger テスト用ロガー
func NewTestLogger(b *bytes.Buffer) (*zap.Logger, error) {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})
	core := zapcore.NewCore(encoder, zapcore.AddSync(b), zapcore.InfoLevel)

	return zap.New(core), nil
}

// SetLogger ロガーをセットする
func SetLogger(l *zap.Logger) {
	Logger = l
}

// ConvertToZapLevel zapcore.Levelに変換する
func ConvertToZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	}
	return zap.InfoLevel
}
