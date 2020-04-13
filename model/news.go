package model

import (
	"os"

	"google.golang.org/api/sheets/v4"
)

// News 新着情報
type News struct {
	NewsItems []NewsItem `json:"newsItems"`
}

// NewsItem 新着情報単体
type NewsItem struct {
	Date string `json:"date"`
	URL  string `json:"url"`
	Text string `json:"text"`
}

func FetchNews(svc *sheets.Service, spreadsheetID string) (*News, error) {
	var err error
	sheetRange := os.Getenv("COVID19_JSON2CSV_SHEET_RANGE_NEWS")
	call := svc.Spreadsheets.Values.Get(spreadsheetID, sheetRange)
	values, err := call.Do()
	if err != nil {
		return nil, err
	}
	news := News{}
	for _, v := range values.Values {
		n := NewsItem{}
		if val, ok := v[0].(string); ok {
			n.Date = val
		}
		if val, ok := v[1].(string); ok {
			n.URL = val
		}
		if val, ok := v[2].(string); ok {
			n.Text = val
		}
		news.NewsItems = append([]NewsItem{n}, news.NewsItems...)
	}

	return &news, nil
}
