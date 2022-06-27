package cab

// resource cab endpoints
const (
	// swagger:operation POST /cabs Cab AddCabRequest
	// ---
	// summary: Register new cab
	// description: Register new cab with cab details
	// responses:
	//   "200":
	//     description: New cab ID
	//     schema:
	//         $ref: '#/definitions/AddCabResponse'
	//   "400":
	//     description: Bad Request
	//   "500":
	//     description: The request was not processed due to an server error.
	AddCabPath = "/cabs"

	// swagger:operation GET /cabs/{cabID} Cab GetCabRequest
	// ---
	// summary: Get cab details
	// description: Get details of an existing cab
	// responses:
	//   "200":
	//     description: Cab details
	//     schema:
	//         $ref: '#/definitions/GetCabResponse'
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	GetCabPath = "/cabs/{cabID}"

	// swagger:operation DELETE /cabs/{cabID} Cab DeleteCabRequest
	// ---
	// summary: Delete an existing cab
	// description: Delete an existing cab
	// responses:
	//   "200":
	//     description: Ok
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	DeleteCabPath = "/cabs/{cabID:[0-9+]}"

	// swagger:operation PUT /cabs/{cabID}/changecity Cab ChangeCityRequest
	// ---
	// summary: Change city of a cab
	// description: Change city of a cab
	// responses:
	//   "200":
	//     description: Ok
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	ChangeCityPath = "/cabs/{cabID:[0-9+]}/changecity"
)
