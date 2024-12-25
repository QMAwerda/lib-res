package db

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/idoko/bucketeer/models"
)

func (db Database) GetAllOrders() (*models.OrderList, error) {
	list := &models.OrderList{}

	rows, err := db.Conn.Query("SELECT * FROM orders ORDER BY id DESC") // DESC - в порядке убывания
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.Id,
			&order.UserFullName,
			&order.Isbn,
			&order.Date,
		)
		if err != nil {
			return list, err
		}
		list.Orders = append(list.Orders, order)
	}

	return list, nil
}

// todo: нужно добавить, чтобы можно было получить айди ордера из таблицы ордера, по введенному фулнейму и isbn
// todo: т.е. на удаление заказа нужно будет брать эти два параметра и искать айди заказа и его в транзакции удалять

// Мне нужно удалить ордер по айди, но для этого нужно
func (db Database) DeleteOrder(orderID uint64) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Получаем ISBN из удаляемого заказа
	var isbn string
	query :=
		`SELECT isbn 
		 FROM orders 
		 WHERE id = $1`

	err = tx.QueryRow(query, orderID).Scan(&isbn)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Удаляем заказ
	deleteQuery :=
		`DELETE 
	 	 FROM orders 
	 	 WHERE id = $1`

	_, err = tx.Exec(deleteQuery, orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Обновляем количество книг в таблице books
	updateQuery :=
		`UPDATE books 
		 SET amount = amount + 1 
		 WHERE isbn = $1`

	_, err = tx.Exec(updateQuery, isbn)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db Database) AddOrder(order *models.Order) error {
	tx, err := db.Conn.Begin() // Начало транзакции
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Откат транзакции в случае паники
		}
	}()

	var createdAt time.Time
	updateBookQuery :=
		`UPDATE books 
		SET amount = amount - 1 
		WHERE isbn = $1`

	_, err = tx.Exec(updateBookQuery, order.Isbn)
	switch err {
	case sql.ErrNoRows: // todo
		return ErrNoMatch
	}
	if err != nil {
		tx.Rollback() // Откат транзакции при ошибке
		return err
	}

	addOrderQuery :=
		`INSERT INTO orders (userFullName, isbn) 
		VALUES ($1, $2) 
		RETURNING id, order_date`

	err = tx.QueryRow(
		addOrderQuery,
		order.UserFullName,
		order.Isbn,
	).Scan(
		&order.Id, // тут была копия в локальную переменную id и из нее уже в ордер
		&createdAt,
	)
	if err != nil {
		tx.Rollback() // Откат транзакции при ошибке
		return err
	}

	order.Date = createdAt.Format("2006-01-02 15:04:05")

	err = tx.Commit() // Завершение транзакции
	if err != nil {
		return err
	}
	return nil
}

func (db Database) GetOrderById(orderId uint64) (models.Result, error) {
	result := models.Result{}

	query :=
		`SELECT o.id, o.userFullName, b.title, b.author, o.order_date
		FROM orders o
		JOIN books b ON o.isbn = b.isbn
		WHERE o.id = $1;`

	row := db.Conn.QueryRow(query, orderId)
	err := row.Scan(
		&result.Id,
		&result.FullName,
		&result.Title,
		&result.Author,
		&result.OrderDate,
	)

	if err != nil {
		fmt.Println("can't get result from db")
	}

	return result, err
}
