package handlers

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ogarraux/airmon/ports"
)

type TimerHandler struct {
	Service ports.AirMonitorService
}

func NewTimerHandler(s ports.AirMonitorService) TimerHandler {
	return TimerHandler{
		Service: s,
	}
}

func (th TimerHandler) Schedule() {
	interval := int64(th.Service.GetConfiguration().GetInterval())

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	fmt.Println("Scheduling collections every " + strconv.Itoa(int(interval)) + " seconds.")
	done := make(chan bool)

	for {
		select {
		case <-done:
			// Nothing ever calls this, this is intended to continue running
			fmt.Println("Done with ticks")
		case t := <-ticker.C:
			fmt.Println("Running collections ", t)
			err := th.Service.CollectAllMeasurements()
			if err != nil {
				// Probably best to just continue to next collection in case of errors
				fmt.Println(err.Error())
				continue
			}
			if MonitorURL := th.Service.GetConfiguration().GetMonitorEndpoint(); MonitorURL != "" {
				// Monitoring Endpoint - don't validate TLS
				transport := &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
				client := &http.Client{Transport: transport}
				_, err := client.Get(MonitorURL)
				if err != nil {
					fmt.Println("Error updating monitor endpoint " + MonitorURL + " Error: " + err.Error())
				} else {
					fmt.Println("Updated monitor endpoint " + MonitorURL)
				}
			}

		}
	}
}
