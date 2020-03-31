package main
import ("fmt"
"net/http")
func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Wel come Go Lang form aasksoft!")

}
func login_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Golan Login Page")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/login", login_handler)
	http.ListenAndServe(":8080",nil)
}