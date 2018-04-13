/*
 *  flower9
 *  16/12/2017
 *
 */
package model

import (
	"time"
	"io/ioutil"
	"io"
	"encoding/json"
)

type SUser struct {
	User_id              string    `json:"user_id",omitempty`
	Referrer_id          string    `json:"referrer_id",omitempty`        //'推荐人',
	User_password        string    `json:"user_password",omitempty`      //'密码，需加密',
	User_name            string    `json:"user_name",omitempty`          //'用户名称',
	User_email           string    `json:"user_email", omitempty`        //'邮件'
	Real_name            string    `json:"real_name",omitempty`          //'真实姓名',
	Pay_password         string    `json:"pay_password",omitempty`       //'提款密码，需加密',
	Phone_number         string    `json:"phone_number",omitempty`       //'手机号码',
	Integrals            float64   `json:"integrals",omitempty`          //'剩余积分',
	Trial_integrals      float64   `json:"trial_Integrals",omitempty`    //'试用积分注册可设置试用积分，不设为0',
	Create_datetime      time.Time `json:"create_datetime",omitempty`    //'创建时间',
	Last_update_datetime time.Time `json:"lastUpdateDatetime",omitempty` //'最后修改时间',
}

func (user *SUser) JSONSTR2OBJ(r io.Reader) error {
	body, err := ioutil.ReadAll(r)

	if nil == err {
		json.Unmarshal(body, user)
	}
	return err
}

func (user SUser) OBJ2JSONSTR() string {
	body, err := json.Marshal(user)

	if nil == err {
		return string(body)
	}
	return ""
}
