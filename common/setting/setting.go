package setting

import (
	"os"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	mode := os.Getenv("MODE")
	var configName string
	switch mode {
	case "DEVELOP":
		configName = "dev"
	case "TESTING":
		configName = "test"
	case "PRODUCTION":
		configName = "prod"
	default:
		configName = "dev"
	}
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AddConfigPath("config/")
	vp.SetConfigType("yml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
