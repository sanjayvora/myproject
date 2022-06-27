package transport

import (
	"encoding/json"
	"fmt"
	spec "myproject/cabmgmt/spec/cab"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CabBL interface {
	AddCab(r *spec.AddCabParam) (*spec.AddCabResponse, error)
	GetCab(r *spec.GetCabRequest) (*spec.GetCabResponse, error)
	DeleteCab(r *spec.DeleteCabRequest) error
	ChangeCity(r *spec.ChangeCityRequest) error
}

type Cab struct {
	CabBL CabBL
}

func (c *Cab) AddHandler(r *mux.Router, basePath string) {
	r.HandleFunc(basePath+spec.AddCabPath, c.AddCab).Methods(http.MethodPost)
	r.HandleFunc(basePath+spec.GetCabPath, c.GetCab).Methods(http.MethodGet)
	r.HandleFunc(basePath+spec.DeleteCabPath, c.DeleteCab).Methods(http.MethodDelete)
	r.HandleFunc(basePath+spec.ChangeCityPath, c.ChangeCity).Methods(http.MethodPut)
}
func (c *Cab) AddCab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "AddCab")

	param := &spec.AddCabParam{}
	err := json.NewDecoder(r.Body).Decode(param)
	if err != nil {
		fmt.Println("Tranport", "AddCab", "body param not found", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.CabBL.AddCab(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *Cab) DeleteCab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "DeleteCab")
	request := &spec.DeleteCabRequest{}
	cabID, errCode := getCabID(mux.Vars(r), "DeleteCab")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.CabID = cabID

	err := c.CabBL.DeleteCab(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *Cab) GetCab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "GetCab")
	request := &spec.GetCabRequest{}
	cabID, errCode := getCabID(mux.Vars(r), "GetCab")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.CabID = cabID

	val := r.URL.Query().Get("getHistory")
	if val == "true" {
		request.GetCabHistory = true
	}

	val = r.URL.Query().Get("count")
	if val != "" {
		request.RecordCount, _ = strconv.Atoi(val)
	}
	result, err := c.CabBL.GetCab(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *Cab) ChangeCity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tranport", "ChangeCity")

	request := &spec.ChangeCityRequest{}
	cabID, errCode := getCabID(mux.Vars(r), "ChangeCity")
	if errCode != 0 {
		w.WriteHeader(errCode)
		return
	}
	request.CabID = cabID
	request.Param = &spec.ChangeCityParam{}
	err := json.NewDecoder(r.Body).Decode(request.Param)
	if err != nil {
		fmt.Println("Tranport", "ChangeCity", "body param not found", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.CabBL.ChangeCity(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getCabID(params map[string]string, method string) (int, int) {
	cabID, err := strconv.Atoi(params["cabID"])
	if err != nil {
		fmt.Println("Tranport", method, "cabID not found in path", err.Error())
		return 0, http.StatusBadRequest
	}
	if cabID == 0 {
		fmt.Println("Tranport", method, "cabID found zero in path")
		return 0, http.StatusNotFound
	}
	return cabID, 0
}
