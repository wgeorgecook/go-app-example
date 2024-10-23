package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func getAllPets(conn *pgx.Conn) ([]Pet, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM pets")
	if err != nil {
		return nil, err
	}

	pets, err := pgx.CollectRows(rows, pgx.RowToStructByName[Pet])
	if err != nil {
		return nil, err
	}

	return pets, nil
}
