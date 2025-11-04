package constants

const (
	TenantID = "8bcd4156-69bc-11ed-bdc7-62d6209fc93a"
	Taxt     = 8.0
)

const (
	ClaimUserId         = "user_id"
	ClaimUserUUID       = "user_uuid"
	ClaimUsername       = "user_name"
	ClaimUserType       = "user_type"
	ClaimPermission     = "user_permission"
	RequestID           = "request_id"
	RequestErrorMessage = "error_message"
)

const (
	TypeFamily         = "family"
	TypeBusiness       = "business"
	TypeIndustrial     = "industrial"
	TypeAdministrative = "administrative"
	VoltageLow         = "low"
	VoltageHigh        = "high"
	VoltageMedium      = "medium"
)

type UnitRecord struct {
	UnitPrice    float64
	Quantity     float64
	VoltageLevel string
}

type Hour struct {
	Start int
	End   int
}

type TimeRange struct {
	Start int
	End   int
}

type Tariff struct {
	Low    []TimeRange
	Normal []TimeRange
	High   []TimeRange
}

var (
	TariffHours = Tariff{
		Low: []TimeRange{
			{Start: 22, End: 24},
			{Start: 0, End: 4},
		},
		Normal: []TimeRange{
			{Start: 4, End: 10},
			{Start: 12, End: 17},
			{Start: 20, End: 22},
		},
		High: []TimeRange{
			{Start: 10, End: 12},
			{Start: 17, End: 20},
		},
	}

	UnitFamilyLevel1 = UnitRecord{UnitPrice: 1984.0, Quantity: 50.0}
	UnitFamilyLevel2 = UnitRecord{UnitPrice: 2050.0, Quantity: 50.0}
	UnitFamilyLevel3 = UnitRecord{UnitPrice: 2380.0, Quantity: 100.0}
	UnitFamilyLevel4 = UnitRecord{UnitPrice: 2998.0, Quantity: 100.0}
	UnitFamilyLevel5 = UnitRecord{UnitPrice: 3350.0, Quantity: 100.0}
	UnitFamilyLevel6 = UnitRecord{UnitPrice: 3460.0, Quantity: 0.0}

	UnitBusinessLowLevel1    = UnitRecord{UnitPrice: 1830, Quantity: 6000, VoltageLevel: "low"}
	UnitBusinessLowLevel2    = UnitRecord{UnitPrice: 3007, Quantity: 6000, VoltageLevel: "medium"}
	UnitBusinessLowLevel3    = UnitRecord{UnitPrice: 5174, Quantity: 6000, VoltageLevel: "high"}
	UnitBusinessHighLevel1   = UnitRecord{UnitPrice: 1535, Quantity: 22000, VoltageLevel: "low"}
	UnitBusinessHighLevel2   = UnitRecord{UnitPrice: 2755, Quantity: 22000, VoltageLevel: "medium"}
	UnitBusinessHighLevel3   = UnitRecord{UnitPrice: 4795, Quantity: 22000, VoltageLevel: "high"}
	UnitBusinessMediumLevel1 = UnitRecord{UnitPrice: 1746, Quantity: 22000, VoltageLevel: "low"}
	UnitBusinessMediumLevel2 = UnitRecord{UnitPrice: 2965, Quantity: 22000, VoltageLevel: "medium"}
	UnitBusinessMediumLevel3 = UnitRecord{UnitPrice: 4963, Quantity: 22000, VoltageLevel: "high"}

	UnitIndustrialLowLevel1    = UnitRecord{UnitPrice: 1241, Quantity: 6000, VoltageLevel: "low"}
	UnitIndustrialLowLevel2    = UnitRecord{UnitPrice: 1896, Quantity: 6000, VoltageLevel: "medium"}
	UnitIndustrialLowLevel3    = UnitRecord{UnitPrice: 3474, Quantity: 6000, VoltageLevel: "high"}
	UnitIndustrialHighLevel1   = UnitRecord{UnitPrice: 1136, Quantity: 110000, VoltageLevel: "low"}
	UnitIndustrialHighLevel2   = UnitRecord{UnitPrice: 1749, Quantity: 110000, VoltageLevel: "medium"}
	UnitIndustrialHighLevel3   = UnitRecord{UnitPrice: 3242, Quantity: 110000, VoltageLevel: "high"}
	UnitIndustrialMediumLevel1 = UnitRecord{UnitPrice: 1178, Quantity: 22000, VoltageLevel: "low"}
	UnitIndustrialMediumLevel2 = UnitRecord{UnitPrice: 1812, Quantity: 22000, VoltageLevel: "medium"}
	UnitIndustrialMediumLevel3 = UnitRecord{UnitPrice: 3348, Quantity: 22000, VoltageLevel: "high"}
	UnitIndustrialMaxLevel1    = UnitRecord{UnitPrice: 1094, Quantity: 110000, VoltageLevel: "low"}
	UnitIndustrialMaxLevel2    = UnitRecord{UnitPrice: 1728, Quantity: 110000, VoltageLevel: "medium"}
	UnitIndustrialMaxLevel3    = UnitRecord{UnitPrice: 3116, Quantity: 110000, VoltageLevel: "high"}

	UnitAdministrativeLowLevel1  = UnitRecord{UnitPrice: 1851, Quantity: 6000, VoltageLevel: "low"}
	UnitAdministrativeHighLevel1 = UnitRecord{UnitPrice: 1977, Quantity: 6000, VoltageLevel: "high"}
)
