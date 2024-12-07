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

-- サンプルデータを挿入
INSERT INTO cm_m_menu (menu_id, menu_name, parent_id, menu_type, url, icon, status, description, created_by, updated_by)
VALUES ('1', 'ホーム', NULL, 'ページ', '/home', 'home_icon', 'active', 'ホームページ', 'system', 'system'),
       ('2', 'ユーザー管理', NULL, 'ページ', '/users', 'user_icon', 'active', 'ユーザー管理ページ', 'system', 'system'),
       ('3', '設定', NULL, 'ページ', '/settings', 'settings_icon', 'active', '設定ページ', 'system', 'system'),
       ('4', 'プロフィール', '2', 'ページ', '/users/profile', 'profile_icon', 'active', 'プロフィールページ', 'system', 'system');