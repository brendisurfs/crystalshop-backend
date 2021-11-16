/*
TODO:
Implement with smart contract, so that the actual ownership isnt being stored in the db.


**/
package graphqlconfig

import gql "github.com/graphql-go/graphql"

type Crystal struct {
	ID      int
	Title   string
	Creator string
	Owner   string
}

type Creator struct {
	Name     string
	Crystals []int
}
type Owner struct {
	Name       string
	Collection []int
}

var crystals []Crystal

// CrystalType - GraphQL version of the Crystal struct.
var CrystalType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Crystal",
		Fields: gql.Fields{
			"title": &gql.Field{
				Type: gql.String,
			},
			"creator": &gql.Field{
				Type: gql.String,
			},
			"owner": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

// CreatorType - GraphQL version of the Creator struct.
// TODO: read top note.
var CreatorType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Creator",
		Fields: gql.Fields{
			"name": &gql.Field{
				Type: gql.String,
			},
			"crystals": &gql.Field{
				Type: gql.NewList(gql.Int),
			},
		},
	},
)

// OwnerType GraphQL version of Owner struct.
// TODO: read top note.
var OwnerType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Owner",
		Fields: gql.Fields{
			"name": &gql.Field{
				Type: gql.String,
			},
			"Collection": &gql.Field{
				Type: gql.NewList(gql.Int),
			},
		},
	},
)

// TODO: work on retrieving crystals from Supabase.
var Fields = gql.Fields{
	"crystal": &gql.Field{
		Type:        CrystalType,
		Description: "Get crystal by id",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.Int,
			},
			"title": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
	},
	"list": &gql.Field{
		Type:        gql.NewList(CrystalType),
		Description: "get all crystals",
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			// fill this section with supabase stuff.
			var crystal Crystal
			crystals = append(crystals, crystal)
			return crystals, nil
		},
	},
}
