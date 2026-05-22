# Atomic Symbol Quiz (爆速！元素記号クイズ)

Go言語（Golang）による超軽量なREST APIバックエンドと、Vanilla JavaScriptによるシングルページアプリケーション（SPA）構成を採用した、リロードレスで快適に動作する元素記号クイズゲームです。

外部のJSONファイルからデータを動的にロードし、フロントエンド側で出題済みの元素を状態管理することで、**「重複なしの全問正解クリア形式」**を実現しています。

---
<img width="1919" height="1035" alt="スクリーンショット 2026-05-23 064350" src="https://github.com/user-attachments/assets/eb71ce3b-ec87-4105-b26b-3f5c834e4d80" />

## 🛠️ システムアーキテクチャ・設計の特徴

本システムは、保守性と拡張性を担保するため、責務が明確に三層に分離された構造科学的設計を採用しています。

[ Data Layer: JSON ] ──(Decode)──> [ Backend Layer: Go (REST API) ]
│
(fetch / JSON)
▼
[ Frontend Layer: JS (SPA) ]


1. **データ層 (`elements.json`)**
   * 原子番号、元素記号、日本語名、カテゴリ（属性）をカプセル化した構造データ。Goのロジックを変更することなく、容易に118元素への拡張や英語名の追加が可能です。
2. **バックエンド層 (`main.go`)**
   * 静的ファイル（`index.html`）の配信、およびクライアントへのデータ供給マシーン（REST API）に特化。高効率なメモリ展開と最速のJSONシリアライズを実現しています。
3. **フロントエンド層 (`index.html`)**
   * 画面遷移によるオーバーヘッド（リロード）を一切排除したSPA構成。APIから取得したデータをプールし、`Array.prototype.splice()` を用いたランダム抽出を行うことで、非破壊的かつ被りのない出題アルゴリズムをブラウザ上で高速に処理します。

---

## 📁 フォルダ構成

```text
atomic-symbol-quiz/
├── main.go          # Go言語：Webサーバー ＆ REST API
├── index.html       # HTML/CSS/JS：フロントエンド（SPA画面）
└── elements.json    # JSON：元素データベース
🚀 動作環境・セットアップ
前提条件
Go言語 (Version 1.20 以上を推奨)

1. 起動方法
プロジェクトのルートディレクトリで以下のコマンドを実行し、Webサーバーを起動します。

Bash
go run main.go
正常に起動すると、ターミナルに以下のログが出力されます。

Plaintext
🎉 元素データを読み込みました（合計: 6 元素）
🚀 サーバーが起動しました: http://localhost:8080
2. ブラウザでアクセス
ローカルサーバーが起動したら、ブラウザから以下のアドレスにアクセスしてください。

クイズ画面（UI）: http://localhost:8080/

🌐 公開APIエンドポイント仕様
バックエンドは、外部連携や将来のモバイルアプリ化を見据え、以下のエンドポイントを標準提供しています。

1. 全元素データ取得 API
URL: /api/elements

Method: GET

Response Content-Type: application/json

ボディ例:

JSON
[
  { "number": 1, "symbol": "H", "name": "水素", "category": "非金属" },
  { "number": 2, "symbol": "He", "name": "ヘリウム", "category": "貴ガス" }
]
2. 単発ランダム元素取得 API (互換用)
URL: /api/elements/random

Method: GET

📝 拡張について
問題（元素データ）を追加したい場合は、elements.json に新たな要素オブジェクトを追記するだけで、システム全体のロジックを変更することなく自動的に出題プールが拡張されます。

JSON
{ "number": 8, "symbol": "O", "name": "酸素", "category": "非金属" }

---
