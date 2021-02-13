package main

import (
	. "github.com/hadrianl/ibapi"
	"golang.org/x/net/context"
	. "ibGoClient/ib"
	"time"
)

func main() {
	log := GetLogger().Sugar()
	defer log.Sync()

	SetAPILogger()

	ibConf := GetIbConfigLocal()

	ibApi := NewIBApi(ibConf, &IBWrapper{})
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	ic := ibApi.GetIBClient()
	ic.SetContext(ctx)

	err := ibApi.Connect()
	CheckNoError(err, "Connect failed:")

	ic.ReqCurrentTime()

	eusdContract := GetForexContract("EUR", "USD", "")
	eusdContract.Symbol = "EUR"

	//ic.ReqHistoricalData()

	ic.ReqAutoOpenOrders(true)
	ic.ReqAccountUpdates(true, "")
	//ic.ReqExecutions(ic.GetReqID(), ExecutionFilter{})

	//ic.Run()
	//err = ic.LoopUntilDone() // block until ctx or ic is done
	//log.Info(err)

	// start to send req and receive msg from tws or gateway after this
	_ = ic.Run()
	<-time.After(time.Second * 10)
	err = ibApi.Disconnect()
	CheckNoError(err, "Failed to disconnect from IB Gateway")

}
