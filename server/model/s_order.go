/*
 *  flower9
 *  16/12/2017
 *
 */
package model

import (
	"time"
	"io"
	"io/ioutil"
	"chaoshanmj_server/src/gopkg.in/square/go-jose.v1/json"
)

type SOrder struct {
	User_id              string    `json:"user_id",omitempty`              //'用户编号',
	Lottery_period       string    `json:"lottery_period",omitempty`       //'开奖期数',
	Seq_no               int       `json:"seq_no",omitempty`               //'流水号,使用function get_order_seq 取值',
	Location_amt1        float64   `json:"location_amt1",omitempty`        //'位置1金额',
	Location_amt2        float64   `json:"location_amt2",omitempty`        //'位置2金额',
	Location_amt3        float64   `json:"location_amt3",omitempty`        //'位置3金额',
	Location_amt4        float64   `json:"location_amt4",omitempty`        //'位置4金额',
	Location_amt5        float64   `json:"location_amt5",omitempty`        //'位置5金额',
	Location_amt6        float64   `json:"location_amt6",omitempty`        //'位置6金额',
	Location_amt7        float64   `json:"location_amt7",omitempty`        //'位置7金额',
	Location_amt8        float64   `json:"location_amt8",omitempty`        //'位置8金额',
	Location_amt9        float64   `json:"location_amt9",omitempty`        //'位置9金额',
	Location_amt10       float64   `json:"location_amt10",omitempty`       //'位置10金额',
	Location_calc_amt1   float64   `json:"location_calc_amt1",omitempty`   //'位置1计算金额',
	Location_calc_amt2   float64   `json:"location_calc_amt2",omitempty`   //'位置2计算金额',
	Location_calc_amt3   float64   `json:"location_calc_amt3",omitempty`   //'位置3计算金额',
	Location_calc_amt4   float64   `json:"location_calc_amt4",omitempty`   //'位置4计算金额',
	Location_calc_amt5   float64   `json:"location_calc_amt5",omitempty`   //'位置5计算金额',
	Location_calc_amt6   float64   `json:"location_calc_amt6",omitempty`   //'位置6计算金额',
	Location_calc_amt7   float64   `json:"location_calc_amt7",omitempty`   //'位置7计算金额',
	Location_calc_amt8   float64   `json:"location_calc_amt8",omitempty`   //'位置8计算金额',
	Location_calc_amt9   float64   `json:"location_calc_amt9",omitempty`   //'位置9计算金额',
	Location_calc_amt10  float64   `json:"location_calc_amt10",omitempty`  //'位置10计算金额',
	Create_datetime      time.Time `json:"create_datetime",omitempty`      //'创建时间',
	Last_update_datetime time.Time `json:"last_update_datetime",omitempty` //'最后修改时间',
}

func (order *SOrder) JSONSTR2OBJ(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if nil == err {
		json.Unmarshal(body, order)
	}
	return err
}

func (order SOrder) OBJ2JSONSTR() string {
	body, err := json.Marshal(order)

	if nil == err {
		return string(body)
	}
	return ""
}
