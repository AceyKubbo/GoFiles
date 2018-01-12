package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi,world")
	//使用package模式来编译go,有包名字要求.一般来说启动的源码包名字都用main,基本上所有项目都这么做
	//1.7以前不用main做package模式的文件名来编译直接编译不出来...
	//就算你用files的模式来编译,你也得先配置工作目录啊
}
