package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/api/response"
	"github.com/richhh7g/mm-api-nfe/internal/usecase/invoice"
)

var (
	ErrKeyAlreadyExists = errors.New("the key already exists")
	ErrKeyNotFound      = errors.New("the key not found")
)

var _ Invoice = &InvoiceImpl{}

type Invoice interface {
	CreateInvoice(w http.ResponseWriter, r *http.Request)
	GetInvoice(w http.ResponseWriter, r *http.Request)
}

type InvoiceParams struct {
	Ctx                    *context.Context
	CheckKeyExistsUseCase  invoice.CheckKeyExistsUseCase
	CreateInvoiceUseCase   invoice.CreateInvoiceUseCase
	GetInvoiceByKeyUseCase invoice.GetInvoiceByKeyUseCase
}

type InvoiceImpl struct {
	ctx                    *context.Context
	checkKeyExistsUseCase  invoice.CheckKeyExistsUseCase
	createInvoiceUseCase   invoice.CreateInvoiceUseCase
	getInvoiceByKeyUseCase invoice.GetInvoiceByKeyUseCase
}

func NewInvoice(input *InvoiceParams) *InvoiceImpl {
	return &InvoiceImpl{
		ctx:                    input.Ctx,
		checkKeyExistsUseCase:  input.CheckKeyExistsUseCase,
		createInvoiceUseCase:   input.CreateInvoiceUseCase,
		getInvoiceByKeyUseCase: input.GetInvoiceByKeyUseCase,
	}
}

// @Summary Create invoice
// @Tags v1, Invoices
// @Accept json
// @Produce json
// @Param invoice body CreateInvoiceBody true "Invoice"
// @Success 201 {object} CreateInvoiceResponse
// @Failure 409 {object} ErrorResponse
// @Router /api/v1/nota-fiscal [POST]
func (c *InvoiceImpl) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input dto.CreateInvoiceBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	keyAlreadyExists := c.checkKeyExistsUseCase.Exec(&input.Chave)
	if keyAlreadyExists {
		response.NewError(ErrKeyAlreadyExists, http.StatusConflict).Send(w)
		return
	}

	invoiceResponse, err := c.createInvoiceUseCase.Exec(&input)

	if err != nil {
		response.NewError(err, http.StatusAccepted).Send(w)
		return
	}

	response.NewSuccess(invoiceResponse, http.StatusCreated).Send(w)
}

// @Summary Get invoice
// @Tags v1, Invoices
// @Accept json
// @Produce json
// @Param chave path string true "Chave"
// @Success 200 {object} GetInvoiceResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/nota-fiscal/{chave} [GET]
func (c *InvoiceImpl) GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID := chi.URLParam(r, "chave")

	invoiceResponse, err := c.getInvoiceByKeyUseCase.Exec(&invoiceID)
	if err != nil {
		response.NewError(ErrKeyNotFound, http.StatusNotFound).Send(w)
		return
	}

	response.NewSuccess(invoiceResponse, http.StatusOK).Send(w)
}
