/*
 *  flower9
 *  18/12/2017
 *
 */
package dao

import (
	"All_win/server/model"
	"log"
	"database/sql"
	"strconv"
	"fmt"
)

//查询用户账户金额
func FindUserAccount(userId string) {

}

//查询单个数字总下注金额
func FindOrderTotalInfo(lottery_period string) *model.SResult {

	s_result := `
SELECT
lottery_period,
	location_amt1,
	location_amt2,
	location_amt3 ,
	location_amt4 ,
	location_amt5 ,
	location_amt6 ,
	location_amt7 ,
	location_amt8 ,
	location_amt9 ,
	location_amt10
FROM lottery_result WHERE lottery_period =?
`

	rows, err := MY_SQL_DBC.Query(s_result, lottery_period)
	defer rows.Close()
	CheckErr(err)

	total_amt := &model.SResult{}
	for rows.Next() {
		rows.Scan(&total_amt.Lottery_period,
			&total_amt.Location_amt1, &total_amt.Location_amt2,
			&total_amt.Location_amt3, &total_amt.Location_amt4,
			&total_amt.Location_amt5, &total_amt.Location_amt6,
			&total_amt.Location_amt7, &total_amt.Location_amt8,
			&total_amt.Location_amt9, &total_amt.Location_amt10)
	}
	return total_amt
}

func FindOrderUserPayInfo(user_id, lottery_period string) *model.SOrder {

	order_amt := &model.SOrder{}
	s_user_str := "SELECT * FROM user_order WHERE USER_ID =? AND lottery_period = ?"
	rows, err := MY_SQL_DBC.Query(s_user_str, user_id, lottery_period)
	defer rows.Close()
	CheckErr(err)

	for rows.Next() {
		rows.Scan(&order_amt.User_id, &order_amt.Lottery_period,
			&order_amt.Location_amt1, &order_amt.Location_amt1,
			&order_amt.Location_amt2, &order_amt.Location_amt3,
			&order_amt.Location_amt4, &order_amt.Location_amt5,
			&order_amt.Location_amt6, &order_amt.Location_amt7,
			&order_amt.Location_amt9, &order_amt.Location_amt10)
	}

	return order_amt
}

//查询历史开奖金额
func FindOrderUserPayListInfo(userId string) []*model.SOrder {

	s_user_order := `
	SELECT
	lottery_period,
	location_amt1,
	location_amt2,
	location_amt3 ,
	location_amt4 ,
	location_amt5 ,
	location_amt6 ,
	location_amt7 ,
	location_amt8 ,
	location_amt9 ,
	location_amt10
	FROM USER_ORDER WHERE USER_ID = ?
	`

	orderList := make([]*model.SOrder, 0)
	rows, err := MY_SQL_DBC.Query(s_user_order, userId)
	if err == nil {
		for rows.Next() {
			order := &model.SOrder{}
			orderList = append(orderList, order)
			rows.Scan(&order.Lottery_period,
				&order.Location_amt1, &order.Location_amt2,
				&order.Location_amt3, &order.Location_amt4,
				&order.Location_amt5, &order.Location_amt6,
				&order.Location_amt7, &order.Location_amt8,
				&order.Location_amt9, &order.Location_amt10)
		}
	}

	return orderList
}

//下注接口
func CreateUserOrderPay(order *model.SOrder) int64 {

	tx, _ := MY_SQL_DBC.Begin()

	//1.查询order表判断有没有数据
	i_err := insertUserOrder(order, tx)
	if i_err != nil {
		tx.Rollback()
		return 0
	}

	//2.查询result表判断有没有数据
	var r_count int64
	tx.QueryRow("SELECT COUNT(1) FROM LOTTERY_RESULT WHERE lottery_period = ?", order.Lottery_period).Scan(&r_count)

	if r_count <= 0 {
		i_err := insertReult(order, tx)
		if nil != i_err {
			tx.Rollback()
			return 0
		}
	} else {
		u_err := updateReult(order, tx)
		if nil != u_err {
			tx.Rollback()
			return 0
		}
	}

	//3.减少用户的金额
	subUserAMt(order, tx)

	tx.Commit()
	return 1
}

func subUserAMt(order *model.SOrder, tx *sql.Tx) error {
	stmt, err := tx.Prepare(`UPDATE T_USER SET integrals = integrals - ? WHERE USER_ID = ?`)
	if err != nil {
		log.Println("update data failed.", err)
		return err
	}
	sum_order_amt :=
		order.Location_amt1 + order.Location_amt2 + order.Location_amt3 + order.Location_amt4 + order.Location_amt5 +
			order.Location_amt6 + order.Location_amt7 + order.Location_amt8 + order.Location_amt9 + order.Location_amt10;
	result, err := stmt.Exec(sum_order_amt, order.User_id)
	if nil != err {
		log.Println("update data failed.", err)
		return err
	}

	row_count, err := result.RowsAffected()
	if nil != err {
		log.Println("update data failed.", err)
		return err
	}

	log.Println("update data count.", strconv.FormatInt(row_count, 5))
	return nil
}

