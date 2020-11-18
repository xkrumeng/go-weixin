package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	// 日志文件目录
	logFilePath := "logs"

	// 日志文件名
	logFileName := "log"

	// 日志文件名
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs 日志分割
	logWriter, _ := rotatelogs.New(
		// 分割后的名称
		fileName+".%Y%m%d.log",

		// 生成软链， 指向最新的日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大的保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//设置日志切割时间间隔(1天）
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	return func(ctx *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		ctx.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := ctx.Request.Method

		// 请求路由
		reqUri := ctx.Request.RequestURI

		// 状态码
		statusCode := ctx.Writer.Status()

		// 请求IP
		clientIP := ctx.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
