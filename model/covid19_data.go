package model

type Covid19Data struct {
	Patients        Patients        `json:"patients"`
	PatientsSummary PatientsSummary `json:"patients_summary"`
	MainSummary     MainSummary     `json:"main_summary"`
	LastUpdate      string          `json:"lastUpdate"`
}
