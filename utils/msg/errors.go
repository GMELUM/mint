package msg

// ErrorInvalidProtocol contains a pre-serialized message pack-format error
// indicating missing required fields or type mismatches.
var ErrorForbidden = serializeJson(Data{
	Error: &ErrorData{
		Code:     0,
		Message:  "Forbidden",
		Critical: true,
	},
})

// ErrorInvalidProtocol contains a pre-serialized message pack-format error
// indicating missing required fields or type mismatches.
var ErrorUnauthorized = serializeJson(Data{
	Error: &ErrorData{
		Code:     1,
		Message:  "Unauthorized",
		Critical: true,
	},
})

// ErrorInvalidProtocol contains a pre-serialized message pack-format error
// indicating missing required fields or type mismatches.
var ErrorInvalidProtocol = serializeJson(Data{
	Error: &ErrorData{
		Code:     3,
		Message:  "Required fields are missing or their type does not match the declared one",
		Critical: true,
	},
})

// ErrorNoAccount contains a pre-serialized message pack-format error
// indicating that the account is not linked to any user page.
var ErrorNoAccount = serializeJson(Data{
	Error: &ErrorData{
		Code:     4,
		Message:  "This account is not linked to any user page",
		Critical: false,
	},
})

// ErrorManyRequest contains a pre-serialized message pack-format error
// indicating too many requests.
var ErrorManyRequest = serializeJson(Data{
	Error: &ErrorData{
		Code:     5,
		Message:  "Too many requests",
		Critical: false,
	},
})

// ErrorInvalidFields contains a pre-serialized message pack-format error
// indicating invalid or missing fields.
var ErrorInvalidFields = serializeJson(Data{
	Error: &ErrorData{
		Code:     6,
		Message:  "Required fields are missing or their type does not match the declared one",
		Critical: true,
	},
})

// ErrorExpiration contains a pre-serialized message pack-format error
// indicating token expiration.
var ErrorExpiration = serializeJson(Data{
	Error: &ErrorData{
		Code:     7,
		Message:  "Token expiration",
		Critical: true,
	},
})

// ErrorServiceWork contains a pre-serialized message pack-format error
// indicating that technical work is underway.
var ErrorServiceWork = serializeJson(Data{
	Error: &ErrorData{
		Code:     8,
		Message:  "Technical work is underway",
		Critical: true,
	},
})

// ErrorOutdatedVersion contains a pre-serialized message pack-format error
// indicating an outdated version of the application.
var ErrorOutdatedVersion = serializeJson(Data{
	Error: &ErrorData{
		Code:     9,
		Message:  "Outdated version of the application",
		Critical: true,
	},
})
