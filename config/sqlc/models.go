// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package tutorial

import (
	"time"
)

type User struct {
	ID        int32
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
