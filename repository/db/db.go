package database

import (
	"apiTest/repository/cache"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Reservation struct {
	id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  int    `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
	Uuid        string `json:"uuid"`
}

func CreateDb() (db *sql.DB) {
	connstr := "postgres://postgres:112233@localhost:1000/data?sslmode=disable"

	db, err := sql.Open("postgres", connstr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createReservationTable(db)
	return db
}

func createReservationTable(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS reservation(
		id BIGSERIAL PRIMARY KEY,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50) NOT NULL,
		email VARCHAR(150) NOT NULL UNIQUE,
		national_id INTEGER NOT NULL,
		ticket_count INTEGER NOT NULL,
		uuid UUID NOT NULL
	);CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func InsertReservation(db *sql.DB, reservation cache.Reservation) int {
	query := `INSERT INTO reservation (first_name,last_name,email,national_id,ticket_count,uuid) VALUES ($1,$2,$3,$4,$5,uuid_generate_v4()) RETURNING id;`

	var pk int
	err := db.QueryRow(query, reservation.FirstName, reservation.LastName, reservation.Email, reservation.NationalID, reservation.TicketCount).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func GetAllReservations(db *sql.DB) ([]Reservation, error) {
	query := `SELECT * FROM reservation;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.id, &reservation.FirstName, &reservation.LastName, &reservation.Email, &reservation.NationalID, &reservation.TicketCount, &reservation.Uuid)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
func GetReservationByID(db *sql.DB, national_id int) (Reservation, error) {
	query := `SELECT * FROM reservation WHERE national_id = $1;`
	var reservation Reservation
	err := db.QueryRow(query, national_id).Scan(&reservation.id, &reservation.FirstName, &reservation.LastName, &reservation.Email, &reservation.NationalID, &reservation.TicketCount, &reservation.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return Reservation{}, nil // اگر رزرو با این شناسه پیدا نشد، مقدار خالی برمی‌گردانیم
		}
		return Reservation{}, err
	}
	return reservation, nil
}
func UpdateReservation(db *sql.DB, national_id int, reservation cache.Reservation) error {
	query := `UPDATE reservation SET first_name = $1, last_name = $2, email = $3, ticket_count = $4 WHERE national_id = $5;`
	_, err := db.Exec(query, reservation.FirstName, reservation.LastName, reservation.Email, reservation.TicketCount, national_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReservation(db *sql.DB, national_id int) error {
	query := `DELETE FROM reservation WHERE national_id = $1;`
	_, err := db.Exec(query, national_id)
	if err != nil {
		return err
	}
	return nil
}
