package main

import (
	"fmt"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	cron "gopkg.in/robfig/cron.v2"
)

type TaxInfo struct {
	TaxInfoDescription string
	url                string
}

func main() {
	c := schedule()
	c.Start()
	<-make(chan int)
}

func schedule() (c *cron.Cron) {
	c = cron.New()
	_, err := c.AddFunc("@every 0h01m0s", load)
	if err != nil {
		panic(err)
	}
	return
}

func load() {

	var downloads = []TaxInfo{
		{"Pseudo Company", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/PSEUDO_COMPANY/KZ_ALL/fileName/list_PSEUDO_COMPANY_KZ_ALL.xlsx"},
		// {"Bankrupt", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/BANKRUPT/KZ_ALL/fileName/list_BANKRUPT_KZ_ALL.xlsx"},
		// {"Inactive", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/INACTIVE/KZ_ALL/fileName/list_INACTIVE_KZ_ALL.xlsx"},
		// {"Invalid Registration", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/INVALID_REGISTRATION/KZ_ALL/fileName/list_INVALID_REGISTRATION_KZ_ALL.xlsx"},
		// {"Viovation Tax Code", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/VIOLATION_TAX_CODE/KZ_ALL/fileName/list_VIOLATION_TAX_CODE_KZ_ALL.xlsx"},
		// {"Tax Arrears 150", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/TAX_ARREARS_150/KZ_ALL/fileName/list_TAX_ARREARS_150_KZ_ALL.xlsx"},
		// {"Wrong Address", "http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/WRONG_ADDRESS/KZ_ALL/fileName/list_WRONG_ADDRESS_KZ_ALL.xlsx"},
	}
	var answers = []string{}
	for download := range downloads {
		f := DownloadTaxinfo(downloads[download].url)
		if f == nil {
			answers = append(answers, "Could not read the bad taxpayer information "+downloads[download].TaxInfoDescription)
			continue
		}
		if errorT := parseAndSendToES(downloads[download].TaxInfoDescription, f); errorT != nil {
			answers = append(answers, "Could not parse or send to ES, the bad taxpayer information "+downloads[download].TaxInfoDescription)
		} else {
			answers = append(answers, "Have succesfully downloaded the bad taxpayer information "+downloads[download].TaxInfoDescription)
		}
	}
	for answer := range answers {
		fmt.Println(answers[answer])
	}

}

func DownloadTaxinfo(url string) *excelize.File {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	f, err := excelize.OpenReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	return f
}
