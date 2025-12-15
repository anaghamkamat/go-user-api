package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"go-user-api/internal/logger"
	"go-user-api/internal/models"
	"go-user-api/internal/service"
)

type UserHandler struct {
	svc service.UserService
	val *validator.Validate
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		svc: s,
		val: validator.New(),
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Error("failed to parse request", zap.Error(err))
		return c.Status(400).JSON(err.Error())
	}

	if err := h.val.Struct(req); err != nil {
		logger.Log.Warn("validation failed", zap.Error(err))
		return c.Status(400).JSON(err.Error())
	}

	user, err := h.svc.Create(c.Context(), req)
	if err != nil {
		logger.Log.Error("failed to create user", zap.Error(err))
		return c.Status(500).JSON(err.Error())
	}

	logger.Log.Info("user created", zap.Int32("id", user.ID))
	return c.Status(201).JSON(user)
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := h.svc.Get(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON("user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	users, err := h.svc.ListPaginated(c.Context(), page, limit)
	if err != nil {
		logger.Log.Error("failed to list users", zap.Error(err))
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(users)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user, err := h.svc.Update(c.Context(), int32(id), req)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.svc.Delete(c.Context(), int32(id)); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.SendStatus(204)
}
