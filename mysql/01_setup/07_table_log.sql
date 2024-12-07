-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_t_log;

-- ログテーブル
CREATE TABLE cm_t_log (
    log_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'ログID',
    program_id VARCHAR(50) NOT NULL COMMENT 'プログラムID',
    log_level VARCHAR(50) NOT NULL COMMENT 'ログレベル',
    message TEXT NOT NULL COMMENT 'メッセージ',
    client_ip VARCHAR(45) DEFAULT NULL COMMENT 'クライアントIPアドレス',
    server_ip VARCHAR(45) DEFAULT NULL COMMENT 'サーバーIPアドレス',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー'
);