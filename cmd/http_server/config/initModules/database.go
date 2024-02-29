package initModules

import (
	"database/sql"
	"log"
	"time"
)

func (cfg *Config) initPostgresSql() {
	connStr := "host=localhost port=5432 user=your_username password=your_password dbname=your_database sslmode=disable"

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
	createTablesQuery := [2]string{
		`
		CREATE TABLE IF NOT EXISTS project (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			create_at TIMESTAMP
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
            create_at TIMESTAMP
        );
    `}

	for _, query := range createTablesQuery {
		_, err := cfg.db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	insertRowToProduct := `INSERT INTO project (name, create_at) VALUES ($1, $2)`
	_, err := cfg.db.Exec(insertRowToProduct, "Migration Project", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	insertRowToGoods := `INSERT INTO goods (name,project_id,description,removed,create_at) VALUES ($1, $2,$3,$4,$5)`
	_, err = cfg.db.Exec(insertRowToGoods, "Migration Goods", 1, "Migration Description", false, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
