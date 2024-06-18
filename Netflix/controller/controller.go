package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb://localhost:27017"
const dbName = "godb"
const collectionName = "movies"

var collection *mongo.Collection // ref to collection to ensure actual write

// run very first time when application started
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)
	// connect to mongo
	// context given when working with different machine as db on different machine
	// context.Background() or context,TODO()
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	fmt.Println("connected to database")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance is ready")
}

// mongoDb Helpers
func insertOneMovie(movie models.Netflix) {
	insertResult, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	fmt.Println(updateResult)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	fmt.Println("Updated a Single Record ", updateResult)
}
func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted a Single Record ", deleteResult)
}
func deleteAllRecords() int64 {
	filter := bson.M{} // no value means select all objects
	deleteAllResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all Record ", deleteAllResult)
	return deleteAllResult.DeletedCount
}
func getAllMovies() []primitive.M {
	filter := bson.M{} // no value means select all objects
	allResult, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for allResult.Next(context.Background()) {
		var movie bson.M
		err := allResult.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer allResult.Close(context.Background())
	fmt.Println("get all Record ", allResult)
	return movies
}

// controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}
func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllRecords()
	json.NewEncoder(w).Encode(count)
}
