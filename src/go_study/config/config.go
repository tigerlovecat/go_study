package config

import (
	"fmt"
	"os"
	"strings"
	"web_framework/storage"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Version defines the LoRa Server version.
var Version string

// C holds the global configuration.
var C Config

var StorageFileType map[string]int = map[string]int{
	".bmp":2,".gif":2,".jpg":2,".pic":2,".png":2,".tif":2,
	".avi":3,".mpg":3,".mov":3,".swf":3,".flv":3,".mp4":3,".mkv":3,
	".wav":4,".aif":4,".mp3":4,".wma":4,".mmf":4}


// Config defines the configuration structure.
type Config struct {
	Name        string
	RunMode     string `mapstructure:"runmode"`
	Address     string `mapstructure:"addr"`
	JwtSecret   string `mapstructure:"jwt_secret"`
	SVer        int    `mapstructure:"sver"`
	P4pclientSver int `mapstructure:"p4pclient_sver"`
	AppSK       string `mapstructure:"app_sk"`
	P4pclientSk string `mapstructure:"p4pclient_sk"`
	YpApiKey    string `mapstructure:"yunpian_apikey"`
	YyetsV3Key  string `mapstructure:"yyets_v3_key"`
	StVer       string `mapstructure:"st_ver"`

	TLS struct {
		Addr string `mapstructure:"addr"`
		Cert string `mapstructure:"cert"`
		Key  string `mapstructure:"key"`
	} `mapstructure:"tls"`

	Log struct {
		LoggerLevel int    `mapstructure:"logger_level"`
		LoggerFile  string `mapstructure:"logger_file"`
		LoggerType  string `mapstructure:"logger_type"`
	} `mapstructure:"log"`

	DB struct {
		Type string `mapstructure:"type"`
		Name string `mapstructure:"name"`
		DSN  string `mapstructure:"dsn"`
		DB   *storage.DBLogger
	} `mapstructure:"db"`

	DockerDB struct {
		Type string `mapstructure:"type"`
		Name string `mapstructure:"name"`
		DSN  string `mapstructure:"dsn"`
		DB   *storage.DBLogger
	} `mapstructure:"docker_db"`

	Redis struct {
		URL string `mapstructure:"url"`
		//Pool *redis.Pool
		RedisPool *storage.RedisPool
	} `mapstructure:"redis"`

	Account struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"account"`

	Lianle struct {
		Account  string `mapstructure:"account"`
		Password string `mapstructure:"password`
		Key      string `mapstructure:"key"`
	}

	Email struct {
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password`
		User     string `mapstructure:"user"`
		From     string `mapstructure:"from"`
		Request  string `mapstructure:"request"`
		Port     int    `mapstructure:"port"`
	}

	Router *gin.Engine
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.InitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "init config error: %s\n", err.Error())
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) InitConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		//gopath := os.Getenv("GOPATH")
		//pathslice := filepath.SplitList(gopath)
		//configPath := filepath.Join(pathslice[0], "src/web_framework/conf")
		//viper.AddConfigPath(configPath) // 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")         // 设置配置文件格式为YAML
	viper.AutomaticEnv()                // 读取匹配的环境变量
	viper.SetEnvPrefix("web_framework") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	} else {
		fmt.Printf("init using config file %s\n", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&C); err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
