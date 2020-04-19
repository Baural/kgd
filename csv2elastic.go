package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type PseudoCompany struct {
	Bin                      string `json:"bin"`
	Rnn                      string `json:"rnn"`
	TaxpayerOrganization     string `json:"taxpayer_organization"`
	TaxpayerName             string `json:"taxpayer_name"`
	OwnerName                string `json:"owner_name"`
	OwnerIin                 string `json:"owner_iin"`
	OwnerRnn                 string `json:"owner_rnn"`
	CourtDecision            string `json:"court_decision"`
	IllegalActivityStartDate string `json:"illegal_activity_start_date"`
}

func main() {

	if err := CSV2Struct("list_PSEUDO_COMPANY_KZ_ALL.csv"); err != nil {
		panic(err)
	}

}

func CSV2Struct(file string) error {

	csvfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// now, safe to move raw CSV data to struct

	var oneRecord PseudoCompany

	// var allRecords []PseudoCompany

	for _, each := range rawCSVdata {

		oneRecord.Bin = each[1]
		oneRecord.Rnn = each[2]
		oneRecord.TaxpayerOrganization = each[3]
		oneRecord.TaxpayerName = each[4]
		oneRecord.OwnerName = each[5]
		oneRecord.OwnerIin = each[6]
		oneRecord.OwnerRnn = each[7]
		oneRecord.CourtDecision = each[8]
		oneRecord.IllegalActivityStartDate = each[9]
		allRecords = append(allRecords, oneRecord)
	}
	fmt.Println(allRecords)
	fmt.Println(allRecords[7].Bin)
	fmt.Println(allRecords[7].Rnn)
	fmt.Println(allRecords[7].TaxpayerOrganization)
	fmt.Println(allRecords[7].TaxpayerName)
	fmt.Println(allRecords[7].OwnerName)
	fmt.Println(allRecords[7].OwnerIin)
	fmt.Println(allRecords[7].OwnerRnn)
	fmt.Println(allRecords[7].CourtDecision)
	fmt.Println(allRecords[7].IllegalActivityStartDate)

	return err

}
