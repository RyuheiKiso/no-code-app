-- データベースを選択
USE sample;

-- 既存のテーブルを削除
DROP TABLE IF EXISTS user_roles;

-- ユーザーと権限の関係テーブル
CREATE TABLE user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
);