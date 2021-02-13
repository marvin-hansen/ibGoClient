package ib

import (
	"github.com/hadrianl/ibapi"
)

const (
	currency = "USD"
	//
	equityExchange = "SMART"
	forexExchange  = "IDEALPRO"
	futureExchange = "GLOBEX"
	//
	indexSecurityType  = "IND"
	equitySecurityType = "STK"
	forexSecurityType  = "CASH"
	futureSecurityType = "FUT"
	contFutureSecType  = "CONTFUT"
)

// Basic IB Contracts
// http://interactivebrokers.github.io/tws-api/basic_contracts.html

//For certain smart-routed stock contracts that have the same symbol, currency and exchange,
// you would also need to specify the primary exchange attribute to uniquely define the contract. T
//his should be defined as the native exchange of a contract, and is good practice to include for all stocks

/*
 * GetEquityContract returns an equity contract object.
 *
 * symbol - string defines the equity i.e. APPL
 *
 * primaryExchange - string defines the ib trading place.
 *
 * Currency is USD by default
 *
 * Note, in the IB API, NASDAQ is always defined as ISLAND
 */
func GetEquityContract(symbol string, primaryExchange string) ibapi.Contract {
	return ibapi.Contract{
		Symbol:          symbol,
		PrimaryExchange: primaryExchange,
		SecurityType:    equitySecurityType,
		Currency:        currency,
		Exchange:        equityExchange,
	}
}

/**
 * 	GetForexContract returns a Forex contract object.
 *
 *  symbol - string defines the fx i.e. EUR
 *
 *  When fxExchange is an empty string, IDEALPRO is used by default.
 *
 * Currency is USD by default
 */
func GetForexContract(symbol string, currency string, fxExchange string) ibapi.Contract {

	var exchange string

	switch {
	case fxExchange == "":
		exchange = forexExchange
	default:
		exchange = fxExchange
	}

	return ibapi.Contract{
		Symbol:       symbol,
		Currency:     currency,
		Exchange:     exchange,
		SecurityType: forexSecurityType,
	}
}

/**
 * 	GetFutureContract returns a Future contract object.
 *  localSymbol - string defines the IB future symbol i.e. ESZ6;
 *
 *  Since a local symbol uniquely defines a future, an expiry is not necessary.
 *
 *  When futExchange is an empty string, GLOBEX is used by default.
 *
 * Currency is USD by default
 */
func GetFutureContract(localSymbol string, futExchange string) ibapi.Contract {

	var exchange string

	switch {
	case futExchange == "":
		exchange = futureExchange
	default:
		exchange = futExchange
	}

	return ibapi.Contract{
		LocalSymbol:  localSymbol,
		Exchange:     exchange,
		SecurityType: futureSecurityType,
		Currency:     currency,
	}
}

/**
 * GetContinuousFutureContract returns a continuous future contract object.
 * symbol - string defines the ib Future i.e. ES
 *
 * futExchange - string  exchange trading the ib future
 *
 * When futExchange is an empty string, GLOBEX is used by default.
 *
 * Currency is USD by default
 *
 * Note, continuous futures cannot be used with real time data or to place orders,
 * but only for historical data.
 */
func GetContinuousFutureContract(symbol string, futExchange string) ibapi.Contract {

	var exchange string

	switch {
	case futExchange == "":
		exchange = futureExchange
	default:
		exchange = futExchange
	}

	return ibapi.Contract{
		Symbol:       symbol,
		Exchange:     exchange,
		Currency:     currency,
		SecurityType: contFutureSecType,
	}
}

/**
 * GetIndexContract returns an index contract object.
 *
 * symbol - string the index symbol
 *
 * currency - string the index currency. No default. Must be provided.
 *
 * exchange - string the index exchange. No default. Must be provided.
 *
 */
func GetIndexContract(symbol string, currency string, exchange string) ibapi.Contract {

	return ibapi.Contract{
		Symbol:       symbol,
		Currency:     currency,
		Exchange:     exchange,
		SecurityType: indexSecurityType,
	}
}
