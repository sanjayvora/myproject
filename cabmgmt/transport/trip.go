package transport

import (
	"encoding/json"
	"fmt"
	spec "myproject/cabmgmt/spec/trip"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TripBL interface {
	AddTrip(r *spec.AddTripParam) (*spec.AddTripResponse, error)
	StartTrip(r *spec.TripRequest) error
	EndTrip(r *spec.TripRequest) error
}

type Trip struct {
	TripBL TripBL
}

func (c *Trip) AddHandler(r *mux.Router, basePath string) {
	r.HandleFunc(basePath+spec.AddTripPath, c.AddTrip).Methods(http.MethodPost)
	r.HandleFunc(basePath+spec.StartTripPath, c.StartTrip).Methods(http.MethodPut)
	r.HandleFunc(basePath+spec.EndTripPath, c.EndTrip).Methods(http.MethodPut)
}
func (c *Trip) AddTrip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "AddTrip")

	param := &spec.AddTripParam{}
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		fmt.Println("Tranport", "AddTrip", "body param not found", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.TripBL.AddTrip(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *Trip) StartTrip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "StartTrip")
	request := &spec.TripRequest{}
	tripID, errCode := getTripID(mux.Vars(r), "StartTrip")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.TripID = tripID
	param, errCode := getTime(r, "EndTrip")
	if errCode != 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Param = param

	err := c.TripBL.StartTrip(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Trip) EndTrip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "EndTrip")
	request := &spec.TripRequest{}
	tripID, errCode := getTripID(mux.Vars(r), "EndTrip")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.TripID = tripID
	param, errCode := getTime(r, "EndTrip")
	if errCode != 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Param = param
	err := c.TripBL.EndTrip(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getTripID(params map[string]string, method string) (int, int) {
	tripID, err := strconv.Atoi(params["tripID"])
	if err != nil {
		fmt.Println("Tranport", method, "tripID not found in path", err.Error())
		return 0, http.StatusBadRequest
	}
	if tripID == 0 {
		fmt.Println("Tranport", method, "tripID found zero in path")
		return 0, http.StatusNotFound
	}
	return tripID, 0
}

func getTime(r *http.Request, method string) (*spec.TripParam, int) {
	param := &spec.TripParam{}
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		fmt.Println("Tranport", method, "body param not found", err.Error())
		return nil, http.StatusBadRequest
	}

	if param.Time == 0 {
		fmt.Println("Tranport", method, "trip start/end time found zero")
		return nil, http.StatusNotFound
	}
	return param, 0
}
