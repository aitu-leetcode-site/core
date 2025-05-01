package httpcodes

import "net/http"

type HTTPCode int

func (h HTTPCode) StatusCode() int {
	return int(h)
}

func (h HTTPCode) Int() int {
	return int(h)
}

type SuccessCode = HTTPCode

const (
	StatusOK                   SuccessCode = http.StatusOK
	StatusCreated              SuccessCode = http.StatusCreated
	StatusAccepted             SuccessCode = http.StatusAccepted
	StatusNonAuthoritativeInfo SuccessCode = http.StatusNonAuthoritativeInfo
	StatusNoContent            SuccessCode = http.StatusNoContent
	StatusResetContent         SuccessCode = http.StatusResetContent
	StatusPartialContent       SuccessCode = http.StatusPartialContent
	StatusMultiStatus          SuccessCode = http.StatusMultiStatus
	StatusAlreadyReported      SuccessCode = http.StatusAlreadyReported
)

type UserErrorCode = HTTPCode

const (
	StatusBadRequest                   UserErrorCode = http.StatusBadRequest
	StatusUnauthorized                 UserErrorCode = http.StatusUnauthorized
	StatusPaymentRequired              UserErrorCode = http.StatusPaymentRequired
	StatusForbidden                    UserErrorCode = http.StatusForbidden
	StatusNotFound                     UserErrorCode = http.StatusNotFound
	StatusMethodNotAllowed             UserErrorCode = http.StatusMethodNotAllowed
	StatusNotAcceptable                UserErrorCode = http.StatusNotAcceptable
	StatusProxyAuthRequired            UserErrorCode = http.StatusProxyAuthRequired
	StatusRequestTimeout               UserErrorCode = http.StatusRequestTimeout
	StatusConflict                     UserErrorCode = http.StatusConflict
	StatusGone                         UserErrorCode = http.StatusGone
	StatusLengthRequired               UserErrorCode = http.StatusLengthRequired
	StatusPreconditionFailed           UserErrorCode = http.StatusPreconditionFailed
	StatusRequestEntityTooLarge        UserErrorCode = http.StatusRequestEntityTooLarge
	StatusRequestURITooLong            UserErrorCode = http.StatusRequestURITooLong
	StatusUnsupportedMediaType         UserErrorCode = http.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable UserErrorCode = http.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed            UserErrorCode = http.StatusExpectationFailed
	StatusTeapot                       UserErrorCode = http.StatusTeapot
	StatusMisdirectedRequest           UserErrorCode = http.StatusMisdirectedRequest
	StatusUnprocessableEntity          UserErrorCode = http.StatusUnprocessableEntity
	StatusLocked                       UserErrorCode = http.StatusLocked
	StatusFailedDependency             UserErrorCode = http.StatusFailedDependency
	StatusTooEarly                     UserErrorCode = http.StatusTooEarly
	StatusUpgradeRequired              UserErrorCode = http.StatusUpgradeRequired
	StatusPreconditionRequired         UserErrorCode = http.StatusPreconditionRequired
	StatusTooManyRequests              UserErrorCode = http.StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge  UserErrorCode = http.StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons   UserErrorCode = http.StatusUnavailableForLegalReasons
)

type ServerErrorCode = HTTPCode

const (
	StatusInternalServerError           ServerErrorCode = http.StatusInternalServerError
	StatusNotImplemented                ServerErrorCode = http.StatusNotImplemented
	StatusBadGateway                    ServerErrorCode = http.StatusBadGateway
	StatusServiceUnavailable            ServerErrorCode = http.StatusServiceUnavailable
	StatusGatewayTimeout                ServerErrorCode = http.StatusGatewayTimeout
	StatusHTTPVersionNotSupported       ServerErrorCode = http.StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates         ServerErrorCode = http.StatusVariantAlsoNegotiates
	StatusInsufficientStorage           ServerErrorCode = http.StatusInsufficientStorage
	StatusLoopDetected                  ServerErrorCode = http.StatusLoopDetected
	StatusNotExtended                   ServerErrorCode = http.StatusNotExtended
	StatusNetworkAuthenticationRequired ServerErrorCode = http.StatusNetworkAuthenticationRequired
)
