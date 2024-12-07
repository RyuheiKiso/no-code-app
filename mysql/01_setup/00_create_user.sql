-- ユーザーを追加
CREATE USER 'dev'@'%' IDENTIFIED BY 'dev';

-- ユーザーに権限を付与
GRANT ALL PRIVILEGES ON *.* TO 'dev'@'%' WITH GRANT OPTION;

-- 権限の反映
FLUSH PRIVILEGES;