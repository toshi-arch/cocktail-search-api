# liquor_api
酒api

docker-compose up -d (docker-composeを立ち上げる)
docker exec -it liquor-mysql bash (Mysqlが立ち上がってるコンテナに入る)
mysql -uroot -p (Mysqlに入る　※パスワードはdocker-composeのenviromentのMYSQL_ROOT_PASSWORD: rootにある)
