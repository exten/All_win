/*
 *  flower9
 *  11/12/2017
 *
 */
package main

import (
	_ "All_win/server/control"
	_ "All_win/server/dao"
	"log"
	"All_win/server/common"
	"net/http"
	"strconv"
)

func init() {

	log.Println("init http router!~~~")
}

func main() {
	
	log.Println("init All-win-server...")
	log.Print(common.S_config)

	//http.ListenAndServe(":"+strconv.Itoa(common.S_config.Point), nil)

	contextedMux := common.AddContextSupport()
	http.ListenAndServe(":"+strconv.Itoa(common.S_config.Point), contextedMux)
}
