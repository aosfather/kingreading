package main

import "fmt"

func main() {
	fmt.Println("hello")
	app := Application{}
	app.Init()  //初始化
	app.start() //服务启动
}

type Application struct {
}

func (this *Application) Init() {

}
func (this *Application) start() {

}
