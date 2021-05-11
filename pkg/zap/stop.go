package zap

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

// const values
const AScanStop = "/JSON/ascan/action/stopAllScans/?"
const SpiderStop = "/JSON/spider/action/stopAllScans/?"
const AjaxSpiderStop = "/JSON/ajaxSpider/action/stopAllScans/?"

func Stop(api, prefix string, options OptionsZAP) {
	req, err := http.NewRequest("GET", api+prefix, nil)
	if err != nil {
		panic(err)
	}
	if options.APIKey != "" {
		req.Header.Add("X-ZAP-API-Key", options.APIKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	log.WithFields(log.Fields{
		"ZAP API": prefix,
	}).Info("Stoped")

	if err != nil {

	}

	defer resp.Body.Close()
}