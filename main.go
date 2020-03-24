package main

import (
	_ "TianGou/boot"
	_ "TianGou/router"
	"github.com/gogf/gf/frame/g"
)


func main() {
	//增加静态资源路径
	s:=g.Server()
	s.SetIndexFolder(true)
	s.AddStaticPath("/js", "public/resource/js")
	s.AddStaticPath("/css", "public/resource/css")
	s.AddStaticPath("/images", "public/resource/images")
	s.AddStaticPath("/fonts", "public/resource/fonts")
	//启动
	s.Run()
}