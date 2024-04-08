package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/richhh7g/mm-api-nfe/internal/api/controller"
)

type InvoiceRoute interface {
	CreateInvoice(w http.ResponseWriter, req *http.Request)
}

type InvoiceRouteImpl struct {
	router            *chi.Mux
	invoiceController controller.Invoice
}

func NewInvoiceRoute(router *chi.Mux, invoiceController controller.Invoice) *InvoiceRouteImpl {
	return &InvoiceRouteImpl{
		router:            router,
		invoiceController: invoiceController,
	}
}

func (r *InvoiceRouteImpl) GetPath() string {
	return "/api/v1/nota-fiscal"
}

func (r *InvoiceRouteImpl) CreateInvoice(w http.ResponseWriter, req *http.Request) {
	r.router.Post(r.GetPath(), r.invoiceController.CreateInvoice)
}
