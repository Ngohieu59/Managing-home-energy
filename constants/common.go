package constants

const (
	TenantID = "8bcd4156-69bc-11ed-bdc7-62d6209fc93a"
	Taxt     = 8.0
)
const (
	ClaimUserId         = "user_id"
	ClaimUserUUID       = "user_uuid"
	ClaimUsername       = "user_name"
	ClaimPermission     = "user_permission"
	RequestID           = "request_id"
	RequestErrorMessage = "error_message"
)

type UnitRecord struct {
	UnitPrice float64
	Quantity  float64
}

var (
	UnitLevel1 = UnitRecord{UnitPrice: 1984.0, Quantity: 52.0}
	UnitLevel2 = UnitRecord{UnitPrice: 2050.0, Quantity: 52.0}
	UnitLevel3 = UnitRecord{UnitPrice: 2380.0, Quantity: 103.0}
	UnitLevel4 = UnitRecord{UnitPrice: 2998.0, Quantity: 103.0}
	UnitLevel5 = UnitRecord{UnitPrice: 3350.0, Quantity: 103.0}
	UnitLevel6 = UnitRecord{UnitPrice: 3460.0, Quantity: 0.0}
)
