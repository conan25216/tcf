package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"strings"
	"encoding/json"
)


type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func GetParameter(ctx *macaron.Context) {
	para := ctx.Query("para")
	stdout,stderr := processor.RunProcessor(para)
	ctx.Resp.Write([]byte("right output: " +stdout+"\nwrong output: "+stderr))
}

func GetIndex(ctx *macaron.Context) {
	ctx.HTML(200, "testinput")
}

func GetTcfParas(ctx *macaron.Context) {
	parastring :=  strings.TrimSpace(ctx.Query("parastring"))

	//test_data := "52 200 10 167  1102 225 1 195 161 110 1 111 136 1"
	stdout,stderr := processor.RunProcessor(parastring)
	result :=  Result{Stdout:stdout,Stderr:stderr}
	resultJson,_:=json.Marshal(result)
	ctx.Resp.Write([]byte(resultJson))
}