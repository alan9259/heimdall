package location

import (
	"miu-auth-api-v1/internal/model"
)

type Store interface {
	Create(*model.Location) error
}
