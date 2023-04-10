package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"pingcheck/flaginit"
	"strings"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func PingCheck(host string, writer *bufio.Writer) {
	var r *exec.Cmd
	if strings.Contains(host, ".") {
		r = exec.Command("ping", "-c", "4", "-i", "0.3", "-W", "5", host)

	} else if strings.Contains(host, ":") {
		r = exec.Command("ping6", "-c", "4", "-i", "0.3", "-W", "5", host)
	}
	//ipv4,参数必须拆分
	err := r.Run()
	if err != nil {
		fmt.Printf("%v down\n", host)
		writer.WriteString(fmt.Sprintf("%v down\n", host))
	} else {
		fmt.Printf("%v up\n", host)
		writer.WriteString(fmt.Sprintf("%v up\n", host))
	}
	wg.Done()
}

// 传入参数 -f 指定待ping的ip文件 -d 指定输出结果

func main() {
	//初始化flag
	checkfile, resutlfile := flaginit.InitFlag()
	//打开文件
	cf, err := os.Open(checkfile)
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()
	rf, err := os.Create(resutlfile)
	if err != nil {
		log.Fatal(err)
	}
	defer rf.Close()
	//读入带缓冲的io
	reader := bufio.NewReader(cf)
	//写入带缓冲的io
	writer := bufio.NewWriter(rf)
	defer writer.Flush()
	for {
		host, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		//去除换行符
		host = strings.Trim(host, "\n")
		wg.Add(1) // 登记1个goroutine
		go PingCheck(host, writer)
	}
	wg.Wait() // 阻塞等待登记的goroutine完成
}
