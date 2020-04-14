package model

import (
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

// Patients 陽性患者の属性（一覧）
type Patients struct {
	Date string    `json:"date"`
	Data []Patient `json:"data"`
}

type Patient struct {
	No               int       `json:"No"`
	AnnouncementDate time.Time `json:"発表日"`
	Residence        string    `json:"居住地"`
	Age              string    `json:"年代"`
	Gender           string    `json:"性別"`
	Note             string    `json:"備考"`
}

func FetchPatients(svc *sheets.Service, spreadsheetID string) (*Patients, error) {
	var err error
	sheetRange := os.Getenv("COVID19_JSON2CSV_SHEET_RANGE_PATIENTS")
	call := svc.Spreadsheets.Values.Get(spreadsheetID, sheetRange)
	values, err := call.Do()
	if err != nil {
		return nil, err
	}
	ps := Patients{}
	for _, v := range values.Values {
		ps.Date = time.Now().Format("2006/01/02 15:04")
		var p Patient
		if val, ok := v[0].(string); ok {
			p.No, err = strconv.Atoi(val)
			if err != nil {
				log.Println(err)
				break
			}
		}
		if val, ok := v[3].(string); ok {
			if val == "" {
				break
			}
			p.Residence = val
		}
		if val, ok := v[4].(string); ok {
			p.AnnouncementDate, err = time.Parse("2006-01-02", val)
			if err != nil {
				log.Println(err)
				break
			}
		}
		if val, ok := v[7].(string); ok {
			p.Age = val
		}
		if val, ok := v[8].(string); ok {
			p.Gender = val
		}
		ps.Data = append(ps.Data, p)
	}
	return &ps, err
}

func setSummaryData(cur time.Time, kv map[string]int) int {
	for v := range kv {
		if cur.Format("2006-01-02") == v {
			return kv[v]
		}
	}
	return 0
}

func (p *Patients) GenSummary() (*PatientsSummary, error) {
	kv := map[string]int{}
	for _, d := range p.Data {
		kv[d.AnnouncementDate.Format("2006-01-02")]++
	}
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	cur := time.Date(2020, 1, 24, 0, 0, 0, 0, loc)
	now := time.Now()
	result := PatientsSummary{}
	result.Date = time.Now().Format("2006/01/02 15:04")
	for {
		var d PatientsSummaryData
		d.Date = cur
		if cur.After(now) {
			break
		}
		d.Value = setSummaryData(cur, kv)
		result.Data = append(result.Data, d)
		cur = cur.AddDate(0, 0, 1)
	}

	return &result, nil
}
