package ctrl

import (
	"github.com/guoyueqiang2013/goim/args"
	"github.com/guoyueqiang2013/goim/service"
	"github.com/guoyueqiang2013/goim/util"
	"log"
	"net/http"
)

var contactService *service.ContactService

func AddFriend(w http.ResponseWriter, r *http.Request) {

	//定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg args.ContactArg
	util.Bind(r,&arg)
	//调用service
	err := contactService.AddFriend(arg.Userid,arg.Dstid)
	//
	if err!=nil{
		util.RespFail(w,err.Error())
	}else{
		util.RespOK(w,nil,"好友添加成功")

	}
}


func LoadFriends(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r,&arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w,users,len(users))

}

func Community(w http.ResponseWriter, r *http.Request) {
	log.Println("Community")
	util.RespFail(w,"test")
}

