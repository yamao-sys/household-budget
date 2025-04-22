#!/bin/bash

# シーディング
# docker-compose exec -d business_api make test-seed-local

# テスト用の環境でAPIサーバを起動
# docker-compose exec -d api_server godotenv -f .env.test.local go run main.go

# テスト実行
docker-compose exec frontend pnpm test:e2e

# テスト用DBのリセット
# docker-compose run --rm db mysql -h db -u root -p -e "DROP DATABASE household_budget_test; CREATE DATABASE household_budget_test;"
# docker-compose run --rm migrations sh -c 'make migrate-test-local'
