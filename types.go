package krakenapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

type Asset struct {
	Altname         string `json:"altname"`          // alternate name
	Aclass          string `json:"aclass"`           // asset class
	Decimals        int    `json:"decimals"`         // scaling decimal places for record keeping
	DisplayDecimals int    `json:"display_decimals"` // scaling decimal places for output display
}

type AssetPair struct {
	Altname           string      `json:"altname"`             // alternate pair name
	AclassBase        string      `json:"aclass_base"`         // asset class of base component
	Base              string      `json:"base"`                // asset id of base component
	AclassQuote       string      `json:"aclass_quote"`        // asset class of quote component
	Quote             string      `json:"quote"`               // asset id of quote component
	Lot               string      `json:"lot"`                 // volume lot size
	PairDecimals      int         `json:"pair_decimals"`       // scaling decimal places for pair
	LotDecimals       int         `json:"lot_decimals"`        // scaling decimal places for volume
	LotMultiplier     int         `json:"lot_multiplier"`      // amount to multiply lot volume by to get currency volume
	LeverageBuy       []int       `json:"leverage_buy"`        // array of leverage amounts available when buying
	LeverageSell      []int       `json:"leverage_sell"`       // array of leverage amounts available when selling
	Fees              [][]float64 `json:"fees"`                // fee schedule array in [volume, percent fee] tuples
	FeesMaker         [][]float64 `json:"fees_maker"`          // maker fee schedule array in [volume, percent fee] tuples (if on maker/taker)
	FeeVolumeCurrency string      `json:"fee_volume_currency"` // volume discount currency
	MarginCall        int         `json:"margin_call"`         // margin call level
	MarginStop        int         `json:"margin_stop"`         // stop-out/liquidation margin level
}

type TradeToday struct {
	Price  float64
	Volume float64
}

type RecentTrade struct {
	Price     float64
	Volume    float64
	Time      float64
	Type      string
	TradeType string
	Misc      string
}

type AskBid struct {
	Price          float64
	WholeLotVolume float64
	LotVolume      float64
}

type TodayH24Float64 [2]float64

type Ticker struct {
	Ask          AskBid          `json:"a"`        // ask array(<price>, <whole lot volume>, <lot volume>)
	Bid          AskBid          `json:"b"`        // bid array(<price>, <whole lot volume>, <lot volume>)
	LastTrade    TradeToday      `json:"c"`        // last trade closed array(<price>, <lot volume>)
	VolumeArray  TodayH24Float64 `json:"v"`        // volume array(<today>, <last 24 hours>)
	VWAP         TodayH24Float64 `json:"p"`        // volume weighted average price array(<today>, <last 24 hours>)
	Trades       [2]int          `json:"t"`        // number of trades array(<today>, <last 24 hours>)
	Low          TodayH24Float64 `json:"l"`        // low array(<today>, <last 24 hours>)
	High         TodayH24Float64 `json:"h"`        // high array(<today>, <last 24 hours>)
	OpeningPrice float64         `json:"o,string"` // today's opening price
}

type OHLCEntry struct {
	Time   float64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	VWAP   float64
	Volume float64
	Count  float64
}

type Trade struct {
	Ordertxid string  `json:"ordertxid"`     // order responsible for execution of trade
	Pair      string  `json:"pair"`          // asset pair
	Time      float64 `json:"time"`          // unix timestamp of trade
	Type      string  `json:"type"`          // type of order (buy/sell)
	Ordertype string  `json:"ordertype"`     // order type
	Price     float64 `json:"price,string"`  // average price order was executed at (quote currency)
	Cost      float64 `json:"cost,string"`   // total cost of order (quote currency)
	Fee       float64 `json:"fee,string"`    // total fee (quote currency)
	Vol       float64 `json:"vol,string"`    // volume (base currency)
	Margin    float64 `json:"margin,string"` //initial margin (quote currency)
	Misc      string  `json:"misc"`          // comma delimited list of miscellaneous info

	Posstatus string   `json:"posstatus"`      // position status (open/closed)
	Cprice    float64  `json:"cprice,string"`  // average price of closed portion of position (quote currency)
	Ccost     float64  `json:"ccost,string"`   // total cost of closed portion of position (quote currency)
	Cfee      float64  `json:"cfee,string"`    // total fee of closed portion of position (quote currency)
	Cvol      float64  `json:"cvol,string"`    // total fee of closed portion of position (quote currency)
	Cmargin   float64  `json:"cmargin,string"` // total margin freed in closed portion of position (quote currency)
	Net       float64  `json:"net,string"`     // net profit/loss of closed portion of position (quote currency, quote currency scale)
	Trades    []string `json:"trades"`         // list of closing trades for position (if available)
}

