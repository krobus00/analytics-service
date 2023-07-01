package constant

type ctxKey string

const (
	KeyDBCtx          ctxKey = "DB"
	KeyUserIDCtx      ctxKey = "USERID"
	KeyDataSourceCtx  ctxKey = "DATA_SOURCE"
	KeyUserIPCtx      ctxKey = "USERIP"
	KeyUserCountryCtx ctxKey = "USERCOUNTRY"
	KeyUserCityCtx    ctxKey = "USERCity"

	SystemID = string("SYSTEM")
	GuestID  = string("GUEST")
)

const (
	SourceDB int = iota
	SourceOS
	SourceRedis
)
