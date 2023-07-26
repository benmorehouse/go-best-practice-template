package models

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamo "github.com/guregu/dynamo"
)

// ExampleRepo will interface what the example repo does to save stuff
type ExampleRepo interface {
	Get(userID string) (Example, error) // Put request
	Put(e Example) error                // Put request
}

// Example provisions an example model
type Example struct {
	UserID    string `dynamo:"user_id" json:"user_id"`       // Hash key, a.k.a. partition key
	CreatedAt int    `dynamo:"created_at" json:"created_at"` // Range key, a.k.a. sort key

	FieldName int `dynamo:"field_name" json:"field_name"`
}

type baseExampleRepo struct {
	table dynamo.Table
}

// ExampleTableName is the name of the example table in Dynamo
const ExampleTableName = "example-table"

// Get will give us back the latest example which was set for this user
func (b *baseExampleRepo) Get(userID string) (Example, error) {
	example := Example{}

	if err := b.table.Get("user_id", userID).
		Range("created_at", dynamo.Greater, 0).One(&example); err != nil {
		log.Println("example_table_find_fail: %w", err)
		return example, err
	}

	return example, nil
}

// Get will give us back the latest example which was set for this user
func (b *baseExampleRepo) Put(e Example) error {
	if err := b.table.Put(e).Run(); err != nil {
		log.Println("example_table_find_fail: %w", err)
		return err
	}

	return nil
}

// NewExampleRepo will return a repo model handler
func NewExampleRepo(awsRegion, dynamoHost string) ExampleRepo {
	db := dynamo.New(session.New(), &aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(dynamoHost),
	})
	return &baseExampleRepo{table: db.Table(ExampleTableName)}
}
