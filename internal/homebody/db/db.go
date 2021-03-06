package db

import (
	"github.com/Gateway/pkg/database"
	"github.com/Gateway/pkg/serializer"
)

type DB struct {
	Client     database.DataBase
	serializer serializer.Serializer
}

func Init(dbType, address string, serializer serializer.Serializer) (*DB, error) {
	client, err := database.Init(dbType, address)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client:     client,
		serializer: serializer,
	}, nil
}
