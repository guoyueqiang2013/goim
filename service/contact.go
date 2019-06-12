package service

import (
	"errors"
	"github.com/guoyueqiang2013/goim/model"
	"time"
)

type ContactService struct {
}

func (service *ContactService) AddFriend(uid, did int64) (err error) {

	//如果加自己
	if uid == did {
		return errors.New("不能添加自己为好友啊")
	}

	//判断该好友是否存在
	tmpUser := model.User{}
	DBEngin.Where("id = ?",did).Get(&tmpUser)
	if tmpUser.Id <= 0{
		return errors.New("该用户不存在")
	}


	//判断是否已经加了好友
	tmp := model.Contact{}
	DBEngin.Where("Ownerid = ?", uid).
		And("dstid = ?", did).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)

	//Get获得1条记录
	//count()
	//如果存在记录说明已经是好友了不加
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}

	//事务,
	session := DBEngin.NewSession()
	defer session.Close()
	session.Begin()
	//插自己的
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  uid,
		Dstobj:   did,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插对方的
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  did,
		Dstobj:   uid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//没有错误
	if e2 == nil && e3 == nil {
		//提交
		session.Commit()
		return nil
	} else {
		//回滚
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DBEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DBEngin.In("id", objIds).Find(&coms)
	return coms
}
