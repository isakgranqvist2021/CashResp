package models

import "github.com/graphql-go/graphql"

func PostModel() *graphql.Field {
	posts := []map[string]interface{}{
		{
			"id":          1,
			"title":       "sometimes by accident",
			"description": "aining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more rece",
		},
		{
			"id":          2,
			"title":       "Finibus Bonorum",
			"description": "Lorem Ipsum is simply dummy text of the printing and typesetting indus",
		},
		{
			"id":          3,
			"title":       "anything embarrassing",
			"description": "re-or-less normal distribution of letters, as opposed to using",
		},
		{
			"id":          4,
			"title":       "te Lorem Ipsum which",
			"description": "e majority have suffered alteration in some form, by injected humour, or randomised words",
		},
		{
			"id":          5,
			"title":       "Richard McClintock",
			"description": " words, combined with a handful of model sentence structures, to g",
		},
	}

	postsType := graphql.NewObject(graphql.ObjectConfig{
		Name: "posts",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"title":       &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
		},
	})

	return &graphql.Field{
		Type:        graphql.NewList(postsType),
		Description: "Get posts",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return posts, nil
		},
	}
}
