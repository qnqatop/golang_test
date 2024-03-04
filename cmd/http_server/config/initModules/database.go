package initModules

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func (cfg *Config) initPostgresSql() {
	connStr := "host=localhost port=5432 user=user password=userPass dbname=base sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	cfg.db = db
	cfg.databasePostgresMigration()
}

func (cfg *Config) GetPostgresDb() *sql.DB {
	return cfg.db
}

func (cfg *Config) databasePostgresMigration() {
	createTablesQuery := []string{
		`
  CREATE TABLE IF NOT EXISTS project (
   id SERIAL PRIMARY KEY,
   name VARCHAR(100),
   created_at TIMESTAMP
  );
  `,
		`
  CREATE TABLE IF NOT EXISTS goods (
   id SERIAL PRIMARY KEY,
   project_id INT,
   name VARCHAR(100),
   description VARCHAR(100),
   priority INT,
   removed BOOLEAN,
   created_at TIMESTAMP,
   FOREIGN KEY (project_id) REFERENCES project(id)
  );
  `,
	}

	for _, query := range createTablesQuery {
		_, err := cfg.db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	insertRowToProject := "INSERT INTO project (name, created_at) VALUES ($1, $2)"
	_, err := cfg.db.Exec(insertRowToProject, "Первая запись", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	insertRowToGoods := "INSERT INTO goods (name, project_id, description, removed, created_at,priority) VALUES ($1, $2, $3, $4, $5,$6)"
	_, err = cfg.db.Exec(insertRowToGoods, "Первая запись Goods", 1, "Migration Description", false, time.Now(), 1)
	if err != nil {
		log.Fatal(err)
	}
}
