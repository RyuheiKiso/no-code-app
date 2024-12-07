-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_users;

-- ユーザーテーブル
CREATE TABLE cm_m_users (
    user_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ユーザーID',
    first_name VARCHAR(50) NOT NULL COMMENT '名',
    last_name VARCHAR(50) NOT NULL COMMENT '姓',
    email VARCHAR(100) NOT NULL COMMENT 'メールアドレス',
    password VARCHAR(255) NOT NULL COMMENT 'パスワード',
    last_login TIMESTAMP NULL COMMENT '最終ログイン日時',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー'
);

-- テーブルにインデックスを追加
CREATE INDEX idx_cm_m_users_email ON cm_m_users (email);

-- データを挿入
INSERT INTO cm_m_users (first_name, last_name, email, password, created_by, updated_by)
VALUES ('太郎', '山田', 'test_1@test.com', 'password', 'system', 'system'),
       ('花子', '山田', 'test_2@test.com', 'password', 'system', 'system');