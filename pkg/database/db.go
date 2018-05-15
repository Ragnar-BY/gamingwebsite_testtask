package database

import (
	"errors"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/player"
)

var (
	// ErrWrongID is error for wrong id.
	ErrWrongID = errors.New("wrong ID")
)

// TODO move this interface to manager.go and don't use double putting.
// DB is interface for database.
type DB interface {
	PlayerByID(id int) (*player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayer(id int, player player.Player) error
}
