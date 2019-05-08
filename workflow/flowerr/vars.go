package flowerr

var ErrUnknow = err(101, "")

var ErrRequest = err(201, "")

var ErrNil = err(301, "nil")

var ErrUnsupportedConnectType = err(302, "err unsupported connect type")

var ErrNode = err(303, "err node")

var NextFlow = err(304, "next flow")

var ScriptTrue = err(401, "script pass check")

var ScriptFalse = err(402, "script no passing")

var ErrScriptResult = err(403, "result value error")

var ErrSubmitFormFieldNil = err(501, "")

var ErrSubmitFormFieldType = err(502, "")

var ErrSubmitFormFieldRegex = err(503, "")

var ErrSubmitFormFieldValue = err(504, "")

var ErrInvalidInstance = err(601, "")
