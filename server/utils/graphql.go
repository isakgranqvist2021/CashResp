package utils

import (
	"github.com/graphql-go/graphql"
)

var Schema *graphql.Schema

func GQLSchema() (*graphql.Schema, error) {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"posts": &graphql.Field{
			Type: &graphql.Scalar{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return []string{"hello", "world", "john", "doe"}, nil
			},
		},
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
