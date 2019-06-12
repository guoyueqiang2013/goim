package service

import (
	"fmt"
	"github.com/guoyueqiang2013/goim/model"
	"github.com/guoyueqiang2013/goim/util"
	"github.com/wendal/errors"
	"math/rand"
	"time"
)

type UserService struct {

}


func (s *UserService)Register(mobile,plainpwd,nickname,avatar,sex string)(user model.User,err error)  {

	//检测手机号是否存在
	tmp := model.User{}
	_,err =DBEngin.Where("mobile=?",mobile).Get(&tmp)
	if err != nil{
		return tmp,err
	}

	if tmp.Id>0{
		return  tmp,errors.New("The mobile is exist!")
	}
	//否则插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d",rand.Int31())
	tmp.Passwd = util.MakePasswd(plainpwd,tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d",rand.Int31())

	_,err = DBEngin.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败

	return tmp,err
}

func (s *UserService)Login(mobile,plainpwd string)(user model.User,err error) {
	//1.通过手机号查询用户
	tmp := model.User{}
	DBEngin.Where("mobile = ?",mobile).Get(&tmp)

	//2.查询到了对比密码
	if tmp.Id == 0{
		return tmp,errors.New("The user IS NOT exist!")
	}

	//3.对比密码是否正确
	if util.ValidatePasswd(plainpwd,tmp.Salt,tmp.Passwd) == false{
		return tmp,errors.New("The user IS NOT exist!")
	}

	//4.刷新token，安全
	str := fmt.Sprintf("%d",time.Now().Unix())
	token :=util.MD5Encode(str)
	tmp.Token = token
	DBEngin.ID(tmp.Id).Cols("token").Update(&tmp)

	return tmp,nil
}

//查找某个用户
func (s *UserService)Find(
	userId int64 )(user model.User) {
	//首先通过手机号查询用户
	tmp :=model.User{}
	DBEngin.ID(userId).Get(&tmp)
	return tmp
}