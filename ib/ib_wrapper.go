package ib

import (
	"sync/atomic"
	"time"

	"go.uber.org/zap"

	. "github.com/hadrianl/ibapi"
)

var log, _ = zap.NewProduction()

var DBG bool = false

type IBWrapper struct {
	orderID int64
}

func (w *IBWrapper) InitWrapper() {

}

func (w *IBWrapper) GetNextOrderID() (i int64) {
	i = w.orderID
	atomic.AddInt64(&w.orderID, 1)
	return
}

func (w IBWrapper) ConnectAck() {
	log.Info("<ConnectAck>...")
}

func (w IBWrapper) ConnectionClosed() {
	log.Info("<ConnectionClosed>...")
}

func (w *IBWrapper) NextValidID(reqID int64) {
	atomic.StoreInt64(&w.orderID, reqID)
	if DBG {
		log.With(zap.Int64("reqID", reqID)).Info("<NextValidID>")
	}
}

func (w IBWrapper) ManagedAccounts(accountsList []string) {
	log.Info("<ManagedAccounts>", zap.Strings("accountList", accountsList))
}

func (w IBWrapper) TickPrice(reqID int64, tickType int64, price float64, attrib TickAttrib) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickPrice>", zap.Int64("tickType", tickType), zap.Float64("price", price))
}

func (w IBWrapper) UpdateAccountTime(accTime time.Time) {
	if DBG {
		log.Info("<UpdateAccountTime>", zap.Time("accountTime", accTime))
	}
}

func (w IBWrapper) UpdateAccountValue(tag string, value string, currency string, account string) {
	log.Info("<UpdateAccountValue>", zap.String("tag", tag), zap.String("value", value), zap.String("currency", currency), zap.String("account", account))
}

func (w IBWrapper) AccountDownloadEnd(accName string) {
	log.Info("<AccountDownloadEnd>", zap.String("accountName", accName))
}

func (w IBWrapper) AccountUpdateMulti(reqID int64, account string, modelCode string, tag string, value string, currency string) {
	log.With(zap.Int64("reqID", reqID)).Info("<AccountUpdateMulti>",
		zap.String("account", account),
		zap.String("modelCode", modelCode),
		zap.String("tag", tag),
		zap.String("value", value),
		zap.String("curreny", currency),
	)
}

func (w IBWrapper) AccountUpdateMultiEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<AccountUpdateMultiEnd>")
}

func (w IBWrapper) AccountSummary(reqID int64, account string, tag string, value string, currency string) {
	log.With(zap.Int64("reqID", reqID)).Info("<AccountSummary>",
		zap.String("account", account),
		zap.String("tag", tag),
		zap.String("value", value),
		zap.String("curreny", currency),
	)

}

func (w IBWrapper) AccountSummaryEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<AccountSummaryEnd>")
}

func (w IBWrapper) VerifyMessageAPI(apiData string) {
	log.Info("<VerifyMessageAPI>", zap.String("apiData", apiData))
}

func (w IBWrapper) VerifyCompleted(isSuccessful bool, err string) {
	log.Info("<VerifyCompleted>", zap.Bool("isSuccessful", isSuccessful), zap.String("error", err))
}

func (w IBWrapper) VerifyAndAuthMessageAPI(apiData string, xyzChallange string) {
	log.Info("<VerifyMessageAPI>", zap.String("apiData", apiData), zap.String("xyzChallange", xyzChallange))
}

func (w IBWrapper) VerifyAndAuthCompleted(isSuccessful bool, err string) {
	log.Info("<VerifyCompleted>", zap.Bool("isSuccessful", isSuccessful), zap.String("error", err))
}

func (w IBWrapper) DisplayGroupList(reqID int64, groups string) {
	log.With(zap.Int64("reqID", reqID)).Info("<DisplayGroupList>", zap.String("groups", groups))
}

func (w IBWrapper) DisplayGroupUpdated(reqID int64, contractInfo string) {
	log.With(zap.Int64("reqID", reqID)).Info("<DisplayGroupUpdated>", zap.String("contractInfo", contractInfo))
}

func (w IBWrapper) PositionMulti(reqID int64, account string, modelCode string, contract *Contract, position float64, avgCost float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<PositionMulti>",
		zap.String("account", account),
		zap.String("modelCode", modelCode),
		zap.Any("contract", contract),
		zap.Float64("position", position),
		zap.Float64("avgCost", avgCost),
	)
}

