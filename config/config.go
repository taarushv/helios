package config

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Helios struct {
        NodeUrl string `yaml:"nodeUrl"`, // node ws url
        Mode string `yaml:"mode"`, // quick or full
    } `yaml:"helios"`

    Postgres struct {
        PostgresHot string `yaml:"postgres_host"`
        PostgresUser string `yaml:"postgres_user"`
        PostgresPassword string `yaml:"postgres_password"`
    } `yaml:"postgres"`

    Redis struct {
        RedisHost string `yaml:"redis_host"`
        RedisUser string `yaml:"redis_user"`
        RedisPassword string `yaml:"redis_password"`
    } `yaml:"redis"`
}

func NewConfig(configPath string) (*Config, err) {
    config := &Config{}

    file, err := os.Open(configPath)

    if err != nil {
        return nil, err
    }

    defer file.Close()

    decoder := yaml.NewDecoder(file)

    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }
    return config, nil
}

func ValidatePath(path string) error {
    s, err := os.Stat(path)
    if err != nil {
        return err
    }
    if s.IsDir() {
        return fmt.Errorf("'%s' is a directory, please specify a file", path)
    }
    return nil
}

func ParseFlags() (string, error) {
    var configPath string

    flag.StringVar(&configPath, "ocnfig", "./helios.yml", "path to yaml config for helios")

    flag.Parse()

    if err := ValidatePath(configPath); err != nil {
        return "", err
    }

    return configPath, nil
}