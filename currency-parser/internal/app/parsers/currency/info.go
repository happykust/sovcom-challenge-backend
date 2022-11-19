package currency

var TickersGroupName = "currency"

var TickersGroupID uint

// 86400*364*2 is a about 2 years
var UnixTimeStartHistory int64

// RUBUSR StartUnixTime EndUnixTime
var CurrencyURL = "https://query1.finance.yahoo.com/v8/finance/chart/%s=X?symbol=RUB%3DX&period1=%s&period2=%s&useYfid=true&interval=1h&includePrePost=true&events=div%7Csplit%7Cearn&lang=en-US&region=US&crumb=undefined&corsDomain=finance.yahoo.com"
