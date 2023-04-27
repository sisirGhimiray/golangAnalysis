package db

import (
	"context"
	"encoding/json"
	"os"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func ConnectToDatabaseAddFile(fileName string,fileSymbol string){

clientOptions:=options.Client().ApplyURI("mongodb://192.168.100.239")

client,err:=mongo.Connect(context.Background(),clientOptions)

if err!=nil{
	panic(err)
}
db:=client.Database("stockDatabase")

fileContent,err:=os.ReadFile(fileName)

if err!=nil{
	panic(err)
}

if fileSymbol=="pre"{

	var preStockData []interface{}
	err=json.Unmarshal(fileContent,&preStockData)
	if err!=nil{
		panic(err)
	}
	premarketCollection:=db.Collection("premarketstocks")
	_,err=premarketCollection.DeleteMany(context.Background(),bson.M{})

	_,err=premarketCollection.InsertMany(context.Background(),preStockData)
	if err!=nil{
		panic(err)
	}

	err=client.Disconnect(context.Background())
}

if fileSymbol=="nif"{
	var nifStockData []interface{}
	err=json.Unmarshal(fileContent,&nifStockData)
	if err!=nil{
		panic(err)
	}
	niftyMarketCollection:=db.Collection("niftymarketstocks")
	_,err=niftyMarketCollection.DeleteMany(context.Background(),bson.M{})

	_,err=niftyMarketCollection.InsertMany(context.Background(),nifStockData)
	if err!=nil{
		panic(err)
	}

	err=client.Disconnect(context.Background())


}



}
