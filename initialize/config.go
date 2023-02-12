package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/namsral/flag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"pay/config"
	"pay/global"
	"strconv"
)

func InitConfig() {
	initEnvConfig()
	initFlagConfig()
	initViperConfig()
}

func initEnvConfig() {
	serverPort, _ := strconv.Atoi(os.Getenv("PORT"))
	userServerPort, _ := strconv.Atoi(os.Getenv("USER-SERVER_PORT"))

	if global.ServerConfig == nil {
		global.ServerConfig = &config.ServerConfig{}
	}

	global.ServerConfig.Name = os.Getenv("SERVER_NAME")
	global.ServerConfig.Host = "0.0.0.0"
	global.ServerConfig.Port = int64(serverPort)

	global.ServerConfig.UserServerConfig.Name = os.Getenv("USER-SERVER_HOST")
	global.ServerConfig.UserServerConfig.Host = os.Getenv("USER-SERVER_NAME")
	global.ServerConfig.UserServerConfig.Port = int64(userServerPort)
}

// initViperConfig 初始化配置 > viper 配置包
func initViperConfig() {
	file := "config.yaml"
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return
	}

	v := viper.New()
	v.SetConfigFile(file)
	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("获取配置异常: %s", err)
	}
	// 映射配置文件
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicf("加载配置异常: %s", err)
	}
	// 动态监听配置
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&global.ServerConfig)
	})
}

func initFlagConfig() {
	portString := flag.String("port", "9704", "server port")
	queueSuffixString := flag.String("suffix", "", "server rabbitmq queue suffix")

	flag.Parse()

	fmt.Println(*portString)
	fmt.Println(*queueSuffixString)

	//fmt.Println(port)
	//fmt.Println(queueSuffix)

	if portString != nil {
		p, _ := strconv.Atoi(*portString)
		global.ServerConfig.Port = int64(p)
	}

}
