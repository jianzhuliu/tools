package main

import (
	"fmt"
	"flag"
	
	"github.com/jianzhuliu/tools/net"
)

var port = flag.Int("port",1024,"preferred server port [1024, 65536]. Default: 1024")

func main(){
	flag.Parse()
	
	checkPortInUse := net.CheckPortInUse(*port)
	fmt.Printf("the port(%d) whether or not in use, %v \n", *port, checkPortInUse)
}