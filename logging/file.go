package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

func ZApLogger() *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   getCurrentDirectory() + "/api.log", // 日志文件路径
		MaxSize:    int(10),                            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: int(5),                             // 日志文件最多保存多少个备份
		MaxAge:     int(7),                             // 文件最多保存多少天
		Compress:   true,                               // 是否压缩
	}
	defer hook.Close()
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 设置初始化字段
	filed := zap.Fields()
	// 构造日志
	logger := zap.New(core).WithOptions(filed)

	// 开启开发模式，堆栈跟踪
	development := zap.Development()
	logger = logger.WithOptions(development)

	// 开启文件及行号
	caller := zap.AddCaller()
	logger = logger.WithOptions(caller)

	return logger
}

func getCurrentDirectory() string {
	path, _ := os.Getwd()
	path = strings.Replace(path, "\\", "/", -1)
	return path
}
