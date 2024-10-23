package db

import "time"

// diet is the database representation
// of a petsapisv1.Diet type.
type diet int32

const (
	DIET_UNSPECIFIED diet = iota
	DIET_NO_RESTRICTIONS
	DIET_GRAIN_FREE
	DIET_RAW
)

// friendlyWith is the database representation
// of a petsapisv1.FriendlyWith type.
type friendlyWith int32

const (
	FRIENDLY_WITH_UNSPECIFIED friendlyWith = iota
	FRIENDLY_WITH_LITTLE_DOGS
	FRIENDLY_WITH_BIG_DOGS
	FRIENDLY_WITH_CATS
	FRIENDLY_WITH_YOUNG_CHILDREN
	FRIENDLY_WITH_OLDER_CHILDREN
)

// Pet is the database representation of a
// petapisv1.Pet type.
type Pet struct {
	ID           int64     `db:"id"`
	Name         string    `db:"name"`
	Birthdate    time.Time `db:"birthdate"`
	Description  string    `db:"description"`
	Diet         string    `db:"diet"`
	FriendlyWith []string  `db:"friendly_with"`
	PictureUrl   string    `db:"picture_url"`
}
