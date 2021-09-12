package utils

import (
	"github.com/graphql-go/graphql"
	"github.com/isakgranqvist2021/cashresp/models"
)

var Schema *graphql.Schema

func GQLSchema() (*graphql.Schema, error) {
	fields := graphql.Fields{
		"posts": models.PostModel(),
		"user":  models.UserModel(),
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func InitQraphQL() error {
	schema, err := GQLSchema()

	if err != nil {
		return err
	}

	Schema = schema
	return nil
}
