package model

import (
	"os"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

// SickbedsSummary 入院患者数と残り病床数
type SickbedsSummary struct {
	Data struct {
		Patients      int `json:"入院患者数"`
		RemainingBeds int `json:"残り病床数"`
	} `json:"data"`
	LastUpdate string `json:"last_update"`
}

func FetchSickbedsSummary(svc *sheets.Service, spreadsheetID string) (*SickbedsSummary, error) {
	var err error
	sheetRange := os.Getenv("COVID19_JSON2CSV_SHEET_RANGE_HOSPITALIZATION")
	call := svc.Spreadsheets.Values.Get(spreadsheetID, sheetRange)
	values, err := call.Do()
	if err != nil {
		return nil, err
	}
	ss := SickbedsSummary{}
	ss.LastUpdate = time.Now().Format("2006/01/02 15:04")
	for i, v := range values.Values[0] {
		if i == 2 {
			if val, ok := v.(string); ok {
				ss.Data.Patients, err = strconv.Atoi(val)
				if err != nil {
					return nil, err
				}

			}
			continue
		}
		if i == 8 {
			if val, ok := v.(string); ok {
				ss.Data.RemainingBeds, err = strconv.Atoi(val)
				if err != nil {
					return nil, err
				}

			}
			continue
		}

	}
	return &ss, nil
}
