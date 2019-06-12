package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func RespOK(writer http.ResponseWriter, data interface{},msg string) {
	Resp(writer, 0, data, msg)
}

func RespFail(writer http.ResponseWriter, msg string) {
	Resp(writer, -1, nil, msg)
}

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header 为json
	//application/json
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}

	//定义一个结构体，将结构体转化成json
	writer.Write([]byte(ret))
}

func RespOkList(w http.ResponseWriter,lists interface{},total interface{}){
	//分页数目,
	RespList(w,0,lists,total)
}
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}
