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

func main_pgx() {

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
	fmt.Printf("执行: %s\n", query)

	// 插入数据并返回ID
	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id;
	`, "test", "test@test.com")
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User inserted. id= %d\n", id)

	// 查询单条数据
	row = db.QueryRow(`
		SELECT name, email
		FROM users
		WHERE id=$1;
	`, id)

	var testName, testEmail string
	err = row.Scan(&testName, &testEmail)
	if err == sql.ErrNoRows {
		fmt.Println("Error, no rows!")
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("User information: name=%s, email=%s\n", testName, testEmail)

	// 创建订单
	userID := 1
	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
			INSERT INTO orders(user_id, amount, description)
			VALUES($1, $2, $3)
		`, userID, amount, desc)

		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Create fake orders.")

	// 查询多条数据
	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1
	`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}
	fmt.Println("Orders: ", orders)
}
