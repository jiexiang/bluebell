package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInConfig error", err)
		return
	}

	// 把读到的配置文件信息，反序列化到Conf中
	if err = unmarshal(); err != nil {
		return
	}

	// 观察配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件变更...")
		if err = unmarshal(); err != nil {
			return
		}
	})
	return
}

// unmarshal 反序列化
func unmarshal() (err error) {
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Println("viper.Unmarshal error", err)
	}
	return
}