package main

// Дз 2 Создать элементарные формы регистрации и авторизации на html.
// Создать эндпоинты и обработчики что бы:
// Регистрироваться( пароли хранить в хешированом виде)
// Авторизоваться.
// Отправлять запросы из этих форм.

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	id       uint64
	name     string
	email    string
	password string
}

var exist bool

// HashPassword - encoding string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - checking if encoded password == password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//CacheUsers - type struct, that has mutex for including map of users
type CacheUsers struct {
	data map[uint64]user
	sync.RWMutex
}

//InitUsers - initialize CacheUsers struct
func InitUsers() CacheUsers {
	return CacheUsers{
		data: make(map[uint64]user),
	}
}

//cache of all users
var cacheUsers = InitUsers()

//returns user from map by id
func (c *CacheUsers) getUser(id uint64) user {
	c.RLock()
	tempUser := c.data[id]
	c.RUnlock()
	return tempUser
}

//adds user to map
func (c *CacheUsers) setUser(u user) {
	c.Lock()
	c.data[u.id] = u
	c.Unlock()
}

//gets length of map of users
func (c *CacheUsers) getLen() uint64 {
	c.Lock()
	len := uint64(len(c.data))
	c.Unlock()
	return len
}

func main() {
	fmt.Println("[Server started]")

	myHash, _ := HashPassword("123")
	firstUser := user{0, "SaYaku", "example@gmail.com", myHash}

	cacheUsers.setUser(firstUser)

	// start page, that uses html-file (localhost/)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})

	http.HandleFunc("/postreg", func(w http.ResponseWriter, r *http.Request) {
		newName := r.FormValue("newname")
		email := r.FormValue("email")
		newPass := r.FormValue("newpass")

		// checking empty boxes
		if newName != "" && email != "" && newPass != "" {

			// checking existing account
			cacheUsers.RLock()
			for _, user := range cacheUsers.data {
				if user.name == newName || user.email == email {
					fmt.Fprintf(w, "Oops! This name or email is already taken!\n")
					exist = true
					break
				}
			}
			cacheUsers.RUnlock()

			//register new user
			if exist == false {
				hash, err := HashPassword(newPass)
				if err != nil {
					fmt.Fprintf(w, "Error in encoding password!\n%v\n", err)
				} else {
					newUser := user{cacheUsers.getLen(), newName, email, hash}
					cacheUsers.setUser(newUser)
					fmt.Fprintf(w, "Successfully registered!\n")
				}
			}
		} else {
			fmt.Fprintf(w, "You haven't filled all the boxes!\n")
		}
		exist = false
		return
	})

	http.HandleFunc("/postlog", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		pass := r.FormValue("pass")

		if name != "" && pass != "" {
			//checking name and password
			cacheUsers.RLock()
			for _, user := range cacheUsers.data {
				if user.name == name && CheckPasswordHash(pass, user.password) {
					fmt.Fprintf(w, "Welcome back, %s\n", name)
					break
				}
			}
			cacheUsers.RUnlock()
		} else {
			fmt.Fprintf(w, "You haven't filled all the boxes!\n")
		}
		return
	})

	http.ListenAndServe(":80", nil)
}
