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
	"github.com/tmtk75/go-oauth2/oauth2"
	"github.com/tmtk75/go-oauth2/oauth2/github"
	"golang.org/x/net/websocket"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//    {{.UserData.name}}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("client", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		// v.name !!
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
		fmt.Println("try redirect1")
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

	callbackURL := os.Getenv("CallbackURL")

	oauth2.WithProviders(
		github.New(oauth2.NewConfig(oauth2.GITHUB, callbackURL+oauth2.GITHUB)),
		// facebook.New(oauth2.NewConfig(oauth2.FACEBOOK, "http://localhost:8080/auth/callback/"+oauth2.FACEBOOK)),
		// google.New(oauth2.NewConfig(oauth2.GOOGLE, "http://localhost:8080/auth/callback/"+oauth2.GOOGLE)),
		// slack.New(oauth2.NewConfig(oauth2.SLACK, "http://localhost:8080/auth/callback/"+oauth2.SLACK)),
	)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	providerName := segs[3]
	switch action {
	case "login":
		loginURL := oauth2.ProviderByName(providerName).Config().AuthCodeURL("state")
		fmt.Println("try redirect2")
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		code := r.FormValue("code")

		provider := oauth2.ProviderByName(providerName)
		profile, err := oauth2.ProfileByCode(provider, code)

		if err != nil {
			log.Println("Failed to Profile", providerName, "-", err)
		}

		// grimmer
		// token, err := provider.Config().Exchange(xoauth2.NoContext, code)
		// fmt.Print("token2:", token)
		// fmt.Print("access token:", profile.Token())
		// profile, err := provider.Profile(t) -> send request

		fmt.Print("github access token:", profile.Token().AccessToken)
		_, err = getStarredInfo(profile.Nickname(), profile.Token().AccessToken)
		if err != nil {
			log.Println("cant not get starred info.")
		}
		fmt.Println("try redirect3")
		saveSession(w, profile)
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
	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   base64.StdEncoding.EncodeToString([]byte(msg)),
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
}

func main() {
	fmt.Println("start main")
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
	r.Static("/styles", "client/styles")

	r.GET("/", gin.WrapH(mustAuth(&templateHandler{filename: "index.html"})))
	r.GET("/login", gin.WrapH(&templateHandler{filename: "templates/login.html"}))
	r.GET("/logout", func(c *gin.Context) {
		fmt.Println("try redirect4")
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

	// r.NoRoute(func(c *gin.Context) {
	// 	fmt.Println("try redirect5")
	// 	w := c.Writer
	// 	w.Header()["Location"] = []string{"/"}
	// 	w.WriteHeader(http.StatusTemporaryRedirect)
	// })

	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	// signalChan := make(chan os.Signal, 1)
	// // signal.Notify(signalChan, os.Interrupt)
	// signal.Notify(signalChan, os.Interrupt, syscall.s
	// 	syscall.SIGHUP,
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGQUIT)

	// go func() {
	// 	for _ = range signalChan {
	// 		fmt.Println("\nReceived an interrupt, stopping services...\n")
	// 		os.Exit(1)
	// 	}
	// }()
	// http://nathanleclaire.com/blog/2014/08/24/handling-ctrl-c-interrupt-signal-in-golang-programs/
	// https://gist.github.com/reiki4040/be3705f307d3cd136e85

	//
	log.Println("Start web server. Port: ", *addr)
	if err := r.Run(*addr); err != nil {
		log.Fatal("Run:", err)
	}

	log.Println("end of main")

}
