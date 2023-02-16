package api

import (
	"context"
	"net/http"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda" // https://pkg.go.dev/github.com/aws/aws-lambda-go
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	config "budget-service/config"
	models "budget-service/models"
)

// Server embodies all the central application structure data.
type Server struct {
	// Config holds the configuration of the application
	Config *config.Configuration

	// GinLambda is our gin adapter to allow lambda and gin
	// to work together in unison.
	GinLambda *ginadapter.GinLambda

	// Router is the router we run our server through
	Router *gin.Engine

	BudgetTable models.BudgetRepo
}

// MainLambdaHandler is our main handler, which links initial requests to
// our lambda function from our gin engine.
func (s *Server) MainLambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return s.GinLambda.ProxyWithContext(ctx, req)
}

// Status will handle a request to append logs from a user's app store
func (s *Server) Status(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Ok bool
	}{
		true,
	})
	return
}

func (s *Server) initRouter() {
	router := gin.Default()
	v1 := router.Group("")
	v1.GET("/status", s.Status)
	v1.GET("/users/:id/budget", s.GetBudget)
	v1.POST("/users/:id/budget", s.PostBudget)

	s.Router = router
}

// Start will start the remote server using our gin adapter
func Start() {
	conf := config.New()
	s := Server{
		Config:      conf,
		BudgetTable: models.NewBudgetRepo(conf.AWSRegion, conf.DynamoHost),
	}
	s.initRouter()

	s.GinLambda = ginadapter.New(s.Router)
	lambda.Start(s.MainLambdaHandler)
}
