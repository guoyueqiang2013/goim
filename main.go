package main

import (
	"fmt"
	"github.com/guoyueqiang2013/goim/ctrl"
	"html/template"
	"log"
	"net/http"
)

//万能模板渲染
func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Printf("Tpl name : %s\n", tplname)
		http.HandleFunc(tplname, func(writer http.ResponseWriter,
			request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})

	}
}

func main() {
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriends)
	http.HandleFunc("/contact/loadcommunity", ctrl.Community)
	http.HandleFunc("/contact/addfriend", ctrl.AddFriend)
	http.HandleFunc("/user/login", ctrl.Login)
	http.HandleFunc("/user/register", ctrl.Register)
	http.HandleFunc("/chat", ctrl.Chat)

	//静态文件目录
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	RegisterView()
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		return
	}

}
