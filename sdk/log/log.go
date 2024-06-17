package log

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

// 定义一个结构体用于实现日志格式化接口
type customFormatter struct{}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取文件名和行号
	_, file, line, ok := runtime.Caller(8)
	if !ok {
		file = "unknown"
		line = 0
	}

	// 时间戳
	timestamp := time.Now().Format(time.RFC3339)

	// 日志级别颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = 37 // 灰色
	case logrus.InfoLevel:
		levelColor = 36 // 蓝色
	case logrus.WarnLevel:
		levelColor = 33 // 黄色
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // 红色
	default:
		levelColor = 37 // 灰色
	}

	// 格式化日志输出
	logMessage := fmt.Sprintf("\033[%dm[%s] %s:%d %s: %s\033[0m\n",
		levelColor, timestamp, file, line, entry.Level.String(), entry.Message)

	return []byte(logMessage), nil
}

// 定义一个全局的日志记录器
var logger = logrus.New()

func init() {
	logger.SetFormatter(&customFormatter{})
	logger.SetReportCaller(true)
	// 你可以在这里设置日志级别
	logger.SetLevel(logrus.DebugLevel)
}

// 包级别的日志函数
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
