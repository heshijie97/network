package flaginit

import (
	"flag"
	"fmt"
)

var checkfile, resultfile string
var help bool

func InitFlag() (checkfile, resultfile string) {
	flag.StringVar(&checkfile, "f", "ipv4", "指定待ping的ip文件")
	flag.StringVar(&resultfile, "d", "result.txt", "指定输出结果")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = func() {
		fmt.Println("usage: pingcheck [-f ipv4] [-d result.txt]")
		//打印所有默认值
		flag.PrintDefaults()
	}
	//解析参数
	flag.Parse()
	//展示帮助信息
	if help {
		flag.Usage()
		return
	}
	return
}
