package invoice

import (
	"context"
	"time"

	"github.com/richhh7g/mm-api-nfe/internal/entity"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/client/database/sqlc_pg"
)

var _ InvoiceDataSource = &InvoiceDataSourceImpl{}

type InvoiceDataSource interface {
	CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error)
	FindInvoiceKeyExists(key entity.Key) bool
	FindInvoiceByKey(key entity.Key) (*entity.Invoice, error)
}

type InvoiceDataSourceImpl struct {
	ctx        context.Context
	repository *sqlc_pg.Queries
}

func NewInvoice(ctx context.Context, client sqlc_pg.DBTX) *InvoiceDataSourceImpl {
	return &InvoiceDataSourceImpl{
		ctx:        ctx,
		repository: sqlc_pg.New(client),
	}
}

func (d *InvoiceDataSourceImpl) CreateInvoice(invoice *entity.Invoice) (*entity.Invoice, error) {
	issuedAt, _ := time.Parse(time.RFC3339, invoice.IssuedAt)
	receivedAt, _ := time.Parse(time.RFC3339, invoice.ReceivedAt)

	params := sqlc_pg.CreateInvoiceParams{
		ID:         invoice.ID.String(),
		Cnpj:       invoice.CNPJ,
		Chave:      invoice.Key,
		EmissaoEm:  issuedAt,
		RecebidoEm: receivedAt,
	}

	_, err := d.repository.CreateInvoice(d.ctx, params)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (d *InvoiceDataSourceImpl) FindInvoiceKeyExists(key entity.Key) bool {
	_, err := d.repository.FindInvoiceKeyExists(d.ctx, string(key))
	return err == nil
}

func (d *InvoiceDataSourceImpl) FindInvoiceByKey(key entity.Key) (*entity.Invoice, error) {
	invoice, err := d.repository.FindInvoiceByKey(d.ctx, string(key))
	if err != nil {
		return nil, err
	}

	id, err := entity.ParseID(invoice.ID)
	if err != nil {
		return nil, err
	}

	return &entity.Invoice{
		ID:         id,
		IssuedAt:   invoice.EmissaoEm.String(),
		ReceivedAt: invoice.RecebidoEm.String(),
		Key:        entity.Key(invoice.Chave),
		CNPJ:       entity.CNPJ(invoice.Cnpj),
	}, nil
}
