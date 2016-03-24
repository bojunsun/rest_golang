package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("rest"))

func generalSession() (name string) {
	name = "rest_common"
	return
}

func checkLogin(r *http.Request) bool {
	session, _ := store.Get(r, generalSession())
	if session.Values["login"] == nil || session.Values["login"].(bool) == false {
		return false
	} else {
		return true
	}

}

func convertJson(input interface{}, w http.ResponseWriter) {
	var js []byte
	js, _ = json.Marshal(input)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", SignupPost).Methods("POST")
	router.HandleFunc("/login", LoginPost).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")
	router.HandleFunc("/getuser", UserselfGet).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}

/*
err的处理是经验，有三类，一类是系统内部api的某些函数，直接把error 省略，
第二，三种是跟用户的输入数据相关的，第二类是合法输入（比如信用卡信息）产生的错误，这些要严格检测错误返回到前端；
第三类是非法输入，比如没有登录就想看我们某些url，或者乱写id，
这些错误要在server包严格挡在外面（详见checkLogin, checkParam, checkId, checkAdmin这些函数）
*/
