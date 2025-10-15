package dto

type EBillMoneyReq struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
}

type EbillMoneyResp struct {
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
