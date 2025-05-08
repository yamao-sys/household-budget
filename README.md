# household-budget

家計簿アプリ

## コマンド類

現状、各サービスごとのデプロイだが、後で整理したい

### マイグレーション

```
cd migrations && gcloud builds submit --config ./cloudbuild.yaml
```

## ハマりポイント

- 本番での環境変数の渡し方が少し煩わしい...
  - Secret Manager に.env.production の元データを作っておき、build 時にそれを.env.production に書き出すことで解決

## 参考

- https://github.com/fullcalendar/fullcalendar-react?tab=readme-ov-file
- https://qiita.com/FumioNonaka/items/1810f7e211638988af62
- https://fullcalendar.io/docs/height
- React Router v7(SPA)でクライアントのミドルウェアを作る
  - https://zenn.dev/coji/articles/react-router-v7-client-middleware
- React Router v7 での request.url
  - https://zenn.dev/atman/scraps/bafae280189ac9
