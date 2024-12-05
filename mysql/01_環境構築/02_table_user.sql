-- データベースを選択
USE sample;

-- テーブルを作成
CREATE TABLE cm_t_users (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ユーザーID',
    first_name VARCHAR(50) NOT NULL COMMENT '名',
    last_name VARCHAR(50) NOT NULL COMMENT '姓',
    email VARCHAR(100) NOT NULL COMMENT 'メールアドレス',
    password VARCHAR(255) NOT NULL COMMENT 'パスワード',
    last_login TIMESTAMP NULL COMMENT '最終ログイン日時',
    permission_code VARCHAR(20) NULL COMMENT '権限コード',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー'
);

-- テーブルにインデックスを追加
CREATE INDEX idx_cm_t_users_email ON cm_t_users (email);