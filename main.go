package main

import (
	"fmt"
	wsjssh "wsj/lib/ssh"
)

func main() {
	fmt.Println("测试不同目录的包")
	wsjssh.Ssh()
}
