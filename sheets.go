package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config, tokenCacheFilename string) *http.Client {
	cacheFile, err := tokenCacheFile(tokenCacheFilename)
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile(filename string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, url.QueryEscape(filename)), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// SheetsSetup sets up the sheet service for the requested sheet ID
func SheetsSetup(id string) (*sheets.Service, int) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config, "sheets.googleapis.com-jp.json")

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	masterSheet := getMasterSheet(srv, id)
	numSheets := len(masterSheet.Sheets)

	return srv, numSheets
}

func getMasterSheet(srv *sheets.Service, id string) *sheets.Spreadsheet {

	masterSheet, err := srv.Spreadsheets.Get(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	return masterSheet
}

// GetDataForSheetIndex returns a range with all data for the requested sheet
func GetDataForSheetIndex(srv *sheets.Service, id string, sheetIndex int, firstCell string, lastColumn string) *sheets.ValueRange {

	sheet := getMasterSheet(srv, id)
	worksheet := sheet.Sheets[sheetIndex]
	readRange := worksheet.Properties.Title + "!" + firstCell + ":" + lastColumn //A2:E"

	resp, err := srv.Spreadsheets.Values.Get(id, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	return resp
}

// GetSheetName returns the name of a worksheet
func GetSheetName(srv *sheets.Service, id string, sheetIndex int) string {

	sheet := getMasterSheet(srv, id)
	worksheet := sheet.Sheets[sheetIndex]
	return worksheet.Properties.Title
}

// SheetsWriteURLData writes all URL data to a sheet
func SheetsWriteURLData(srv *sheets.Service, id string, sheetIndex int, values *sheets.ValueRange) {

	sheet := getMasterSheet(srv, id)
	worksheet := sheet.Sheets[sheetIndex]
	writeRange := worksheet.Properties.Title + "!C3"

	fmt.Print("Writing data to sheet " + worksheet.Properties.Title + "...")

	_, err := srv.Spreadsheets.Values.Update(id, writeRange, values).ValueInputOption("RAW").Do()

	if err != nil {
		log.Fatalf("\nUnable to write data to sheet. %v", err)
	} else {
		fmt.Print(" done")
	}
}
