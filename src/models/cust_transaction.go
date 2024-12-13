package models

import (
	"context"
	"errors"
	"studi_kasus_xyz/entities"
	"time"

	"github.com/google/uuid"
)

func GetTransactionFromCustID(ctx context.Context, id_cust int) (results []entities.CustTransaction, err error) {
	if id_cust < 1 {
		return nil, errors.New("id customer invalid")
	}

	results = []entities.CustTransaction{}

	rows, err := dbPool.QueryContext(ctx,
		`SELECT 
			contract_no,
			id_customer,
			otr,
			asset_name,
			admin_fee_idr,
			interest,
			installment,
			created_at
		WHERE id_customer = ?`,
		id_cust,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var created_at time.Time
		var row entities.CustTransaction
		err := rows.Scan(
			&row.ContractNo,
			&row.IdCustomer,
			&row.Otr,
			&row.AssetName,
			&row.AdminFeeIdr,
			&row.Interest,
			&row.Installment,
			&created_at,
		)
		if err != nil {
			return nil, err
		}

		row.CreatedAt = created_at.Local().Format(time.DateTime)

		results = append(results, row)
	}

	return results, nil
}

func InsertCustTransaction(ctx context.Context, data entities.CustTransactionInsert) error {
	if data.IdCustomer < 1 {
		return errors.New("id customer invalid")
	}

	contract_no := uuid.New().String()

	_, err := dbPool.ExecContext(ctx,
		`INSERT INTO 
		(contract_no, id_customer, otr, asset_name, admin_fee_idr, interest, installment)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)`,
		contract_no, data.IdCustomer, data.Otr, data.AssetName, data.AdminFeeIdr, data.Interest, data.Installment,
	)
	if err != nil {
		return err
	}

	return nil
}
