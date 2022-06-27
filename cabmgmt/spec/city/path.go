package city

// resource city endpoints
const (
	// swagger:operation POST /cities City AddCityRequest
	// ---
	// summary: Add new city
	// description: Add new city with details
	// responses:
	//   "200":
	//     description: New city ID
	//     schema:
	//         $ref: '#/definitions/AddCityResponse'
	//   "400":
	//     description: Bad Request
	//   "500":
	//     description: The request was not processed due to an server error.
	AddCityPath = "/cities"

	// swagger:operation GET /cities City GetCityRequest
	// ---
	// summary: Add new city
	// description: Add new city with details
	// responses:
	//   "200":
	//     description: New city ID
	//     schema:
	//         $ref: '#/definitions/GetCityResponse'
	//   "400":
	//     description: Bad Request
	//   "500":
	//     description: The request was not processed due to an server error.
	GetCityPath = "/cities"

	// swagger:operation DELETE /cities/{cityID} City DeleteCityRequest
	// ---
	// summary: Delete an existing city
	// description: Delete an existing city. City can only be deleted if there are no more cabs present in the city
	// responses:
	//   "200":
	//     description: Ok
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	DeleteCityPath = "/cities/{cityID:[0-9+]}"

	// swagger:operation GET /cities/demand City GetDemandRequest
	// ---
	// summary: Get high demand cities
	// description: Get high demand cities and peek time.
	// responses:
	//   "200":
	//     description: City details
	//     schema:
	//         $ref: '#/definitions/DemandResponse'
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	GetDemandPath = "/cities/demand"
)
