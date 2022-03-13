package internal

import (
	"net/http"
)

type IParameters interface {
	EsaAPIParameter() *EsaAPIParameter
}

type IPaginationParameters interface {
	PageValue() (int, bool)
	PerPageValue() (int, bool)
}

type IResponse interface {
	SetRateLimitInfo(h http.Header)
}
