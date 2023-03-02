package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/paihari/vending-machine-golang-graphql/base"
	"github.com/paihari/vending-machine-golang-graphql/graph"


	// "github.com/99designs/gqlgen/graphql"
	// "github.com/99designs/gqlgen/graphql/executor"
	
	
	// "github.com/99designs/gqlgen/graphql/resolver"
	// "github.com/99designs/gqlgen/graphql/schema"
	// "github.com/vektah/gqlparser/v2/ast"
	// "github.com/vektah/gqlparser/v2/validator"

)
// type ExportDirective struct{}

// func (d *ExportDirective) DirectiveName() string {
// 	return "export"
// }

// func (d *ExportDirective) VisitFieldDefinition(field *ast.FieldDefinition, _ *ast.Schema, _ interface{}) validator.NextCheck {
// 	// Add custom resolver for the field
// 	oldResolve := field.Resolve
// 	field.Resolve = func(ctx context.Context, obj interface{}, execInfo graphql.ResolveInfo, responseKey string) (interface{}, error) {
// 	  // Call the original resolver function
// 	  value, err := oldResolve(ctx, obj, execInfo, responseKey)
// 	  if err != nil {
// 		return nil, err
// 	  }
  
// 	  // Export the result to the context under the specified name
// 	  exportName := field.Directives.ForName("export").Arguments[0].Value.Raw
// 	  graphql.AddResultContext(ctx, exportName, value)
  
// 	  // Return the original value
// 	  return value, nil
// 	}
  
// 	return validator.ContinueCheck
//   }



const defaultPort = "8080"

func main() {

 // Define the export directive and add it to your schema
	// cfg := schema.Config{
	// 	Directives: schema.DirectiveRoot{
	// 	Export: &ExportDirective{},
	// 	},
	// }


	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := godotenv.Load("./../.env"); if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	
	Database := base.Connect()

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{
				DB: Database}}))

	   
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
