package models

import "github.com/graphql-go/graphql"

func UserModel() *graphql.Field {
	users := []map[string]interface{}{
		{
			"id":   1,
			"name": "Isak",
			"age":  20,
		},
		{
			"id":   2,
			"name": "John",
			"age":  21,
		},
	}

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
			"age":  &graphql.Field{Type: graphql.Int},
		},
	})

	return &graphql.Field{
		Type:        userType,
		Description: "Get one user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)

			if !ok {
				return nil, nil
			}

			for i := 0; i < len(users); i++ {
				if users[i]["id"] == id {
					return users[i], nil
				}
			}

			return nil, nil
		},
	}
}
