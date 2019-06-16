package main
import(
	"fmt"
	"log"
	"time"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
ctx := r.Context()
log.Printf("handler started")
defer log.Printf("handler end")
select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w,"hello")
	case <- ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w,err.Error(),http.StatusInternalServerError)

  }
}

func main(){

http.HandleFunc("/",handler)
log.Fatal(http.ListenAndServe("192.168.0.166:8081",nil))

}
