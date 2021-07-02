package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") //指定配置文件名(不需要带后缀)
	viper.SetConfigType("yaml")   //指定配置文件类型
	viper.AddConfigPath(".")      //指定查找配置文件的路径(相对路径)
	err = viper.ReadInConfig()    //读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed , err:%v", err)
		return
	}
	viper.WatchConfig() //支持配置自动热加载信息
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
