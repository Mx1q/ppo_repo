package tests

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testDbInstance *pgxpool.Pool

func TestMain(m *testing.M) {
	testDB := SetupTestDatabase(DbName)
	defer testDB.TearDown()
	testDbInstance = testDB.DbInstance
	err := ExecuteSQLsFromDir(testDbInstance, TestDataDir)
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(m.Run())
}
