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

func (m MarketStatus) String() string {
	switch m {
	case INVALID:
		return "invalid"
	case CLOSE:
		return "close"
	case PRE:
		return "pre"
	case REG:
		return "reg"
	case POST:
		return "post"
	}
	return "unknown"
}
