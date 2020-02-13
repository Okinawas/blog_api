package codes

type Code string

cosnt (
  OK             Code = "OK"

  InvalidSyntax  Code = "invalid_syntax"
  BadParams      Code = "bad_params"
  EmptyBody      Code = "empty_body"
  InvalidRequest Code = "inavalid_request"

  Unauthorized   Code = "unauthorized"

  NotFound       Code = "not_found"

  Database       Code = "database_error"
  Redis          Code = "redis_error"
  Internal       Code = "internal_error"
  ExternalAPI    Code = "external_api_error"
  Google         Code = "google_error"
  Apple          Code = "apple_error"
  AWS            Code = "aws_error"

  Forbidden      Code = "forbidden"

  Unknown        Code = "unknown"
)
