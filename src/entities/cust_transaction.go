package entities

type CustTransaction struct {
	ContractNo  string `json:"contract_no"`
	IdCustomer  int    `json:"id_customer"`
	Otr         string `json:"otr"`
	AssetName   string `json:"asset_name"`
	AdminFeeIdr int    `json:"admin_fee_idr"`
	Interest    int    `json:"interest"`
	Installment int    `json:"installment"`
	CreatedAt   string `json:"created_at"`
}

type CustTransactionInsert struct {
	IdCustomer  int    `json:"id_customer"`
	Otr         string `json:"otr" validate:"required"`
	AssetName   string `json:"asset_name" validate:"required"`
	AdminFeeIdr int    `json:"admin_fee_idr" validate:"required"`
	Interest    int    `json:"interest" validate:"required"`
	Installment int    `json:"installment" validate:"required"`
}
