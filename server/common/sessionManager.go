/*
 *  flower9
 *  15/12/2017
 *
 */
package common

import (
	"sync"
	"time"
	"net/http"
	"log"
	"crypto"
	"encoding/hex"
	"crypto/rand"
)

func init() {
	log.Println("init timer sync...")
}

var timeOutLimit string = "30m"
var _session = make(map[string]*Session)

type Session struct {
	Store map[string]interface{}
	timeOut time.Time
}

var s_lock = &sync.Mutex{}
var Sesion = &Session{make(map[string]interface{}), time.Now()}

/*----------------------------------------------+
|											 	|
|	将用户添加到Session中去						|
|	每一次请求都生成新的session以防session劫持		|
|												|
+----------------------------------------------*/
func (s *Session) AddSession(w http.ResponseWriter, source_id string) (bool, *Session) {
	s_lock.Lock()
	defer s_lock.Unlock()

	session_id := randNum()
	//写cookie到客户端
	cookie := http.Cookie{
		Name:     SESSION_ID,
		Value:    session_id,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 1)}

	//写cookie到客户端
	http.SetCookie(w, &cookie)

	//是否为新建session值对象
	isNew := false
	//获得上一次的session对象
	smap := _session[source_id]
	outLimit, _ := time.ParseDuration(timeOutLimit)

	if nil == smap {
		isNew = true
		smap = &Session{make(map[string]interface{}), time.Now().Add(outLimit)}
	}

	_session[session_id] = smap

	log.Println("session_id is :=> ",session_id, " source_id is :=> ", source_id)

	return isNew, smap
}



//用户退出操作
func (s *Session) DelSession(w http.ResponseWriter, source_id string) {
	s_lock.Lock()
	defer s_lock.Unlock()

	session_id := ""
	//写cookie到客户端
	cookie := http.Cookie{
		Name:     SESSION_ID,
		Value:    session_id,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 1)}

	//写cookie到客户端
	http.SetCookie(w, &cookie)

	//删除服务器上德cookie
	delete(_session, source_id)

	log.Println("login out delete session_id is :=> ",session_id, " source_id is :=> ", source_id)

}


//AddSession 不能调用 GetSession 方法 会出现思索
func getSession(r *http.Request) *Session {
	c, err := r.Cookie(SESSION_ID)
	if nil != err {
		log.Println("read session error!")
	}
	v, _ := _session[c.Value]
	return v
}

/*----------------------------------------------+
|											 	|
|	获得当前用户的Session							|
|	可以发布为外部可见方法，获取到当先的session		|
|												|
+----------------------------------------------*/
func (s *Session) GetSession(r *http.Request) *Session {
	s_lock.Lock()
	defer s_lock.Unlock()

	return getSession(r)
}

/*----------------------------------------------+
|											 	|
|	消除过期的Session								|
|												|
|												|
+----------------------------------------------*/
func (s *Session) destorySession(id string) {
	delete(_session, id)
}

func randNum() string {
	//生成32位随机数-备用
	rands := make([]byte, 128)
	rand.Read(rands)

	md5 := crypto.MD5.New()
	md5.Write(rands)
	return hex.EncodeToString((md5.Sum([]byte("9527"))))
}

func TimerBackRunDestorySession(duration time.Duration) {
	time.AfterFunc(duration, func() {
		for k, v := range _session {
			log.Print(k, v)
			if time.Now().Sub(v.timeOut) > 0 {
				delete(_session, k)
				log.Println("销毁SESSION =>", k)
			}
		}
		TimerBackRunDestorySession(duration)
	})
}
