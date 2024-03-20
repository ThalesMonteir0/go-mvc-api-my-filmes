package request

type UserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Name            string `json:"name" validate:"required, min=2"`
}

//func (e *UserRequest) ValidateUserRequest() error {
//	validate := validator.New()
//	err := validate.Struct(e)
//	return err
//}
