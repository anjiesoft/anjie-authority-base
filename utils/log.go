package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func init() {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseColorLevelEncoder, //设置颜色
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}) // 实现两个判断日志等级的interface

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	fatalLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.FatalLevel
	})

	// 获取 error日志文件的io.Writer 抽象 getWriter() 在下方实现
	path := GetConfigString("log.path")
	errorFile := GetConfigString("log.errorFile")
	errorWriter := getWriter(path + "/" + errorFile) // 最后创建具体的Logger

	//infoFile := GetConfigString("log.infoFile")
	fatalWriter := getWriter(path + "/" + errorFile)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel), //打印到控制台
		zapcore.NewCore(encoder, fatalWriter, fatalLevel),
		zapcore.NewCore(encoder, errorWriter, errorLevel),
	)
	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑

	Logger = log.Sugar()
}

func getWriter(filename string) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    GetConfigInt("log.maxSize"),
		MaxBackups: GetConfigInt("log.maxBackups"),
		MaxAge:     GetConfigInt("log.maxAge"),
		Compress:   false,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
