package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
	// 第1引数はネットワークアドレス、第2引数はリクエストを処理するハンドラ。
	// ネットワークアドレスが""の場合、ポート80がデフォルトになる、ハンドラがnilの場合、デフォルトのマルチプレクサDefaultServeMuxが使われる。
}
