package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//全局变量 Conf，用来保存程序的所有配置信息
var Conf = new(AppConfig) //指针，所以当下面反序列化更改都会同步，程序里马上就能知道最新的变量是什么

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdelConns int    `mapstructure:"max_idel_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigName("config") //指定配置文件名(不需要带后缀)
	viper.SetConfigType("yaml")   //指定配置文件类型（专用于从远程获取配置信息时指定配置文件类型的）
	//viper.SetConfigFile("config.yaml") //也可以使用这种方式
	viper.AddConfigPath("./settings") //指定查找配置文件的路径(相对路径)
	err = viper.ReadInConfig()        //读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed , err:%v", err)
		return
	}
	// 把读取到的配置信息发序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	viper.WatchConfig() //支持配置自动热加载信息
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		//这里在执行以下反序列化操作，把最新的配置文件也序列化到 Conf 变量里面去
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
