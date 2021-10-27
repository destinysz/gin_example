package init

import (
	"gin_example/common/utils"
	"gin_example/global"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func initLog() {
	createLogDir()
	setLogLevel()
	if global.LogSetting.ShowLine {
		options = append(options, zap.AddCaller())
	}
	global.Log = zap.New(getZapCore(), options...)
	return

}

func createLogDir() {
	if ok, _ := utils.PathExists(global.LogSetting.RootDir); !ok {
		_ = os.Mkdir(global.LogSetting.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.LogSetting.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05" + "]"))
	}
	// encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	// 	encoder.AppendString(os.Getenv("MODE") + "." + l.String())
	// }

	// 设置编码器
	if global.LogSetting.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.LogSetting.RootDir + "/" + global.LogSetting.Filename,
		MaxSize:    global.LogSetting.MaxSize,
		MaxBackups: global.LogSetting.MaxBackups,
		MaxAge:     global.LogSetting.MaxAge,
		Compress:   global.LogSetting.Compress,
	}

	return zapcore.AddSync(file)
}
