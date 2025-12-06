package models

//go get github.com/shopspring/decimal < necessário instalar
import "github.com/shopspring/decimal"

type inf_diario struct {
	TpFundoClasse    *string
	CnpjFundoClasse  *string
	IdSubclasse      *string
	DtComptc         *string
	VlTotal          *decimal.Decimal
	VlQuota          *decimal.Decimal
	VlPatrimLiq      *decimal.Decimal
	CaptcDia         *decimal.Decimal
	ResgDia          *decimal.Decimal
	NrCotst          *decimal.Decimal
}