package handlers

import (
	graphqlconfig "brendisurfs/crystalshop-backend/graphql_config"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type RequestQuery struct {
	Color string
}

// CrystalHandler - handles our graphql crystal connection from the frontend.
func CrystalHandler() gin.HandlerFunc {
	//Need query params
	queryParams := c.Query("query")

	rootQuery := gql.ObjectConfig{Name: "RootQuery", Fields: graphqlconfig.Fields}
	schemaConfig := gql.SchemaConfig{
		Query: gql.NewObject(rootQuery),
	}
	schema, err := gql.NewSchema(schemaConfig)
	if err != nil {
		log.Println(err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	// set up our params to query our db.
	params := gql.Params{Schema: schema, RequestString: queryParams}
	req := gql.Do(params)
	if req.HasErrors() {
		log.Fatalf("failed to execute graphql Do operation, errors: %+v", req.Errors)
	}

	reqJSON, _ := json.Marshal(req)
	fmt.Printf("%s\n", reqJSON)
	
	return fun(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
