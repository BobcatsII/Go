package setting

import (
	"github.com/spf13/viper"
	"go/types"
	"syscall"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

//编写组件
//初始化项目配置的基础属性
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}


