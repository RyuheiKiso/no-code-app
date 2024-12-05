-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS cm_m_user_roles;

-- ユーザーと権限の関係テーブル
CREATE TABLE cm_m_user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES cm_m_users(user_id),
    FOREIGN KEY (role_id) REFERENCES cm_m_roles(role_id)
);