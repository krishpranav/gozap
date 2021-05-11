package zap

import (
	"strings"
)


func Spider(urls, apis string, options OptionsZAP) {
	Run(urls, apis, SpiderAPI, options)
}

func AjaxSpider(urls, apis string, options OptionsZAP) {
	Run(urls, apis, AjaxSpiderAPI, options)
}


func ActiveScan(urls, apis string, options OptionsZAP) {
	Run(urls, apis, AScanAPI, options)
}

func StopSpider(apis string, options OptionsZAP) {
	arrayAPIs := strings.Split(apis, ",")
	for _, v := range arrayAPIs {
		Stop(v, SpiderStop, options)
	}
}

func StopActiveScan(apis string, options OptionsZAP) {
	arrayAPIs := strings.Split(apis, ",")
	for _, v := range arrayAPIs {
		Stop(v, AScanStop, options)
	}

}


func StopAjaxSpider(apis string, options OptionsZAP) {
	arrayAPIs := strings.Split(apis, ",")
	for _, v := range arrayAPIs {
		Stop(v, AjaxSpiderStop, options)
	}
}