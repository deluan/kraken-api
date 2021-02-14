// Provides a wrapper to query kraken api endpoints
package krakenapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	URL_ROOT = "https://api.kraken.com"

	URL_PUBLIC_TIME          = "/0/public/Time"
	URL_PUBLIC_ASSETS        = "/0/public/Assets"
	URL_PUBLIC_ASSET_PAIRS   = "/0/public/AssetPairs"
	URL_PUBLIC_TICKER        = "/0/public/Ticker"
	URL_PUBLIC_OHLC          = "/0/public/OHLC"
	URL_PUBLIC_ORDER_BOOK    = "/0/public/Depth"
	URL_PUBLIC_RECENT_TRADES = "/0/public/Trades"
	URL_PUBLIC_SPREAD        = "/0/public/Spread"

	URL_PRIVATE_BALANCE         = "/0/private/Balance"
	URL_PRIVATE_TRADE_BALANCE   = "/0/private/TradeBalance"
	URL_PRIVATE_OPEN_ORDERS     = "/0/private/OpenOrders"
	URL_PRIVATE_CLOSED_ORDERS   = "/0/private/ClosedOrders"
	URL_PRIVATE_QUERY_ORDERS    = "/0/private/QueryOrders"
	URL_PRIVATE_TRADES_HISTORY  = "/0/private/TradesHistory"
	URL_PRIVATE_QUERY_TRADES    = "/0/private/QueryTrades"
	URL_PRIVATE_OPEN_POSITIONS  = "/0/private/OpenPositions"
	URL_PRIVATE_LEDGERS         = "/0/private/Ledgers"
	URL_PRIVATE_QUERY_LEDGERS   = "/0/private/QueryLedgers"
	URL_PRIVATE_TRADE_VOLUME    = "/0/private/TradeVolume"
	URL_PRIVATE_ADD_ORDER       = "/0/private/AddOrder"
	URL_PRIVATE_CANCEL_ORDER    = "/0/private/CancelOrder"
	URL_PRIVATE_WITHDRAW        = "/0/private/Withdraw"
	URL_PRIVATE_WITHDRAW_STATUS = "/0/private/WithdrawStatus"
)

type KrakenApi struct {
	Key       string
	secret    string
	ApiRoot   string
	UserAgent string
	Client    *http.Client
}

// Create a new KrakenApi client
// Returns a pointer to KrakenApi
func New(key string, secret string) *KrakenApi {
	client := &http.Client{}
	user_agent := "kraken-api"

	return &KrakenApi{key, secret, URL_ROOT, user_agent, client}
}

func parse(resp []byte, struct_type interface{}) (interface{}, error) {
	var response KrakenResponse

	if struct_type != nil {
		response.Result = struct_type
	}

	err := json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request! (%s)", response.Error)
	}

	return response.Result, nil
}

func (api *KrakenApi) Query(url_path string, params url.Values, with_signature bool) ([]byte, error) {
	headers := map[string]string{}
	method := "GET"

	if with_signature {
		secret, _ := base64.StdEncoding.DecodeString(api.secret)
		params.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

		signature := createKrakenSignature(url_path, params, secret)

		headers["API-Key"] = api.Key
		headers["API-Sign"] = signature

		method = "POST"
	}

	if api.UserAgent != "" {
		headers["User-Agent"] = api.UserAgent
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"

	return executeHttpQuery(method, api.ApiRoot+url_path, headers, params)
}
