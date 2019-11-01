package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"strings"
)


func GetParameter(ctx *macaron.Context) {
	para := ctx.Query("para")
	stdout := processor.RunProcessor(para)
	ctx.Resp.Write([]byte("response: " +stdout))
}

func GetIndex(ctx *macaron.Context) {
	ctx.HTML(200, "testinput")
}

func GetTcfParas(ctx *macaron.Context) {
	para1 :=  strings.TrimSpace(ctx.Query("para1"))
	para2 :=  strings.TrimSpace(ctx.Query("para2"))
	para3 :=  strings.TrimSpace(ctx.Query("para3"))
	stdout := processor.RunProcessor(para1+para2+para3)
	ctx.Resp.Write([]byte("response: " +stdout))
}