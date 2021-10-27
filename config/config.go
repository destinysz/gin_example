package config

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MysqlSetting struct {
	UserName            string
	Password            string
	Host                string
	Port                string
	Database            string
	TablePrefix         string
	Charset             string
	LogMode             string
	EnableFileLogWriter bool
	LogFilename         string
}

type LogSetting struct {
	Level      string
	RootDir    string
	Filename   string
	Format     string
	ShowLine   bool
	MaxBackups int
	MaxSize    int // MB
	MaxAge     int // day
	Compress   bool
}
