package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	rows := make([][]interface{}, 0)
	// Adding the first line of the sheet
	rows = append(rows, []interface{}{"First Name", "Last Name", "Age"})

	age := "19"
	first_name := "sergio"
	last_name := "ramos"

	// Append the data to add to the sheet
	rows = append(rows, []interface{}{first_name, last_name, age})

	// ref link for getting token - https://www.prudentdevs.club/gsheets-go
	b, err := ioutil.ReadFile("./secret.json")
	if err != nil {
		fmt.Errorf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		fmt.Errorf("error in parsing secret: %v", err)
	}

	client := config.Client(context.Background())

	// if your sheet url is:https://docs.google.com/spreadsheets/d/1890/edit#gid=0 , id is 1890
	spreadsheetId := "1890"

	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		fmt.Errorf("error in retrieving sheets client: %v", err)
	}

	// Modify this to your Needs
	rangeData := "sheet1!A1:E"

	rs, err := srv.Spreadsheets.Values.Append(spreadsheetId, rangeData, &sheets.ValueRange{
		Values: rows,
	}).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs.HTTPStatusCode)
}
