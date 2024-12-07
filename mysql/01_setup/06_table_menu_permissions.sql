-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_menu_permissions;

-- メニューマスタ権限テーブル
CREATE TABLE cm_m_menu_permissions (
    permission_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '権限ID',
    menu_id VARCHAR(50) NOT NULL COMMENT 'メニューID',
    role VARCHAR(50) NOT NULL COMMENT 'ロール',
    can_view BOOLEAN DEFAULT FALSE COMMENT '表示権限',
    can_edit BOOLEAN DEFAULT FALSE COMMENT '編集権限',
    can_delete BOOLEAN DEFAULT FALSE COMMENT '削除権限',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    created_by VARCHAR(50) NOT NULL COMMENT '作成ユーザー',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    updated_by VARCHAR(50) NOT NULL COMMENT '更新ユーザー',
    FOREIGN KEY (menu_id) REFERENCES cm_m_menu(menu_id)
);