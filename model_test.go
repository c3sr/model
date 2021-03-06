package model

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/c3sr/config"
	mongodb "github.com/c3sr/database/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	config.Init(
		config.AppName("rai"),
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}

func TestConnection(t *testing.T) {
	db, err := mongodb.NewDatabase(config.App.Name)
	assert.NoError(t, err)
	assert.NotEmpty(t, db)
	defer db.Close()
}

func TestInsertRanking(t *testing.T) {
	db, err := mongodb.NewDatabase(config.App.Name)
	assert.NoError(t, err)
	assert.NotEmpty(t, db)
	defer db.Close()

	tbl, err := mongodb.NewTable(db, "test")
	assert.NoError(t, err)
	assert.NotEmpty(t, tbl)

	// tbl.Create(nil)

	err = tbl.Insert(Base{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
	})
	assert.NoError(t, err)

}
