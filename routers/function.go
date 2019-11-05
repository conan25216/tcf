package routers

import (
	"gopkg.in/macaron.v1"
	"tcf/processor"
	"encoding/json"
	"reflect"
	"errors"
	"fmt"
	"strings"
	"strconv"
)


type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type Params struct{
	Age string `json:"age"`
	Gender string `json:"gender"`
	Cpt string `json:"cpt"`
	Rbp string
	Sc string `json:"sc"`
	Fbs string `json:"fbs"`
	Err string `json:"err"`
	Mhr string `json:"mhr"`
	Eia string `json:"eia"`
	St string `json:"st"`
	Slope string `json:"slope"`
	Mvcf string `json:"mvcf"`
	Tst string `json:"tst"`
	Hdd string `json:"hdd"`
}


func GetIndex(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}

func GetTcfParas(ctx *macaron.Context) {
	params := new(Params)
	values,_ := ctx.Req.Body().Bytes()
	json.Unmarshal(values,params)
	paramString, err := params.ParasMerge()
	paramString = strings.TrimSpace(paramString)
	if err != nil  {
		return //if miss value we dont run process, just return empty string
	}
	fmt.Println("paramString is ", paramString)
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
		case reflect.String: // security check
			intValue,err := strconv.Atoi(f.String())
			fmt.Println("checked value is", intValue)
			if err != nil || intValue <0 || intValue > 1000{
				return "", err
			}
			if f.String() == "" {
				return "",errors.New("miss value error!")
			}
			params+=f.String()
			params+=" "
		}
	}
	return params,nil
	//return p.Age+" "+p.Gender+" "+p.Rbp+" "+p.Mhr+" "+p.Sc+" "+p.Fbs+" "+p.St+" "+p.Tst+" "+
	//	p.Hdd+" "+p.Cpt+" "+p.Err+" "+p.Eia+" "+p.Slope+" "+p.Mvcf
}