type PublicOrder struct {
	Price  float64
	Volume float64
	Time   float64
}

type PublicOrderBook struct {
	Asks []PublicOrder `json:"asks"`
	Bids []PublicOrder `json:"bids"`
}

type Spread struct {
	Time float64
	Bid  float64
	Ask  float64
}

type OrderResult struct {
	Descr struct {
		Order string
	}
	Txid []string
}

type CancelResult struct {
	Count     int
	IsPending bool
}

type TradeBalance struct {
	Eb float64 `json:"eb,string"` // equivalent balance (combined balance of all currencies)
	Tb float64 `json:"tb,string"` // trade balance (combined balance of all equity currencies)
	M  float64 `json:"m,string"`  // margin amount of open positions
	N  float64 `json:"n,string"`  // unrealized net profit/loss of open positions
	C  float64 `json:"c,string"`  // cost basis of open positions
	V  float64 `json:"v,string"`  // current floating valuation of open positions
	E  float64 `json:"e,string"`  // equity = trade balance + unrealized net profit/loss
	Mf float64 `json:"mf,string"` // free margin = equity - initial margin (maximum margin available to open new positions)
	Ml float64 `json:"ml,string"` // margin level = (equity / initial margin) * 100
}

type Order struct {
	RefId      string     `json:"refid"`             // Referral order transaction id that created this order
	Userref    int64      `json:"userref"`           // user reference id
	Status     string     `json:"status"`            // status of order: pending / open / closed / canceled / expired
	Opentm     float64    `json:"opentm"`            // unix timestamp of when order was placed
	Starttm    float64    `json:"starttm"`           // unix timestamp of order start time (or 0 if not set)
	Expiretm   float64    `json:"expiretm"`          // unix timestamp of order end time (or 0 if not set)
	Descr      OrderDescr `json:"descr"`             // order description info
	Vol        float64    `json:"vol,string"`        // volume of order (base currency unless viqc set in oflags)
	VolExec    float64    `json:"vol_exec,string"`   // volume executed (base currency unless viqc set in oflags)
	Cost       float64    `json:"cost,string"`       // total cost (quote currency unless unless viqc set in oflags)
	Fee        float64    `json:"fee,string"`        // total fee (quote currency)
	Price      float64    `json:"price,string"`      // average price (quote currency unless viqc set in oflags)
	Stopprice  float64    `json:"stopprice,string"`  // stop price (quote currency, for trailing stops)
	Limitprice float64    `json:"limitprice,string"` // triggered limit price (quote currency, when limit based order type triggered)
	Misc       string     `json:"misc"`              // comma delimited list of miscellaneous info (stopped, touched, liquidated, partial)
	Oflags     string     `json:"oflags"`            // comma delimited list of order flags (viqc, fcib, fciq, nompp)
	Trades     []string   `json:"trades"`            // array of trade ids related to order (if trades info requested and data available)
	Closetm    float64    `json:"closetm"`           // unix timestamp of when order was closed
	Reason     string     `json:"reason"`            // Closed orders: additional info on status (if any)
}

type OrderDescr struct {
	Pair      string  `json:"pair"`          // asset pair
	Type      string  `json:"type"`          // type of order (buy/sell)
	Ordertype string  `json:"ordertype"`     // order type (market/limit/...)
	Price     float64 `json:"price,string"`  // primary price
	Price2    float64 `json:"price2,string"` // secondary price
	Leverage  string  `json:"leverage"`      // amount of leverage (can be "none")
	Order     string  `json:"order"`         // order description
	Close     string  `json:"close"`         // conditional close order description (if conditional close set)
}

type OpenOrders struct {
	Open map[string]Order
}

type ClosedOrders struct {
	Closed map[string]Order
	Count  int
}

type QueryOrder map[string]Order

type TradeHistoryResult struct {
	Trades map[string]Trade
	Count  int
}

