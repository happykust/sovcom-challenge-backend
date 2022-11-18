package account

type AccountVerifyRequest struct {
	Id                uint   `form:"id" json:"id" binding:"required"`
	AdditionalContact string `form:"additional_contact" json:"additional_contact" binding:"required"`
}

type AccountVerifyResponse struct {
	Message            string `form:"message" json:"message" binding:"required"`
	MeetingInformation string `form:"message" json:"message" binding:"required"`
	PersonalAssistant  string `form:"message" json:"message" binding:"required"`
}
