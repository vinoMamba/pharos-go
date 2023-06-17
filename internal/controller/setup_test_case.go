package controller

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/pharos-go/config"
	"github.com/vinoMamba/pharos-go/config/queries"
	"github.com/vinoMamba/pharos-go/internal/database"
)

var (
	r *gin.Engine
	q *queries.Queries
	c context.Context
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	r = gin.New()
	config.LoadAppConfig()
	database.Connect()

	q = database.NewQuery()
	c = context.Background()

	// if err := q.DeleteAllUsers(c); err != nil {
	// 	t.Fatal(err)
	// }
	//
	// if err := q.DeleteAllValidationCodes(c); err != nil {
	// 	t.Fatal(err)
	// }

	return func(t *testing.T) {
	}
}
