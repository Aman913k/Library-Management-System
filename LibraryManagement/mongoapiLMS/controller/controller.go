package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/Aman913k/mongoapiLMS/model"

	"github.com/Aman913k/MONGOAPILMS/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://amanrana9133:aman1n1@cluster0.6ohqxgo.mongodb.net/?retryWrites=true&w=majority"
const dbName = "library"
const colName = "booklist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection is success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}

func insertOneBook(book model.Library) {
	fmt.Println("Inserting book:", book)
	inserted, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id", inserted.InsertedID)
}

func updateOneBook(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"available": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneBook(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book got delete with delete count: ", deleteCount)
}

func deleteAllBook() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of books delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllBooks() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var books []primitive.M

	for cur.Next(context.Background()) {
		var book bson.M
		err := cur.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	defer cur.Close(context.Background())
	return books
}

// Actual Controller - file

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allBooks := getAllBooks()
	json.NewEncoder(w).Encode(allBooks)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var book model.Library
	_ = json.NewDecoder(r.Body).Decode(&book)
	insertOneBook(book)
	json.NewEncoder(w).Encode(book)
}

func MarkAsAvailable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneBook(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "Delete")

	params := mux.Vars(r)
	deleteOneBook(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methos", "Delete")

	count := deleteAllBook()
	json.NewEncoder(w).Encode(count)
}
