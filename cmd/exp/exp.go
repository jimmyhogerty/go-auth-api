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

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	// us := models.UserService{
	// 	DB: db,
	// }

	// user, err := us.Create("Jimbo@jimbo.com", "onetwothree")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)

	// // Create a table...
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id SERIAL PRIMARY KEY,
	// 		name TEXT,
	// 		email TEXT UNIQUE NOT NULL
	// 	);

	// 	CREATE TABLE IF NOT EXISTS orders (
	// 		id SERIAL PRIMARY KEY,
	// 		user_id INT NOT NULL,
	// 		amount INT,
	// 		description TEXT
	// 	);
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables created.")

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS tweets (
	// 		id SERIAL PRIMARY KEY,
	// 		user_id INT NOT NULL,
	// 		tweet TEXT
	// 	);
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tweet table created.")

	// Insert some data...
	// name := "Jimbo2"
	// email := "jslice2@jimbo.com"
	// row := db.QueryRow(`
	// INSERT INTO users(name, email)
	// VALUES ($1, $2) RETURNING id;`, name, email)
	// row.Err()
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User inserted. id = ", id)

	// id := 1
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id=$1`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error, no rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	tweet := fmt.Sprintf("I am tweet number #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO tweets(user_id, tweet)
	// 		VALUES($1, $2)`, userID, tweet)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	// type Tweet struct {
	// 	ID     int
	// 	UserID int
	// 	tweet  string
	// }

	// var tweets []Tweet
	// rows, err := db.Query(`
	// 	SELECT * FROM tweets
	// 	WHERE user_id=$1`, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var tweet Tweet
	// 	err := rows.Scan(&tweet.UserID, &tweet.ID, &tweet.tweet)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	tweets = append(tweets, tweet)
	// }

	// type Order struct {
	// 	ID          int
	// 	UserID      int
	// 	Amount      int
	// 	Description string
	// }
	// var orders []Order
	// rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id=$1`, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserID = userID
	// 	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Tweets:", tweets)
}
