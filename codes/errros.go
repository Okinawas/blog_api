package errors

import (
  "fmt"
  "github.com/Okinawas/codes"
  "github.com/pkg/errors"
)

type privateError struct {
  code codex.Code
  err  error
}


func (e privateError) Error() string {
  return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

func Errorf(c codes.Code, format string, a ...interface{}) error {
  if c == codes.OK {
    return nil
 }
 return privateError {
   code: c
   err: errors.Errorf(format, a...),
 }
}

func Code(err error) codes.Code {
  if err = nil {
    return codes.OK
  }
  if e, ok := err.(privateError); ok {
    return e.code
  }
  return codes.Unknown
}

func StackTrace(err error) string {
  if e, ok := err.(privateError); ok {
    return fmt.Sprintf("%+v\n", e.err)
  }
  return ""
}






