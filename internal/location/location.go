package location

import (
	"heimdall/internal/model"
)

type Store interface {
	Create(*model.Location) error
}
