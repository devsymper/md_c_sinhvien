package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Todo struct {
	Name  string `json:"name"`
	State bool   `json:"state"`
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
	Name string `json:"name"`
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

func (tdl *Todolist) SetStatus(NameCheck string) error {
	tdl.m.Lock()
	defer tdl.m.Unlock()
	var t *Todo
	var place int = -1
	for i, j := range tdl.list {
		if j.Name == NameCheck {
			t = &j
			place = i
			break
		}
	}
	if place == -1 {
		return echo.ErrNotFound
	}
	if t.State {
		err := echo.ErrBadRequest
		err.Message = "Todo is already completed."
		return err
	}
	tdl.list[place].State = true
	return nil
}

func (tdl *Todolist) Del(NameCheck string) error {
	tdl.m.Lock()
	defer tdl.m.Unlock()
	var place int = -1
	for i, j := range tdl.list {
		if j.Name == NameCheck {
			place = i
			break
		}
	}
	if place == -1 {
		return echo.ErrNotFound
	}
	tdl.list = append(tdl.list[:place], tdl.list[place+1:]...)
	return nil
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
	e.PATCH("/:name/state", func(c echo.Context) error {
		name := c.Param("name")
		err := tdl.SetStatus(name)
		if err != nil {
			c.Error(err)
			return err
		}
		return nil
	})
	e.DELETE("/:name", func(c echo.Context) error {
		name := c.Param("name")
		err := tdl.Del(name)
		if err != nil {
			c.Error(err)
			return err
		}
		return nil
	})
	e.Start(":1811")
}
