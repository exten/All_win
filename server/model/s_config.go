package model

import "time"

type S_config struct {
	Point           int           `json:"point",omitempty`              //服务器启动的端口
	OpenLtteryLimit time.Duration `json:"open_lttery_limit", omitempty` //获取开奖号码的时间	间隔为秒
	SessionTimeOut  time.Duration `json:"session_time_out",omitempty`   //session过期时间		间隔为分钟
	DbHost          string        `json:"db_host",omitempty`            //数据库主机
	DbPoint         int           `json:"db_point",omitempty`           //数据库端口
	DbSchema        string        `json:"db_schema",omitempty`          //数据库 库
	DbUser          string        `json:"db_user",omitempty`            //数据库用户
	DbPswd          string                                              //数据库密码安全控制不配置，作为启动参数传入
}
