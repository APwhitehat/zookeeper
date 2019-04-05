package bimock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type sensorData struct {
	RideID string
	Lat    float64
	Lon    float64
}

func readData(r *http.Request, s *Server, rideID string, collection *(mongo.Collection)) {
	// TODO
	decoder := json.NewDecoder(r.Body)
	var data sensorData
	err := decoder.Decode(&data)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse JSON")
	}
	dataToInsert := insData{
		TimeStamp: int64(time.Now().Unix()),
		Lat:       data.Lat,
		Lon:       data.Lon}
	fmt.Println(data)
	insertToDB(collection, dataToInsert)
	dist := getDistance(collection)
	fmt.Println("Distance (in kms): ", dist)
	//TODO: Use this data generated by simulator to store in database or do some operation
}