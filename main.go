package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Element はJSONの構造に対応する構造体
type Element struct {
	Number   int    `json:"number"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// 全元素データを保持するスライス
var elements []Element

// JSONファイルからデータを読み込む関数
func loadElementsFromJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&elements)
}

func main() {
	// 乱数のシード設定
	rand.Seed(time.Now().UnixNano())

	// 起動時にJSONをロード
	if err := loadElementsFromJSON("elements.json"); err != nil {
		fmt.Printf("JSONの読み込みに失敗しました: %v\n", err)
		return
	}
	fmt.Printf("🎉 元素データを読み込みました（合計: %d 元素）\n", len(elements))

	// ---------------------------------------------------
	// 🌐 ルーティング設定
	// ---------------------------------------------------
	
	// 1. トップページ（ブラウザに画面を返す）
	http.HandleFunc("/", handleHome)

	// 2. API（JavaScriptに全データをJSONで返す）
	http.HandleFunc("/api/elements", handleGetElements)

	// 3. API（一応残しておくランダム単発用）
	http.HandleFunc("/api/elements/random", handleGetRandomElement)

	fmt.Println("🚀 サーバーが起動しました: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------
// 🛠️ 各ハンドラー関数の定義（ここが正しく分かれている必要があります）
// ---------------------------------------------------

// トップページでindex.htmlファイルをブラウザに返す
func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

// 全元素データを一括で返すAPI
func handleGetElements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elements)
}

// ランダムに1件返すAPI
func handleGetRandomElement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(elements) == 0 {
		http.Error(w, `{"error": "No data"}`, http.StatusNotFound)
		return
	}
	randomIndex := rand.Intn(len(elements))
	json.NewEncoder(w).Encode(elements[randomIndex])
}