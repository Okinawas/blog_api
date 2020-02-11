package controllers

type Context interfaec {
  Param(string) string
  Bind(interface{}) error
  Status(int)
  JSON(int, interface{})
}

