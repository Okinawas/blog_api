package controllers

import (
  "github.com/Okinawas/blog_api/api/domain"
  "github.com/Okinawas/blog_api/api/interfaces/database"
  "github.com/Okinawas/blog_api/api/usecase"
  "strconv"
)

type UserController struct {
  Interactor usecase.UserInteractor
}

func (controller *UserController) Create(c Context) {
  u := domain.User{}
  c.Bind(&u)
  err := controller.Interactor.Add(u)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }

  c.JSON(201)
}

func (controller *UserController) Show(c Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := controller.Interactor.UserById(id)
  if err != nil {
    c.JSON(500, NewError(err))
    return
  }
  c.JSON(200, user)
}

