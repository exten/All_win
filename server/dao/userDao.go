/*
 *  flower9
 *  18/12/2017
 *
 */
package dao

import (
	"All_win/server/model"
	_"database/sql"
	"log"
)

func FindUserInfoByUid(user *model.SUser)  {

	s_user_str := "SELECT USER_ID, USER_NAME, INTEGRALS FROM T_USER WHERE USER_ID = ?"
	rows, err := MY_SQL_DBC.Query(s_user_str, user.User_id)
	defer rows.Close()
	CheckErr(err)

	for rows.Next() {
		rows.Scan(&user.User_id, &user.User_name, &user.Integrals)
	}
}





func FindUserInfoByEntity(user *model.SUser) (bool, string) {

	s_user_str := `SELECT * FROM T_USER WHERE USER_ID =? AND USER_PASSWORD = ?`
	rows, err := MY_SQL_DBC.Query(s_user_str, user.User_id, user.User_password)
	defer rows.Close()
	CheckErr(err)


	for rows.Next() {
		rows.Scan(&user.User_id, &user.User_name, &user.Integrals)
		return true, user.User_id
	}
	return false, ""
}


//通过model创建数据
func CreateUserEntity(user *model.SUser) int64 {
	tx, _ := MY_SQL_DBC.Begin()

	// 插入一条新数据
	create_user := `
	INSERT INTO
	T_USER(
	user_id,
	referrer_id,
	user_password,
	user_name,
	real_name,
	pay_password,
	phone_number,
	integrals,
	trial_integrals)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := tx.Prepare(create_user)

	if err != nil {
		log.Printf("insert data failed:", err.Error())
		tx.Rollback()
		return -1
	}

	result, err := stmt.Exec(user.User_id, user.Referrer_id, user.User_password, user.User_name,
		user.Real_name, user.Pay_password, user.Phone_number, user.Integrals, user.Trial_integrals)

	if err != nil {
		log.Println("delete data failed:", err.Error())
		tx.Rollback()
		return -1
	}

	//获得影响的行数--用于修改
	/*num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return 0
	}*/

	//获得插入后的id
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("fetch last insert id failed:", err.Error())
		tx.Rollback()
		return -1
	}
	log.Printf("insert new record", id)

	tx.Commit()
	return id
}
