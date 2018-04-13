/*
 *  flower9
 *  17/12/2017
 *
 */
package common

import (
	"net/http"
	"log"
	"flag"
	"os"
	"context"
	"All_win/server/model"
	"encoding/json"
	"time"
)

func init() {
	//通过启动参数获得配置文件的位置从而解析出
	dbpsw := flag.String("psw", "123456", "http server point")
	filePath := flag.String("conf", "./config.json", "system config file path")

	flag.Parse()

	S_config.DbPswd = *dbpsw

	file, _ := os.Open(*filePath)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(S_config)
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("S_config=>", S_config)

	//初始化开奖定时器
	if S_config.OpenLtteryLimit > 0{
		t1 := S_config.OpenLtteryLimit * time.Second
		TimerBackRunOpenLotteryTicket(t1)
	}


	//初始化Session的过期时间
	if S_config.SessionTimeOut > 0 {
		t2 := S_config.SessionTimeOut * time.Minute
		TimerBackRunDestorySession(t2)
	}
}

//上下文 Nginx 和前台请求时使用的 路由路径
const (
	SERVER_CONTEXT string = "/allwing/service"
	SESSION_ID     string = "SID_"
)

var (
	mux      = http.NewServeMux()
	S_config = &model.S_config{}
	payTimes = []int{2, 7, 12, 17, 22, 27, 32, 37, 42, 47, 52, 57}
)

//构造请求路径
func BuildHandel(path string, handel func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(SERVER_CONTEXT+path, handel)
}

func AddContextSupport() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie(SESSION_ID)
		if cookie != nil {
			ctx := context.WithValue(r.Context(), SESSION_ID, cookie.Value)
			// WithContext returns a shallow copy of r with its context changed
			// to ctx. The provided ctx must be non-nil.
			mux.ServeHTTP(w, r.WithContext(ctx))
		} else {
			mux.ServeHTTP(w, r)
		}
	})
}


