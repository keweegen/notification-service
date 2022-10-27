package config

import "github.com/spf13/viper"

type Config struct {
    Database             Database             `yaml:"database"`
    MessageBroker        MessageBroker        `yaml:"messageBroker"`
    NotificationChannels NotificationChannels `yaml:"notificationChannels"`
}

type Database struct {
    Host     string `yaml:"host"`
    Port     uint   `yaml:"port"`
    Name     string `yaml:"name"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
}

type MessageBroker struct {
    Addr     []string `yaml:"addr"`
    Password string   `yaml:"password"`
}

type NotificationChannels struct {
    Telegram Telegram `yaml:"telegram"`
    Email    Email    `yaml:"email"`
}

type Telegram struct {
    Host   string `yaml:"host"`
    APIKey string `yaml:"apiKey"`
}

type Email struct {
    Host     string `yaml:"host"`
    Port     uint   `yaml:"port"`
    From     string `yaml:"from"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
}

func Read() (*Config, error) {
    viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var cfg *Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return cfg, nil
}
