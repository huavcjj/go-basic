package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Webアプリケーション：クライアントのHTTPリクエストに応答し、HTTPレスポンスとしてクライアントにHTMLを送り返すコンピュータプログラムを指す。Webサーバ。

// リクエストメソッド
// GET：指定されたリソースを返すようサーバに指示する。
// HEAD：GETと似ているが、メッセージ本体は返さない。レスポンスヘッダを入力したいけれども、重いメッセージ本体のデータをどう処理するかはサーバに任される。
// POST：メッセージ本体のデータをURIで示されたリソースに渡すようサーバに指示する。サーバがメッセージ本体をどう処理するかはサーバに任される。
// PUT：メッセージ本体のデータをURIで指定されたリソースとするようサーバに指示する。URIで示されたリソースがすでに存在する場合は、そのデータが置き換えられる。リソースが存在しない場合は、URIで指定された示された場所に新たなリソースが作成する。
// DELETE：URIによって指定されたリソースを削除するようサーバに指示する。
// TRACE：リクエストに送り返すようサーバに指示する。これにより、途中のサーバがリクエストをどのように変更したかを知ることができる。
// OPTIONS：サーバがサポートするHTTPメソッドのリストを返すようサーバに指示する。
// CONNECT：クライアントとのネットワーク接続を準備するようサーバに指示する。このメソッドは主にHTTPSのためのSSLトンネリングを準備するために使われる。
// PATCH：URIで指定されたリソースをメッセージ本体のデータを従って変更するようにサーバに指示する。

// よく使われるHTTPリクエストヘッダ
// Accept：クライアントがHTTPレスポンスの一部として受け取ることができるコンテンツタイプ。例えば[Accept: text/html]によってクライアントがサーバに伝えるのは、レスポンスの本体のコンテンツタイプをHTMLとして受け取リたいという要求である。
// Accept-Charset：サーバから受け取る際に要求される文字セット。例えば[Accept-Charset: utf-8]によってクライアントがサーバに伝えるのは、レスポンスの本体の文字がUTF-8でエンコードされることという要求である。
// Authorization：ベージック認証の証明書をサーバに送る際に使われる。
// Cookie：クライアントは、呼び出し先のサーバが設定したクッキーを送りなければならない。サーバがそれまでにブラウザに対して２つのクッキーを設定したとすれば、ヘッダーフィールドCookieには二つのクッキーがすべて含まれており、それぞれが[;]で区切った名前と値の組で構成される。例：Cookie: name1=value1; name2=value2
// Content-Length：リクエストの本体の長さをバイト単位で示す。
// Content-Type：リクエスト本体のコンテンツタイプ。POSTまたはPUTが送信された場合、デフォルトのコンテンツタイプはapplication/x-www-form-urlencodedとなる。ファイルをアップロードする場合は、multipart/form-dataとしなければならない。
// Host：サーバ名とポート番号。例：Host: www.example.com:80
// Referrer：リクエストされたページにリンクしていた以前のページのURL。例：Referer: http://www.example.com/index.html
// User-Agent：呼び出しているクライアントに関する説明。
