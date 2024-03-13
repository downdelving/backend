package id

import (
	"github.com/downdelving/backend/pkg/model"
)

type GetResponse struct {
	Account model.Account `json:"account"`
}
