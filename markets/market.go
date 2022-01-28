package markets

type Market int64

const (
	UNDEFINED Market = iota
	US_STOCK
	CA_STOCK
)

type MarketStatus int64

const (
	INVALID MarketStatus = iota
	CLOSE
	PRE
	REG
	POST
)
