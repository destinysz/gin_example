package init

import (
	"fmt"
	"gin_example/global"
	"gin_example/internal/model"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func initMysql() error {
	// https://gorm.io/docs/gorm_config.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		global.MysqlSetting.UserName,
		global.MysqlSetting.Password,
		global.MysqlSetting.Host,
		global.MysqlSetting.Port,
		global.MysqlSetting.Database,
		global.MysqlSetting.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.MysqlSetting.TablePrefix,
			SingularTable: true,
		},
		Logger: getGormLogger(), // 使用自定义 Logger
	})
	if err != nil {
		// todo 写入信号关闭服务好 还是请求的时候抛出异常好
		return err
	}

	model.AddCustomCallbacks(db)
	global.DB = db
	return nil
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel
	switch global.MysqlSetting.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                   // 慢 SQL 阈值
		LogLevel:                  logMode,                                  // 日志级别
		IgnoreRecordNotFoundError: false,                                    // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.MysqlSetting.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.MysqlSetting.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.LogSetting.RootDir + "/" + global.MysqlSetting.LogFilename,
			MaxSize:    global.LogSetting.MaxSize,
			MaxBackups: global.LogSetting.MaxBackups,
			MaxAge:     global.LogSetting.MaxAge,
			Compress:   global.LogSetting.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
