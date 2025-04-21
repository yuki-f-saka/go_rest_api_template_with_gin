# Go REST API Template with Gin

A template for building REST APIs in Go using the Gin framework and following Clean Architecture principles.

## Overview

This project serves as a template for creating REST APIs in Go, implementing Clean Architecture patterns with the Gin web framework. It provides a solid foundation for building scalable and maintainable web applications.

## Features

- Clean Architecture implementation
- RESTful API using Gin framework
- Structured logging
- Configuration management
- Error handling
- Middleware support
- Database integration (prepared)
- API documentation
- Testing setup

## Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

## Project Structure

```
TBD
```

## Getting Started

1. Clone the repository
```bash
git clone <repository-url>
```

2. Install dependencies
```bash
go mod download
```

3. Run the application
```bash
go run main.go
```

## Architecture

This project follows Clean Architecture principles, separating the codebase into distinct layers:

- **Entities**: Core business entities
- **Use Cases**: Application business rules
- **Interface Adapters**: Controllers, presenters, and gateways
- **Frameworks & Drivers**: External frameworks and tools

## API Documentation

TBD

## Testing

To run tests:
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

# Go REST API Template with Gin (日本語)

Ginフレームワークとクリーンアーキテクチャを使用したGo言語によるREST APIテンプレート

## 概要

このプロジェクトは、クリーンアーキテクチャパターンを実装し、Gin Webフレームワークを使用してGo言語でREST APIを作成するためのテンプレートです。スケーラブルで保守性の高いWebアプリケーションを構築するための堅固な基盤を提供します。

## 特徴

- クリーンアーキテクチャの実装
- GinフレームワークによるRESTful API
- 構造化ログ
- 設定管理
- エラーハンドリング
- ミドルウェアサポート
- データベース統合（準備済み）
- APIドキュメント
- テスト環境

## 前提条件

- Go 1.21以上
- Docker（オプション、コンテナ化用）

## プロジェクト構成

```
TBD
```

## 始め方

1. リポジトリのクローン
```bash
git clone <repository-url>
```

2. 依存関係のインストール
```bash
go mod download
```

3. アプリケーションの実行
```bash
go run main.go
```

## アーキテクチャ

このプロジェクトはクリーンアーキテクチャの原則に従い、コードベースを以下の層に分離しています：

- **エンティティ**: コアビジネスエンティティ
- **ユースケース**: アプリケーションのビジネスルール
- **インターフェースアダプター**: コントローラー、プレゼンター、ゲートウェイ
- **フレームワークとドライバー**: 外部フレームワークとツール

## APIドキュメント

TBD

## テスト

テストの実行方法：
```bash
go test ./...
```

## 貢献

貢献を歓迎します！Pull Requestを気軽に提出してください。

## ライセンス

このプロジェクトはMITライセンスの下で提供されています - 詳細はLICENSEファイルを参照してください。