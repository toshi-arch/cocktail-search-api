# liquor_api
酒api

## 手順

docker-composeを立ち上げる
```
docker-compose up -d 
```

Mysqlが立ち上がってるコンテナに入る
```
docker exec -it liquor-mysql bash 
```

Mysqlに入る※パスワードはdocker-composeのenviromentのMYSQL_ROOT_PASSWORD: rootにある

```
mysql -uroot -p
```
