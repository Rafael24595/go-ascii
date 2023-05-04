package repository

import (
	"context"
	"encoding/base64"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/dto"
	"go-ascii/src/domain/ascii"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommandRepositoryMongo struct {
	queryRepository QueryRepository
	collection mongo.Collection
}

func NewCommandRepositoryMongo(queryRepository QueryRepository) CommandRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongodb:27017"))
	if err != nil { 
		panic(err)
	}
	collection := client.Database("go-ascii").Collection("ascii")
	return &CommandRepositoryMongo{queryRepository: queryRepository, collection: *collection}
}

func (this CommandRepositoryMongo) OnLoad() bool {
	this.fillQuery()
	return true
}

func (this CommandRepositoryMongo) fillQuery() {
	cursor, err := this.collection.Find(context.TODO(), bson.D{})
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

        this.InsertQuery(image)
    }
}

func (this CommandRepositoryMongo) OnExit() bool {
	return true
}

func (this *CommandRepositoryMongo) InsertAscii(image ascii.ImageAscii) string {
	response := dto.NewAsciiResponse(image.GetName(), image.GetExtension(), request_state.STORED, this.encodeFrames(image))
	_, err := this.collection.InsertOne(context.Background(), response)
	if err != nil { 
		panic(err)
	}
	this.InsertQuery(image)
	return image.GetName()
}

func (this CommandRepositoryMongo) InsertQuery(image ascii.ImageAscii) {
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
