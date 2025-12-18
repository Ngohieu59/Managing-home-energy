package dto

type StaffReportListReq struct {
	StartDate string  `json:"StartDate"`
	EndDate   string  `json:"EndDate"`
	Type      string  `json:"Type"`
	City      string  `json:"City"`
	Ward      string  `json:"Ward"`
	Threshold float64 `json:"Threshold"`
}

type ReportEUseReq struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Type      string `json:"Type"`
	City      string `json:"City"`
	Ward      string `json:"Ward"`
}

type UserReport struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Money     float64 `json:"money"`
	ElectUsed float64 `json:"electused"`
}

type Vol struct {
	VolLevel  string  `json:"vollevel"`
	ElectUsed float64 `json:"electused"`
}

type TotalLevel struct {
	VolLevel string `json:"vol-level"`
	NumUsers int    `json:"num-users"`
}
type FamilyResp struct {
	TotalMember []*TotalLevel `json:"TotalMember"`
}

type VolResp struct {
	Level []*Vol `json:"level"`
}

type StaffReportListResp struct {
	Filter   *StaffReportListReq `json:"Filter"`
	ListUser []*UserReport       `json:"ListUser"`
}

type ReportEUseResp struct {
	Filter *ReportEUseReq `json:"Filter"`
	Data   any            `json:"SumInfo"`
}
