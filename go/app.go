package main
import(
    "net/http"
    "runtime"
    "log"
    "io"
)
func handleRequest(w http.ResponseWriter, req*http.Request) {
    io.WriteString(w, "helloworld!")
}
func main() {
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8088",nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
