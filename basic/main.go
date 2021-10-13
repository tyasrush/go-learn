package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"encoding/json"
)

type BinanceExchangeInfo struct {
	Symbols []struct {
		Symbol string                   `json:"symbol"`
		Filter []map[string]interface{} `json:"filters"`
	} `json:"symbols"`
}

type IndodaxTickerResponse struct {
	Ticker struct {
		Last string `json:"last"`
	} `json:"ticker"`
}

type TradingRule struct {
	TradingPair               string  `json:"trading_pair"`
	MinOrderSize              float64 `json:"min_order_size"`
	MaxOrderSize              float64 `json:"max_order_size"`
	MinPriceIncrement         float64 `json:"min_price_increment"`
	MinBaseAmountIncrement    float64 `json:"min_base_amount_increment"`
	MinQuoteAmountIncrement   float64 `json:"min_quote_amount_increment"`
	MaxPriceSignificantDigits float64 `json:"max_price_significant_digits"`
	MinNotionalSize           float64 `json:"min_notional_size"`
	MinOrderValue             float64 `json:"min_order_value"`
	SupportsLimitOrders       bool    `json:"supports_limit_orders"`
	SupportsMarketOrders      bool    `json:"support_limit_orders"`
}

func main() {
	n := time.Now()
	monthInt := fmt.Sprintf("%d", int(n.Month()))
	if n.Month() < 10 {
		monthInt = fmt.Sprintf("0%d", int(n.Month()))
	}
	strDateJoin := fmt.Sprintf("%d%s%d", n.Year(), monthInt, n.Day())
	fmt.Println("time - ", strDateJoin)

	logger := log.New(os.Stdout, "app: ", log.Lshortfile)
	logger.Print("testing trace")

	var varitf interface{}
	varitf = 1

	fmt.Println(varitf)

	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("not oke gaes")
	}

	fmt.Printf("pc - %v file - %s line - %d\n", pc, file, line)

	stringJSON := `
	{
  "timezone": "UTC",
  "serverTime": 1634079711932,
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 1200
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "SECOND",
      "intervalNum": 10,
      "limit": 50
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "DAY",
      "intervalNum": 1,
      "limit": 160000
    },
    {
      "rateLimitType": "RAW_REQUESTS",
      "interval": "MINUTE",
      "intervalNum": 5,
      "limit": 6100
    }
  ],
  "exchangeFilters": [],
  "symbols": [
    {
      "symbol": "BTCUSDT",
      "status": "TRADING",
      "baseAsset": "BTC",
      "baseAssetPrecision": 8,
      "quoteAsset": "USDT",
      "quotePrecision": 8,
      "quoteAssetPrecision": 8,
      "baseCommissionPrecision": 8,
      "quoteCommissionPrecision": 8,
      "orderTypes": [
        "LIMIT",
        "LIMIT_MAKER",
        "MARKET",
        "STOP_LOSS_LIMIT",
        "TAKE_PROFIT_LIMIT"
      ],
      "icebergAllowed": true,
      "ocoAllowed": true,
      "quoteOrderQtyMarketAllowed": true,
      "isSpotTradingAllowed": true,
      "isMarginTradingAllowed": true,
      "filters": [
        {
          "filterType": "PRICE_FILTER",
          "minPrice": "0.01000000",
          "maxPrice": "1000000.00000000",
          "tickSize": "0.01000000"
        },
        {
          "filterType": "PERCENT_PRICE",
          "multiplierUp": "5",
          "multiplierDown": "0.2",
          "avgPriceMins": 5
        },
        {
          "filterType": "LOT_SIZE",
          "minQty": "0.00001000",
          "maxQty": "9000.00000000",
          "stepSize": "0.00001000"
        },
        {
          "filterType": "MIN_NOTIONAL",
          "minNotional": "10.00000000",
          "applyToMarket": true,
          "avgPriceMins": 5
        },
        {
          "filterType": "ICEBERG_PARTS",
          "limit": 10
        },
        {
          "filterType": "MARKET_LOT_SIZE",
          "minQty": "0.00000000",
          "maxQty": "118.44991687",
          "stepSize": "0.00000000"
        },
        {
          "filterType": "MAX_NUM_ORDERS",
          "maxNumOrders": 200
        },
        {
          "filterType": "MAX_NUM_ALGO_ORDERS",
          "maxNumAlgoOrders": 5
        }
      ],
      "permissions": [
        "SPOT",
        "MARGIN"
      ]
    },
    {
      "symbol": "BNBUSDT",
      "status": "TRADING",
      "baseAsset": "BNB",
      "baseAssetPrecision": 8,
      "quoteAsset": "USDT",
      "quotePrecision": 8,
      "quoteAssetPrecision": 8,
      "baseCommissionPrecision": 8,
      "quoteCommissionPrecision": 8,
      "orderTypes": [
        "LIMIT",
        "LIMIT_MAKER",
        "MARKET",
        "STOP_LOSS_LIMIT",
        "TAKE_PROFIT_LIMIT"
      ],
      "icebergAllowed": true,
      "ocoAllowed": true,
      "quoteOrderQtyMarketAllowed": true,
      "isSpotTradingAllowed": true,
      "isMarginTradingAllowed": true,
      "filters": [
        {
          "filterType": "PRICE_FILTER",
          "minPrice": "0.10000000",
          "maxPrice": "100000.00000000",
          "tickSize": "0.10000000"
        },
        {
          "filterType": "PERCENT_PRICE",
          "multiplierUp": "5",
          "multiplierDown": "0.2",
          "avgPriceMins": 5
        },
        {
          "filterType": "LOT_SIZE",
          "minQty": "0.00100000",
          "maxQty": "900000.00000000",
          "stepSize": "0.00100000"
        },
        {
          "filterType": "MIN_NOTIONAL",
          "minNotional": "10.00000000",
          "applyToMarket": true,
          "avgPriceMins": 5
        },
        {
          "filterType": "ICEBERG_PARTS",
          "limit": 10
        },
        {
          "filterType": "MARKET_LOT_SIZE",
          "minQty": "0.00000000",
          "maxQty": "39516.20646212",
          "stepSize": "0.00000000"
        },
        {
          "filterType": "MAX_NUM_ORDERS",
          "maxNumOrders": 200
        },
        {
          "filterType": "MAX_NUM_ALGO_ORDERS",
          "maxNumAlgoOrders": 5
        }
      ],
      "permissions": [
        "SPOT",
        "MARGIN"
      ]
    }
  ]
}`

	var (
		eir                 BinanceExchangeInfo
		binanceTradingRules []TradingRule
	)
	err := json.Unmarshal([]byte(stringJSON), &eir)
	if err != nil {
		fmt.Println("error = ", err)
	}

	for _, item := range eir.Symbols {
		var tradingRule = TradingRule{
			TradingPair: item.Symbol,
		}
		for _, filterItem := range item.Filter {
			if filterItem["filterType"] == "PRICE_FILTER" {
				value, err := strconv.ParseFloat(filterItem["tickSize"].(string), 10)
				if err != nil {
					log.Fatal(err)
				}
				tradingRule.MinPriceIncrement = value
			}

			if filterItem["filterType"] == "LOT_SIZE" {

				value, err := strconv.ParseFloat(filterItem["stepSize"].(string), 10)
				if err != nil {
					log.Fatal(err)
				}
				tradingRule.MinBaseAmountIncrement = value
				value1, err := strconv.ParseFloat(filterItem["minQty"].(string), 10)
				if err != nil {
					log.Fatal(err)
				}

				tradingRule.MinOrderSize = value1
			}

			if filterItem["filterType"] == "MIN_NOTIONAL" {
				value, err := strconv.ParseFloat(filterItem["minNotional"].(string), 10)
				if err != nil {
					log.Fatal(err)
				}
				tradingRule.MinNotionalSize = value
			}
		}

		binanceTradingRules = append(binanceTradingRules, tradingRule)
	}

	// trading_pair = rule.get("symbol")
	// filters = rule.get("filters")
	// price_filter = [f for f in filters if f.get("filterType") == "PRICE_FILTER"][0]
	// lot_size_filter = [f for f in filters if f.get("filterType") == "LOT_SIZE"][0]
	// min_notional_filter = [f for f in filters if f.get("filterType") == "MIN_NOTIONAL"][0]

	// min_order_size = Decimal(lot_size_filter.get("minQty"))
	// tick_size = price_filter.get("tickSize")
	// step_size = Decimal(lot_size_filter.get("stepSize"))
	// min_notional = Decimal(min_notional_filter.get("minNotional"))

	// retval.append(
	// 	TradingRule(trading_pair,
	// 				min_order_size=min_order_size,
	// 				min_price_increment=Decimal(tick_size),
	// 				min_base_amount_increment=Decimal(step_size),
	// 				min_notional_size=Decimal(min_notional)))

	// binanceBytesTradingRules, err := json.Marshal(binanceTradingRules)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("json marshal - ", eir)
	// fmt.Println("trade rules - ", string(binanceBytesTradingRules))

	iddxTickerResp := `{
  "ticker": {
    "high": "822840000",
    "low": "770030000",
    "vol_btc": "102.49640046",
    "vol_idr": "83019289867",
    "last": "803415000",
    "buy": "803415000",
    "sell": "803416000",
    "server_time": 1634082788
  }
}`

	var (
		iddxResp IndodaxTickerResponse
		// iddxTradingRules TradingRule
	)
	err = json.Unmarshal([]byte(iddxTickerResp), &iddxResp)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("response - ", iddxResp)

	var (
		msgChan    = make(chan string)
		errorChan  = make(chan error)
		intCursor  = 0
		syncG      sync.WaitGroup
		errorRetry = make(chan int, 1)
		retryCount = 0
	)

	ctx, cancel := context.WithCancel(context.Background())

	TestGoFunc(intCursor, msgChan, errorChan)

	syncG.Add(1)
	go func() {
		fmt.Println("goroutine executed - ", time.Now().String())
		defer syncG.Done()
		for {
			select {
			case errretry := <-errorRetry:
				if errretry == 2 {
					cancel()
					return
				}
			case <-ctx.Done():
				fmt.Println("context done")
				syncG.Done()
				return
			case errMsg := <-errorChan:
				fmt.Println("error: ", errMsg)
				// syncG.Done()
				intCursor = 0
				retryCount++
				errorRetry <- retryCount
				time.Sleep(2 * time.Second)
				errorChan = make(chan error)
				// syncG.Done()
				// syncG.Add(1)
				TestGoFunc(intCursor, msgChan, errorChan)
				// return
			case msg := <-msgChan:
				fmt.Println("message: ", msg)
			}
		}
	}()

	syncG.Wait()
}

func TestGoFunc(ic int, msgChan chan<- string, errMsg chan<- error) {
	go func() {
		for {
			if ic <= 2 {
				ic++
				msgChan <- fmt.Sprintf("nah kan, message ke - %d time - %s", ic, time.Now().String())
				time.Sleep(3 * time.Second)
			} else {
				errMsg <- errors.New("udah abis gaes")
			}
		}
	}()
}
