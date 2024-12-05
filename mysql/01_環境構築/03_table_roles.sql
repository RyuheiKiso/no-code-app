-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS roles;

-- 権限テーブル
CREATE TABLE roles (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);