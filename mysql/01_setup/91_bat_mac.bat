@echo off

REM MySQLコンテナを起動
docker run --name mysql_container -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_DATABASE=sample -e MYSQL_USER=dev -e MYSQL_PASSWORD=dev -p 3306:3306 -d mysql:latest

REM MySQLコンテナが起動するまで待機
timeout /t 30

REM MySQLコンテナに接続してスクリプトを実行
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/00_create_user.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/01_create_database.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/02_table_user.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/03_table_roles.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/04_table_user_roles.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/05_table_menu.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/06_table_menu_permissions.sql
docker exec -i mysql_container mysql -u root -prootpassword < ../01_setup/07_table_log.sql

echo 環境構築が完了しました。
pause