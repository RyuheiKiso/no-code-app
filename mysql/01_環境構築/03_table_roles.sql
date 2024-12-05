-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_roles;

-- 権限テーブル
CREATE TABLE cm_m_roles (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL
);