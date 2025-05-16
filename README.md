# GPT CLI

このレポジトリはオニオンアーキテクチャの理解を深めるための学習目的のレポジトリです。

ほぼ、AI による生成を行っています。

## 学習目的

このプロジェクトは、以下の目的で作成されています：

1. **オニオンアーキテクチャの理解**
   - クリーンアーキテクチャの実践的な学習
   - 依存関係の適切な管理方法の習得
   - レイヤー分離の重要性の理解

## アーキテクチャ

このプロジェクトは、関心の分離と依存性逆転を重視したオニオンアーキテクチャパターンに従っています。アーキテクチャは3つの主要な層に分かれています：

### ドメイン層（コア）

ビジネスロジックとエンティティを含む最も内側の層です。

#### モデル
- **チャットモード用** (`domain/model/chat_model.go`)
  - `ChatMessage`: チャットメッセージの構造体
    - `Role`: 送信者の役割（例：「user」、「assistant」）
    - `Content`: メッセージ内容
  - `ChatSession`: チャット会話の構造体
    - `Messages`: ChatMessage配列
    - `AddMessage()`: メッセージ追加メソッド

- **生成モード用** (`domain/model/generate_model.go`)
  - `GenerationRequest`: テキスト生成リクエストの構造体
    - `Prompt`: 生成のためのプロンプト
    - `Temperature`: 生成の多様性パラメータ
    - `MaxTokens`: 最大トークン数

#### リポジトリインターフェース
- **チャットモード用** (`domain/repository/chat_repository.go`)
  - `ChatRepository`
    - `CreateChatCompletion()`: チャットメッセージの送受信
- **生成モード用** (`domain/repository/generate_repository.go`)
  - `GenerationRepository`
    - `CreateCompletion()`: テキスト生成

### ユースケース層

アプリケーションのビジネスロジックを実装し、データの流れを調整します。

- **チャットモード用** (`usecase/chat_usecase.go`)
  - `ChatUseCase`
    - `StartChat()`: チャットセッション初期化
    - `SendMessage()`: メッセージ送信と応答処理

- **生成モード用** (`usecase/generate_usecase.go`)
  - `GenerationUseCase`
    - `GenerateText()`: テキスト生成リクエスト管理

### インフラストラクチャ層

外部依存関係と実装を処理する最も外側の層です。

- **チャットモード用** (`infrastructure/openai/chat_client.go`)
  - `ChatClient`
    - `CreateChatCompletion()`: チャットAPI呼び出し
    - `HandleStream()`: チャットストリーミング処理

- **生成モード用** (`infrastructure/openai/generate_client.go`)
  - `GenerationClient`
    - `CreateCompletion()`: テキスト生成API呼び出し
    - `HandleStream()`: 生成ストリーミング処理

## 依存関係の流れ

```
+----------------------------------+
|  +--------------------------+    |
|  |      Domain Layer        |    |
|  |  +------------------+    |    |
|  |  |      Model       |    |    |
|  |  |  +--------+      |    |    |
|  |  |  |Chat    |      |    |    |
|  |  |  |Gen     |      |    |    |
|  |  |  +--------+      |    |    |
|  |  +------------------+    |    |
|  |  +------------------+    |    |
|  |  |Repository        |    |    |
|  |  |Interface         |    |    |
|  |  |-ChatRepo         |    |    |
|  |  |-GenRepo          |    |    |
|  |  +------------------+    |    |
|  +--------------------------+    |
|  +--------------------------+    |
|  |     UseCase Layer        |    |
|  |  +------------------+    |    |
|  |  |   ChatUseCase    |    |    |
|  |  |-StartChat        |    |    |
|  |  |-SendMessage      |    |    |
|  |  +------------------+    |    |
|  |  +------------------+    |    |
|  |  |  GenUseCase      |    |    |
|  |  |-GenerateText     |    |    |
|  |  +------------------+    |    |
|  +--------------------------+    |
|  +--------------------------+    |
|  | Infrastructure Layer     |    |
|  |  +------------------+    |    |
|  |  |OpenAIChatClient  |    |    |
|  |  |-CreateChatComp   |    |    |
|  |  |-HandleStream     |    |    |
|  |  +------------------+    |    |
|  |  +------------------+    |    |
|  |  |OpenAIGenClient   |    |    |
|  |  |-CreateComp       |    |    |
|  |  |-HandleStream     |    |    |
|  |  +------------------+    |    |
|  +--------------------------+    |
+----------------------------------+
```

## 利点

1. **保守性**
   - 関心の明確な分離
   - 個々のコンポーネントの理解と修正が容易

2. **テスト容易性**
   - モック実装の作成が容易
   - 各層の独立したテストが可能

3. **柔軟性**
   - 実装の交換が容易（例：異なるAIサービスへの切り替え）
   - 新機能の追加が簡単

4. **スケーラビリティ**
   - コンポーネント間の明確な境界
   - 機能の拡張が容易

## はじめ方

### ビルド方法

プロジェクトのビルドには、以下のコマンドを使用します：

```bash
# ビルド
make build

# クリーン（実行ファイルの削除）
make clean
```

ビルドが成功すると、`gptcli` という実行ファイルが生成されます。

### 設定

APIキーの設定：

```bash
gptcli config set api-key YOUR_API_KEY
```

### 使用方法

チャットモード：
```bash
gptcli chat --model gpt-3.5-turbo
```

生成モード：
```bash
gptcli generate "プロンプト" --model gpt-3.5-turbo
``` 