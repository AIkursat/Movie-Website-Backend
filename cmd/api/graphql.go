package main

import (
	"backend/models"
	"io"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// That's a GraphQL schema

var movies []*models.Movie

var fields = graphql.Fields{
	"movie": &graphql.Field{
		Type: movieType,
		Description: "Get movie by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error){
			id, ok := p.Args["id"].(int)
			if ok{
				for _, movie := range movies {
					if movie.ID == id {
						return movie, nil 
					}	
				}
			}
			return nil, nil
		},
	},
     "list" : &graphql.Field{
		Type : graphql.NewList(movieType),
		Description: "Get all movies",
		Resolve: func(params graphql.ResolveParams) (interface{}, error){
			return movies, nil
		},
	 },
}

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
          "id": &graphql.Field{
			Type: graphql.Int,
		  },
		  "title": &graphql.Field{
			Type: graphql.String,
		  },
		  "description": &graphql.Field{
			Type: graphql.String,
		  },
		  "year": &graphql.Field{
			Type: graphql.Int,
		  },
		  "release_Date": &graphql.Field{
			Type: graphql.DateTime,
		  },
		  "runtime": &graphql.Field{
			Type: graphql.Int,
		  },
		  "rating": &graphql.Field{
			Type: graphql.Int,
		  },
		  "mpaa_rating": &graphql.Field{
			Type: graphql.String,
		  },
		  "created_at": &graphql.Field{
			Type: graphql.DateTime,
		  },
		  "updated_at": &graphql.Field{
			Type: graphql.DateTime,
		  },
		}, 
	},
)

func (app *application) moviesGraphQL(w http.ResponseWriter, r *http.Request){
    movies, _ = app.models.DB.All() // "_" means Ignore the error 


	 q, _ := io.ReadAll(r.Body) 
	 query := string(q)

	 log.Println(query)
}