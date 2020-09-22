"# tools" 

1、download
git clone https://github.com/jianzhuliu/tools.git

2、go get -v github.com/jianzhuliu/tools

3、demo.go
package main 

import (
	"net/http"
	"github.com/jianzhuliu/tools/browser"
)

func main(){
	addr := "127.0.0.1:8080"
	go browser.OpenWithNotice(addr)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World !!!"))
	})
	
	http.ListenAndServe(addr,nil)
}

4、go run demo.go

5、you'll see, browser is open and visit 127.0.0.1:8080