func (w IBWrapper) PositionMultiEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<PositionMultiEnd>")
}

func (w IBWrapper) UpdatePortfolio(contract *Contract, position float64, marketPrice float64, marketValue float64, averageCost float64, unrealizedPNL float64, realizedPNL float64, accName string) {
	log.Info("<UpdatePortfolio>",
		zap.String("localSymbol", contract.LocalSymbol),
		zap.Float64("position", position),
		zap.Float64("marketPrice", marketPrice),
		zap.Float64("averageCost", averageCost),
		zap.Float64("unrealizedPNL", unrealizedPNL),
		zap.Float64("realizedPNL", realizedPNL),
		zap.String("accName", accName),
	)
}

func (w IBWrapper) Position(account string, contract *Contract, position float64, avgCost float64) {
	log.Info("<UpdatePortfolio>",
		zap.String("account", account),
		zap.Any("contract", contract),
		zap.Float64("position", position),
		zap.Float64("avgCost", avgCost),
	)
}

func (w IBWrapper) PositionEnd() {
	log.Info("<PositionEnd>")
}

func (w IBWrapper) Pnl(reqID int64, dailyPnL float64, unrealizedPnL float64, realizedPnL float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<PNL>",
		zap.Float64("dailyPnL", dailyPnL),
		zap.Float64("unrealizedPnL", unrealizedPnL),
		zap.Float64("realizedPnL", realizedPnL),
	)
}

func (w IBWrapper) PnlSingle(reqID int64, position int64, dailyPnL float64, unrealizedPnL float64, realizedPnL float64, value float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<PNLSingle>",
		zap.Float64("dailyPnL", dailyPnL),
		zap.Float64("unrealizedPnL", unrealizedPnL),
		zap.Float64("realizedPnL", realizedPnL),
		zap.Float64("value", value),
	)
}

func (w IBWrapper) OpenOrder(orderID int64, contract *Contract, order *Order, orderState *OrderState) {
	log.With(zap.Int64("orderID", orderID)).Info("<OpenOrder>",
		zap.Any("contract", contract),
		zap.Any("order", order),
		zap.Any("orderState", orderState),
	)
}

func (w IBWrapper) OpenOrderEnd() {
	log.Info("<OpenOrderEnd>")

}

func (w IBWrapper) OrderStatus(orderID int64, status string, filled float64, remaining float64, avgFillPrice float64, permID int64, parentID int64, lastFillPrice float64, clientID int64, whyHeld string, mktCapPrice float64) {
	log.With(zap.Int64("orderID", orderID)).Info("<OrderStatus>",
		zap.String("status", status),
		zap.Float64("filled", filled),
		zap.Float64("remaining", remaining),
		zap.Float64("avgFillPrice", avgFillPrice),
	)
}

func (w IBWrapper) ExecDetails(reqID int64, contract *Contract, execution *Execution) {
	log.With(zap.Int64("reqID", reqID)).Info("<ExecDetails>",
		zap.Any("contract", contract),
		zap.Any("execution", execution),
	)
}

func (w IBWrapper) ExecDetailsEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<ExecDetailsEnd>")
}

func (w IBWrapper) DeltaNeutralValidation(reqID int64, deltaNeutralContract DeltaNeutralContract) {
	log.With(zap.Int64("reqID", reqID)).Info("<DeltaNeutralValidation>",
		zap.Any("deltaNeutralContract", deltaNeutralContract),
	)
}

func (w IBWrapper) CommissionReport(commissionReport CommissionReport) {
	log.Info("<CommissionReport>",
		zap.Any("commissionReport", commissionReport),
	)
}

func (w IBWrapper) OrderBound(reqID int64, apiClientID int64, apiOrderID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<OrderBound>",
		zap.Int64("apiClientID", apiClientID),
		zap.Int64("apiOrderID", apiOrderID),
	)
}

func (w IBWrapper) ContractDetails(reqID int64, conDetails *ContractDetails) {
	log.With(zap.Int64("reqID", reqID)).Info("<ContractDetails>",
		zap.Any("conDetails", conDetails),
	)
}

func (w IBWrapper) ContractDetailsEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<ContractDetailsEnd>")
}

func (w IBWrapper) BondContractDetails(reqID int64, conDetails *ContractDetails) {
	log.With(zap.Int64("reqID", reqID)).Info("<BondContractDetails>",
		zap.Any("conDetails", conDetails),
	)
}

