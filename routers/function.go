package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
)


func GetParameter(ctx *macaron.Context) {
	para := ctx.Query("para")
	stdout,stderr := processor.RunProcessor(para)
	ctx.Resp.Write([]byte("right output: " +stdout+"\nwrong output: "+stderr))
}

func GetIndex(ctx *macaron.Context) {
	ctx.HTML(200, "testinput")
}

func GetTcfParas(ctx *macaron.Context) {
	//para1 :=  strings.TrimSpace(ctx.Query("para1"))
	//para2 :=  strings.TrimSpace(ctx.Query("para2"))
	//para3 :=  strings.TrimSpace(ctx.Query("para3"))
	test_data := "52 200 10 167  1102 225 1 195 161 110 1 111 136 1"
	stdout := processor.RunProcessor(test_data)
	ctx.Resp.Write([]byte("response: " +stdout))
}