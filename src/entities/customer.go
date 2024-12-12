package entities

type Customer struct {
	Id               int    `json:"id"`
	Nik              string `json:"nik"`
	FullName         string `json:"full_name"`
	LegalName        string `json:"legal_name"`
	DateOfBirth      string `json:"date_of_birth"`
	LocationOfBirth  string `json:"location_of_birth"`
	KtpPicture       string `json:"ktp_picture"`
	SelfiePicture    string `json:"selfie_picture"`
	MonthlySalaryIdr int    `json:"monthly_salary_idr"`
	Month1Limit      int    `json:"1st_month_limit"`
	Month2Limit      int    `json:"2nd_month_limit"`
	Month3Limit      int    `json:"3rd_month_limit"`
	Month4Limit      int    `json:"4th_month_limit"`
}

type CustomerInsert struct {
	Nik              string `json:"nik" validate:"required"`
	FullName         string `json:"full_name" validate:"required"`
	LegalName        string `json:"legal_name" validate:"required"`
	DateOfBirth      string `json:"date_of_birth" validate:"required"`
	LocationOfBirth  string `json:"location_of_birth" validate:"required"`
	KtpPicture       string `json:"ktp_picture" validate:"required"`
	SelfiePicture    string `json:"selfie_picture" validate:"required"`
	MonthlySalaryIdr int    `json:"monthly_salary_idr" validate:"required"`
	Month1Limit      int    `json:"1st_month_limit" validate:"required"`
	Month2Limit      int    `json:"2nd_month_limit" validate:"required"`
	Month3Limit      int    `json:"3rd_month_limit" validate:"required"`
	Month4Limit      int    `json:"4th_month_limit" validate:"required"`
}

type CustomerUpdate struct {
	Id               int     `json:"id" validate:"required"`
	Nik              *string `json:"nik" validate:"omitempty"`
	FullName         *string `json:"full_name" validate:"omitempty"`
	LegalName        *string `json:"legal_name" validate:"omitempty"`
	DateOfBirth      *string `json:"date_of_birth" validate:"omitempty"`
	LocationOfBirth  *string `json:"location_of_birth" validate:"omitempty"`
	KtpPicture       *string `json:"ktp_picture" validate:"omitempty"`
	SelfiePicture    *string `json:"selfie_picture" validate:"omitempty"`
	MonthlySalaryIdr *int    `json:"monthly_salary_idr" validate:"omitempty"`
	Month1Limit      *int    `json:"1st_month_limit" validate:"omitempty"`
	Month2Limit      *int    `json:"2nd_month_limit" validate:"omitempty"`
	Month3Limit      *int    `json:"3rd_month_limit" validate:"omitempty"`
	Month4Limit      *int    `json:"4th_month_limit" validate:"omitempty"`
}