func (w IBWrapper) SymbolSamples(reqID int64, contractDescriptions []ContractDescription) {
	log.With(zap.Int64("reqID", reqID)).Info("<SymbolSamples>",
		zap.Any("contractDescriptions", contractDescriptions),
	)
}

func (w IBWrapper) SmartComponents(reqID int64, smartComps []SmartComponent) {
	log.With(zap.Int64("reqID", reqID)).Info("<SmartComponents>",
		zap.Any("smartComps", smartComps),
	)
}

func (w IBWrapper) MarketRule(marketRuleID int64, priceIncrements []PriceIncrement) {
	log.Info("<MarketRule>",
		zap.Int64("marketRuleID", marketRuleID),
		zap.Any("priceIncrements", priceIncrements),
	)
}

func (w IBWrapper) RealtimeBar(reqID int64, time int64, open float64, high float64, low float64, close float64, volume int64, wap float64, count int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<RealtimeBar>",
		zap.Int64("time", time),
		zap.Float64("open", open),
		zap.Float64("high", high),
		zap.Float64("low", low),
		zap.Float64("close", close),
		zap.Int64("volume", volume),
		zap.Float64("wap", wap),
		zap.Int64("count", count),
	)
}

func (w IBWrapper) HistoricalData(reqID int64, bar *BarData) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalData>",
		zap.Any("bar", bar),
	)

}

func (w IBWrapper) HistoricalDataEnd(reqID int64, startDateStr string, endDateStr string) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalDataEnd>",
		zap.String("startDate", startDateStr),
		zap.String("endDate", endDateStr),
	)
}

func (w IBWrapper) HistoricalDataUpdate(reqID int64, bar *BarData) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalDataUpdate>",
		zap.Any("bar", bar),
	)
}

func (w IBWrapper) HeadTimestamp(reqID int64, headTimestamp string) {
	log.With(zap.Int64("reqID", reqID)).Info("<HeadTimestamp>",
		zap.String("headTimestamp", headTimestamp),
	)
}

func (w IBWrapper) HistoricalTicks(reqID int64, ticks []HistoricalTick, done bool) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalTicks>",
		zap.Any("ticks", ticks),
		zap.Bool("done", done),
	)
}

func (w IBWrapper) HistoricalTicksBidAsk(reqID int64, ticks []HistoricalTickBidAsk, done bool) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalTicksBidAsk>",
		zap.Any("ticks", ticks),
		zap.Bool("done", done),
	)
}

func (w IBWrapper) HistoricalTicksLast(reqID int64, ticks []HistoricalTickLast, done bool) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalTicksLast>",
		zap.Any("ticks", ticks),
		zap.Bool("done", done),
	)
}

func (w IBWrapper) TickSize(reqID int64, tickType int64, size int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickSize>",
		zap.Int64("tickType", tickType),
		zap.Int64("size", size),
	)
}

func (w IBWrapper) TickSnapshotEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickSnapshotEnd>")
}

func (w IBWrapper) MarketDataType(reqID int64, marketDataType int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<MarketDataType>",
		zap.Int64("marketDataType", marketDataType),
	)
}

func (w IBWrapper) TickByTickAllLast(reqID int64, tickType int64, time int64, price float64, size int64, tickAttribLast TickAttribLast, exchange string, specialConditions string) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickByTickAllLast>",
		zap.Int64("tickType", tickType),
		zap.Int64("time", time),
		zap.Float64("price", price),
		zap.Int64("size", size),
	)
}

func (w IBWrapper) TickByTickBidAsk(reqID int64, time int64, bidPrice float64, askPrice float64, bidSize int64, askSize int64, tickAttribBidAsk TickAttribBidAsk) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickByTickBidAsk>",
		zap.Int64("time", time),
		zap.Float64("bidPrice", bidPrice),
		zap.Float64("askPrice", askPrice),
		zap.Int64("bidPrice", bidSize),
		zap.Int64("askPrice", askSize),
	)
}

func (w IBWrapper) TickByTickMidPoint(reqID int64, time int64, midPoint float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickByTickMidPoint>",
		zap.Int64("time", time),
		zap.Float64("midPoint", midPoint),
	)
}

func (w IBWrapper) TickString(reqID int64, tickType int64, value string) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickString>",
		zap.Int64("tickType", tickType),
		zap.String("value", value),
	)
}

