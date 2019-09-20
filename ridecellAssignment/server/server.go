package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"example.com/ridecellAssignment/handlers"
	"encoding/json"
	"log"
	"strconv"
	"example.com/ridecellAssignment/localdb/entities"
	"example.com/ridecellAssignment/handlers/requests"
)

func StartServer()  {
	 port :=":9900"
	router := httprouter.New()
	router.GET("/api/v1/:cityId/:streetId/parkingSpots",GetParkingSpot)
	router.GET("/api/v1/:cityId/:streetId/nearByParkingSpots",GetNearByParkingSpots)
	router.POST("/api/v1/:cityId/:streetId/:spotId/reserveParking",PostReserveParking)
	router.POST("/api/v1/:cityId/:streetId/:spotId/cancelParking",PostCancelParking)
	log.Printf("Server started and listening at %s",port)
	http.ListenAndServe(port,router)
}


func GetParkingSpot(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityId := params.ByName("cityId")
	streetId := params.ByName("streetId")
	status := request.URL.Query().Get("status")
	spots, err := handlers.HandleGetParkingSpots(cityId, streetId,status)
	writer.Header().Set("Content-Type","application/json")
	if err !=nil{
		writer.WriteHeader(err.HTTPCode)
		json.NewEncoder(writer).Encode(spots)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(spots)
}


func GetNearByParkingSpots(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityId := params.ByName("cityId")
	streetId := params.ByName("streetId")
	lat := request.URL.Query().Get("lat")
	latF, e := strconv.ParseFloat(lat, 64)
	writer.Header().Set("Content-Type","application/json")
	if e!=nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(entities.Error{
			Code:"PS-23",
			Message:"Invalid lat value",
		})
	}
	lng := request.URL.Query().Get("lng")
	lngF, e := strconv.ParseFloat(lng, 64)
	if e!=nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(entities.Error{
			Code:"PS-23",
			Message:"Invalid lng value",
		})
	}
	radius := request.URL.Query().Get("radius")
	radiusF, e := strconv.ParseFloat(radius, 64)
	if e!=nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(entities.Error{
			Code:"PS-23",
			Message:"Invalid radius value",
		})
	}
	spots, err := handlers.HandleGetNearByParkingSpots(cityId,streetId,latF,lngF,radiusF)

	if err !=nil{
		writer.WriteHeader(err.HTTPCode)
		json.NewEncoder(writer).Encode(spots)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(spots)
}

func PostReserveParking(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityId := params.ByName("cityId")
	streetId := params.ByName("streetId")
	parkingSpotId := params.ByName("spotId")
	decoder := json.NewDecoder(request.Body)
	spotRequest := &requests.ReserveParkingSpotRequest{}
	decodeErr := decoder.Decode(spotRequest)
	if decodeErr!=nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(entities.Error{Code:"PS-34",Message:decodeErr.Error()})
		return
	}
	writer.Header().Set("Content-Type","application/json")
	response, err := handlers.HandleReserveParkingSpot(cityId, streetId, parkingSpotId)
	if err !=nil{
		writer.WriteHeader(err.HTTPCode)
		json.NewEncoder(writer).Encode(err)
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func PostCancelParking(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityId := params.ByName("cityId")
	streetId := params.ByName("streetId")
	parkingSpotId := params.ByName("spotId")
	writer.Header().Set("Content-Type","application/json")
	response, err := handlers.HandleCancelParkingSpot(cityId, streetId, parkingSpotId)
	if err !=nil{
		writer.WriteHeader(err.HTTPCode)
		json.NewEncoder(writer).Encode(err)
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}