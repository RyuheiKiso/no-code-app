# ORM 使用方法

このドキュメントでは、`go/internal/10_utils/database/orm.go` ファイルに定義されている ORM の使用方法について説明します。

## インポート

まず、必要なパッケージをインポートします。

```go
import (
    "log"
    "go/internal/10_utils/database/orm"
)
```

## ORM の初期化

次に、データベース接続情報を使用して ORM を初期化します。

```go
dataSourceName := "user:password@tcp(localhost:3306)/dbname"
ormInstance, err := orm.NewORM(dataSourceName)
if err != nil {
    log.Fatalf("Failed to initialize ORM: %v", err)
}
defer ormInstance.Close()
```

## データの作成

新しいレコードを作成するには、`Create` メソッドを使用します。

```go
query := "INSERT INTO users (name, email) VALUES (?, ?)"
result, err := ormInstance.Create(query, "John Doe", "john@example.com")
if err != nil {
    log.Fatalf("Failed to create record: %v", err)
}
log.Printf("Record created with ID: %d", result.LastInsertId())
```

## データの読み取り

レコードを読み取るには、`Read` メソッドを使用します。

```go
query := "SELECT id, name, email FROM users WHERE id = ?"
rows, err := ormInstance.Read(query, 1)
if err != nil {
    log.Fatalf("Failed to read record: %v", err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name, email string
    if err := rows.Scan(&id, &name, &email); err != nil {
        log.Fatalf("Failed to scan row: %v", err)
    }
    log.Printf("ID: %d, Name: %s, Email: %s", id, name, email)
}
```

## データの更新

レコードを更新するには、`Update` メソッドを使用します。

```go
query := "UPDATE users SET email = ? WHERE id = ?"
result, err := ormInstance.Update(query, "new-email@example.com", 1)
if err != nil {
    log.Fatalf("Failed to update record: %v", err)
}
log.Printf("Number of records updated: %d", result.RowsAffected())
```

## データの削除

レコードを削除するには、`Delete` メソッドを使用します。

```go
query := "DELETE FROM users WHERE id = ?"
result, err := ormInstance.Delete(query, 1)
if err != nil {
    log.Fatalf("Failed to delete record: %v", err)
}
log.Printf("Number of records deleted: %d", result.RowsAffected())
```

## トランザクションの使用

トランザクションを使用するには、`BeginTransaction`、`CommitTransaction`、および `RollbackTransaction` メソッドを使用します。

```go
tx, err := ormInstance.BeginTransaction()
if err != nil {
    log.Fatalf("Failed to begin transaction: %v", err)
}

query := "UPDATE users SET email = ? WHERE id = ?"
_, err = tx.Exec(query, "transaction-email@example.com", 1)
if err != nil {
    tx.Rollback()
    log.Fatalf("Failed to execute query: %v", err)
}

if err := ormInstance.CommitTransaction(tx); err != nil {
    log.Fatalf("Failed to commit transaction: %v", err)
}
```
