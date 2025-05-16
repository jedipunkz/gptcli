# GPT CLI プロンプト

## プロジェクト概要

このプロジェクトは、オニオンアーキテクチャを学習するためのOpenAI GPTモデルとの対話用CLIアプリケーションです。

## アーキテクチャ

### ドメイン層
- **モデル** (`domain/model/`)
  - `chat_model.go`: チャット関連のモデル定義
  - `generate_model.go`: 生成関連のモデル定義

- **リポジトリインターフェース** (`domain/repository/`)
  - `chat_repository.go`: チャット用リポジトリインターフェース
  - `generate_repository.go`: 生成用リポジトリインターフェース

### ユースケース層
- **チャット** (`usecase/chat_usecase.go`)
- **生成** (`usecase/generate_usecase.go`)

### インフラストラクチャ層
- **OpenAIクライアント** (`infrastructure/openai/`)
  - `chat_client.go`: チャット用クライアント
  - `generate_client.go`: 生成用クライアント

## ビルドと実行

```bash
# ビルド
make build

# クリーン
make clean
```

## 使用方法

### チャットモード
```bash
gptcli chat --model gpt-3.5-turbo
```

### 生成モード
```bash
gptcli generate "プロンプト" --model gpt-3.5-turbo
```

## 学習目標

1. オニオンアーキテクチャの理解と実装
2. 依存関係の適切な管理
3. テスト可能なコード設計
4. クリーンなコード構造の維持

## 要件
### model 名指定

引数で model 名を指定できるように

#### 利用可能なモデル
- 標準モデル:
  - `gpt-4-turbo-preview`
  - `gpt-4`
  - `gpt-4-32k`
  - `gpt-3.5-turbo`
  - `gpt-3.5-turbo-16k`

- エイリアス:
  - `o4` → `gpt-4-turbo-preview`
  - `o3` → `gpt-3.5-turbo`
  - `o3-mini` → `gpt-3.5-turbo`

### 技術要件
- 言語: Go
- 依存パッケージ:
  - github.com/sashabaranov/go-openai
  - github.com/spf13/cobra (CLIフレームワーク)
  - github.com/spf13/viper (設定管理)

## コマンド仕様
```bash
# チャットモード
gptcli chat [--model MODEL] [--system-prompt PROMPT]

# テキスト生成
gptcli generate "プロンプト" [--model MODEL] [--temperature TEMP] [--max-tokens MAX]

# 設定
gptcli config set api-key KEY
gptcli config set default-model MODEL
```

## 出力形式
- チャットモード: 対話形式で出力
- テキスト生成: 生成されたテキストのみを出力
- エラー時: 適切なエラーメッセージを表示

## セキュリティ要件
- APIキーは環境変数または設定ファイルで安全に管理
- 対話履歴はローカルに保存（オプション）

## 今後の拡張性
- 画像生成APIのサポート
- ストリーミング出力のサポート
- プロンプトテンプレートの管理
- 対話履歴のエクスポート/インポート 