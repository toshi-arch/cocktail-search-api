
## TODO
記事の内容とは関係ない(と思ったが、記事の内容に合わせれば自然と解決すると思われる)TODO

- レシーバメソッドを作る
- sql.goの実体はcocktail_repositoryだと思う
- gorm.goとsql.goは役割が違うので別のフォルダ
    - gorm.goは記事でいうところのdbconfig.goだし、sql.goは記事でいうところのtesk_repository.goの内容


# フォルダ構成をよりきれいにする
https://zenn.dev/tonbi_attack/articles/cf412d07baa10a#task_repository.go-%E3%81%AE%E5%A4%89%E6%9B%B4%E7%82%B9 を理解する

## 記事の方針(最初はこっちをやる)

### 初見の概念
- レシーバメソッド
- New~()のメソッド
- モデルに対してメソッドを生やす概念


main
- dbを作成
- repository(db)を作成
- handler(repository)を作成
- router.GET("/",handler.GetAll())のようにしてhandlerを呼ぶ

handler
- Handler{repository}構造体を作る
- handlerをレシーバとするメソッドとしてhandler.GetAll(c *gin.Context)を作る
- GetAllの中でrepository.GetAll(c *gin.Context)を使う

repository
- repositroy{db}構造体を作る
- respositoryをレシーバとするメソッドとしてrepository.GetAll(db)を作る
- GetAllの中でdbに接続しレコードを取得する

db
- 今のgorm.goを使ってdbのインスタンスを作成する

## それを踏まえた方針(発展なのでとりあえず今はしなくていい)

### 注意点
- 依存性逆転は無視する
- リポジトリが増えた時の対応はその時考える
- handlerが増えたらどうしようは後から考える

main
- dbを作成
- handler := Handler(db)を作成
- router.GET("/",handler(パッケージ名じゃなくて上で作成したインスタンス名).GetAll())のようにしてhandlerを呼ぶ

handler
- Handler{service}構造体を作る
- handlerをレシーバとするメソッドとしてhandler.GetAll(c *gin.Context)を作る
- GetAllの中でservice.GetAll(c *gin.Context)を使う

service
- Service{repostory}構造体を作る
- serviceをレシーバとするメソッドとしてservice.GetAll(c *gin.Context)を作る
- GetAllの中でrepository.GetAllを使う

repository
- repositroy{db}構造体を作る
- respositoryをレシーバとするメソッドとしてrepository.GetAll(db)を作る
- GetAllの中でdbに接続しレコードを取得する

db
- 今のgorm.goを使ってdbのインスタンスを作成する


