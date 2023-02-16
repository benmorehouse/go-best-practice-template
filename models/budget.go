package models

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	dynamo "github.com/guregu/dynamo"
)

// BudgetRepo will interface what the budget repo does to save stuff
type BudgetRepo interface {
	Get(userID string) (Budget, error) // Put request
}

type Budget struct {
	UserID    string `dynamo:"user_id" json:"user_id"`       // Hash key, a.k.a. partition key
	CreatedAt int    `dynamo:"created_at" json:"created_at"` // Range key, a.k.a. sort key

	Food     int `dynamo:"food" json:"food""`
	Health   int `dynamo:"health" json:"health""`
	Shopping int `dynamo:"shopping" json:"shopping""`
	Travel   int `dynamo:"travel" json:"travel""`
	Other    int `dynamo:"everything_else" json:"other""`
}

type baseBudgetRepo struct {
	table dynamo.Table
}

// BudgetTableName is the name of the budget table in Dynamo
const BudgetTableName = "budget-table"

// Get will give us back the latest budget which was set for this user
func (b *baseBudgetRepo) Get(userID string) (Budget, error) {
	budget := Budget{}

	if err := b.table.Get("user_id", userID).
		Range("created_at", dynamo.Greater, 0).One(&budget); err != nil {
		log.Println("budget_table_find_fail: %w", err)
		return budget, err
	}

	return budget, nil
}

// NewBudgetRepo will return a repo model handler
func NewBudgetRepo(awsRegion, dynamoHost string) BudgetRepo {
	db := dynamo.New(session.New(), &aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(dynamoHost),
	})
	return &baseBudgetRepo{table: db.Table(BudgetTableName)}
}
