package errmsg

const (
	ErrorMsgUnexpected                     = "unexpected error"
	ErrorMsgNotFound                       = "record not found"
	ErrorMsgCantScanQueryResult            = "can't scan query results"
	ErrorMsgCantExecuteCommand             = "can't execute command"
	ErrorMsgCantGetInsertedID              = "can't get inserted user ID"
	ErrorMsgMysqlQuery                     = "mysql query error"
	ErrorMsgPhoneNumberIsNotUnique         = "Phone number is not unique"
	ErrorMsgInvalidFields                  = "invalid fields"
	ErrorMsgUsernameRequired               = "username is required"
	ErrorMsgUsernameLength                 = "username must be between 3 and 50 characters"
	ErrorMsgPhoneNumberRequired            = "phone number is required"
	ErrorMsgPhoneNumberIsNotValid          = "phone number is not valid"
	ErrorMsgPasswordRequired               = "password is required"
	ErrorMsgPasswordLength                 = "password must be between 4 and 72 characters"
	ErrorMsgPhoneNumberOrPasswordIncorrect = "phone number or password not correct"
	ErrorMsgInvalidRequest                 = "invalid request payload"
	ErrorMsgTokenNotValid                  = "token is not valid"
	ErrorMsgRefreshTokenNotValid           = "refresh token is not valid"
)
