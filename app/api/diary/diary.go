package diary

import (
	"TianGou/app/model/diary"
	_return "TianGou/app/model/return"
	"TianGou/app/utils"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"math/rand"
)

func Get(r *ghttp.Request) {
	all, err := diary.Model.All()
	if err==nil {
		_return.Ok(all[0].Content)
		r.Response.Writeln(_return.Ok(all[rand.Intn(len(all))]))
	}

}
func Post(r *ghttp.Request)  {
	c:=r.GetFormString("content")
	all, _ := diary.Model.All()
	for _, entity := range all {
		s := utils.Params()
		//str hash 值
		hash := s.Simhash(entity.Content)
		////str2 hash 值
		hash2 := s.Simhash(c)
		////计算相似度
		sm := s.Similarity(hash, hash2)
		//距离
		ts := s.HammingDistance(hash, hash2)

		if 0.95<sm&&ts.Int64()<=50{
			fmt.Println(sm)
			fmt.Println(ts)
			rs:= _return.ErrAddMsg("已存在相似度高的日记")
			 rs.Code=101
			r.Response.Writeln(rs)
			return
		}
	}
	entity := diary.Entity{Time: r.GetFormString("time"), Content: c}
	_, err := diary.Model.Insert(entity)
	if err==nil {
		r.Response.Writeln(_return.OkNoData())
	}else {
		r.Response.Writeln(_return.Err())
	}
}