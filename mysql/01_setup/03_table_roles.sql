-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_roles;

-- 権限テーブル
CREATE TABLE cm_m_roles (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー'
);

-- サンプルデータを挿入
INSERT INTO cm_m_roles (role_id, role_name, created_by, updated_by)
VALUES (1, 'Admin', 'system', 'system'),
       (2, 'User', 'system', 'system'),
       (3, 'Guest', 'system', 'system');