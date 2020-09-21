package main

import (
	"fmt"
	"flag"
	"net/http"
	"github.com/jianzhuliu/tools/browser"
)

var (
	port int 
	host string
)

func init(){
	flag.IntVar(&port,"port",8787,"set port")
	flag.StringVar(&host,"host","127.0.0.1","set host")
	
	flag.Parse()
}

func main(){
	addr := fmt.Sprintf("%s:%d",host,port)
	
	//配置路由
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello World!!!")
	}))
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "your are visiting %s", r.URL.Path)
	})
	
	//另起 goroutine 用于浏览器自动打开
	go browser.OpenWithNotice(addr)
		
	if err := http.ListenAndServe(addr,nil); err != nil {
		fmt.Println("serve start fail -- ", err)
	}
}


