# Config パッケージ

このパッケージは、YAML形式の設定ファイルを読み込み、設定を管理するための機能を提供します。

## 構造体

### Config

設定を保持する構造体です。必要な設定項目を追加できます。

```go
type Config struct {
    Database struct {
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        Name     string `yaml:"name"`
    } `yaml:"database"`
}
```

## 関数

### LoadConfig

指定されたパスから設定ファイルを読み込みます。

```go
func LoadConfig(configPath string) (*Config, error)
```

- `configPath`: 設定ファイルのパス

### GetEnv

環境変数を取得します。環境変数が存在しない場合はデフォルト値を返します。

```go
func GetEnv(key string, defaultValue string) string
```

- `key`: 環境変数のキー
- `defaultValue`: 環境変数が存在しない場合に使用するデフォルト値

### LoadConfigFromEnv

環境変数 `CONFIG_PATH` から設定ファイルのパスを取得し、設定ファイルを読み込みます。

```go
func LoadConfigFromEnv() (*Config, error)
```

### WatchConfig

設定ファイルのホットリロードを行います。設定ファイルが変更された場合に自動的に再読み込みされます。

```go
func WatchConfig(configPath string, reloadInterval time.Duration, onChange func(*Config))
```

- `configPath`: 設定ファイルのパス
- `reloadInterval`: 設定ファイルの変更をチェックする間隔
- `onChange`: 設定ファイルが変更された際に呼び出されるコールバック関数

### validateConfig

設定ファイルのバリデーションを行います。

```go
func validateConfig(config *Config) error
```

- `config`: 検証する設定構造体

## 使用例

### 設定ファイルの読み込み

```go
config, err := config.LoadConfig("path/to/config.yaml")
if err != nil {
    log.Fatalf("設定ファイルの読み込みに失敗しました: %v", err)
}
```

### 環境変数から設定ファイルの読み込み

```go
config, err := config.LoadConfigFromEnv()
if err != nil {
    log.Fatalf("設定ファイルの読み込みに失敗しました: %v", err)
}
```

### 設定ファイルのホットリロード

```go
config.WatchConfig("path/to/config.yaml", time.Second*10, func(newConfig *config.Config) {
    log.Println("設定ファイルが変更されました")
    // 新しい設定を適用する処理をここに追加
})
```
