# household-budget

月ごとの収支を可視化する家計簿アプリ

https://github.com/user-attachments/assets/9fde9fcb-9206-4a7c-bd91-056f04b0dd65

## 技術構成

### フロントエンド

- vite
- react(v19)
- react-router(v7)
- openapi-typescript
- openapi-fetch
- tailwindcss
- tanstack/react-query
- eslint
- prettier
- playwright

### バックエンド

- go(v1.24 系)
- gorm
- sql-migrate
- echo
- oapi-codegen
- air
- ozzo-validation
- godotenv
- go-txdb
- stretchr/testify
- go-randomdata
- factory-go/factory

### インフラ

- TiDB
- Artifact Registry
- Cloud Run
- Cloud Storage

## 機能

### 認証

|          | URI      | 権限     |
| -------- | -------- | -------- |
| 会員登録 | /sign_up | 権限なし |
| ログイン | /sign_in | 権限なし |

### 収支管理

|                        | URI                    | 権限     |
| ---------------------- | ---------------------- | -------- |
| 月の収支カレンダー表示 | /monthly_budget        | 認証済み |
| 月の収支詳細           | /monthly_budget/:month | 認証済み |

## 技術選定方針

### REST API or GraphQL

以下 2 点より GraphQL を採用するメリットが小さいと考え、REST API に決定

- 機能数の増加ペースが緩やかで(エンドポイントの増加ペースが緩やかで)
- 取得フィールドに柔軟性を持たせる重要性が小さそうだった

### フロントエンド

以下より、Vite・React(React Router v7)の SPA Mode を選定

- ページ数の拡大ペースが緩やか
- スムーズな画面遷移をしたい(UX 面)
- SSR 不要
- Tanstack Query と合わせてキャッシュをうまく活用できそう

レイアウトはユーティリティが豊富な Tailwind CSS を活用

OpenAPI からコードを自動生成する仕組みを作るため、openapi-typescript とそれを使用して openapi-fetch による fetch

また、フォームの状態管理はそこまでフォームが多くないため、react-hook-form 等は使用せず、useState で行う

#### 状態管理

- API によるリソース取得・非同期の状態管理は Tanstack Query
- それ以外は Context と useState

### バックエンド

- コンテナサイズを小さくできることから Go を選定
- ORM は Ruby on Rails の ActiveRecord に似た形で書け、アソシエーションも扱いやすい GORM
- OpenAPI 定義からボイラーコードを自動生成でき、バリデーションや認証チェックをしやすい oapi-codegen
- 構造体ベースでバリデーションチェックがかけられる ozzo-validation
- テストごとに DB データをクリアするため、go-txdb
- Rspec の FactoryBot に似た形でテストデータを用意できる factory-go

## 設計方針

### フロントエンド

- features に機能ごとの components と hooks を作成
- 可能な限りロジックは hooks に
- API クライアント関連はドメインごとに services 配下に作成し、Tanstack Query をキャッシュキーの一意性とともに扱いやすく
- コンポーネントは専属のものは features に、アプリケーション内で共通で使用するものは root の components 配下に作成
- ページ遷移時の共通処理(認証チェック)はクライアントミドルウェアに

### バックエンド

- Handler → Service のレイヤードアーキテクチャ
- ロジックは Service に寄せる

## テスト方針

- データ周りや処理パターンの網羅性はバックエンド側のテストでカバレッジを担保
  - codecov でカバレッジの増減を可視化
- ハッピーパスのシステムテスト(E2E)で機能が仕様通りに動作することを担保
  - フロントエンド側のリファクタリングがしやすい
  - ライブラリのバージョンアップ時のリグレッションテストにもなる
- VRT は共通部が少ないことからメリットが小さいと考え、現状はなし

## ハマりポイント

- Vite の本番での環境変数の渡し方が少し煩わしい...
  - Secret Manager に.env.production の元データを作っておき、build 時にそれを.env.production に書き出すことで解決

## 参考

- https://github.com/fullcalendar/fullcalendar-react?tab=readme-ov-file
- https://qiita.com/FumioNonaka/items/1810f7e211638988af62
- https://fullcalendar.io/docs/height
- React Router v7(SPA)でクライアントのミドルウェアを作る
  - https://zenn.dev/coji/articles/react-router-v7-client-middleware
- React Router v7 での request.url
  - https://zenn.dev/atman/scraps/bafae280189ac9
- Github Actions で Cloud Run へのデプロイ
  - https://cloud.google.com/blog/ja/products/devops-sre/deploy-to-cloud-run-with-github-actions/
- Github Actions で特定のディレクトリ配下に差分のあった時のみ Job を実行する
  - https://zenn.dev/aishift/articles/f24bef7836aaec
- Tanstack Query Tips
  - https://zenn.dev/taisei_13046/books/133e9995b6aadf/viewer/d0d84c
