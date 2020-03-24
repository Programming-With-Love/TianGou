package router

import (
	"TianGou/app/api/diary"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.html")
	})
	s.BindHandler("GET:/diary", func(r *ghttp.Request){
		diary.Get(r)
	})
	s.BindHandler("POST:/diary", func(r *ghttp.Request){
		diary.Post(r)
	})

}
