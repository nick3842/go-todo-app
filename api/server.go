// package api

// import (
// 	"net/http"
// 	"github.com/google/uuid"
// 	"github.com/gorilla/mux"
// )

// type todo struct {
// 	ID uuid.UUID `json:"id"`
// 	Name string `json:"name"`
// }

// type Server struct {
// 	*mux.Router

// 	todos []todo
// }

// func NewServer() *Server {
// 	s := &Server {
// 		Router : mux.NewRouter(),
// 		todos: []todo{},
// 	}
// 	return s
// }

// func (s *Server) createTodo() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 			var newTodo todo
// 			if err := json.NewDecoder(r.Body).Decode(&i); err!=nil {
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}

// 			newTodo.ID = uuid.New()
// 			s.todos = append(s.todos, newTodo)

// 			w.Header().Set("Content-Type", "application/json")
// 			if err := json.NewEncoder(w).Encode(newTodo); err !=nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 		}
// }

// func (s *Server) removeTodo() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		idStr, _ := mux.Vars(r)["id"]
// 		id, err := uuid.Parse(idStr)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 		}

// 		for i, item := range s.todos {
// 			if item.ID == id {
// 				s.todos = append(s.todos[:i], s.todos[i+1:]...)
// 				break
// 			}
// 		}
// 	}
// }

// func (s *Server) listTodos() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(s.todos); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

