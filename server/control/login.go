/*
 *  flower9
 *  11/12/2017
 *
 */
package control

import (
	"net/http"
	"crypto"
	"log"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
	"All_win/server/common"
	"All_win/server/model"
	"All_win/server/dao"
)

func init() {
	log.Println("login init ...")
	common.BuildHandel("/login", doLogin)
	common.BuildHandel("/loginOut", doLoginOut)
}

func doLogin(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if nil != err {
		w.Write([]byte("0"))
		return
	}

	user := &model.SUser{}
	json.Unmarshal(body, user)

	log.Println("person :=> ", user)

	//需要查询数据库
	md5 := crypto.MD5.New()
	md5.Write([]byte(user.User_password))
	psd := hex.EncodeToString(md5.Sum([]byte("FUCK HACK")))
	user.User_password = psd
	//差数据库
	isOk, user_id := findeLoginUser(user)

	if isOk {

		var s_sid = r.RemoteAddr
		c, _ := r.Cookie(common.SESSION_ID)
		if nil != c {
			s_sid = c.Value
		}
		//记录session
		isNew, session := common.Sesion.AddSession(w, s_sid)
		if isNew {
			session.Store["uid"] = user_id
		}

		//1:SUCCES
		w.Write([]byte("1"))
	} else {
		//0:FAIL
		w.Write([]byte("0"))
	}
}


//登出注销操作
func doLoginOut(w http.ResponseWriter, r *http.Request) {
	var s_sid = r.RemoteAddr
	c, _ := r.Cookie(common.SESSION_ID)
	if nil != c {
		s_sid = c.Value
	}

	common.Sesion.DelSession(w, s_sid)
}



func findeLoginUser(user *model.SUser) (bool, string) {
	return dao.FindUserInfoByEntity(user)
}
