package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"studi_kasus_xyz/entities"
)

func GetCustDataByID(ctx context.Context, id int) (result *entities.Customer, err error) {
	if id < 1 {
		return nil, errors.New("id customer invalid")
	}

	var cust_data entities.Customer
	err = dbPool.QueryRowContext(ctx,
		`SELECT 
			id,
			nik,
			full_name,
			legal_name,
			date_of_birth,
			location_of_birth,
			ktp_picture,
			selfie_picture,
			monthly_salary_idr,
			1st_month_limit,
			2nd_month_limit,
			3rd_month_limit,
			4th_month_limit
		FROM customer
		WHERE id = ?`,
		id,
	).Scan(
		cust_data.Id,
		cust_data.Nik,
		cust_data.FullName,
		cust_data.LegalName,
		cust_data.DateOfBirth,
		cust_data.LocationOfBirth,
		cust_data.KtpPicture,
		cust_data.SelfiePicture,
		cust_data.MonthlySalaryIdr,
		cust_data.Month1Limit,
		cust_data.Month2Limit,
		cust_data.Month3Limit,
		cust_data.Month4Limit,
	)
	if err != nil {
		return nil, err
	}

	return &cust_data, nil
}

func InsertCustData(ctx context.Context, data entities.CustomerInsert) error {
	// TODO
	return nil
}

func UpdateCustData(ctx context.Context, fields entities.CustomerUpdate) error {
	if fields.Id < 1 {
		return errors.New("id customer invalid")
	}

	v := reflect.ValueOf(fields)

	updates := []string{}
	args := []any{}

	for i := range v.NumField() {
		field_type := v.Field(i).Type().String()
		json_tag := v.Type().Field(i).Tag.Get("json")

		if json_tag == "id" {
			continue
		}

		switch field_type {
		case "*string":
			{
				val := v.Field(i).Interface().(*string)
				if val != nil {
					updates = append(updates, json_tag+" = ?")
					args = append(args, *val)
				}
			}
		case "*int":
			{
				val := v.Field(i).Interface().(*int)
				if val != nil {
					updates = append(updates, json_tag+" = ?")
					args = append(args, *val)
				}
			}
		}
	}

	if len(updates) == 0 || len(args) == 0 {
		return errors.New("no field updated")
	}
	args = append(args, fields.Id)

	result, err := dbPool.ExecContext(ctx, fmt.Sprintf(`UPDATE customer %s WHERE id = ?`, strings.Join(updates, ", ")), args...)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("customer not found")
	}

	return nil
}
