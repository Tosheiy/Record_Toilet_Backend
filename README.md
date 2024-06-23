# Record_Toilet_Backend
トイレの記録を行うRESTAPIを実装したバックエンド
TLSを用いてHTTPSに対応

## 技術スタック
* Gin
* sqlite3
* TLS

## データベース
### toilet_records
| 　　　　カラム名 　　　　| 　　　　　　　説明 　　　　　　　　　|
|:-----------|:------------|
| id |(INTEDER)　一意に割り振られる |
| description |(TEXT)　説明     |
| created_at |(DATETIME) いつPOSTメソッドにより作成されたか自動で決定  |
| length |(INTEGER) トイレにいた長さ  |
| location |(TEXT) 場所       |
| feeling |(INTEGER) 0,1,2の3段階評価     |
---
主にユーザーが設定する必要があるのが
description, length, location, feelingの４つ
## APIリファレンス
### GET "/toilet"
データベースにあるすべての記録を返す
### POST "/toilet"
description, length, location, feelingの４つをbodyに含める。
データベースに id, created_atを追加した状態で追加。
```
{
   "description": "this is test", 
   "length": 3,
   "location": "home",
   "feeling": 2
}
```
### GET "/toilet/:id"
一致する id の記録を一つ返す
### PUT "/toilet/:id"
一致する　id の記録を更新
### DELETE "/toilet/:id"
一致する　id の記録を削除
