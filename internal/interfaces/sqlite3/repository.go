package sqlite3

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/config"
	"github.com/tksasha/balance/internal/interfaces/sqlite3/category"
	"github.com/tksasha/balance/internal/interfaces/sqlite3/item"
	"github.com/tksasha/balance/internal/models"
)

type repository struct {
	db *sql.DB
}

func Open(config *config.Config) *repository {
	db, err := sql.Open("sqlite3", config.SQLite3Config.DB)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal(err)
	}

	return &repository{db}
}

func (repo *repository) NewItemRepository() models.ItemRepository {
	return item.NewRepository(repo.db)
}

func (repo *repository) NewCategoryRepository() models.CategoryRepository {
	return category.NewRepository(repo.db)
}
