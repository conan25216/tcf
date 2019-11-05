package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"encoding/json"
	"reflect"
	"fmt"
	"errors"
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
	paramString, err := params.ParasMerge()
	if err != nil  {
		return //if miss value we dont run process, just return empty string
	}
	stdout,stderr := processor.RunProcessor(paramString)
	result :=  Result{Stdout:stdout,Stderr:stderr}
	resultJson,_:=json.Marshal(result)
	ctx.Resp.Write([]byte(resultJson))
}

func (p *Params)ParasMerge()(params string,err error){
	v := reflect.ValueOf(*p)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		//	//fmt.Println("f.kind is", f)
		case reflect.String:
			fmt.Println(f.String())
			if f.String() == "" {
				return "",errors.New("miss value error!")
			}
			params+=f.String()
		}
	}
	return params,nil
	//return p.Age+" "+p.Gender+" "+p.Rbp+" "+p.Mhr+" "+p.Sc+" "+p.Fbs+" "+p.St+" "+p.Tst+" "+
	//	p.Hdd+" "+p.Cpt+" "+p.Err+" "+p.Eia+" "+p.Slope+" "+p.Mvcf
}