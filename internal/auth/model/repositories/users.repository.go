package repositories

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/auth/model/entities"
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type UsersRepository interface {
	Insert(ctx context.Context, identity *entities.Identity, user *entities.User) (types.Id, error)
	GetByEmailOrUsername(ctx context.Context, email string, username string) (*entities.User, error)
}
