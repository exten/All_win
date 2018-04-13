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
	"encoding/json"
)

type SResult struct {
	Lottery_period       string    `json:"lottery_period",omitempty`       //'开奖期数',
	Create_datetime      time.Time `json:"create_datetime",omitempty`      //'开奖时间',
	Last_update_datetime time.Time `json:"last_update_datetime",omitempty` //'开奖时间',
	Status               string    `json:"status",omitempty`               //'开奖状态',
	Location_num1        float64   `json:"location_num1",omitempty`       //'位置1数字',
	Location_num2        float64   `json:"location_num2",omitempty`       //'位置2数字',
	Location_num3        float64   `json:"location_num3",omitempty`       //'位置3数字',
	Location_num4        float64   `json:"location_num4",omitempty`       //'位置4数字',
	Location_num5        float64   `json:"location_num5",omitempty`       //'位置5数字',
	Location_num6        float64   `json:"location_num6",omitempty`       //'位置6数字',
	Location_num7        float64   `json:"location_num7",omitempty`       //'位置7数字',
	Location_num8        float64   `json:"location_num8",omitempty`       //'位置8数字',
	Location_num9        float64   `json:"location_num9",omitempty`       //'位置9数字',
	Location_num10       float64   `json:"location_num10",omitempty`      //'位置10数字',
	Location_amt1        float64   `json:"location_amt1",omitempty`       //'位置1金额',
	Location_amt2        float64   `json:"location_amt2",omitempty`       //'位置2金额',
	Location_amt3        float64   `json:"location_amt3",omitempty`       //'位置3金额',
	Location_amt4        float64   `json:"location_amt4",omitempty`       //'位置4金额',
	Location_amt5        float64   `json:"location_amt5",omitempty`       //'位置5金额',
	Location_amt6        float64   `json:"location_amt6",omitempty`       //'位置6金额',
	Location_amt7        float64   `json:"location_amt7",omitempty`       //'位置7金额',
	Location_amt8        float64   `json:"location_amt8",omitempty`       //'位置8金额',
	Location_amt9        float64   `json:"location_amt9",omitempty`       //'位置9金额',
	Location_amt10       float64   `json:"location_amt10",omitempty`      //'位置10金额',
	Location_calc_amt1   float64   `json:"location_calc_amt1",omitempty`  //'位置1计算金额',
	Location_calc_amt2   float64   `json:"location_calc_amt2",omitempty`  //'位置2计算金额',
	Location_calc_amt3   float64   `json:"location_calc_amt3",omitempty`  //'位置3计算金额',
	Location_calc_amt4   float64   `json:"location_calc_amt4",omitempty`  //'位置4计算金额',
	Location_calc_amt5   float64   `json:"location_calc_amt5",omitempty`  //'位置5计算金额',
	Location_calc_amt6   float64   `json:"location_calc_amt6",omitempty`  //'位置6计算金额',
	Location_calc_amt7   float64   `json:"location_calc_amt7",omitempty`  //'位置7计算金额',
	Location_calc_amt8   float64   `json:"location_calc_amt8",omitempty`  //'位置8计算金额',
	Location_calc_amt9   float64   `json:"location_calc_amt9",omitempty`  //'位置9计算金额',
	Location_calc_amt10  float64   `json:"location_calc_amt10",omitempty` //'位置10计算金额',
	Commission_amt       float64   `json:"commission_amt",omitempty`       // '佣金',
}




func (result * SResult) JSONSTR2OBJ(r io.Reader) error{
	body, err := ioutil.ReadAll(r)

	if nil == err {
		json.Unmarshal(body, result)
	}
	return err
}

func (result SResult) OBJ2JSONSTR() string{
	body, err := json.Marshal(result)

	if nil == err {
		return string(body)
	}
	return ""
}
