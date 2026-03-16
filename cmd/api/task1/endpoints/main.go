package main

import (
	"go-activities/cmd/api/task1"
	"net/http"
)

func handlePOSTFacilities(w http.ResponseWriter, r *http.Request) {
	// Implement your POST /facilities handler logic here
	w.WriteHeader(http.StatusNotImplemented)
}

func handlePOSTBookings(w http.ResponseWriter, r *http.Request) {
	// Implement your POST /bookings handler logic here
	w.WriteHeader(http.StatusNotImplemented)
}

func handleGETBookingsByID(w http.ResponseWriter, r *http.Request) {
	// Implement your GET /bookings?facility_id={ID} handler logic here
	w.WriteHeader(http.StatusNotImplemented)
}

func handleDELETEBooking(w http.ResponseWriter, r *http.Request) {
	// Implement your DELETE /bookings/{ID} handler logic here
	w.WriteHeader(http.StatusNotImplemented)
}

func main() {
	mux := http.NewServeMux()

	RegisterHandlers(mux)

	http.ListenAndServe(":8081", mux)
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/facilities", task1.GetFacilities)
	mux.HandleFunc("/facilities", handlePOSTFacilities) // Placeholder for POST /facilities handler
	mux.HandleFunc("/bookings", handlePOSTBookings)     // Placeholder for POST /bookings handler
	mux.HandleFunc("/bookings", handleGETBookingsByID)  // Placeholder for GET /bookings?facility_id={ID} handler
	mux.HandleFunc("/bookings/", handleDELETEBooking)   // Placeholder for DELETE /bookings/{ID} handler
}
