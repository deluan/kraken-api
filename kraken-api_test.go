package krakenapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

var api = CreatePrivateApiClient()

type Config struct {
	Key    string
	Secret string
}

func LoadConfiguration(path string, config interface{}) (interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func CreatePrivateApiClient() *KrakenApi {
	var config Config

	_, err := LoadConfiguration("config.json", &config)
	if err != nil {
		panic(err)
	}

	return New(config.Key, config.Secret)
}

func TestApiServerTime(t *testing.T) {
	log.Println("TestApiServerTime...")
	resp, err := api.ApiServerTime()
	if err != nil {
		panic(err)
	}

	log.Println(resp)
}

func TestApiAssets(t *testing.T) {
	log.Println("TestApiAssets...")
	assets, err := api.ApiAssets()
	if err != nil {
		panic(err)
	}

	for k, v := range assets {
		log.Printf("%s: %v\n", k, v)
	}
}

func TestApiAssetPairs(t *testing.T) {
	log.Println("TestApiAssetPairs...")
	pairs, err := api.ApiAssetPairs("", "XXBTZEUR")
	if err != nil {
		panic(err)
	}

	for name, pair := range pairs {
		log.Println(name)
		log.Println(pair)
	}
}

func TestApiAssetPairsAll(t *testing.T) {
	log.Println("TestApiAssetPairs...")
	pairs, err := api.ApiAssetPairs("", "")
	if err != nil {
		panic(err)
	}

	for name, pair := range pairs {
		log.Println(name)
		log.Println(pair)
	}
}

func TestApiTicker(t *testing.T) {
	log.Println("TestApiTicker...")

	pairs := [...]string{"DASHEUR", "XXBTZEUR"}

	tickers, err := api.ApiTicker(pairs[:])
	if err != nil {
		panic(err)
	}

	for name, ticker := range tickers {
		log.Println(name)
		log.Println(ticker)
	}
}

func TestApiOHLC(t *testing.T) {
	log.Println("TestApiOHLC...")

	last, data, err := api.ApiOHLC("XXBTZEUR", 0, 0)
	if err != nil {
		panic(err)
	}

	log.Printf("last: %f\n", last)
	for unit := range data {
		log.Println(data[unit])
	}
}

func TestApiDepth(t *testing.T) {
	log.Println("TestApiDepth...")

	data, err := api.ApiDepth("XXBTZEUR", 0)
	if err != nil {
		panic(err)
	}

	for k, v := range data {
		log.Printf("%s:\n", k)

		log.Println(v.Asks)
		log.Println(v.Bids)
	}
}

func TestApiSpread(t *testing.T) {
	log.Println("TestApiSpread...")

	data, last, err := api.ApiSpread("XXBTZEUR", "")
	if err != nil {
		panic(err)
	}

	log.Println(data)
	log.Println(last)
}

func TestApiTrades(t *testing.T) {
	log.Println("TestApiTrades...")

	data, last, err := api.ApiTrades("XXBTZEUR", "")
	if err != nil {
		panic(err)
	}

	log.Println(data)
	log.Println(last)
}

func TestApiBalance(t *testing.T) {
	log.Println("TestApiBalance...")

	balance, err := api.ApiBalance()
	if err != nil {
		panic(err)
	}

	for name, balance := range balance {
		log.Printf("%s: %f\n", name, balance)
	}
}

func TestApiTradeBalance(t *testing.T) {
	log.Println("TestApiTradeBalance...")

	balance, err := api.ApiTradeBalance("ZEUR")
	if err != nil {
		panic(err)
	}

	log.Println(balance)
}

func TestApiTradesHistory(t *testing.T) {
	log.Println("TestApiTradesHistory...")

	trades, err := api.ApiTradesHistory("all", true, "", "", 0)
	if err != nil {
		panic(err)
	}

	for k, v := range trades {
		log.Printf("%s: %v\n", k, v)
	}
}

func TestApiQueryTrades(t *testing.T) {
	log.Println("TestApiQueryTrades...")

	txids := "TM5PQX-GLKZS-25MJUV"

	trades, err := api.ApiQueryTrades(txids, false)
	if err != nil {
		panic(err)
	}

	for k, v := range trades {
		log.Printf("%s: %v\n", k, v)
	}
}

func TestApiOpenPositions(t *testing.T) {
	log.Println("TestApiOpenPositions...")

	positions, err := api.ApiOpenPositions("", true)
	if err != nil {
		panic(err)
	}

	for k, pos := range positions {
		log.Printf("%s: status: %s value: %5.3f net: %s\n", k, pos.Posstatus, pos.Value, pos.Net)
	}
}

func TestApiLedgers(t *testing.T) {
	log.Println("TestApiLedgers...")

	ledgers, err := api.ApiLedgers("", "", "", "", 0)
	if err != nil {
		panic(err)
	}

	for k, v := range ledgers {
		log.Printf("%s: %v\n", k, v)
	}
}

func TestApiQueryLedgers(t *testing.T) {
	log.Println("TestApiQueryLedgers...")

	ledgers, err := api.ApiQueryLedgers("LL4UG5-DMFOH-SNGNT6")
	if err != nil {
		panic(err)
	}

	for k, v := range ledgers {
		log.Printf("%s: %v\n", k, v)
	}
}

func TestApiTradeVolume(t *testing.T) {
	log.Println("TestApiTradeVolume...")

	info, err := api.ApiTradeVolume("XXBTZEUR", true)
	if err != nil {
		panic(err)
	}

	log.Println(info)
}

func TestApiAddOrder(t *testing.T) {
	log.Println("TestApiAddOrder...")

	order, err := api.ApiAddOrder(
		"XXBTZEUR", // pair
		"buy",      // buy/sell
		"limit",    // ordertype
		1,          // price
		0,          // price2
		0.1,        // volume
		"")
	if err != nil {
		panic(err)
	}

	log.Println(order)
	log.Println(order.Txid)

	log.Println("TestCancelOrder...")

	for _, txid := range order.Txid {
		cancel_result, err := api.ApiCancelOrder(txid)
		if err != nil {
			panic(err)
		}
		log.Println(cancel_result)
	}
}

func DumpOrder(order_id string, order *Order) {
	log.Printf("%s: refid%s userref%d status:%s descr:%s opentm:%f closetm:%f\n",
		order_id,
		order.RefId,
		order.Userref,
		order.Status,
		order.Descr.Order,
		order.Opentm,
		order.Closetm,
	)
}

func TestApiOpenOrders(t *testing.T) {
	log.Println("TestApiOpenOrder...")

	orders, err := api.ApiOpenOrders(false, "")
	if err != nil {
		panic(err)
	}

	for k, v := range orders.Open {
		DumpOrder(k, &v)
	}
}

func TestApiClosedOrders(t *testing.T) {
	log.Println("TestApiClosedOrders...")

	orders, err := api.ApiClosedOrders(false, "", "", "", 0, "both")
	if err != nil {
		panic(err)
	}

	for k, v := range orders.Closed {
		DumpOrder(k, &v)
	}
}

func TestApiQueryOrders(t *testing.T) {
	log.Println("TestApiQueryOrders...")

	txids := "OBS47P-P5LV7-WAE3OV,OMDL5O-AEILS-MYEXLF"

	orders, err := api.ApiQueryOrders(false, "", txids)
	if err != nil {
		panic(err)
	}

	for k, v := range *orders {
		DumpOrder(k, &v)
	}
}
