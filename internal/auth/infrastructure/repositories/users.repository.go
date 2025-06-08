package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"github.com/GanderBite/reservation-api/internal/auth/model/entities"
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type PostgresUsersRepository struct {
	db *sql.DB
}

func NewPostgresUsersRepository(db *sql.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{
		db: db,
	}
}

func (repo *PostgresUsersRepository) Insert(
	ctx context.Context,
	identity *entities.Identity,
	user *entities.User,
) (types.Id, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}

	// Handle rollback on panic or error
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("tx rollback failed: %v", rbErr)
			}
		}
	}()

	const identityQuery = `
		INSERT INTO identities (email, password)
		VALUES ($1, $2)
		RETURNING id
	`

	var identityID uuid.UUID
	err = tx.QueryRowContext(ctx, identityQuery, identity.Email, identity.Password).Scan(&identityID)
	if err != nil {
		return uuid.Nil, err
	}

	const userQuery = `
		INSERT INTO users (name, identity_id)
		VALUES ($1, $2)
	`

	_, err = tx.ExecContext(ctx, userQuery, user.Name, identityID)
	if err != nil {
		return uuid.Nil, err
	}

	if err = tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (repo *PostgresUsersRepository) GetByEmailOrUsername(
	ctx context.Context,
	email string,
	username string,
) (*entities.User, error) {
	const query = `
		SELECT u.id, u.name, u.identity_id
		FROM users u
		INNER JOIN identities i ON u.identity_id = i.id
		WHERE u.name = $1 OR i.email = $2
		LIMIT 1
	`

	row := repo.db.QueryRowContext(ctx, query, username, email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.IdentityId)
	if err == sql.ErrNoRows {
		return nil, nil // no user or identity matched
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
