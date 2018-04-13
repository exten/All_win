/*
 *  flower9
 *  11/12/2017
 *
 */
package control

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto"
	"encoding/hex"
	"log"
	"All_win/server/model"
	"All_win/server/common"
	"All_win/server/dao"
	"strconv"
)

func init() {
	log.Println("login register ...")
	common.BuildHandel("/register", doRegister)
}

func doRegister(w http.ResponseWriter, r *http.Request) {
	var status int64
	body, err := ioutil.ReadAll(r.Body)

	//服务解析错误
	if nil != err {
		w.Write([]byte("-1"))
	}else {
		user := &model.SUser{}
		json.Unmarshal(body, user)

		//对密码加密
		md5 := crypto.MD5.New()
		md5.Write([]byte(user.User_password))
		psd := hex.EncodeToString(md5.Sum([]byte("FUCK HACK")))
		user.User_password = psd
		log.Printf(" register password can't print in log!! => ", psd+"-")

		//数据库操作
		status = createUser(user)
		log.Print(strconv.FormatInt(status, 10))
		w.Write([]byte(strconv.FormatInt(status, 10)))
	}

}




//注册用户
func createUser(user *model.SUser) int64 {

	_user:= &model.SUser{}

	//查询是否已经存在此用户
	_user.User_id = user.User_id
	dao.FindUserInfoByUid(_user)
	if len(_user.User_name) > 0 {
		return -2
	}

	//验证推荐人是否存在
	_user.User_name = ""
	_user.User_id = user.Referrer_id
	dao.FindUserInfoByUid(_user)
	if len(_user.User_name) <= 0 {
		return -3
	}

	//账号注册
	user_id := dao.CreateUserEntity(user)
	//数据库注册错误
	if user_id < 0 {
		log.Println("register error !!!!!")
		return -1
	}

	log.Printf("insert into db recode id is:%v", user_id)
	return user_id
}