func insertUserOrder(order *model.SOrder, tx *sql.Tx) error {
	c_order_t := `
	INSERT INTO user_order(
	user_id,
	seq_no,
	lottery_period,
	location_amt1,
	location_amt2,
	location_amt3 ,
	location_amt4 ,
	location_amt5 ,
	location_amt6 ,
	location_amt7 ,
	location_amt8 ,
	location_amt9 ,
	location_amt10)
	VALUES(
		?, get_order_seq(?, ?), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	`

	c_stmt, err := tx.Prepare(c_order_t)
	log.Printf(c_order_t)
	if err != nil {
		log.Printf("insert data failed:", err.Error())
		return err
	}

	c_result, err := c_stmt.Exec(
		order.User_id, order.User_id,
		order.Lottery_period, order.Lottery_period,
		order.Location_amt1, order.Location_amt2,
		order.Location_amt3, order.Location_amt4,
		order.Location_amt5, order.Location_amt6,
		order.Location_amt7, order.Location_amt8,
		order.Location_amt9, order.Location_amt10)

	if err != nil {
		log.Println("delete data failed:", err.Error(), c_result)
		return err
	}

	//更新后得到更新ID
	id, err := c_result.LastInsertId()
	if err != nil {
		log.Println("fetch last insert id failed:", err.Error(), id)
		return err
	}

	return nil
}

func updateUserOrder(order *model.SOrder, tx *sql.Tx) error {

	u_result_t := `
	UPDATE user_order
	SET
	location_amt1 = location_amt1 + ?,
	location_amt2 = location_amt2 + ?,
	location_amt3 = location_amt3 + ?,
	location_amt4 = location_amt4 + ?,
	location_amt5 = location_amt5 + ?,
	location_amt6 = location_amt6 + ?,
	location_amt7 = location_amt7 + ?,
	location_amt8 = location_amt8 + ?,
	location_amt9 = location_amt9 + ?,
	location_amt10 = location_amt10 + ?
	WHERE lottery_period = ? and user_id = ?
	`

	u_stmt, err := tx.Prepare(u_result_t)
	log.Printf(u_result_t)
	if err != nil {
		log.Printf("insert data failed:", err.Error())
		return err
	}

	u_result, err := u_stmt.Exec(
		order.Location_amt1, order.Location_amt2,
		order.Location_amt3, order.Location_amt4,
		order.Location_amt5, order.Location_amt6,
		order.Location_amt7, order.Location_amt8,
		order.Location_amt9, order.Location_amt10,
		order.Lottery_period, order.User_id)

	if err != nil {
		log.Println("delete data failed:", err.Error(), u_result)
		return err
	}

	//获得影响的行数--用于修改
	num, err := u_result.RowsAffected()
	if err != nil {
		log.Println("fetch row affected failed:", err.Error(), num)
		return err
	}
	return nil
}

func insertReult(order *model.SOrder, tx *sql.Tx) error {
	c_order_t := `
	INSERT INTO lottery_result(
	lottery_period,
	status,
	location_amt1,
	location_amt2,
	location_amt3 ,
	location_amt4 ,
	location_amt5 ,
	location_amt6 ,
	location_amt7 ,
	location_amt8 ,
	location_amt9 ,
	location_amt10)
	VALUES(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	`

	c_stmt, err := tx.Prepare(c_order_t)
	log.Printf(c_order_t)
	if err != nil {
		log.Printf("insert data failed:", err.Error())
		return err
	}

	c_result, err := c_stmt.Exec(order.Lottery_period, "0",
		order.Location_amt1, order.Location_amt2,
		order.Location_amt3, order.Location_amt4,
		order.Location_amt5, order.Location_amt6,
		order.Location_amt7, order.Location_amt8,
		order.Location_amt9, order.Location_amt10)

	if err != nil {
		log.Println("delete data failed:", err.Error(), c_result)
		return err
	}

	//更新后得到更新ID
	id, err := c_result.LastInsertId()
	if err != nil {
		log.Println("fetch last insert id failed:", err.Error(), id)
		return err
	}

	return nil
}

func updateReult(order *model.SOrder, tx *sql.Tx) error {
	u_result_t := `
	UPDATE lottery_result
	SET
	location_amt1 = location_amt1 + ?,
	location_amt2 = location_amt2 + ?,
	location_amt3 = location_amt3 + ?,
	location_amt4 = location_amt4 + ?,
	location_amt5 = location_amt5 + ?,
	location_amt6 = location_amt6 + ?,
	location_amt7 = location_amt7 + ?,
	location_amt8 = location_amt8 + ?,
	location_amt9 = location_amt9 + ?,
	location_amt10 = location_amt10 + ?
	WHERE lottery_period = ? and status = '0'
	`

	u_stmt, err := tx.Prepare(u_result_t)
	log.Printf(u_result_t)
	if err != nil {
		log.Printf("insert data failed:", err.Error())
		return err
	}

	u_result, err := u_stmt.Exec(
		order.Location_amt1, order.Location_amt2,
		order.Location_amt3, order.Location_amt4,
		order.Location_amt5, order.Location_amt6,
		order.Location_amt7, order.Location_amt8,
		order.Location_amt9, order.Location_amt10,
		order.Lottery_period)

	if err != nil {
		log.Println("delete data failed:", err.Error(), u_result)
		return err
	}

	//获得影响的行数--用于修改
	num, err := u_result.RowsAffected()
	if err != nil {
		log.Println("fetch row affected failed:", err.Error(), num)
		return err
	}
	return nil
}

func GetlotteryPeriod() int {
	row := MY_SQL_DBC.QueryRow(`SELECT period
										FROM pk10_period
										WHERE start_bet_datetime <= now()
										AND end_bet_datetime >= now()`)
	period := 0
	row.Scan(&period)
	return period
}



//计算佣金
func CalcAmt(period string) {
	handle, err := MY_SQL_DBC.Prepare(`call calc(?)`)
	if err != nil {
		panic(err.Error())
	}
	defer handle.Close()

	//call procedure
	var result sql.Result
	result, err = handle.Exec(period)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}
