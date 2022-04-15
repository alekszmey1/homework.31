package app

import (
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/go-chi/chi"
	"homework.31/pkg/controller"
	"homework.31/pkg/repository"
	"homework.31/pkg/usecase"
	"net/http"
)

func Run() {
	repository := repository.NewRepository()
	usecase := usecase.NewUsecase(repository)
	controller := controller.NewController(usecase)
	mux := http.NewServeMux()
	mux.HandleFunc("/create", controller.CreateUser)
	mux.HandleFunc("/make_friends", controller.MakeFriends)
	mux.HandleFunc("/delete", controller.DeleteUser)
	mux.HandleFunc("/get_friends", controller.GetFriends)
	mux.HandleFunc("/put", controller.UpdateAge)
	http.ListenAndServe("localhost:8080", mux)
}

//	mux.HandleFunc("/get", srv.GetAll	)

/*router := chi.NewRouter()
router.Use(.Logger)
router.Use(.Recoverer)*/
//router.Post("/create", controller.CreateUser)
/*http.ListenAndServe(":3000", router)
}*/
