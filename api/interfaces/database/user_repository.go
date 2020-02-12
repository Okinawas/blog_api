package database

import "github.com/Okinawas/blog_api/api/domain"


type UserRepository struct {
  SqlHandler
}


func (repo *UsrRepository) Store(u domain.User)(id int, err error) {
  result, err := repo.Execute(
    "INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName,
  )
  if err != nil {
    return
  }

  id64, err := result.LastInsertId()
  if err != nil {
    return
  }

  id = int(id64)
  return
}


func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
  row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)

  defer row.Close()
  if err != nil {
    return
  }

  var id int
  var firstName string
  var lastNmae string

  if err = row.Scan(&id, &firstName, &lastName); err != nil {
    return
  }

  user.ID = id
  user.FirstName = firstName
  user.LastName = lastName
  return
}

func (repo *UserRepository) FindAll() (users domain.Usres, err error) {
  rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
  defer rows.Close()
  if err != nil {
    return
  }

  for rows.Next(){
    var id int
    var firstName string
    var lastName string
    if err := rows.Scan(&id, &firstName, &lastName); err != nil {
      continue
    }

    users := domain.User{
      ID:        id,
      FirstName: firstName,
      LastName:  lastName,
    }
    users = append(users, user)
  }

  return
}

