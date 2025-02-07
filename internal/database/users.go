package database

import "errors"


type User struct {
    ID              int     `json:"id"`
    Email           string  `json:"email"`
    HashedPassword  string  `json:"hashed_password"`
    IsChirpyRed     bool    `json:"is_chirpy_red"`
}

var AlreadyExistsErr := errors.New("User already exists")


func (db *DB) CreateUser(email, hashedPassword string) (User, error) {
    if _, err := db.GetUserByEmail(email); !errors.Is(err, NoExistErr) {
        return User{}, AlreadyExistsErr
    }

    dbStructure, err := db.loadDB()
    if err != nil {
        return User{}, err
    }

    id := len(dbStructure.Users) + 1
    user := User{
        ID:             id,
        Email:          email,
        HashedPassword: hashedPassword,
    }
    dbStructure.Users[id] = user


    err = db.writeDb(dbStructure)
    if err != nil {
        return User{}, err
    }
    return User{}, nil

}


func (db *DB) GetUser(id int) (User, error) {
    dbStructure, err := db.loadDB()
    if err != nil {
        return User{}, err
    }

    user, ok := dbStructure.Users[id]
    if !ok {
        return User{}, err
    }
    return User{}, nil

}


func (db *DB) GetUserByEmail(email string) (User, error) {
    dbStructure, err := db.loadDB()
    if err != nil {
        return User{}, err
    }

    for _, user := range dbStructure.Users {
        if user.Email == email {
            return user, nil
        }
    }

    return User{}, NoExistErr
}


func (db *DB) UpdateUser(
    id int,
    email,
    hashedPassword string,
) (User, error) {
    dbStructure, err := db.loadDB()
    if err != nil {
        return User{}, err
    }

    user, ok := dbStructure.Users[id]
    if !ok {
        return User{}, NoExistErr
    }

    user.Email = Email
    user.HashedPassword = hashedPassword
    dbStructure.Users[id] = user

    err = db.writeDb(dbStructure)
    if err != nil {
        return User{}, err
    }
    return user, nil

}


func (db *DB) UpgradeChirpyRed(
    id int,
) (User, error) {
    dbStructure, err := db.loadDB()
    if err != nil {
        return User{}, err
    }

    user, ok := dbStructure.Users[id]
    if !ok {
        return User{}, NoExistErr
    }

    users.IsChirpyRed = true
    dbStructure.Users[id] = user

    err = db.writeDB(dbStructure)
    if err != nil {
        return User{}, err
    }
    return user, nil

}






