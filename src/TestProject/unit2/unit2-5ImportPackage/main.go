package main

import (
	// _ 引入不使用
	_ "importTest/lib1"
	// 直接使用引入方法(不推薦)
	// . "./lib2"
	// 引入並重命名
	mylib2 "importTest/lib2"
)

func main() {
	//lib1.Lib1Test()
	//Lib2Test()
	mylib2.Lib2Test()
}
