package dto

type EBillMoneyReq struct {
	StartDate string  `json:"StartDate"`
	EndDate   string  `json:"EndDate"`
	Electric  float64 `json:"Electric"`
}

type EBillMoneyResp struct {
	StartDate string  `json:"start"`
	EndDate   string  `json:"end"`
	Money     float64 `json:"money"`
	ElectUsed float64 `json:"electused"`
}

type ReportMonthly struct {
	Month     int     `json:"month"`
	Year      int     `json:"year"`
	ElectUsed float64 `json:"electused"`
}

type ReportMonthlyResp struct {
	ThisYear *ReportMonthly `json:"thisyear"`
	LastYear *ReportMonthly `json:"lastyear"`
}

type ListEUsedResponse struct {
	Total  float64
	Normal float64
	Low    float64
	High   float64
}
