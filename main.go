package main

import (
	"context"
	"fmt"
	awesomev1 "github.com/abitofhelp/awesome/client/client_go"
)

func main() {

	if client, err := awesomev1.NewAwesomeServiceClient("localhost", 20580); err == nil {
		if report, err := client.FindReportByPetName(context.Background(), "Lassie"); err == nil {
			fmt.Printf("\nFound a report for '%s'\n", report.Report.Pet.Name)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}
