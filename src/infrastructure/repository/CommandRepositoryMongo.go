package repository

import (
	"os"
	"time"
	"strings"
	"context"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/domain/ascii"
)

type CommandRepositoryMongo struct {
	queryRepository QueryRepository
	collection mongo.Collection
}

func NewCommandRepositoryMongo(queryRepository QueryRepository) CommandRepository {
	connection := getConnectionUri()
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connection)
	client, err := mongo.Connect(ctx, options)
	if err != nil { 
		panic(err)
	}
	collection := client.Database("go-ascii").Collection("ascii")
	return &CommandRepositoryMongo{queryRepository: queryRepository, collection: *collection}
}

func getConnectionUri() string {
	user := os.Getenv("ASCII_MONGODB_USERNAME")
	password := os.Getenv("ASCII_MONGODB_PASSWORD")
	server := os.Getenv("ASCII_MONGODB_SERVER")
	port := os.Getenv("ASCII_MONGODB_PORT")

	var connection strings.Builder
	connection.WriteString("mongodb://")
	connection.WriteString(user)
	connection.WriteString(":")
	connection.WriteString(password)
	connection.WriteString("@")
	connection.WriteString(server)
	connection.WriteString(":")
	connection.WriteString(port)
	return connection.String()
}

func (this CommandRepositoryMongo) OnLoad() bool {
	this.fillQuery()
	return true
}

func (this CommandRepositoryMongo) fillQuery() {
	cursor, err := this.collection.Find(context.TODO(), bson.M{"status": bson.M{ "$ne": request_state.DELETED }})
	if err != nil {
        panic(err)
    }
	for cursor.Next(context.TODO()) {
        var dto dto.AsciiResponse
        err := cursor.Decode(&dto)
        if err != nil {
            panic(err)
        }

		image := ascii.NewImageAscii(dto.Name, dto.Extension, dto.Status, this.decodeFrames(dto))

        this.ToQuery(image)
    }
}

func (this CommandRepositoryMongo) OnExit() bool {
	this.cleanDeleted()
	return true
}

func (this CommandRepositoryMongo) cleanDeleted() {
	_, err := this.collection.DeleteMany(context.TODO(), bson.M{"status": request_state.DELETED})
	if err != nil {
        panic(err)
    }
}

func (this *CommandRepositoryMongo) Insert(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), request_state.STORED, this.encodeFrames(image))
	_, err := this.collection.InsertOne(context.Background(), response)
	if err != nil { 
		panic(err)
	}
	this.ToQuery(image)
	return image.GetName()
}

func (this *CommandRepositoryMongo) Delete(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), request_state.DELETED, this.encodeFrames(image))
	filter := bson.M{"name": image.GetName()}
	_, err := this.collection.ReplaceOne(context.Background(), filter, response)
	if err != nil { 
		panic(err)
	}
	this.ToQuery(image)
	return image.GetName()
}

func (this CommandRepositoryMongo) ToQuery(image ascii.ImageAscii) {
	this.queryRepository.InsertCommand(image)
}

func (this CommandRepositoryMongo) encodeFrames(image ascii.ImageAscii) (encodedFrames []string) {
	for _, frame := range image.GetFrames() {
		encode:= base64.StdEncoding.EncodeToString([]byte(frame))
		encodedFrames = append(encodedFrames, string(encode))
	}
	return
}

func (this CommandRepositoryMongo) decodeFrames(dto dto.AsciiResponse) (encodedFrames []string) {
	for _, frame := range dto.Frames {
		encode, err := base64.StdEncoding.DecodeString(frame)
		if err != nil { 
			panic(err)
		}
		encodedFrames = append(encodedFrames, string(encode))
	}
	return
}
