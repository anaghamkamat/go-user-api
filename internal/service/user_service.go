package service

import (
	"context"
	"time"

	db "go-user-api/db/sqlc"
	"go-user-api/internal/models"
	"go-user-api/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, req models.CreateUserRequest) (db.User, error)
	Get(ctx context.Context, id int32) (models.UserResponse, error)
	List(ctx context.Context) ([]models.UserResponse, error)
	Update(ctx context.Context, id int32, req models.CreateUserRequest) (db.User, error)
	Delete(ctx context.Context, id int32) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Create(ctx context.Context, req models.CreateUserRequest) (db.User, error) {
	dob, _ := time.Parse("2006-01-02", req.DOB)
	return s.repo.Create(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
}

func (s *userService) Get(ctx context.Context, id int32) (models.UserResponse, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  models.CalculateAge(u.Dob),
	}, nil
}

func (s *userService) List(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var res []models.UserResponse
	for _, u := range users {
		res = append(res, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  models.CalculateAge(u.Dob),
		})
	}
	return res, nil
}

func (s *userService) Update(ctx context.Context, id int32, req models.CreateUserRequest) (db.User, error) {
	dob, _ := time.Parse("2006-01-02", req.DOB)
	return s.repo.Update(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
}

func (s *userService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
