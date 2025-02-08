package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"goApi/config"
	"goApi/repository/cache"
	"log"
	"sync"
	"time"
)

type Reservation struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  int    `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
	Uuid        string `json:"uuid"`
}

var (
	db   *sql.DB
	once sync.Once
)

func Main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal(err)
	}
	once.Do(func() {
		var err error
		connstr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err = sql.Open("postgres", connstr)
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
		log.Println("Database Connetcted")
	})
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(time.Minute * 5)
	createReservationTable()
}

func GetDB() *sql.DB {
	return db
}

func createReservationTable() {

	query := `CREATE TABLE IF NOT EXISTS reservation(
		id BIGSERIAL PRIMARY KEY,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50) NOT NULL,
		email VARCHAR(150) NOT NULL,
		national_id INTEGER NOT NULL UNIQUE,
		ticket_count INTEGER NOT NULL,
		uuid UUID NOT NULL
	);CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	seedData()
}

func seedData() {
	queries := []string{
		"INSERT INTO reservation (first_name,last_name,email,national_id,ticket_count,uuid) VALUES ('ali','ali','ali@gmail.com',11111111,1,uuid_generate_v4()) ON CONFLICT DO NOTHING;",
		"INSERT INTO reservation (first_name,last_name,email,national_id,ticket_count,uuid) VALUES ('roz','roz','roz@gmail.com',11111112,2,uuid_generate_v4()) ON CONFLICT DO NOTHING;"}
	for _, query := range queries {
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		stmt.Exec()
	}
}

func InsertReservation(reservation cache.Reservation) error {
	stmt, err := db.Prepare(`INSERT INTO reservation (first_name,last_name,email,national_id,ticket_count,uuid) VALUES ($1,$2,$3,$4,$5,uuid_generate_v4());`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(reservation.FirstName, reservation.LastName, reservation.Email, reservation.NationalID, reservation.TicketCount)
	if err != nil {
		return err
	}
	return nil
}

func GetAllReservations(pageSize, page int) (map[string]interface{}, error) {

	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	var total int
	Countstmt, err := db.Prepare(`SELECT COUNT(*) FROM reservation;`)
	if err != nil {
		return nil, err
	}
	Countstmt.QueryRow().Scan(&total)
	totalPage := (total + pageSize - 1) / pageSize
	offset := (page - 1) * pageSize
	stmt, err := db.Prepare(`SELECT * FROM reservation ORDER BY id LIMIT $1 OFFSET $2`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		err := rows.Scan(&reservation.ID, &reservation.FirstName, &reservation.LastName, &reservation.Email, &reservation.NationalID, &reservation.TicketCount, &reservation.Uuid)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	response := map[string]interface{}{
		"data":       reservations,
		"total_page": totalPage,
		"page":       page,
		"total":      total,
	}
	return response, nil
}

func GetReservationByID(national_id string) (Reservation, error) {
	stmt, err := db.Prepare(`SELECT * FROM reservation WHERE national_id = $1;`)
	if err != nil {
		return Reservation{}, err
	}
	defer stmt.Close()
	var reservation Reservation
	err = stmt.QueryRow(national_id).Scan(&reservation.ID, &reservation.FirstName, &reservation.LastName, &reservation.Email, &reservation.NationalID, &reservation.TicketCount, &reservation.Uuid)
	if err != nil {
		return Reservation{}, err
	}
	return reservation, nil
}
func UpdateReservation(national_id string, reservation cache.Reservation) error {
	stmt, err := db.Prepare(`UPDATE reservation SET first_name = $1, last_name = $2, email = $3, ticket_count = $4 WHERE national_id = $5;`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(reservation.FirstName, reservation.LastName, reservation.Email, reservation.TicketCount, national_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReservation(national_id string) error {
	stmt, err := db.Prepare(`DELETE FROM reservation WHERE national_id = $1;`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(national_id)
	if err != nil {
		return err
	}
	return nil
}

func SearchReservation(value string, page, pageSize int) (map[string]interface{}, error) {

	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	var total int
	SearchValue := value + "%"
	Countstmt, err := db.Prepare(`SELECT COUNT(*) FROM reservation WHERE first_name ILIKE $1;`)
	if err != nil {
		return nil, err
	}
	Countstmt.QueryRow(SearchValue).Scan(&total)
	totalPage := (total + pageSize - 1) / pageSize
	offset := (page - 1) * pageSize
	stmt, err := db.Prepare(`SELECT * FROM reservation WHERE first_name ILIKE $1 LIMIT $2 OFFSET $3;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(SearchValue, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		if err := rows.Scan(&reservation.ID, &reservation.FirstName, &reservation.LastName, &reservation.Email, &reservation.NationalID, &reservation.TicketCount, &reservation.Uuid); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	response := map[string]interface{}{
		"data":       reservations,
		"total_page": totalPage,
		"page":       page,
		"total":      total,
	}
	return response, nil
}
