package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"google.golang.org/api/option"
)

// Firestore client
var client *firestore.Client

// Initialize Firestore client
func initFirestore() {
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	sa := option.WithCredentialsFile("data/sa.json")

	var err error
	client, err = firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
}

// GraphQL schema and resolver
func setupGraphQLSchema() graphql.Schema {
	// Define a Firestore document type
	documentType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Document",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"field1": &graphql.Field{
				Type: graphql.String,
			},
			"field2": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// Define the query type
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"document": &graphql.Field{
				Type:        documentType,
				Description: "Get a document by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}

					doc, err := client.Collection("tfmv").Doc(id).Get(context.Background())
					if err != nil {
						return nil, err
					}

					data := doc.Data()
					data["id"] = doc.Ref.ID
					return data, nil
				},
			},
		},
	})

	// Create the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	return schema
}

func main() {
	// Initialize Firestore
	initFirestore()

	// Set up Gin router
	r := gin.Default()

	// Set up GraphQL schema
	schema := setupGraphQLSchema()

	// Define the GraphQL handler
	r.POST("/graphql", func(c *gin.Context) {
		var params struct {
			Query string `json:"query"`
		}
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})

		if len(result.Errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": result.Errors})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
