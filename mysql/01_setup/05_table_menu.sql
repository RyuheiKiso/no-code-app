-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_menu;

-- メニューマスタテーブル
CREATE TABLE cm_m_menu (
    menu_id VARCHAR(50) PRIMARY KEY COMMENT 'メニューID',
    menu_name VARCHAR(100) NOT NULL COMMENT 'メニュー名',
    parent_id VARCHAR(50) DEFAULT NULL COMMENT '親メニューID',
    menu_type VARCHAR(50) DEFAULT NULL COMMENT 'メニュータイプ',
    url VARCHAR(255) DEFAULT NULL COMMENT 'URL',
    icon VARCHAR(50) DEFAULT NULL COMMENT 'アイコン',
    status VARCHAR(20) DEFAULT 'active' COMMENT 'ステータス',
    description TEXT DEFAULT NULL COMMENT '説明',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー',
    FOREIGN KEY (parent_id) REFERENCES cm_m_menu(menu_id)
);

