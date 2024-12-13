package controllers

import (
	"os"
	"reflect"
	"strconv"
	"studi_kasus_xyz/entities"
	"studi_kasus_xyz/models"
	"studi_kasus_xyz/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCustDataByID(c *fiber.Ctx) error {
	id := c.QueryInt("id_user")
	if id < 1 {
		return utils.Response(c, 400, "[Bad request]", "invalid id")
	}

	response, err := models.GetCustDataByID(c.Context(), id)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}
	return utils.Response(c, 200, "[Success]", response)
}

func InsertCustData(c *fiber.Ctx) error {
	IdUser, err := strconv.Atoi(c.FormValue("id_user"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid id user")
	}
	MonthlySalaryIdr, err := strconv.Atoi(c.FormValue("monthly_salary_idr"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid monthly_salary_idr")
	}
	Month1Limit, err := strconv.Atoi(c.FormValue("1st_month_limit"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid 1st_month_limit")
	}
	Month2Limit, err := strconv.Atoi(c.FormValue("2nd_month_limit"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid 2nd_month_limit")
	}
	Month3Limit, err := strconv.Atoi(c.FormValue("3rd_month_limit"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid 3rd_month_limit")
	}
	Month4Limit, err := strconv.Atoi(c.FormValue("4th_month_limit"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid 4th_month_limit")
	}

	mark_delete_on_fail := []string{}
	defer func() {
		for _, val := range mark_delete_on_fail {
			if val != "" {
				_ = os.Remove("./uploads/" + val)
			}
		}
	}()

	KtpFile, err := c.FormFile("ktp_picture")
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}
	ktpuuid := uuid.New().String()
	err = c.SaveFile(KtpFile, "./uploads/"+ktpuuid)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}
	mark_delete_on_fail = append(mark_delete_on_fail, ktpuuid)

	SelfieFile, err := c.FormFile("selfie_picture")
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}
	selfieuuid := uuid.New().String()
	err = c.SaveFile(SelfieFile, "./uploads/"+selfieuuid)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}
	mark_delete_on_fail = append(mark_delete_on_fail, selfieuuid)

	input := entities.CustomerInsert{
		IdUser:           IdUser,
		Nik:              c.FormValue("nik"),
		FullName:         c.FormValue("full_name"),
		LegalName:        c.FormValue("legal_name"),
		DateOfBirth:      c.FormValue("date_of_birth"),
		LocationOfBirth:  c.FormValue("location_of_birth"),
		KtpPicture:       ktpuuid,
		SelfiePicture:    selfieuuid,
		MonthlySalaryIdr: MonthlySalaryIdr,
		Month1Limit:      Month1Limit,
		Month2Limit:      Month2Limit,
		Month3Limit:      Month3Limit,
		Month4Limit:      Month4Limit,
	}

	err = utils.ValidateStruct(input)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = models.InsertCustData(c.Context(), input)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}

	mark_delete_on_fail = []string{} // delete nothing on success
	return utils.Response(c, 200, "[Success]", nil)
}

func UpdateCustData(c *fiber.Ctx) error {
	var input entities.CustomerUpdate

	v := reflect.ValueOf(input)
	vptr := reflect.ValueOf(&input)

	for i := range v.NumField() {
		field_type := v.Field(i).Type().String()
		json_tag := v.Type().Field(i).Tag.Get("json")

		if json_tag == "id_user" {
			continue
		}

		switch field_type {
		case "*string":
			{
				val := c.FormValue(json_tag)
				if val != "" {
					vptr.Elem().Field(i).Set(reflect.ValueOf(&val))
				}
			}
		case "*int":
			{
				val, err := strconv.Atoi(c.FormValue(json_tag))
				if err != nil {
					return utils.Response(c, 400, "[Bad Request]", "invalid "+json_tag)
				}
				vptr.Elem().Field(i).Set(reflect.ValueOf(&val))
			}
		}
	}

	Id, err := strconv.Atoi(c.FormValue("id_user"))
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", "invalid id user")
	}

	mark_delete_on_fail := []string{}
	defer func() {
		for _, val := range mark_delete_on_fail {
			if val != "" {
				_ = os.Remove("./uploads/" + val)
			}
		}
	}()

	var ktpuuid *string
	var selfieuuid *string
	KtpFile, err := c.FormFile("ktp_picture")
	if err == nil {
		temp := uuid.New().String()
		ktpuuid = &temp
		err = c.SaveFile(KtpFile, "./uploads/"+*ktpuuid)
		if err != nil {
			return utils.Response(c, 400, "[Bad Request]", err.Error())
		}
		mark_delete_on_fail = append(mark_delete_on_fail, *ktpuuid)
	}

	SelfieFile, err := c.FormFile("selfie_picture")
	if err == nil {
		temp := uuid.New().String()
		selfieuuid = &temp
		err = c.SaveFile(SelfieFile, "./uploads/"+*selfieuuid)
		if err != nil {
			return utils.Response(c, 400, "[Bad Request]", err.Error())
		}
		mark_delete_on_fail = append(mark_delete_on_fail, *selfieuuid)
	}

	input.IdUser = Id
	input.KtpPicture = ktpuuid
	input.SelfiePicture = selfieuuid

	err = utils.ValidateStruct(input)
	if err != nil {
		return utils.Response(c, 400, "[Bad Request]", err.Error())
	}

	err = models.UpdateCustData(c.Context(), input)
	if err != nil {
		return utils.Response(c, 500, "[Error]", err.Error())
	}

	mark_delete_on_fail = []string{} // delete nothing on success
	return utils.Response(c, 200, "[Success]", nil)
}
