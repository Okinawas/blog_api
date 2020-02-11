# TODOアプリ

## Dockerコマンド

### コンテナのビルド
`docker-compose build --no-cache `
※ オプションなしで`build`すると、キャッシュされたdocker-composeを読み込むので、キャッシュを参照しないでビルドする必要がある。


### コンテナを起動
`docker-compose up -d`
※ `-d`は、バックグラウンドで起動するオプション


### コンテナの状態を見る
`docker-compose ps`

下記のように`state`が`Up`になっていれば、起動されている。

```shell
❯ docker-compose ps
   Name                 Command               State                 Ports              
   ---------------------------------------------------------------------------------------
   mysql_todo   docker-entrypoint.sh mysql ...   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
   todo_app     bash                             Up      0.0.0.0:8080->8080/tcp           
```




## Golangコマンド
### goファイルをビルド
`go run {ファイル名}`

ex) `go run main.go`

#### 上記のdockerコンテナ内でのビルドの仕方
1. コンテナのシェルに入る
`docker exec -it  todo_app /bin/sh`

2. ビルドコマンドを実行
`go run main.go`







