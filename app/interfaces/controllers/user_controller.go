package controllers

import (
  "github.com/Okinawas/blog_api/app/domain"
  "github.com/Okinawas/blog_api/app/interfaces/database"
  "github.com/Okinawas/blog_api/app/usecase"
  "github.com/Okinawas/blog_api/errors"
  "github.com/Okinawas/blog_api/codes"
  "strconv"
)

type UserController struct {
  Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
  return &UserController{
    Interactor: usecase.UserInteractor{
      UserRepository: &database.UserRepository{
        SqlHandler: sqlHandler,
      },
    },
  }
}

//TODO: `NewError`のエラーハンドリングを書き換え
func (controller *UserController) Create(c Context) {
  u := domain.User{}
  c.Bind(&u)
  err := controller.Interactor.Add(u)
  if err != nil {
    c.JSON(500, errors.Errorf(codes.Database, "Failed to create user: %s", err))
    return
  }

  // TODO: （調査）structで定義したinterface{}は握りつぶしても良いのか
  c.JSON(201)
}

func (controller *UserController) Index(c Context) {
  users, err := controller.Interactor.Users()
  if err != nil {
    c.JSON(500, errors.Errorf(codes.NotFound, "Failed to find user: %s", err))
    return
  }

  c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := controller.Interactor.UserById(id)
  if err != nil {
    c.JSON(500, errors.Errorf(codes.NotFInd, "Failed to create user: %s", err))
    return
  }
  c.JSON(200, user)
}

