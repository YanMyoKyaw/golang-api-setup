package router
import (
    	"github.com/gorilla/mux"
    	c "test/controller"
)
var Router *mux.Router
func init() {
	Router = mux.NewRouter()
	Router.HandleFunc("/user/login",c.Login).Methods("PUT")
	Router.HandleFunc("/user/register",c.Register).Methods("POST")
	Router.HandleFunc("/insert",c.Insert).Methods("PUT")
	
}