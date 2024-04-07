package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/api/response"
	"github.com/richhh7g/mm-api-nfe/internal/usecase/invoice"
)

var _ Invoice = &InvoiceImpl{}

type Invoice interface {
	CreateInvoice(w http.ResponseWriter, r *http.Request)
}

type InvoiceImpl struct {
	ctx                  *context.Context
	createInvoiceUseCase invoice.CreateInvoiceUseCase
}

func NewInvoice(ctx *context.Context, createInvoiceUseCase invoice.CreateInvoiceUseCase) *InvoiceImpl {
	return &InvoiceImpl{
		ctx:                  ctx,
		createInvoiceUseCase: createInvoiceUseCase,
	}
}

func (h *InvoiceImpl) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	invoiceResponse, err := h.createInvoiceUseCase.Exec(&input)

	if err != nil {
		response.NewError(err, http.StatusAccepted).Send(w)
		return
	}

	response.NewSuccess(invoiceResponse, http.StatusCreated).Send(w)
}
