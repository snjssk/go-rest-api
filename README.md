# go-rest-api

REST APIを実装するサンプル。  
複雑なことはせずできるだけシンプルな作りとしてみる

## overview
- DBはsqlite3
- GET /item の場合はselect文を発行して結果をjsonで返す
- GET /item/{id} の場合はidを絞り込みして返す

## module
### mux 
- https://github.com/gorilla/mux

### ini
- https://github.com/go-ini/ini
