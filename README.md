# ginのサンプル

Postgresqlが必要

Dockerでの起動方法
```bash
$ docker run --name postgres -p 5432:5432 -e POSTGRES_USER=go -e POSTGRES_PASSWORD=go -d postgres:10.5
$ psql -h localhost -U go
```

テーブル作成
```sql
create table user_account (
    Id text primary key,
    Name text not null,
    Pass text not null
);
```