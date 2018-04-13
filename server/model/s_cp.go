/*
 *  flower9
 *  24/12/2017
 *
 */
package model

type CP struct {
	Rows int64         `json:"rows",omitempty`
	Code string        `json:"code",omitempty`
	Info string        `json:"info",omitempty`
	Data []interface{} `json:"data",omitempty`
}

type PK10 struct {
	Expect        string `json:"expect",omitempty`
	Opencode      string `json:"opencode",omitempty`
	Opentime      string `json:"opentime",omitempty`
	Opentimestamp string `json:"opentimestamp",omitempty`
}
