package main

import (
	"context" // Context object for Do() methodsd
	"fmt"     // Format and print cluster data
	"log"     // Log errors and quit
	"reflect" // Get object methods and attributes

	//"strconv"

	// Convert _id int to string
	"time" // Set a timeout for the connection

	// Import the Olivere Golang driver for Elasticsearch
	"github.com/olivere/elastic"
)

const (
	index = "pseudo_company"
)

// Film represents a movie with some properties.
type ElasticIndex struct {
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

	// Instantiate a client instance of the elastic library
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second), // quit trying after 5 seconds
	)

	// Check and see if olivere's NewClient() method returned an error
	if err != nil {
		fmt.Println("elastic.NewClient() ERROR: %v", err)
		log.Fatalf("quiting connection..")
	} else {
		// Print client information
		fmt.Println("client:", client)
		fmt.Println("client TYPE:", reflect.TypeOf(client))
	}

	// Declare a timeout context for the API calls
	ctx, stop := context.WithTimeout(context.Background(), 3*time.Second)
	defer stop()

	// Check if the Elasticsearch index already exists
	exist, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v", err)

		// Index some documents if the index exists
	} else if exist {

		// Instantiate new Elasticsearch documents from the ElasticIndex struct
		newDocs := []ElasticIndex{
			{Bin: "000140013420", Rnn: "600400113669", TaxpayerOrganization: "товарищество с ограниченной ответственностью 'Горизонт'", TaxpayerName: "", OwnerName: "ЕРЛАН ТЕКЕЕВ БАЗАРБАЕВИЧ", OwnerIin: "810618300930", OwnerRnn: "600400113669", CourtDecision: "Приговор от 02.04.2008 года.", IllegalActivityStartDate: "2008-04-02"},
			{Bin: "000240013048", Rnn: "302000068756", TaxpayerOrganization: "Товарищество с ограниченной ответственностью 'Актив'", TaxpayerName: "", OwnerName: "ШОКИР УСМАНОВ ЯКУБОВИЧ", OwnerIin: "", OwnerRnn: "302000068756", CourtDecision: "Письмо ОБЛ НД вх.№9676 от 31.05.2011", IllegalActivityStartDate: "2000-08-07"},
			{Bin: "000340012073", Rnn: "600900158040", TaxpayerOrganization: "ТОО  АВТО ТРЕК (11)	", TaxpayerName: "", OwnerName: "", OwnerIin: "", OwnerRnn: "600900158040", CourtDecision: "не является плательщиком НДС", IllegalActivityStartDate: "2004-10-01"},
		}

		// Declare an integer for the doc _id
		// var id int = 1 // Omit this if you want dynamically generated _id

		// Iterate over the docs and index them one-by-one
		for _, doc := range newDocs {
			_, err = client.Index().
				Index(index).
				Type("companies"). // unique doctype now deprecated
				Id(doc.Bin).       // Increment _id counter same like BIn
				BodyJson(doc).     // pass struct instance to BodyJson

				// Omit this if you want dynamically generated _id
				// Id(strconv.Itoa(id)). // Convert int to string

				Do(ctx) // Initiate API call with context object

			// Check for errors in each iteration
			if err != nil {
				log.Fatalf("client.Index() ERROR: %v", err)
			} else {
				fmt.Println("\nElasticsearch document indexed:", doc)
				fmt.Println("doc object TYPE:", reflect.TypeOf(doc))
			}

		} // end of doc iterator

		_, err = client.Flush(index).WaitIfOngoing(true).Do(ctx)
	} else {
		fmt.Println("client.Index() ERROR: the index", index, "does not exist")
	}
}
