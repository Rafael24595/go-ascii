package repository

import (
	"time"
	"strconv"
	"strings"
	"context"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/log"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/commons/constants/log-categories"
)

const CommandRepositoryMongoKey = "CommandRepositoryMongo"

type CommandRepositoryMongo struct {
	queryRepository QueryRepository
	collection mongo.Collection
}

func NewCommandRepositoryMongo(queryRepository QueryRepository, args map[string]string) CommandRepository {
	connection := getConnectionUri(args)
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

func getConnectionUri(args map[string]string) string {
	user := args["ASCII_MONGODB_USERNAME"]
	password := args["ASCII_MONGODB_PASSWORD"]
	server := args["ASCII_MONGODB_SERVER"]
	port := args["ASCII_MONGODB_PORT"]

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

func (this CommandRepositoryMongo) DependencyName() string {
	return CommandRepositoryMongoKey
}

func (this CommandRepositoryMongo) OnLoad() bool {
	log.Log(log_categories.INFO, "Initializing \"" + this.DependencyName() + "\" dependency...")
	this.fillQuery()
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was initialized successfully.")
	return true
}

func (this CommandRepositoryMongo) fillQuery() {
	log.Log(log_categories.INFO, "[" + this.DependencyName() + "] => Loading registries from data base...")
	cursor, err := this.collection.Find(context.TODO(), bson.M{"status": bson.M{ "$ne": request_state.DELETED }})
	if err != nil {
        panic(err)
    }
	cont := 0
	for cursor.Next(context.TODO()) {
        var dto dto.AsciiResponse
        err := cursor.Decode(&dto)
        if err != nil {
            panic(err)
        }

		timestamp := time.Unix(0, int64(dto.Timestamp) * int64(time.Millisecond))
		image := ascii.NewImageAscii(dto.Name, dto.Extension, dto.Status, timestamp, this.decodeFrames(dto))

        this.ToQuery(image)
		cont++
    }
	log.Log(log_categories.INFO, "[" + this.DependencyName() + "] => Finished: " + strconv.Itoa(cont) + " registries loaded.")
}

func (this CommandRepositoryMongo) OnExit() bool {
	log.Log(log_categories.INFO, "Exiting \"" + this.DependencyName() + "\" dependency...")
	this.cleanDeleted()
	log.Log(log_categories.INFO, "\"" + this.DependencyName() + "\" dependency was exited successfully.")
	return true
}

func (this CommandRepositoryMongo) cleanDeleted() {
	_, err := this.collection.DeleteMany(context.TODO(), bson.M{"status": request_state.DELETED})
	if err != nil {
        panic(err)
    }
	log.Log(log_categories.INFO, "Registries with \"" + request_state.DELETED + "\" status stored in database was deleted succesfully.")
}

func (this *CommandRepositoryMongo) Insert(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), request_state.STORED, image.GetTimestamp().UnixMilli(), this.encodeFrames(image))
	_, err := this.collection.InsertOne(context.Background(), response)
	if err != nil { 
		panic(err)
	}
	this.ToQuery(image)
	return image.GetName()
}

func (this *CommandRepositoryMongo) Modify(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), image.GetStatus(), image.GetTimestamp().UnixMilli(), this.encodeFrames(image))
	filter := bson.M{"name": image.GetName()}
	_, err := this.collection.ReplaceOne(context.Background(), filter, response)
	if err != nil { 
		panic(err)
	}
	this.ToQuery(image)
	return image.GetName()
}

func (this *CommandRepositoryMongo) Delete(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), image.GetStatus(), image.GetTimestamp().UnixMilli(), this.encodeFrames(image))
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
