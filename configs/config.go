package configs

import (
	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Environment string `mapstructure:"environment"`
	Port        int    `mapstructure:"port"`
	Database    struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// 设置默认值
	viper.SetDefault("environment", "development")
	viper.SetDefault("port", 8080)
	viper.SetDefault("database.dsn", ":memory:")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果没有找到配置文件，使用默认值
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}