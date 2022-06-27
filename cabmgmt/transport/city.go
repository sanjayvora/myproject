package transport

import (
	"encoding/json"
	"fmt"
	spec "myproject/cabmgmt/spec/city"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CityBL interface {
	AddCity(r *spec.AddCityParam) (*spec.AddCityResponse, error)
	GetCities(r *spec.GetCityRequest) (*spec.GetCityResponse, error)
	DeleteCity(r *spec.DeleteCityRequest) error
	GetDemand(r *spec.GetDemandRequest) (*spec.DemandResponse, error)
}

type City struct {
	CityBL CityBL
}

func (c *City) AddHandler(r *mux.Router, basePath string) {
	r.HandleFunc(basePath+spec.AddCityPath, c.AddCity).Methods(http.MethodPost)
	r.HandleFunc(basePath+spec.GetCityPath, c.GetCities).Methods(http.MethodGet)
	r.HandleFunc(basePath+spec.DeleteCityPath, c.DeleteCity).Methods(http.MethodDelete)
	r.HandleFunc(basePath+spec.GetDemandPath, c.GetDemand).Methods(http.MethodGet)
}
func (c *City) AddCity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "AddCity")

	param := &spec.AddCityParam{}
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		fmt.Println("Tranport", "AddCity", "body param not found", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.CityBL.AddCity(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *City) DeleteCity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "DeleteCity")
	request := &spec.DeleteCityRequest{}
	cityID, errCode := getCityID(mux.Vars(r), "DeleteCity")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.CityID = cityID

	err := c.CityBL.DeleteCity(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *City) GetCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "GetCities")
	request := &spec.GetCityRequest{}
	result, err := c.CityBL.GetCities(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *City) GetDemand(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "GetDemand")

	request := &spec.GetDemandRequest{}
	val := r.URL.Query().Get("count")
	if val != "" {
		request.Count, _ = strconv.Atoi(val)
	}

	result, err := c.CityBL.GetDemand(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func getCityID(params map[string]string, method string) (int, int) {
	cityID, err := strconv.Atoi(params["cityID"])
	if err != nil {
		fmt.Println("Tranport", method, "cityID not found in path", err.Error())
		return 0, http.StatusBadRequest
	}
	if cityID == 0 {
		fmt.Println("Tranport", method, "cityID found zero in path")
		return 0, http.StatusNotFound
	}
	return cityID, 0
}
