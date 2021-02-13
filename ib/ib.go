package ib

import (
	. "github.com/hadrianl/ibapi"
)

func NewIBApi(ibConfig IBConfig, ibWrapper IbWrapper) *IB {
	ib := &IB{
		host:     ibConfig.IbHost,
		port:     ibConfig.IbPort,
		clientID: ibConfig.ClientID,
	}

	ibClient := NewIbClient(ibWrapper)
	ib.Client = ibClient
	ib.Wrapper = ibWrapper

	return ib
}

func (ib *IB) GetIBClient() *IbClient {
	return ib.Client
}

// Connect to TWS or Gateway
func (ib *IB) Connect() error {

	if err := ib.Client.Connect(ib.host, ib.port, ib.clientID); err != nil {
		return err
	}

	if err := ib.Client.HandShake(); err != nil {
		return err
	}

	err := ib.Client.Run()
	return err
}

func (ib *IB) Disconnect() error {
	err := ib.Client.Disconnect()
	return err
}
