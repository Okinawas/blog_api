package controllers

import (
  "github/nntakuya/matsun_blogs/app/domain"
  "github/nntakuya/matsun_blogs/app/interfaces/database"
  "github/nntakuya/matsun_blogs/app/usecase"
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

