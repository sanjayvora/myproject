// Package classification Cab Management Service
//
// Cab Management Service (REST)
//
//     Schemes: https
//     Host: cabmgmt.com
//     BasePath: /cabmgmt/v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta

package main

import (
	"encoding/json"
	"fmt"
	"myproject/cabmgmt/bl"
	"myproject/cabmgmt/dl"
	"myproject/cabmgmt/svcparams"
	"myproject/cabmgmt/transport"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received ping ...")
	err := json.NewEncoder(w).Encode("success")
	if err != nil {
		fmt.Println(err.Error(), "Error sending response for Ping")
	}
}

func main() {
	fmt.Println("Starting Cab Management Service")
	dl := &dl.DataLayer{}
	cabTransport := transport.Cab{CabBL: bl.NewCabBL(dl)}
	cityTransport := transport.City{CityBL: bl.NewCityBL(dl)}
	tripTransport := transport.Trip{TripBL: bl.NewTripBL(dl)}
	go func() {
		//Use 0 so that Port is picked automatically
		listener, err := net.Listen("tcp", "localhost:"+fmt.Sprintf("%d", svcparams.Port))
		if err != nil {
			fmt.Println("Error in starting rest server", err.Error())
			return
		}
		port := listener.Addr().(*net.TCPAddr).Port
		fmt.Println("Rest server started on", port)

		router := mux.NewRouter()
		router.HandleFunc("/ping", Ping).Methods(http.MethodGet)
		cabTransport.AddHandler(router, svcparams.BasePath)
		cityTransport.AddHandler(router, svcparams.BasePath)
		tripTransport.AddHandler(router, svcparams.BasePath)

		err = http.Serve(listener, router)
		fmt.Println(err, "Error in starting rest server")
		os.Exit(1)

	}()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	<-errs
	fmt.Println("Exiting Cab Management Service")
}
