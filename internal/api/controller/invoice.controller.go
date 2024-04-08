package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/api/response"
	"github.com/richhh7g/mm-api-nfe/internal/usecase/invoice"
)

var (
	ErrKeyAlreadyExists = errors.New("the key already exists")
)

var _ Invoice = &InvoiceImpl{}

type Invoice interface {
	CreateInvoice(w http.ResponseWriter, r *http.Request)
}

type InvoiceImpl struct {
	ctx                   *context.Context
	checkKeyExistsUseCase invoice.CheckKeyExistsUseCase
	createInvoiceUseCase  invoice.CreateInvoiceUseCase
}

func NewInvoice(ctx *context.Context, checkKeyExistsUseCase invoice.CheckKeyExistsUseCase, createInvoiceUseCase invoice.CreateInvoiceUseCase) *InvoiceImpl {
	return &InvoiceImpl{
		ctx:                   ctx,
		checkKeyExistsUseCase: checkKeyExistsUseCase,
		createInvoiceUseCase:  createInvoiceUseCase,
	}
}

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
