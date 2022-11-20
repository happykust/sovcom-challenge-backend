package account

type AccountGetOneNVUserRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

type AccountGetOneNVUserResponse struct {
	UserName           string `form:"user_name" json:"user_name" binding:"required"`
	Email              string `form:"email" json:"email" binding:"required"`
	FirstName          string `form:"first_name" json:"first_name" binding:"required"`
	LastName           string `form:"last_name" json:"last_name" binding:"required"`
	RegistrationStatus string `form:"registration_status" json:"registration_status" binding:"required"`
	ReferralCode       string `form:"referral_code" json:"referral_code" binding:"required"`
	Ban                bool   `form:"ban" json:"ban" binding:"required"`
	PersonalAssistant  uint   `form:"personal_assistant" json:"personal_assistant" binding:"required"`
	MeetingInformation string `form:"meeting_information" json:"meeting_information" binding:"required"`
	AdditionalContact  string `form:"additional_contact" json:"additional_contact" binding:"required"`
}
