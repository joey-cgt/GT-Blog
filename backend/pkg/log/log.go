package log

import (
	"gt-blog/backend/pkg/config"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger // 全局日志实例

// 初始化日志配置（必须在使用日志前调用）
func Init(cfg config.LogConfig) *zap.Logger {
	// 1. 设置日志级别
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	// 2. 设置输出目标（控制台或文件）
	var writer zapcore.WriteSyncer
	if cfg.OutputPath == "stdout" {
		// 输出到控制台
		writer = zapcore.AddSync(os.Stdout)
	} else {
		// 输出到文件（带滚动功能）
		// 确保目录存在
		if err := os.MkdirAll(filepath.Dir(cfg.OutputPath), 0755); err != nil {
			panic("创建日志目录失败: " + err.Error())
		}

		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.OutputPath, // 日志文件路径
			MaxSize:    cfg.MaxSize,    // 单个文件最大尺寸（MB）
			MaxBackups: cfg.MaxBackups, // 最大备份数
			MaxAge:     cfg.MaxAge,     // 最大保留天数
			Compress:   cfg.Compress,   // 是否压缩备份
		})
	}

	// 3. 设置日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 级别大写（INFO/ERROR）
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 时间格式 ISO8601
		EncodeDuration: zapcore.StringDurationEncoder, // 时长字符串化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 调用者路径简短显示
	}

	// 4. 创建核心配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON格式输出（生产环境推荐）
		writer,
		level,
	)

	// 5. 创建日志实例（开发模式添加调用栈和文件名/行号）
	var logger *zap.Logger
	if cfg.Level == "debug" {
		// 开发模式：启用调用栈和文件名行号
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		// 生产模式：仅在错误级别启用调用栈
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.PanicLevel))
	}

	// 6. 设置全局日志实例
	globalLogger = logger
	return logger
}

// 以下为便捷日志输出函数（封装zap的方法）

func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	globalLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}

// Sync 刷新日志缓存（程序退出前调用）
func Sync() error {
	return globalLogger.Sync()
}

// 辅助函数：创建通用字段（简化调用）
func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func String(key, value string) zap.Field {
	return zap.String(key, value)
}

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func Errorf(err error) zap.Field {
	return zap.Error(err)
}
