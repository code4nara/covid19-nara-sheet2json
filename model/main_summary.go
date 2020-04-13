package model

import (
	"os"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

// MainSummary 入院者数
type MainSummary struct {
	Date     string  `json:"date"`
	Attr     string  `json:"attr"`
	Value    int     `json:"value"`
	Children []Child `json:"children"`
}

// Child KeyValue形式のデータ
type Child struct {
	Attr     string  `json:"attr"`
	Value    int     `json:"value"`
	Children []Child `json:"children,omitempty"`
}

func FetchMainSummary(svc *sheets.Service, spreadsheetID string) (*MainSummary, error) {
	var err error
	sheetRange := os.Getenv("COVID19_JSON2CSV_SHEET_RANGE_HOSPITALIZATION")
	call := svc.Spreadsheets.Values.Get(spreadsheetID, sheetRange)
	values, err := call.Do()
	if err != nil {
		return nil, err
	}
	ms := MainSummary{}
	row := values.Values[0]
	nums := []int{}
	for i, c := range row {
		if i == 0 {
			continue
		}
		if v, ok := c.(string); ok {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			nums = append(nums, val)
		}
	}
	ms.Date = time.Now().Format("2006/01/02 15:04")
	ms.Attr = "検査実施人数"
	ms.Value = 0
	ms.Children = []Child{
		{
			Attr:  "陽性患者数",
			Value: nums[0],
			Children: []Child{
				{
					Attr:  "入院患者数",
					Value: nums[1],
					Children: []Child{
						{
							Attr:  "症状のない方",
							Value: nums[3],
						},
						{
							Attr:  "症状のある方",
							Value: nums[2],
						},
					},
				},
			},
		},
		{
			Attr:  "退院した方",
			Value: nums[5],
		},
		{
			Attr:  "亡くなられた方",
			Value: nums[4],
		},
	}
	return &ms, nil
}
