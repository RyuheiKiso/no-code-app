package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// 設定を保持する構造体
type Config struct {
	// 必要な設定項目をここに追加
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// 設定ファイルを読み込む関数
// configPath: 設定ファイルのパス
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	// 設定ファイルを読み込む
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	// YAMLファイルをパースして構造体にマッピング
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	// 設定のバリデーションを実行
	err = validateConfig(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 環境変数を取得する関数
// key: 環境変数のキー
// defaultValue: 環境変数が存在しない場合に使用するデフォルト値
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("環境変数 %s が見つかりません。デフォルト値 %s を使用します。", key, defaultValue)
		return defaultValue
	}
	return value
}

// 環境変数から設定ファイルを読み込む関数
func LoadConfigFromEnv() (*Config, error) {
	// 環境変数から設定ファイルのパスを取得
	configPath := GetEnv("CONFIG_PATH", "config.yaml")
	return LoadConfig(configPath)
}

// 設定ファイルのホットリロードを行う関数
// configPath: 設定ファイルのパス
// reloadInterval: 設定ファイルの変更をチェックする間隔
// onChange: 設定ファイルが変更された際に呼び出されるコールバック関数
func WatchConfig(configPath string, reloadInterval time.Duration, onChange func(*Config)) {
	var lastModTime time.Time

	for {
		// 設定ファイルの最終更新時刻を取得
		fileInfo, err := os.Stat(configPath)
		if err != nil {
			log.Printf("設定ファイルの監視中にエラーが発生しました: %v", err)
			time.Sleep(reloadInterval)
			continue
		}

		// 設定ファイルが更新された場合
		if fileInfo.ModTime().After(lastModTime) {
			config, err := LoadConfig(configPath)
			if err != nil {
				log.Printf("設定ファイルの再読み込み中にエラーが発生しました: %v", err)
			} else {
				lastModTime = fileInfo.ModTime()
				onChange(config)
			}
		}

		time.Sleep(reloadInterval)
	}
}

// 設定ファイルのバリデーションを行う関数
// config: 検証する設定構造体
func validateConfig(config *Config) error {
	if config.Database.User == "" {
		return errors.New("データベースユーザーが設定されていません")
	}
	if config.Database.Password == "" {
		return errors.New("データベースパスワードが設定されていません")
	}
	if config.Database.Host == "" {
		return errors.New("データベースホストが設定されていません")
	}
	if config.Database.Port == 0 {
		return errors.New("データベースポートが設定されていません")
	}
	if config.Database.Name == "" {
		return errors.New("データベース名が設定されていません")
	}
	return nil
}
