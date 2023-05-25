package logs

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

func Log() (logger *zap.Logger) {
	// 获取程序运行的绝对路径
	absPath, _ := filepath.Abs(".")
	// 创建一个文件输出器
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(absPath, "publish.logs"), // 日志文件名
		MaxSize:    10,                                     // 每个日志文件最大大小，单位：MB
		MaxBackups: 3,                                      // 保留的旧日志文件数
		MaxAge:     7,                                      // 保留的旧日志文件的最大年龄，单位：天
	})
	// 设置日志级别和编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器
		fileWriter,                            // 输出器
		zap.ErrorLevel,                        // 日志级别
	)
	// 创建日志记录器
	logger = zap.New(core)
	return logger
}
