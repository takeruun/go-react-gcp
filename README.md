# これは
Google Cloud App Engine を使用して、React + Go で作成したサンプルアプリケーションです。

## 構成
バックエンド：Go
フロントエンド：React ＋ TypeScript

build ファイルは、バックエンドの dist/ に配置して、/ にアクセスされると dist/assets/index.html を返すようにする

## デプロイ
1. Google Cloud Platform のプロジェクトを作成しておく
2. gcloud 初期化
```bash
gcloud init
```
3. フロントエンドのビルド
```bash
cd frontend
yarn build
```
4. バックエンドのデプロイ
```bash
cd backend
gcloud app deploy
```

