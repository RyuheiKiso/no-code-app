package logger

import (
	"encoding/json" // JSONエンコーディング/デコーディング用パッケージ
	"fmt"           // フォーマットI/O用パッケージ
	"io"            // I/Oインターフェース用パッケージ
	"log"           // ログ出力用パッケージ
	"os"            // OS機能用パッケージ
	"time"          // 時間操作用パッケージ
)

// ログレベルを表す型
type LogLevel int

// ログレベルの定数
const (
	// デバッグレベル
	DEBUG LogLevel = iota
	// 情報レベル
	INFO
	// 警告レベル
	WARN
	// エラーレベル
	ERROR
)

// ログ出力形式を表す型
type LogOutputFormat int

// ログ出力形式の定数
const (
	// テキスト形式
	TEXT LogOutputFormat = iota
	// JSON形式
	JSON
	// CSV形式
	CSV
)

// グローバル変数の宣言
var (
	// 現在のログレベル
	currentLogLevel = INFO
	// 現在のログ出力形式
	currentLogOutputFormat = TEXT
	// ロガーのマップ
	loggers = make(map[LogLevel]*log.Logger)
	// ログファイル
	logFile *os.File
	// ログ出力先のマップ
	logOutputs = make(map[LogLevel][]io.Writer)
	// ログプレフィックスのマップ
	logPrefixes = make(map[LogLevel]string)
	// ログフォーマット
	logFormat = "[%s] [%s] %s"
	// ログレベルごとのカラー
	logColors = map[LogLevel]string{
		DEBUG: "\033[34m", // 青
		INFO:  "\033[32m", // 緑
		WARN:  "\033[33m", // 黄
		ERROR: "\033[31m", // 赤
	}
	// ログフィルターキーワード
	logFilterKeyword string
)

// 初期化関数
func init() {
	// 各ログレベルに対してデフォルトのロガーと出力先を設定
	for _, level := range []LogLevel{DEBUG, INFO, WARN, ERROR} {
		// デフォルトのロガーを標準出力に設定
		loggers[level] = log.New(os.Stdout, "", log.LstdFlags)
		// デフォルトの出力先を標準出力に設定
		logOutputs[level] = []io.Writer{os.Stdout}
	}
}

// ログレベルを設定
func SetLogLevel(level LogLevel) {
	// 現在のログレベルを更新
	currentLogLevel = level
}

// ログ出力形式を設定
func SetLogOutputFormat(format LogOutputFormat) {
	currentLogOutputFormat = format
}

// ログ出力先を設定
// level: ログレベル
// w: ログ出力先のio.Writer
func SetLogOutput(level LogLevel, w io.Writer) {
	// 指定されたログレベルのロガーが存在する場合
	if logger, exists := loggers[level]; exists {
		// 出力先を追加
		logOutputs[level] = append(logOutputs[level], w)
		// 複数の出力先をまとめる
		multiWriter := io.MultiWriter(logOutputs[level]...)
		// ロガーの出力先を更新
		logger.SetOutput(multiWriter)
	}
}

// ログファイルを設定
// level: ログレベル
// filePath: ログファイルのパス
func SetLogFile(level LogLevel, filePath string) error {
	// 既存のログファイルがある場合
	if logFile != nil {
		// ログファイルを閉じる
		logFile.Close()
	}
	var err error
	// ログファイルを開く
	logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// エラーが発生した場合、エラーを返す
	if err != nil {
		return err
	}
	// ログ出力先にファイルを追加
	SetLogOutput(level, logFile)
	return nil
}

// ログファイルをローテーション
// level: ログレベル
// filePath: ログファイルのパス
func RotateLogFile(level LogLevel, filePath string) error {
	// 既存のログファイルがある場合
	if logFile != nil {
		// ログファイルを閉じる
		logFile.Close()
	}
	// バックアップファイルのパスを生成
	backupPath := fmt.Sprintf("%s.%s", filePath, time.Now().Format("20060102T150405"))
	// ログファイルをバックアップ
	err := os.Rename(filePath, backupPath)
	// エラーが発生した場合、エラーを返す
	if err != nil {
		return err
	}
	// 新しいログファイルを設定
	return SetLogFile(level, filePath)
}

// ログプレフィックスを設定
// level: ログレベル
// prefix: ログプレフィックス
func SetLogPrefix(level LogLevel, prefix string) {
	// 指定されたログレベルのプレフィックスを設定
	logPrefixes[level] = prefix
}

// ログフォーマットを設定
// format: ログフォーマット
func SetLogFormat(format string) {
	// ログフォーマットを更新
	logFormat = format
}

// ログカラーを設定
// level: ログレベル
// color: ログカラー
func SetLogColor(level LogLevel, color string) {
	// 指定されたログレベルのカラーを設定
	logColors[level] = color
}

// ログフィルターキーワードを設定
// keyword: ログフィルターキーワード
func SetLogFilterKeyword(keyword string) {
	// ログフィルターキーワードを設定
	logFilterKeyword = keyword
}

