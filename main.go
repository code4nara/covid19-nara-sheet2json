package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/miiton/covid19-nara-sheet2json/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func getClient() (*http.Client, error) {
	data, err := ioutil.ReadFile("./tmp/credentials.json")
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext), nil
}

// data.json を生成
func genData(svc *sheets.Service, spreadsheetID string) {
	var data model.Covid19Data
	patients, err := model.FetchPatients(svc, spreadsheetID)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	mainSummary, err := model.FetchMainSummary(svc, spreadsheetID)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	data.Patients = *patients
	patientsSummary, err := data.Patients.GenSummary()
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	data.PatientsSummary = *patientsSummary
	data.MainSummary = *mainSummary
	j, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(os.Getenv("COVID19_JSON2CSV_OUTPUT_DATA"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(j)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
}

// sickbeds_summary.json を生成
func genSickbeds(svc *sheets.Service, spreadsheetID string) {
	data, err := model.FetchSickbedsSummary(svc, spreadsheetID)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	j, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(os.Getenv("COVID19_JSON2CSV_OUTPUT_SICKBEDS_SUMMARY"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(j)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

}

// news.json を生成
func genNews(svc *sheets.Service, spreadsheetID string) {
	data, err := model.FetchNews(svc, spreadsheetID)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	j, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(os.Getenv("COVID19_JSON2CSV_OUTPUT_NEWS"), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(j)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	spreadsheetID := os.Getenv("COVID19_JSON2CSV_SHEET_ID")
	client, err := getClient()
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	svc, err := sheets.New(client)
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
	genData(svc, spreadsheetID)
	genSickbeds(svc, spreadsheetID)
	genNews(svc, spreadsheetID)
}
