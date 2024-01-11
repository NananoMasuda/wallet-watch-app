# 家計簿アプリ
## 使用言語
- golang
- vue.js
- mysql

## ゴール
- カレンダー画面をホームにした、家計簿アプリを作る
  
## 追加したい機能
  - ログイン機能（個人用だけど、勉強のため）
  - 収入・支出計算機能
  - 固定費を一気に（毎月）追加する機能　など

## TODO
- [ ] ログイン機能実装


## 環境構築　注意点
- .envファイルをbackendコンテナ内にコピーすること（ビルド時の上位階層のCOPY方法がわからなかった。もっといい方法があれば変更する）
  - wallet-watch-app配下で、```docker cp .env wwa-backend:/go/src/github.com/NananoMasuda/wallet-watch-app/backend/.env```