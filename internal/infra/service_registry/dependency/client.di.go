package dependency

import (
	"github.com/google/wire"
	documentValidate "github.com/richhh7g/mm-api-nfe/internal/infra/data/client/http/document_validate"
)

var DocumentValidateClientDI = wire.NewSet(
	documentValidate.NewClient,
	wire.Bind(new(documentValidate.Client), new(*documentValidate.ClientImpl)),
)
