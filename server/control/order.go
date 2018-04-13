/*
 *  flower9
 *  11/12/2017
 *
 */
package control

import (
	"net/http"
	"log"
	"All_win/server/common"
	"All_win/server/model"
	"All_win/server/dao"
	"encoding/json"
	"math/rand"
	"strconv"
)

func init() {
	log.Println("order init ...")
	common.BuildHandel("/userOrderpay", userOrderPay)
	common.BuildHandel("/findUserInfo", findUserInfo)
	common.BuildHandel("/findCurrNumInfo", findCurrNumInfo)
	common.BuildHandel("/findOrderTotalInfo", findOrderTotalInfo)
	common.BuildHandel("/findOrderUserPayInfo", findOrderUserPayInfo)
	common.BuildHandel("/findOrderUserPayListInfo", findOrderUserPayListInfo)
}

func userOrderPay(w http.ResponseWriter, r *http.Request) {

	session := common.Sesion.GetSession(r);
	uid := session.Store["uid"]
	log.Println("order get user id :=> ", uid)
	if len(uid.(string)) <= 0 {
		//需要重新登录
		w.Write([]byte("-1"))
	}

	period_id := strconv.Itoa(dao.GetlotteryPeriod())
	//时间段内不可购买
	if len(period_id) < 3 {
		w.Write([]byte("-2"))
	}

	order := &model.SOrder{}
	err := order.JSONSTR2OBJ(r.Body)
	if err != nil {
		//序列化错误
		w.Write([]byte("-1"))
	}
	order.User_id = uid.(string)
	//去数据库查询期号
	order.Lottery_period = period_id

	//验证金额是否超出
	user := &model.SUser{}
	user.User_id = uid.(string)
	dao.FindUserInfoByUid(user)

	sum_order_amt :=
		order.Location_amt1 + order.Location_amt2 + order.Location_amt3 + order.Location_amt4 + order.Location_amt5 +
			order.Location_amt6 + order.Location_amt7 + order.Location_amt8 + order.Location_amt9 + order.Location_amt10;
	if sum_order_amt > user.Integrals {
		//余额不足
		w.Write([]byte("0"))
	}

	dao.CreateUserOrderPay(order)

}

//查询用户的金额
func findUserInfo(w http.ResponseWriter, r *http.Request) {
	user := &model.SUser{}
	err := user.JSONSTR2OBJ(r.Body)

	session := common.Sesion.GetSession(r);
	if nil == session {
		w.Write([]byte("-1"))
		return
	}

	uid := session.Store["uid"]
	user.User_id = uid.(string)
	if err == nil {
		dao.FindUserInfoByUid(user)
	}

	w.Write([]byte(user.OBJ2JSONSTR()))
}

//当期开奖号码
func findCurrNumInfo(w http.ResponseWriter, r *http.Request) {
	/*resp, err := http.Get("http://r.apiplus.net/newly.do?token=df77fd1d5c296dcdf4d7d6c689c55dd1&code=bjpk10&rows=5&format=json")
	defer resp.Body.Close()
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			// handle error
			cp := &model.CP{}
			json.Unmarshal(body, cp)

			log.Println(cp.Data)
			log.Println(string(body))
		}
	}*/

	result_num := make([]int, 0)
	for i := 0; i < 10; i++ {
		result_num = append(result_num, rand.Intn(10))
	}

	str, _ := json.Marshal(result_num)
	//1.更新结果表的开奖号码 和期号
	log.Println("查询当期开奖号...", str)
	<-common.Async.AddSync()
	w.Write([]byte(str))

	//2.如果拿到 开奖号码 触发计算佣金存储过程 替换result_num[2]
	//dao.CalcAmt(string(result_num[2]))
}

//总下注金额
func findOrderTotalInfo(w http.ResponseWriter, r *http.Request) {
	lottery_period := strconv.Itoa(dao.GetlotteryPeriod())
	result := dao.FindOrderTotalInfo(lottery_period)
	w.Write([]byte(result.OBJ2JSONSTR()))
}

//当前用户总下注金额
func findOrderUserPayInfo(w http.ResponseWriter, r *http.Request) {
	session := common.Sesion.GetSession(r);
	uid := session.Store["uid"]
	log.Println("order get user id :=> ", uid)

	lottery_period := strconv.Itoa(dao.GetlotteryPeriod())
	order := dao.FindOrderUserPayInfo(uid.(string), lottery_period)
	w.Write([]byte(order.OBJ2JSONSTR()))
}

//查询当天下注信息
func findOrderUserPayListInfo(w http.ResponseWriter, r *http.Request) {
	session := common.Sesion.GetSession(r);

	//找不到需要提示登录返回值
	uid := session.Store["uid"]
	user_order_list := dao.FindOrderUserPayListInfo(uid.(string))

	uol, _ := json.Marshal(user_order_list)
	log.Printf("find user user_order_list...", string(uol))

	w.Write(uol)
}
