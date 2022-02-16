package constantUtils

const (
	//MESSAGES

	GenericNotFoundMessage            = "Any %_% was found"
	GenericFoundSuccessMessage        = "%_% found with successful"
	GenericAlreadyExistsMessage       = "This %_% already exists on our database"
	GenericCreateErrorMessage         = "It was not possible to create this %_%"
	GenericCreateSuccessMessage       = "%_% created with successful"
	GenericUpdateErrorMessage         = "It was not possible to update this %_%"
	GenericUpdateSuccessMessage       = "%_% updated with successful"
	GenericDeleteErrorMessage         = "It was not possible to update this %_%"
	GenericInternalServerErrorMessage = "It was not possible to perform this action"
	GenericInvalidFieldMessage        = "%_% cannot be null"
	SetCacheFailed                    = "It was not possible to save datas in cache."
	CpfCnpjInvalidMessage             = "This CPF or CNPJ is not valid!"
	RoleInvalidMessage                = "This Role is not valid!"
	EmailInvalidMessage               = "This Email is not valid!"
	GenericUnauthorizedMessage        = "The token informed is not valid!"

	//STATUS

	StatusSuccess             = "success"
	StatusConflict            = "conflict"
	StatusInternalServerError = "internal_server_error"
	StatusNotFound            = "not_found"
	StatusBadRequest          = "bad_request"
	StatusUnauthorized        = "unauthorized"

	//ROUTES

	PathUserIdParam           = "/:userId"
	PathAddressIdParam        = "/:addressId"
	PathUpdateUserMainAddress = "/:addressId/address/:userId/user"
	PathRewardIdParam         = "/:rewardId"

	//CONDITIONS

	IdCondition     = "id = ?"
	UserIdCondition = "user_id = ?"

	//PARAMS

	RedirectFrontErrorParams = "?status=%s&message=%s&data=%s"
)
