package main
import(
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"

)
type user struct{
	Name string
	Age string
}
func snackWells( res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("hmac.html")
	if err != nil{
		log.Fatalln("Something went wrong: ", err)
	}
	x := user{
		Name: req.FormValue("Name"),
		Age: req.FormValue("Age"),
	}
	b, err := json.Marshal(x)
	if err != nil{
		fmt.Println("Error: ", err)
	}
	y :=base64.StdEncoding.EncodeToString(b)
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String()+"|"+getCode(id.String()),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	cookie.Value = cookie.Value + "|" + y + "|" + getCode(y)
	http.SetCookie(res, cookie)
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("H3110w0rld"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func main(){
	http.HandleFunc("/", snackWells)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
