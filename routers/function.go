package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"encoding/json"
	"fmt"
)


type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type Params struct{
	Age string `json:"age"`
	Gender string `json:"gender"`
	Rbp string `json:"rbp"`
	Mhr string `json:"mhr"`
	Sc string `json:"sc"`
	Fbs string `json:"fbs"`
	St string `json:"st"`
	Tst string `json:"tst"`
	Hdd string `json:"hdd"`
	Cpt string `json:"cpt"`
	Err string `json:"err"`
	Eia string `json:"eia"`
	Slope string `json:"slope"`
	Mvcf string `json:"mvcf"`
}

func GetParameter(ctx *macaron.Context) {
	para := ctx.Query("para")
	stdout,stderr := processor.RunProcessor(para)
	ctx.Resp.Write([]byte("right output: " +stdout+"\nwrong output: "+stderr))
}

func GetIndex(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}

func GetTcfParas(ctx *macaron.Context) {
	params := new(Params)
	values,_ := ctx.Req.Body().Bytes()
	json.Unmarshal(values,params)
	fmt.Println("age is",params.Age)
	fmt.Println("mvcf is",params.Mvcf)
	abc := params.ParasMerge()
	fmt.Println("params string is ", abc)
	stdout,stderr := processor.RunProcessor(abc)
	result :=  Result{Stdout:stdout,Stderr:stderr}
	resultJson,_:=json.Marshal(result)
	ctx.Resp.Write([]byte(resultJson))
}

func (p *Params)ParasMerge()(params string){
	//v := reflect.ValueOf(p)
	//count := v.NumField()
	//
	//for i := 0; i < count; i++ {
	//	f := v.Field(i)
	//	switch f.Kind() {
	//	case reflect.String:
	//		fmt.Println(f.String())
	//		params+=f.String()
	//	}
	//}
	//return
	return p.Age+" "+p.Gender+" "+p.Rbp+" "+p.Mhr+" "+p.Sc+" "+p.Fbs+" "+p.St+" "+p.Tst+" "+
		p.Hdd+" "+p.Cpt+" "+p.Err+" "+p.Eia+" "+p.Slope+" "+p.Mvcf
}