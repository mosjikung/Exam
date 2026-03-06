package handler

import (
	"errors"
	"strconv"

	"product-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	svc service.ProductService
}

func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/products", h.List)
	api.Post("/products", h.Add)
	api.Delete("/products/:id", h.Delete)
}

func (h *ProductHandler) List(c *fiber.Ctx) error {
	products, err := h.svc.ListProducts()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(products)
}

func (h *ProductHandler) Add(c *fiber.Ctx) error {
	var body struct {
		ProductCode string `json:"product_code"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "request body ไม่ถูกต้อง")
	}

	p, err := h.svc.AddProduct(body.ProductCode)
	if err != nil {
		var ve *service.ErrValidation
		if errors.As(err, &ve) {
			return fiber.NewError(fiber.StatusUnprocessableEntity, ve.Message)
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(p)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id ต้องเป็นตัวเลข")
	}

	if err := h.svc.RemoveProduct(uint(id)); err != nil {
		var ve *service.ErrValidation
		if errors.As(err, &ve) {
			return fiber.NewError(fiber.StatusNotFound, ve.Message)
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{"message": "ลบข้อมูลสำเร็จ"})
}
