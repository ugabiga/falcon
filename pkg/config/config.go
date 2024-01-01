package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	DBDriverName       string `mapstructure:"DB_DRIVER_NAME"`
	DBSource           string `mapstructure:"DB_SOURCE"`
	SessionSecretKey   string `mapstructure:"SESSION_SECRET_KEY"`
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	WebURL             string `mapstructure:"WEB_URL"`
	EncryptionKey      string `mapstructure:"ENCRYPTION_KEY"`
}

func NewConfigWithoutSetting() *Config {
	return &Config{}
}

func NewConfig() (*Config, error) {
	return newConfig()
}

func newConfig() (*Config, error) {
	config := &Config{}
	if err := config.Load(nil, nil); err != nil {

		if config.LoadAutomaticEnv() != nil {
			log.Fatalf("failed loading config: %v", err)
			return nil, err
		}

		return config, nil
	}

	return config, nil
}

func (c *Config) LoadAutomaticEnv() error {
	v := viper.New()

	//for loop for config struct
	result := extractMapStructureTags(*c)
	for _, value := range result {
		if err := v.BindEnv(value, value); err != nil {
			return err
		}
	}

	if err := v.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}

func (c *Config) Load(path, name *string) error {
	if path == nil {
		projectRoot, err := findProjectRoot()
		if err != nil {
			return err
		}
		path = &projectRoot
	}

	if name == nil {
		name = &[]string{"config"}[0]
	}

	viper.SetConfigName(*name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		// Move one level up the directory hierarchy
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break // Reached the root directory
		}
		dir = parentDir
	}

	return "", os.ErrNotExist
}

func extractMapStructureTags(configStruct interface{}) map[string]string {
	t := reflect.TypeOf(configStruct)
	if t.Kind() != reflect.Struct {
		panic("Input must be a struct")
	}

	result := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			result[field.Name] = tag
		}
	}

	return result
}
