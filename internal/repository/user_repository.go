package repository

import (
	"context"

	db "go-user-api/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	Get(ctx context.Context, id int32) (db.User, error)
	List(ctx context.Context) ([]db.User, error)

	// ✅ Pagination bonus
	ListPaginated(ctx context.Context, limit, offset int32) ([]db.User, error)

	Update(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	Delete(ctx context.Context, id int32) error
}

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) UserRepository {
	return &userRepository{q}
}

func (r *userRepository) Create(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *userRepository) Get(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUser(ctx, id)
}

func (r *userRepository) List(ctx context.Context) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}

// ✅ Pagination implementation
func (r *userRepository) ListPaginated(
	ctx context.Context,
	limit, offset int32,
) ([]db.User, error) {
	return r.q.ListUsersPaginated(ctx, db.ListUsersPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *userRepository) Update(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.q.UpdateUser(ctx, arg)
}

func (r *userRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
