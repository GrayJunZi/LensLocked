package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (config PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.Database, config.SSLMode)
}

func main() {

	config := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "root",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	// 连接数据库
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id serial PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// 插入数据
	name := "admin"
	email := "admin@email.com"
	_, err = db.Exec(`
		INSERT INTO users (name, email)
		VALUES ($1, $2);
	`, name, email)

	if err != nil {
		panic(err)
	}
	fmt.Println("User created.")

	// SQL 注入
	name = "',''); DROP TABLE users; --"
	query := fmt.Sprintf(`
		INSERT INTO users (name, email)
		VALUES ('%s', '%s')
	`, name, email)
	fmt.Printf("执行: %s", query)
}
