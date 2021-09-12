package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/isakgranqvist2021/cashresp/utils"
)

func GraphQL(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		log.Fatal(err)
		return c.JSON("an error occured while parsing graphql query")
	}

	r := graphql.Do(graphql.Params{
		Schema:        *utils.Schema,
		RequestString: fmt.Sprintf("%v", data["query"]),
	})

	if len(r.Errors) > 0 {
		return c.JSON(r.Errors[len(r.Errors)-1])
	}

	return c.JSON(r)
}
