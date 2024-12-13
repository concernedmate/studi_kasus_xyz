package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"studi_kasus_xyz/entities"
	"time"
)

func GetCustDataByID(ctx context.Context, id int) (result *entities.Customer, err error) {
	if id < 1 {
		return nil, errors.New("id customer invalid")
	}

	var cust_data entities.Customer
	var date_of_birth time.Time
	err = dbPool.QueryRowContext(ctx,
		`SELECT 
			id_user,
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
		WHERE id_user = ?`,
		id,
	).Scan(
		&cust_data.IdUser,
		&cust_data.Nik,
		&cust_data.FullName,
		&cust_data.LegalName,
		&date_of_birth,
		&cust_data.LocationOfBirth,
		&cust_data.KtpPicture,
		&cust_data.SelfiePicture,
		&cust_data.MonthlySalaryIdr,
		&cust_data.Month1Limit,
		&cust_data.Month2Limit,
		&cust_data.Month3Limit,
		&cust_data.Month4Limit,
	)
	if err != nil {
		return nil, err
	}
	cust_data.DateOfBirth = date_of_birth.Local().Format(time.DateOnly)

	return &cust_data, nil
}

func InsertCustData(ctx context.Context, data entities.CustomerInsert) error {
	if data.IdUser < 1 {
		return errors.New("id user invalid")
	}

	_, err := time.Parse(time.DateOnly, data.DateOfBirth)
	if err != nil {
		return errors.New("date of birth invalid")
	}

	_, err = dbPool.ExecContext(ctx,
		`INSERT INTO customer
		( 
			id_user, nik, full_name, legal_name, date_of_birth, location_of_birth, ktp_picture, selfie_picture, 
			monthly_salary_idr, 1st_month_limit, 2nd_month_limit, 3rd_month_limit, 4th_month_limit
		) 
		VALUES
		(
			?, ?, ?, ?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?
		)`,
		data.IdUser, data.Nik, data.FullName, data.LegalName,
		data.DateOfBirth, data.LocationOfBirth, data.KtpPicture, data.SelfiePicture,
		data.MonthlySalaryIdr, data.Month1Limit, data.Month2Limit, data.Month3Limit,
		data.Month4Limit,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCustData(ctx context.Context, fields entities.CustomerUpdate) error {
	if fields.IdUser < 1 {
		return errors.New("id customer invalid")
	}

	if fields.DateOfBirth != nil {
		_, err := time.Parse(time.DateOnly, *fields.DateOfBirth)
		if err != nil {
			return errors.New("date of birth invalid")
		}
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
	args = append(args, fields.IdUser)

	result, err := dbPool.ExecContext(ctx, fmt.Sprintf(`UPDATE customer SET %s WHERE id_user = ?`, strings.Join(updates, ", ")), args...)
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
