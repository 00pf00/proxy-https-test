package main

import (
	"00pf00/proxy-https-test/pkg/conf"
	"00pf00/proxy-https-test/pkg/log"
	"flag"
	"fmt"
)

func main() {
	url := flag.String("c","LS","LS ,BASH,LOGS,LOGSF")
	httpsconf := flag.String("conf","./https.toml","The path of the configuration file")
	flag.Parse()
	err := conf.InitConf(*httpsconf)
	if err != nil {
		fmt.Printf("load config fail error=%v\n",err)
		return
	}
	log.InitLogs()

	
}
