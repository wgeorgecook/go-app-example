package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func getAllPets(conn *pgx.Conn) ([]Pet, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM pets")
	if err != nil {
		return nil, err
	}

	var pets []Pet
	for rows.Next() {
		var id int64
		var name string
		var birthdate time.Time
		var description string
		var diet string
		var friendlyWith []string
		var imageUrl string
		if err := rows.Scan(
			&id, &name, &birthdate, &description, &diet, &friendlyWith, &imageUrl,
		); err != nil {
			return nil, err
		}
		pet := Pet{
			name,
			birthdate,
			description,
			diet,
			friendlyWith,
			imageUrl,
		}
		pets = append(pets, pet)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return pets, nil
}
