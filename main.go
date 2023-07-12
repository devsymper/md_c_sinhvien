package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Todo struct {
	Name  string `json:"Name"`
	State bool   `json:"State"`
}

type Todolist struct {
	list []Todo
	m    sync.Mutex
}

func Listconstructor() Todolist {
	return Todolist{
		list: make([]Todo, 0),
		m:    sync.Mutex{},
	}
}

func (tdl *Todolist) showtodo() []Todo {
	return tdl.list
}

type CreateRQ struct {
	Name string `json:"title"`
}

func (tdl *Todolist) Create(createTodoRequest CreateRQ) Todo {
	tdl.m.Lock()
	defer tdl.m.Unlock()
	newTodo := Todo{
		Name:  createTodoRequest.Name,
		State: false,
	}

	tdl.list = append(tdl.list, newTodo)

	return newTodo
}

func main() {
	tdl := Listconstructor()
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		list := tdl.showtodo()
		return c.JSON(http.StatusOK, list)
	})
	e.POST("/create", func(c echo.Context) error {
		request := CreateRQ{}
		err := c.Bind(&request)
		if err != nil {
			return err
		}

		Todo := tdl.Create(request)
		return c.JSON(http.StatusCreated, Todo)
	})
	e.Start(":1811")
}
