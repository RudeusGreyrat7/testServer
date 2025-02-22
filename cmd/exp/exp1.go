package main

import (
	"fmt"

	"github.com/LENSLOCKED/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// docker compose up -d
// docker exec -it lenslocked-db-1 /usr/bin/psql -U mat -d lenslocked
// docker ps

// go get github.com/jackc/pgx/v5/pgxpool@v5.7.1

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("jon7@snow.com", "123snow")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id SERIAL PRIMARY KEY,
	// 		name TEXT,
	// 		email TEXT UNIQUE NOT NULL
	// 	);

	// 	CREATE TABLE IF NOT EXISTS orders(
	// 		id SERIAL PRIMARY KEY,
	// 		user_id INT NOT NULL,
	// 		amount INT,
	// 		description TEXT
	// 	);
	// `)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Tables created")

	// name := "Jon Ex"
	// email := "jon@snow.com"
	// _, err = db.Exec(`
	// INSERT INTO users (name, email)
	// VALUES ($1, $2);`, name, email)
	// VALUES ('Jon Snow', 'jon@snow.com') - можно вставить вместо VALUES...
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. Id =", id)

	// id := 1
	// row := db.QueryRow(`
	// SELECT name, email
	// FROM users
	// WHERE id=$1`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name - %s | email - %s\n", name, email)

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err = db.Exec(`
	// 	INSERT INTO orders (user_id, amount, description)
	// 	VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println("Create fake orders.")

	// type Order struct {
	// 	ID          int
	// 	userID      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order
	// userID := 1
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
	// 	order.userID = userID
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
	// fmt.Println("Orders:", orders)

	// !!!
	// подключаем ORM, используя gorm
	/*
		ORM (Object-Relational Mapping) — это метод программирования, который позволяет разработчикам взаимодействовать с базами данных, используя объектно-ориентированные подходы. Основная идея ORM заключается в том, чтобы предоставить способ работы с данными в базе как с объектами, а не через традиционные SQL-запросы. Это упрощает работу с данными и абстрагирует разработчика от деталей взаимодействия с базой данных.

		Для вставки данных в базу на языке Go без использования ORM, может потребоваться написать SQL-запрос вручную:
		db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
		С ORM это можно сделать на уровне работы с объектами:
		user := User{Name: "John", Age: 30}
		db.Create(&user)
	*/

}