func (w IBWrapper) TickGeneric(reqID int64, tickType int64, value float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickGeneric>",
		zap.Int64("tickType", tickType),
		zap.Float64("value", value),
	)
}

func (w IBWrapper) TickEFP(reqID int64, tickType int64, basisPoints float64, formattedBasisPoints string, totalDividends float64, holdDays int64, futureLastTradeDate string, dividendImpact float64, dividendsToLastTradeDate float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickEFP>",
		zap.Int64("tickType", tickType),
		zap.Float64("basisPoints", basisPoints),
	)
}

func (w IBWrapper) TickReqParams(reqID int64, minTick float64, bboExchange string, snapshotPermissions int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickReqParams>",
		zap.Float64("minTick", minTick),
		zap.String("bboExchange", bboExchange),
		zap.Int64("snapshotPermissions", snapshotPermissions),
	)
}
func (w IBWrapper) MktDepthExchanges(depthMktDataDescriptions []DepthMktDataDescription) {
	log.Info("<MktDepthExchanges>",
		zap.Any("depthMktDataDescriptions", depthMktDataDescriptions),
	)
}

/*Returns the order book.

tickerId -  the request's identifier
position -  the order book's row being updated
operation - how to refresh the row:
	0 = insert (insert this new order into the row identified by 'position')
	1 = update (update the existing order in the row identified by 'position')
	2 = delete (delete the existing order at the row identified by 'position').
side -  0 for ask, 1 for bid
price - the order's price
size -  the order's size*/
func (w IBWrapper) UpdateMktDepth(reqID int64, position int64, operation int64, side int64, price float64, size int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<UpdateMktDepth>",
		zap.Int64("position", position),
		zap.Int64("operation", operation),
		zap.Int64("side", side),
		zap.Float64("price", price),
		zap.Int64("size", size),
	)
}

func (w IBWrapper) UpdateMktDepthL2(reqID int64, position int64, marketMaker string, operation int64, side int64, price float64, size int64, isSmartDepth bool) {
	log.With(zap.Int64("reqID", reqID)).Info("<UpdateMktDepthL2>",
		zap.Int64("position", position),
		zap.String("marketMaker", marketMaker),
		zap.Int64("operation", operation),
		zap.Int64("side", side),
		zap.Float64("price", price),
		zap.Int64("size", size),
		zap.Bool("isSmartDepth", isSmartDepth),
	)
}

func (w IBWrapper) TickOptionComputation(reqID int64, tickType int64, impliedVol float64, delta float64, optPrice float64, pvDiviedn float64, gamma float64, vega float64, theta float64, undPrice float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<TickOptionComputation>",
		zap.Int64("tickType", tickType),
		zap.Float64("impliedVol", impliedVol),
		zap.Float64("delta", delta),
		zap.Float64("optPrice", optPrice),
		zap.Float64("pvDiviedn", pvDiviedn),
		zap.Float64("gamma", gamma),
		zap.Float64("vega", vega),
		zap.Float64("theta", theta),
		zap.Float64("undPrice", undPrice),
	)
}

func (w IBWrapper) FundamentalData(reqID int64, data string) {
	log.With(zap.Int64("reqID", reqID)).Info("<FundamentalData>",
		zap.String("data", data),
	)
}

func (w IBWrapper) ScannerParameters(xml string) {
	log.Info("<ScannerParameters>",
		zap.String("xml", xml),
	)

}

func (w IBWrapper) ScannerData(reqID int64, rank int64, conDetails *ContractDetails, distance string, benchmark string, projection string, legs string) {
	log.With(zap.Int64("reqID", reqID)).Info("<ScannerData>",
		zap.Int64("rank", rank),
		zap.Any("conDetails", conDetails),
		zap.String("distance", distance),
		zap.String("benchmark", benchmark),
		zap.String("projection", projection),
		zap.String("legs", legs),
	)
}

func (w IBWrapper) ScannerDataEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<ScannerDataEnd>")
}

func (w IBWrapper) HistogramData(reqID int64, histogram []HistogramData) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistogramData>",
		zap.Any("histogram", histogram),
	)
}

func (w IBWrapper) RerouteMktDataReq(reqID int64, contractID int64, exchange string) {
	log.With(zap.Int64("reqID", reqID)).Info("<RerouteMktDataReq>",
		zap.Int64("contractID", contractID),
		zap.String("exchange", exchange),
	)
}

