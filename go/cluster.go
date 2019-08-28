package main
import(
    "net/http"
    "log"
    "io"
)
func handleRequest(w http.ResponseWriter, req*http.Request) {
    io.WriteString(w, "helloworld!")
}
func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8088",nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