type OpenPosition struct {
	Ordertxid  string  `json:"ordertxid"`         // order responsible for execution of trade
	Posstatus  string  `json:"posstatus"`         // position status
	Pair       string  `json:"pair"`              // asset pair
	Time       float64 `json:"time"`              // unix timestamp of trade
	Type       string  `json:"type"`              // type of order used to open position (buy/sell)
	Ordertype  string  `json:"ordertype"`         // order type used to open position
	Cost       float64 `json:"cost,string"`       // opening cost of position (quote currency unless viqc set in oflags)
	Fee        float64 `json:"fee,string"`        // opening fee of position (quote currency)
	Vol        float64 `json:"vol,string"`        // position volume (base currency unless viqc set in oflags)
	VolClosed  float64 `json:"vol_closed,string"` // position volume closed (base currency unless viqc set in oflags)
	Margin     float64 `json:"margin,string"`     // initial margin (quote currency)
	Value      float64 `json:"value,string"`      // current value of remaining position (if docalcs requested. quote currency)
	Net        string  `json:"net"`               // unrealized profit/loss of remaining position (if docalcs requested. quote currency, quote currency scale)
	Misc       string  `json:"misc"`              // comma delimited list of miscellaneous info
	Terms      string  `json:"terms"`             // terms
	Oflags     string  `json:"oflags"`            // comma delimited list of order flags / viqc = volume in quote currency
	Rollovertm float64 `json:"rollovertm,string"`
}

type Ledger struct {
	Refid   string  `json:"refid"` // reference id
	Time    float64 `json:"time"`  // unx timestamp of ledger
	Type    string  `json:"type"`
	Aclass  string  `json:"aclass"`
	Asset   string  `json:"asset"`
	Amount  float64 `json:",string"` // transaction amount
	Fee     float64 `json:",string"` // transaction fee
	Balance float64 `json:",string"` // balance
}

type LedgerResponse struct {
	Ledger map[string]Ledger
}

type TradeVolume struct {
	Currency  string                    `json:"cuurrency"`     // volume currency
	Volume    float64                   `json:"volume,string"` // current discount volume
	Fees      map[string]TradeVolumeFee `json:"fees"`          // array of asset pairs and fee tier info (if requested)
	FeesMaker map[string]TradeVolumeFee `json:"fees_maker"`    // array of asset pairs and maker fee tier info (if requested) for any pairs on maker/taker schedule
}

type TradeVolumeFee struct {
	Fee        float64 `json:"fee,string"`        // current fee in percent
	Minfee     float64 `json:"minfee,string"`     // minimum fee for pair (if not fixed fee)
	Maxfee     float64 `json:"maxfee,string"`     // maximum fee for pair (if not fixed fee)
	Nextfee    float64 `json:"nextfee,string"`    // next tier's fee for pair (if not fixed fee.  nil if at lowest fee tier)
	Nextvolume float64 `json:"nextvolume,string"` // volume level of next tier (if not fixed fee.  nil if at lowest fee tier)
	Tiervolume float64 `json:"tiervolume,string"` // volume level of current tier (if not fixed fee.  nil if at lowest fee tier)
}

func (t *AskBid) UnmarshalJSON(b []byte) error {
	var out []string

	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	if len(out) != 3 {
		return fmt.Errorf("Invalid number of entries")
	}

	t.Price, err = strconv.ParseFloat(out[0], 64)
	if err != nil {
		return err
	}

	t.WholeLotVolume, err = strconv.ParseFloat(out[1], 64)
	if err != nil {
		return err
	}

	t.LotVolume, err = strconv.ParseFloat(out[2], 64)
	if err != nil {
		return err
	}

	return nil
}

func (t *TradeToday) UnmarshalJSON(b []byte) error {
	var out []string

	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	if len(out) != 2 {
		return fmt.Errorf("Invalid number of entries")
	}

	t.Price, err = strconv.ParseFloat(out[0], 64)
	if err != nil {
		return err
	}

	t.Volume, err = strconv.ParseFloat(out[1], 64)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodayH24Float64) UnmarshalJSON(b []byte) error {
	var out []string

	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	if len(out) != 2 {
		return fmt.Errorf("Invalid number of entries")
	}

	t[0], err = strconv.ParseFloat(out[0], 64)
	if err != nil {
		return err
	}

	t[1], err = strconv.ParseFloat(out[1], 64)
	if err != nil {
		return err
	}

	return nil
}

func (t *PublicOrder) UnmarshalJSON(b []byte) error {
	var out []interface{}

	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	t.Price, err = strconv.ParseFloat(out[0].(string), 64)
	if err != nil {
		return err
	}

	t.Volume, err = strconv.ParseFloat(out[1].(string), 64)
	if err != nil {
		return err
	}

	t.Time = out[2].(float64)

	return nil
}