func (w IBWrapper) RerouteMktDepthReq(reqID int64, contractID int64, exchange string) {
	log.With(zap.Int64("reqID", reqID)).Info("<RerouteMktDepthReq>",
		zap.Int64("contractID", contractID),
		zap.String("exchange", exchange),
	)
}

func (w IBWrapper) SecurityDefinitionOptionParameter(reqID int64, exchange string, underlyingContractID int64, tradingClass string, multiplier string, expirations []string, strikes []float64) {
	log.With(zap.Int64("reqID", reqID)).Info("<SecurityDefinitionOptionParameter>",
		zap.String("exchange", exchange),
		zap.Int64("underlyingContractID", underlyingContractID),
		zap.String("tradingClass", tradingClass),
		zap.String("multiplier", multiplier),
		zap.Strings("expirations", expirations),
		zap.Float64s("strikes", strikes),
	)
}

func (w IBWrapper) SecurityDefinitionOptionParameterEnd(reqID int64) {
	log.With(zap.Int64("reqID", reqID)).Info("<SecurityDefinitionOptionParameterEnd>")
}

func (w IBWrapper) SoftDollarTiers(reqID int64, tiers []SoftDollarTier) {
	log.With(zap.Int64("reqID", reqID)).Info("<SoftDollarTiers>",
		zap.Any("tiers", tiers),
	)
}

func (w IBWrapper) FamilyCodes(famCodes []FamilyCode) {
	log.Info("<FamilyCodes>",
		zap.Any("famCodes", famCodes),
	)
}

func (w IBWrapper) NewsProviders(newsProviders []NewsProvider) {
	log.Info("<NewsProviders>",
		zap.Any("newsProviders", newsProviders),
	)
}

func (w IBWrapper) TickNews(tickerID int64, timeStamp int64, providerCode string, articleID string, headline string, extraData string) {
	log.With(zap.Int64("tickerID", tickerID)).Info("<TickNews>",
		zap.Int64("timeStamp", timeStamp),
		zap.String("providerCode", providerCode),
		zap.String("articleID", articleID),
		zap.String("headline", headline),
		zap.String("extraData", extraData),
	)
}

func (w IBWrapper) NewsArticle(reqID int64, articleType int64, articleText string) {
	log.With(zap.Int64("reqID", reqID)).Info("<NewsArticle>",
		zap.Int64("articleType", articleType),
		zap.String("articleText", articleText),
	)
}

func (w IBWrapper) HistoricalNews(reqID int64, time string, providerCode string, articleID string, headline string) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalNews>",
		zap.String("time", time),
		zap.String("providerCode", providerCode),
		zap.String("articleID", articleID),
		zap.String("headline", headline),
	)
}

func (w IBWrapper) HistoricalNewsEnd(reqID int64, hasMore bool) {
	log.With(zap.Int64("reqID", reqID)).Info("<HistoricalNewsEnd>",
		zap.Bool("hasMore", hasMore),
	)
}

func (w IBWrapper) UpdateNewsBulletin(msgID int64, msgType int64, newsMessage string, originExch string) {
	log.With(zap.Int64("msgID", msgID)).Info("<UpdateNewsBulletin>",
		zap.Int64("msgType", msgType),
		zap.String("newsMessage", newsMessage),
		zap.String("originExch", originExch),
	)
}

func (w IBWrapper) ReceiveFA(faData int64, cxml string) {
	log.Info("<ReceiveFA>",
		zap.Int64("faData", faData),
		zap.String("cxml", cxml),
	)
}

func (w IBWrapper) CurrentTime(t time.Time) {
	log.Info("<CurrentTime>",
		zap.Time("time", t),
	)
}

func (w IBWrapper) Error(reqID int64, errCode int64, errString string) {
	log.With(zap.Int64("reqID", reqID)).Info("<Error>",
		zap.Int64("errCode", errCode),
		zap.String("errString", errString),
	)
}

func (w IBWrapper) CompletedOrder(contract *Contract, order *Order, orderState *OrderState) {
	log.Info("<CompletedOrder>",
		zap.Any("contract", contract),
		zap.Any("order", order),
		zap.Any("orderState", orderState),
	)
}

func (w IBWrapper) CompletedOrdersEnd() {
	log.Info("<CompletedOrdersEnd>:")
}

func (w IBWrapper) ReplaceFAEnd(reqID int64, text string) {
	log.With(zap.Int64("reqID", reqID)).Info("<ReplaceFAEnd>", zap.String("text", text))
}
