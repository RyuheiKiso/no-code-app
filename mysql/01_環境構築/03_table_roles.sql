-- データベースを選択
USE sample;

-- 権限テーブル
CREATE TABLE roles (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);