package todo

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"go-api/domain"
)

type ListTodosResponse_v0 struct {
	TodoList	[]*Todo		`json:"todoList"`
	LastID		int		`json:"last_id, omitempty"`
	Message		string		`json:"message,omitempty"`
}

type CreateTodoRequest_v0 struct {
	Todo NewTodo	`json:"todo"`
}

type CreateTodosResponse_v0 struct {
	Todo Todo		`json:"todo,omitempty"`
	Message string	`json:"message,omitempty"`
	Success bool	`json:"success"`
}

type ErrorResponse_v0 struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
}

func (resource *TodoResource) DecodeRequestBody(w http.ResponseWriter, req *http.Request, target interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		resource.RenderError(w, req, http.StatusBadRequest, fmt.Sprintf("Request body parse error: %v", err.Error()))
		return err
	}
	return nil
}

func (resource *TodoResource) RenderError(w http.ResponseWriter, req *http.Request, status int, message string) {
	http.Redirect(w, req, "/erro", http.StatusFound)
}

func (resource *TodoResource) RenderJsonError(w http.ResponseWriter, req *http.Request, status int, message string) {
	resource.Renderer.JSON(w, http.StatusOK, ErrorResponse_v0{
			Message: message,
			Success: false,
	})
}


/*
HandleListTodo_v0 lists todos
Example: 
Url: http://192.168.99.100:6060/api/todo
Header: Accept: application/json;version=1.0,*\/* -> sem o '\'
*/
func (resource *TodoResource) HandleListTodo_v0(w http.ResponseWriter, req *http.Request) {
	repo := resource.TodoRepository()

	//_ := req.FormValue("last_id")
	
	user := domain.GetAuthenticatedUserCtx(req)

	todoList, err := repo.FindAll(user.GetID())
	if err != nil {
		resource.RenderJsonError(w, req, http.StatusBadRequest, fmt.Sprintf("Erro retrieving list: %v", err.Error()))
		return
	}
	
	lastID := 0
	if len(todoList) > 0 {
		lastID = int(todoList[len(todoList)-1].GetID())
	}	
	
	resource.Renderer.JSON(w, http.StatusOK, ListTodosResponse_v0{
		TodoList:   todoList,
		LastID:  lastID,
		Message: "Todos list retrieved",
	})
}

func (resource *TodoResource) HandleIndexTodo_v0(w http.ResponseWriter, req *http.Request) {
	repo := resource.TodoRepository()

	//_ := req.FormValue("last_id")
	user := domain.GetAuthenticatedUserCtx(req)

	todoList, err := repo.FindAll(user.GetID())
	if err != nil {
		resource.RenderJsonError(w, req, http.StatusBadRequest, fmt.Sprintf("Erro retrieving list: %v", err.Error()))
		return
	}

	lastID := 0	
	if len(todoList) > 0 {
		lastID = int(todoList[len(todoList)-1].GetID())
	}
	
	resource.Renderer.HTML(w, http.StatusOK, "todos/index", ListTodosResponse_v0{
		TodoList:   todoList,
		LastID:  	lastID,
		Message: 	"Todos list retrieved",
	})
}

/*
Example: 
Url: http://192.168.99.100:6060/api/todo/create
Header: Accept: application/json;version=1.0,*\/* -> sem o '\'
Post:
	{
		"Todo" : {"name": "wut is dis"}
	}
*/
func (resource *TodoResource) HandleCreateTodo_v0(w http.ResponseWriter, req *http.Request) {

	repo := resource.TodoRepository()

	var body CreateTodoRequest_v0
	err := resource.DecodeRequestBody(w, req, &body)
	if err != nil {
		resource.Renderer.JSON(w, http.StatusOK, ErrorResponse_v0{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	
	user := domain.GetAuthenticatedUserCtx(req)

	var newTodo = Todo{
		Name: body.Todo.Name,
		Due:    time.Now(),
		Status:   TodoStatusPending,
		UserID: user.GetID(),
	}

	// ensure that todo obj is valid
	if !newTodo.IsValid() {
		resource.Renderer.JSON(w, http.StatusOK, ErrorResponse_v0{
			Message: "Todo is not valid",
			Success: false,
		})
		return
	}

	err = repo.CreateTodo(&newTodo)
	if err != nil {
		resource.Renderer.JSON(w, http.StatusOK, ErrorResponse_v0{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	
	resource.Renderer.JSON(w, http.StatusOK, CreateTodosResponse_v0{
		Todo:   newTodo,
		Message: "Lul dude",
		Success: true,
	})
}