// ログメッセージを出力
// level: ログレベル
// message: ログメッセージ
func logMessage(level LogLevel, message string) {
	switch currentLogOutputFormat {
	case TEXT:
		logMessageAsText(level, message)
	case JSON:
		logMessageAsJSON(level, message)
	case CSV:
		logMessageAsCSV(level, message)
	}
}

// テキスト形式でログメッセージを出力
// level: ログレベル
// message: ログメッセージ
func logMessageAsText(level LogLevel, message string) {
	// 現在のログレベル以上の場合
	if level >= currentLogLevel {
		// フィルターキーワードが設定されている場合
		if logFilterKeyword != "" && !containsKeyword(message, logFilterKeyword) {
			// キーワードが含まれていない場合、ログを出力しない
			return
		}
		// タイムスタンプを取得
		timestamp := time.Now().Format(time.RFC3339)
		// ログレベルを文字列に変換
		levelStr := getLevelString(level)
		// プレフィックスを取得
		prefix := logPrefixes[level]
		// カラーを取得
		color := logColors[level]
		// フォーマットされたメッセージを生成
		formattedMessage := fmt.Sprintf(logFormat, timestamp, levelStr, message)
		// 指定されたログレベルのロガーが存在する場合
		if logger, exists := loggers[level]; exists {
			// カラーとプレフィックスを付けてメッセージを出力
			logger.Printf("%s%s%s\033[0m", color, prefix, formattedMessage)
		}
	}
}

// JSON形式でログメッセージを出力
// level: ログレベル
// message: ログメッセージ
func logMessageAsJSON(level LogLevel, message string) {
	// 現在のログレベル以上の場合
	if level >= currentLogLevel {
		// タイムスタンプを取得
		timestamp := time.Now().Format(time.RFC3339)
		// ログレベルを文字列に変換
		levelStr := getLevelString(level)
		// ログエントリをマップとして作成
		logEntry := map[string]string{
			"timestamp": timestamp,
			"level":     levelStr,
			"message":   message,
		}
		// ログエントリをJSON形式に変換
		jsonMessage, _ := json.Marshal(logEntry)
		// JSON形式のメッセージを出力
		logMessageAsText(level, string(jsonMessage))
	}
}

// CSV形式でログメッセージを出力
// level: ログレベル
// message: ログメッセージ
func logMessageAsCSV(level LogLevel, message string) {
	// 現在のログレベル以上の場合
	if level >= currentLogLevel {
		// タイムスタンプを取得
		timestamp := time.Now().Format(time.RFC3339)
		// ログレベルを文字列に変換
		levelStr := getLevelString(level)
		// CSV形式のメッセージを生成
		csvMessage := fmt.Sprintf("%s,%s,%s", timestamp, levelStr, message)
		// 指定されたログレベルのロガーが存在する場合
		if logger, exists := loggers[level]; exists {
			// メッセージを出力
			logger.Println(csvMessage)
		}
	}
}

// メッセージがキーワードを含むか確認
// message: ログメッセージ
// keyword: フィルターキーワード
func containsKeyword(message, keyword string) bool {
	// キーワードが含まれているか確認
	return keyword == "" || (keyword != "" && contains(message, keyword))
}

// メッセージがキーワードを含むか確認（ヘルパー関数）
// message: ログメッセージ
// keyword: フィルターキーワード
func contains(message, keyword string) bool {
	// メッセージの先頭にキーワードが含まれているか確認
	return len(message) >= len(keyword) && message[:len(keyword)] == keyword
}

// ログレベルを文字列に変換
// level: ログレベル
func getLevelString(level LogLevel) string {
	switch level {
	case DEBUG:
		// デバッグレベル
		return "DEBUG"
	case INFO:
		// 情報レベル
		return "INFO"
	case WARN:
		// 警告レベル
		return "WARN"
	case ERROR:
		// エラーレベル
		return "ERROR"
	default:
		// 不明なレベル
		return "UNKNOWN"
	}
}

// デバッグログを出力
// message: ログメッセージ
func Debug(message string) {
	// デバッグレベルのメッセージを出力
	logMessage(DEBUG, message)
}

// インフォログを出力
// message: ログメッセージ
func Info(message string) {
	// 情報レベルのメッセージを出力
	logMessage(INFO, message)
}

// 警告ログを出力
// message: ログメッセージ
func Warn(message string) {
	// 警告レベルのメッセージを出力
	logMessage(WARN, message)
}

// エラーログを出力
// message: ログメッセージ
func Error(message string) {
	// エラーレベルのメッセージを出力
	logMessage(ERROR, message)
}

// ログメッセージをJSON形式で出力
// level: ログレベル
// message: ログメッセージ
func LogMessageAsJSON(level LogLevel, message string) {
	// 現在のログレベル以上の場合
	if level >= currentLogLevel {
		// タイムスタンプを取得
		timestamp := time.Now().Format(time.RFC3339)
		// ログレベルを文字列に変換
		levelStr := getLevelString(level)
		// ログエントリをマップとして作成
		logEntry := map[string]string{
			"timestamp": timestamp,
			"level":     levelStr,
			"message":   message,
		}
		// ログエントリをJSON形式に変換
		jsonMessage, _ := json.Marshal(logEntry)
		// JSON形式のメッセージを出力
		logMessage(level, string(jsonMessage))
	}
}
