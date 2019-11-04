package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"strings"
	"fmt"
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
	age :=  strings.TrimSpace(ctx.Query("age"))
	gender :=  strings.TrimSpace(ctx.Query("gender"))
	cpt :=  strings.TrimSpace(ctx.Query("cpt"))
	rbp :=  strings.TrimSpace(ctx.Query("rbp"))
	sc :=  strings.TrimSpace(ctx.Query("sc"))
	fbs :=  strings.TrimSpace(ctx.Query("fbs"))
	err :=  strings.TrimSpace(ctx.Query("err"))
	mhr :=  strings.TrimSpace(ctx.Query("mhr"))
	eia :=  strings.TrimSpace(ctx.Query("eia"))
	st :=  strings.TrimSpace(ctx.Query("st"))
	slope :=  strings.TrimSpace(ctx.Query("slope"))
	mvcf :=  strings.TrimSpace(ctx.Query("mvcf"))
	tst :=  strings.TrimSpace(ctx.Query("tst"))
	hdd :=  strings.TrimSpace(ctx.Query("hdd"))
	//para3 :=  strings.TrimSpace(ctx.Query("para3"))
	fmt.Printf("age:%s\ngender:%s\ncpt:%s\nrbp:%s\nsc:%s\nfbs:%s\n" +
		"err:%s\nmhr:%s\neia:%s\nst:%s\nslope:%s\nmvcf:%s\ntst:%s\nhdd:%s\n",age,gender,
		cpt,rbp,sc,fbs,err,mhr,eia,st,slope,mvcf,tst,hdd)
	test_data := "52 200 10 167  1102 225 1 195 161 110 1 111 136 1"
	stdout,stderr := processor.RunProcessor(test_data)
	ctx.Resp.Write([]byte("right output: " +stdout+"\nwrong output: "+stderr))
}