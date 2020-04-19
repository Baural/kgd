package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type PseudoCompany struct {
	Bin                         string `json:"bin"`
	Court_decision              string `json:"court_decision"`
	Illegal_activity_start_date string `json:"illegal_activity_start_date": "01.03.2006"`
	Owner_iin                   string `json:"owner_iin": "780309301217"`
	Owner_name                  string `json:"owner_name"`
	Owner_rnn                   string `json:"owner_rnn"`
	Rnn                         string `json:"rnn"`
	Taxpayer_name               string `json:"taxpayer_name"`
	Taxpayer_organization       string `json:"taxpayer_organization"`
}

func main() {

	kgd := []string{"list_PSEUDO_COMPANY_KZ_ALL.xlsx", "list_WRONG_ADDRESS_KZ_ALL.xlsx", "list_BANKRUPT_KZ_ALL.xlsx", "list_INACTIVE_KZ_ALL.xlsx", "list_INVALID_REGISTRATION_KZ_ALL.xlsx", "list_VIOLATION_TAX_CODE_KZ_ALL.xlsx", "list_TAX_ARREARS_150_KZ_ALL.xlsx"}
	for _, element := range kgd {
		fileURL := fmt.Sprintf("http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/PSEUDO_COMPANY/KZ_ALL/fileName/%s", element)

		if err := DownloadFile(element, fileURL); err != nil {
			panic(err)
		}
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the filew
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
