package msg

import (
    "github.com/gin-gonic/gin"
    jsoniter "github.com/json-iterator/go"
)

type ErrorData struct {
    Code     int    `json:"code"`
    Message  string `json:"message"`
    Critical bool   `json:"critical,omitempty"`
}

type Data struct {
    ResponseID uint64     `json:"response_id,omitempty"`
    Error      *ErrorData `json:"error,omitempty"`
    Response   any        `json:"response,omitempty"`
}

var ContentType = "application/json; charset=utf-8"

func serializeJson(data Data) []byte {
    message, err := jsoniter.Marshal(&data)
    if err != nil {
        return []byte{}
    }
    return message
}

func Send(ctx *gin.Context, data any) {
    ctx.Data(200, ContentType, serializeJson(Data{
        Response: data,
    }))
}

func CustomError(ctx *gin.Context, code int, data string) {
    ctx.Data(200, ContentType, serializeJson(Data{
        Error: &ErrorData{
            Code:     code,
            Message:  data,
            Critical: true,
        },
    }))
}

// BadRequest sends a response with an error message indicating a bad request.
// It uses the CustomError function to set the error with a specific code and message.
func BadRequest(ctx *gin.Context, data string) {
    CustomError(ctx, 2, data)
}

// Forbidden sends a response with an error message indicating forbidden access.
// The error is sent as a JSON formatted response using the provided context.
func Forbidden(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorForbidden)
    ctx.Abort()
}

// Unauthorized sends a response with an error message indicating unauthorized access.
// The error is sent as a JSON formatted response using the provided context.
func Unauthorized(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorUnauthorized)
    ctx.Abort()
}

// InvalidProtocol sends a response with an error message indicating protocol issues.
// The error is sent as a JSON formatted response using the provided context.
func InvalidProtocol(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorInvalidProtocol)
    ctx.Abort()
}

// NoAccount sends a response with an error message indicating that the account is not linked.
// The error is sent as a JSON formatted response using the provided context.
func NoAccount(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorNoAccount)
    ctx.Abort()
}

// ManyRequest sends a response with an error message indicating too many requests.
// The error is sent as a JSON formatted response using the provided context.
func ManyRequest(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorManyRequest)
    ctx.Abort()
}

// InvalidFields sends a response with an error message indicating invalid fields.
// The error is sent as a JSON formatted response using the provided context.
func InvalidFields(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorInvalidFields)
    ctx.Abort()
}

// Expiration sends a response with an error message indicating token expiration.
// The error is sent as a JSON formatted response using the provided context.
func Expiration(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorExpiration)
    ctx.Abort()
}

// ServiceWork sends a response with an error message indicating technical work.
// The error is sent as a JSON formatted response using the provided context.
func ServiceWork(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorServiceWork)
    ctx.Abort()
}

// OutdatedVersion sends a response with an error message indicating an outdated version.
// The error is sent as a JSON formatted response using the provided context.
func OutdatedVersion(ctx *gin.Context) {
    ctx.Data(200, ContentType, ErrorOutdatedVersion)
    ctx.Abort()
}