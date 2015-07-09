package errors

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidParam  = &RequestError{ErrorString: "Invalid Parameter", ErrorCode: http.StatusBadRequest}
	ErrInternalError = &RequestError{ErrorString: "Internal Error", ErrorCode: http.StatusInternalServerError}
	ErrNotFound      = &RequestError{ErrorString: "Request Not Found", ErrorCode: http.StatusNotFound}
	ErrUnauthorized  = &RequestError{ErrorString: "Unauthorized", ErrorCode: http.StatusUnauthorized}

	ErrNoIb           error = errors.New("Imageboard id required")
	ErrNoThread       error = errors.New("Thread id required")
	ErrCommentLong    error = errors.New("Comment too long")
	ErrCommentShort   error = errors.New("Comment too short")
	ErrNoComment      error = errors.New("Comment is required")
	ErrTitleLong      error = errors.New("Title too long")
	ErrTitleShort     error = errors.New("Title too short")
	ErrNoTitle        error = errors.New("Title is required")
	ErrNameLong       error = errors.New("Name too long")
	ErrNameShort      error = errors.New("Name too short")
	ErrNoTagId        error = errors.New("Tag id required")
	ErrNoTagType      error = errors.New("Tag type required")
	ErrTagLong        error = errors.New("Tag too long")
	ErrTagShort       error = errors.New("Tag too short")
	ErrNoTagName      error = errors.New("Tag name required")
	ErrDuplicateTag   error = errors.New("Duplicate tag")
	ErrNoImage        error = errors.New("Image is required for new threads")
	ErrDuplicateImage error = errors.New("Duplicate image")
	ErrNoImageId      error = errors.New("Image id required")
	ErrInvalidCookie  error = errors.New("Invalid cookie")
	ErrNoCookie       error = errors.New("Cookie required")
	ErrInvalidKey     error = errors.New("Invalid key")
	ErrNoKey          error = errors.New("Antispam key required")
	ErrThreadClosed   error = errors.New("Thread is closed")
	ErrIpParse        error = errors.New("Input IP cannot be parsed")
)

type RequestError struct {
	ErrorString string
	ErrorCode   int
}

func (err *RequestError) Code() int {
	return err.ErrorCode
}

func (err *RequestError) Error() string {
	return err.ErrorString
}

func ErrorMessage(error_type *RequestError, args ...map[string]interface{}) (code int, message map[string]interface{}) {
	code = error_type.Code()
	message = map[string]interface{}{"error_message": error_type.Error()}

	return
}
