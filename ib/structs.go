package ib

import (
	. "github.com/hadrianl/ibapi"
)

type IBConfig struct {
	IbHost   string
	IbPort   int
	ClientID int64
}

type IB struct {
	Client   *IbClient
	Wrapper  IbWrapper
	host     string
	port     int
	clientID int64
}
