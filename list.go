package main

import (
	"fmt"
)

func main() {

	kgd := []string{"list_PSEUDO_COMPANY_KZ_ALL.xlsx", "list_WRONG_ADDRESS_KZ_ALL.xlsx", "list_BANKRUPT_KZ_ALL.xlsx", "list_INACTIVE_KZ_ALL.xlsx", "list_INVALID_REGISTRATION_KZ_ALL.xlsx", "list_VIOLATION_TAX_CODE_KZ_ALL.xlsx", "list_TAX_ARREARS_150_KZ_ALL.xlsx"}
	for _, element := range kgd {
		// fmt.Println(i, element)
		fmt.Println("http://kgd.gov.kz/mobile_api/services/taxpayers_unreliable_exportexcel/PSEUDO_COMPANY/KZ_ALL/fileName/" + element)
	}

}
