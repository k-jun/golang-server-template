# golang-server-template

## Architecture
Clean Architectureを採用

```
.
├── README.md
├── env.sh
├── go.mod
├── go.sum
├── infrastructure
│   ├── datastore(DB関連)
│   └── routers(各種Router)
├── interface
│   └── controllers(各種Controller))
├── main.go
└── sqlboiler.toml
```

## ORM

GolangMigrate(https://github.com/golang-migrate/migrate)を採用

```sh
brew install golang-migrate
```

### SQLファイルを生成

```sh
migrate create -ext sql -dir ./infrastructure/datastore//migrations [filename]
```

命名規則
- `CreateUser` Userモデル作成
- `AlterUser` Userモデルの変更

### Migration Exec

```
migrate -path ./infrastructure/datastore/mysql/migrations  -database 'postgresql://[username]:[password]@localhost:5432/[database_name]?sslmode=disable' up
migrate -path ./infrastructure/datastore/mysql/migrations  -database 'postgresql://[username]:[password]@localhost:5432/[database_name]?sslmode=disable' down
```
数字を最後につけることで任意の回数migrationを実行

### Migration Fix

```
migrate -path ./infrastructure/datastore/migrations  -database postgresql://[username]:[password]@localhost:5432/[database_name]?sslmode=disable force 20190326174351
```
指定したversionに移動する。Migrationファイルは実行しない。 SQLの書き方を間違えてエラーがmigrationで出た時などに任意のVersionに戻るときに使う。

### ORM

sqlboiler(https://github.com/volatiletech/sqlboiler)を採用

```sh
go get -u -t github.com/volatiletech/sqlboiler

# Also install the driver of your choice, there exists psql, mysql, mssql
# These are separate binaries.
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

```
sqlboiler mysql -o ./infrastructure/datastore/mysql/models
```
