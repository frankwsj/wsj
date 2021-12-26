package main

import (
	"fmt"
	"wsj/lib/conf"
	_ "wsj/lib/ssh"
)

func main() {
	fmt.Println("测试不同目录的包")
	//wsjssh.Test_SSH_run()
	conf.GetConf()
}
