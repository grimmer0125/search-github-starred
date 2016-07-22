// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.LoadHTMLGlob("templates/*.tmpl.html")

// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl.html", nil)
// 	})

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"

	"github.com/tmtk75/go-oauth2/oauth2"
	// "github.com/tmtk75/go-oauth2/oauth2/facebook"
	"github.com/tmtk75/go-oauth2/oauth2/github"
	// "github.com/tmtk75/go-oauth2/oauth2/google"
	// "github.com/tmtk75/go-oauth2/oauth2/slack"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("client/partials", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		var v map[string]interface{}
		log.Println("authCookie: ", authCookie)
		b, _ := base64.StdEncoding.DecodeString(authCookie.Value)
		err := json.Unmarshal([]byte(b), &v)
		if err != nil {
			log.Fatalf("Failed to Unmarshal: %v\n", err)
		}
		data["UserData"] = v
	}
	data["Providers"] = oauth2.Providers()
	t.templ.Execute(w, data)
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func mustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func init() {
	oauth2.WithProviders(
		github.New(oauth2.NewConfig(oauth2.GITHUB, "http://localhost:8080/auth/callback/"+oauth2.GITHUB)),
		// facebook.New(oauth2.NewConfig(oauth2.FACEBOOK, "http://localhost:8080/auth/callback/"+oauth2.FACEBOOK)),
		// google.New(oauth2.NewConfig(oauth2.GOOGLE, "http://localhost:8080/auth/callback/"+oauth2.GOOGLE)),
		// slack.New(oauth2.NewConfig(oauth2.SLACK, "http://localhost:8080/auth/callback/"+oauth2.SLACK)),
	)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		loginURL := oauth2.ProviderByName(provider).Config().AuthCodeURL("state")
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		p := oauth2.ProviderByName(provider)
		// p := oauth2.ProviderByName("github")
		u, err := oauth2.ProfileByCode(p, r.FormValue("code"))
		if err != nil {
			log.Println("Failed to Profile", provider, "-", err)
		}

		saveSession(w, u)
		w.Header()["Location"] = []string{"/"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unsupported action: %s", action)
	}
}

func saveSession(w http.ResponseWriter, u oauth2.Profile) {
	msg, _ := json.Marshal(map[string]interface{}{
		"name": u.Name(),
	})
	log.Println("msg: ", string(msg))
	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   base64.StdEncoding.EncodeToString([]byte(msg)),
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	port2 := fmt.Sprintf(":%s", port)

	var addr = flag.String("addr", port2, "application address")
	flag.Parse()

	//
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/build-client", "build-client")

	r.GET("/", gin.WrapH(mustAuth(&templateHandler{filename: "index.html"})))
	r.GET("/login", gin.WrapH(&templateHandler{filename: "login.html"}))
	r.GET("/logout", func(c *gin.Context) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w := c.Writer
		w.Header()["Location"] = []string{"/login"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	r.GET("/auth/*action", gin.WrapF(loginHandler))
	r.GET("/clock", gin.WrapH(websocket.Handler(func(ws *websocket.Conn) {
		for {
			fmt.Fprint(ws, "{When:'"+time.Now().Format(time.RFC3339)+"'}")
			time.Sleep(1 * time.Second)
		}
	})))

	//
	log.Println("Start web server. Port: ", *addr)
	if err := r.Run(*addr); err != nil {
		log.Fatal("Run:", err)
	}
}
