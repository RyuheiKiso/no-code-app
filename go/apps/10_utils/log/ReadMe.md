# ログ出力モジュール

このディレクトリには、ログ出力を管理するための共通モジュールが含まれています。ログはテキスト形式、JSON形式、CSV形式で出力することができます。

## ファイル構成

- `logger.go`: ログ出力を管理するための主要なコードが含まれています。

## 使用方法

### ログレベルの設定

`SetLogLevel` 関数を使って、ログレベルを設定します。

```go
import "path/to/logger"

func main() {
    // ログレベルをDEBUGに設定
    logger.SetLogLevel(logger.DEBUG)
}
```

### ログ出力形式の設定

SetLogOutputFormat

 関数を使って、ログ出力形式を設定します。サポートされている形式は、テキスト形式、JSON形式、CSV形式です。

```go
import "path/to/logger"

func main() {
    // ログ出力形式をJSONに設定
    logger.SetLogOutputFormat(logger.JSON)
}
```

### ログメッセージの出力

ログメッセージを出力するには、

Debug

、

Info

、

Warn

、

Error

 関数を使用します。

```go
import "path/to/logger"

func main() {
    // デバッグメッセージを出力
    logger.Debug("これはデバッグメッセージです")
    // 情報メッセージを出力
    logger.Info("これは情報メッセージです")
    // 警告メッセージを出力
    logger.Warn("これは警告メッセージです")
    // エラーメッセージを出力
    logger.Error("これはエラーメッセージです")
}
```

### ログファイルの設定

SetLogFile

 関数を使って、ログをファイルに出力することができます。

```go
import "path/to/logger"

func main() {
    // ログファイルを設定
    err := logger.SetLogFile(logger.INFO, "app.log")
    if err != nil {
        log.Fatalf("ログファイルの設定に失敗しました: %v", err)
    }
}
```

### ログ出力先の追加

SetLogOutput

 関数を使って、ログ出力先を追加することができます。

```go
import (
    "os"
    "path/to/logger"
)

func main() {
    // ログ出力先に標準エラー出力を追加
    logger.SetLogOutput(logger.ERROR, os.Stderr)
}
```

### ログプレフィックスの設定

SetLogPrefix

 関数を使って、ログプレフィックスを設定することができます。

```go
import "path/to/logger"

func main() {
    // ログプレフィックスを設定
    logger.SetLogPrefix(logger.INFO, "[INFO]")
}
```

### ログフォーマットの設定

SetLogFormat

 関数を使って、ログフォーマットを設定することができます。

```go
import "path/to/logger"

func main() {
    // ログフォーマットを設定
    logger.SetLogFormat("[%s] [%s] %s")
}
```

### ログカラーの設定

SetLogColor

 関数を使って、ログカラーを設定することができます。

```go
import "path/to/logger"

func main() {
    // ログカラーを設定
    logger.SetLogColor(logger.ERROR, "\033[31m") // 赤
}
```

### ログフィルターキーワードの設定

SetLogFilterKeyword

 関数を使って、ログフィルターキーワードを設定することができます。

```go
import "path/to/logger"

func main() {
    // ログフィルターキーワードを設定
    logger.SetLogFilterKeyword("重要")
}
```
