package ctrl

import (
	"fmt"
	"github.com/guoyueqiang2013/goim/model"
	"github.com/guoyueqiang2013/goim/service"
	"github.com/guoyueqiang2013/goim/util"
	"math/rand"
	"net/http"
)

var userService service.UserService

//测试  curl http://127.0.0.1:8080/user/login -X POST -d "mobile=10000&passwd=111111"
func Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil{
		util.RespFail(w, "Fail to ParseForm!")
		return
	}
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(w, "Fail to login!")
	} else {
		util.RespOK(w, user,"")
	}

}

//测试  curl http://127.0.0.1:8080/user/register -X POST -d "mobile=10000&passwd=111111"
func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		util.RespFail(w, "Fail to ParseForm!")
		return
	}
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_MEN

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.Resp(w, -1, nil, "the mobile is exist")
	} else {
		util.Resp(w, 0, user, "")
	}
}
