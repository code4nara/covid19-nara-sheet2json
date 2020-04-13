package model

import "time"

type PatientsSummary struct {
	Date string                `json:"date"`
	Data []PatientsSummaryData `json:"data"`
}

type PatientsSummaryData struct {
	Date  time.Time `json:"日付"`
	Value int       `json:"小計"`
}
