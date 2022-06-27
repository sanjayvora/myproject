package trip

// resource trip endpoints
const (
	// swagger:operation POST /trips Trip AddTripRequest
	// ---
	// summary: Add Booking
	// description: Add new booking
	// responses:
	//   "200":
	//     description: New trip ID
	//     schema:
	//         $ref: '#/definitions/AddTripResponse'
	//   "400":
	//     description: Bad Request
	//   "500":
	//     description: The request was not processed due to an server error.
	AddTripPath = "/trips"

	// swagger:operation PUT /trips/{tripID}/start Trip StartTripRequest
	// ---
	// summary: Start a trip
	// description: Start a given trip
	// responses:
	//   "200":
	//     description: Ok
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	StartTripPath = "/trips/{tripID:[0-9+]}/start"

	// swagger:operation PUT /trips/{tripID}/end Trip EndTripRequest
	// ---
	// summary: End a trip
	// description: End a given trip
	// responses:
	//   "200":
	//     description: Ok
	//   "400":
	//     description: Bad Request
	//   "404":
	//     description: Resource not found
	//   "500":
	//     description: The request was not processed due to an server error.
	EndTripPath = "/trips/{tripID:[0-9+]}/end"
)
