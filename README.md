# gogst (Golang global stock time)

This lib will tell you what is current market status of your stock.

## Example
```
import "github.com/bandprotocol/gogst/markets"

...

status, _ := markets.GetMarketStatus("AAPL")
if status == markets.PRE || status == markets.POST {
    fmt.Println("Extended time")
} else {
    fmt.Printf("%s time\n", status)
}
```
