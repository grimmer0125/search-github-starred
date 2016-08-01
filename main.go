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

	"github.com/sfreiberg/gotwilio"

	"github.com/gin-gonic/gin"
	"github.com/tmtk75/go-oauth2/oauth2"
	"github.com/tmtk75/go-oauth2/oauth2/github"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//    {{.UserData.name}}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("/login or /")

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
	log.Println("/-1")
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/auth*")

	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	providerName := segs[3]
	switch action {
	case "login":
		loginURL := oauth2.ProviderByName(providerName).Config().AuthCodeURL("state")
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		code := r.FormValue("code")

		provider := oauth2.ProviderByName(providerName)
		profile, err := oauth2.ProfileByCode(provider, code)

		if err != nil {
			log.Println("Failed to Profile !!!! ", providerName, "-", err)
			// c.JSON(401, gin.H{"status": "unauthorized"})

			// Write([]byte) (int, error)
			// w.WriteHeader(http.StatusUnauthorized)

			// return
		} else {

			account := profile.Nickname()

			if account != "" && profile.Token() != nil && profile.Token().AccessToken != "" {

				fmt.Print("github access token:", profile.Token().AccessToken)

				SetUserOrJustUpdateToken(account, profile.Token().AccessToken)
				// prepareUserStarredRepo(profile.Nickname(), profile.Token().AccessToken)

				saveSession(w, profile)

			} else {

				log.Println("Got Profile but info something worng !!!! ")

				// w.WriteHeader(http.StatusUnauthorized)
				// return
			}
		}

		//		saveSession(w, profile)
		w.Header()["Location"] = []string{"/"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unsupported action: %s", action)
	}
}

func saveSession(w http.ResponseWriter, u oauth2.Profile) {
	msg, _ := json.Marshal(map[string]interface{}{
		"name":      u.Nickname(),
		"authToken": u.Token().AccessToken,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   base64.StdEncoding.EncodeToString([]byte(msg)),
		Path:    "/",
		Expires: time.Now().AddDate(1, 0, 0), //Add(24 * time.Hour),
	})
}

var userMap map[string]*GitHubUser
var mux sync.Mutex

func init() {
	log.Println("main init !!!!!!")

	callbackURL := os.Getenv("CallbackURL")

	oauth2.WithProviders(
		github.New(oauth2.NewConfig(oauth2.GITHUB, callbackURL+oauth2.GITHUB)),
	)
	// prepare userMap
	userMap = make(map[string]*GitHubUser)
	_ = userMap
}

// func prepareUserStarredRepo(account string, token string) {
// account           string
// accessToken       string

// 之後再做1.
// 1. 如果再run當中的就不要再new/run了,
//
// 2. 假設前一個token還可以用,
// 那此時第二個同一user的token好像會已經先傳回去? 還像同時兩個token也還好,

// user := GitHubUser{account, token, NOTSTART, 0, 0}
// setupUserToMap(account, &user)

// go user.GetStarredInfo()

// _, err = getStarredInfo(profile.Nickname(), profile.Token().AccessToken)
// if err != nil {
// 	log.Println("cant not get starred info.")
// }
// }

func getReposActionHandler(c *gin.Context) {

	log.Println("getReposActionHandler")
	r := c.Request

	if authCookie, err := r.Cookie("auth"); err == nil {
		var v map[string]interface{}
		b, _ := base64.StdEncoding.DecodeString(authCookie.Value)
		err := json.Unmarshal([]byte(b), &v)

		if err != nil {
			log.Fatalf("Failed to Unmarshal: %v\n", err)

		} else if account2, ok2 := v["name"]; ok2 == true {
			account := account2.(string)
			if user, _ := GetUser(account); user != nil {
				log.Println("found out account:", user.Account)
				user.Status = NOTSTART
				SetUser(user.Account, *user)
			}
		}
	}

	cleanCookieAndToLoginPage(c)
	// c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "",
	// })
}

func getReposHandler(c *gin.Context) {

	log.Println("/repos")
	r := c.Request
	ok := false
	message := "no valid auth"

	// if userMap == nil {
	// 	log.Println("usermap is nil")
	// }

	if authCookie, err := r.Cookie("auth"); err == nil {
		var v map[string]interface{}
		b, _ := base64.StdEncoding.DecodeString(authCookie.Value)
		err := json.Unmarshal([]byte(b), &v)

		if err != nil {
			log.Fatalf("Failed to Unmarshal: %v\n", err)

		} else if account2, ok2 := v["name"]; ok2 == true {
			account := account2.(string)
			if user, _ := GetUser(account); user != nil {
				log.Println("found out account:", user.Account)

				tokenInCookie := v["authToken"].(string)

				// try to compare tokens
				for _, token := range user.Tokens {
					if token == tokenInCookie {
						log.Println("found out the same token")
						ok = true
						break
					}
				}

				if ok {
					// start logic here
					if user.Status == NOTSTART {
						// TODO: if two clients use the same account simutaneously, still have race condition problems
						go user.GetStarredInfo(tokenInCookie)

						// prepareUserStarredRepo(account, tokenInCookie)
					}

					c.JSON(200, gin.H{
						"status":        user.Status,
						"numOfStarred":  user.NumOfStarred,
						"githubAccount": account,
					})
				} else {
					log.Println("can not found out the same token in DB")
				}

			} else {
				log.Println("does not have the same key, force logout")

				http.SetCookie(c.Writer, &http.Cookie{
					Name:   "auth",
					Value:  "",
					Path:   "/",
					MaxAge: -1,
				})
			}
		}
	} else {
		message = "does not have cookie"
	}

	if ok == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
	}
}

func cleanCookieAndToLoginPage(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w := c.Writer
	// w.Header()["Location"] = []string{"/login"}
	w.Header().Set("location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func sendTwilioAlert(repo string) {
	fmt.Println("send twilio alert")
	accountSid := "AC83651bf8e21c30b313a44ccb97db3688"
	authToken := "ee7f411bf34d3424bc8cd4c934193079"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+12016279052"
	to := "+886963052251"
	message := "Over 10k limit for Algolia api:" + repo
	twilio.SendSMS(from, to, message, "", "")
}

func main() {

	// testRedis()
	// os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
	// os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))

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
	r.GET("/logout", cleanCookieAndToLoginPage)
	r.GET("/auth/*action", gin.WrapF(loginHandler))
	// r.GET("/clock", gin.WrapH(websocket.Handler(func(ws *websocket.Conn) {
	// 	for {
	// 		fmt.Fprint(ws, "{When:'"+time.Now().Format(time.RFC3339)+"'}")
	// 		time.Sleep(1 * time.Second)
	// 	}
	// })))
	r.GET("/repos/*action", getReposActionHandler)
	r.GET("/repos", getReposHandler)

	log.Println("Start web server. Port: ", *addr)
	if err := r.Run(*addr); err != nil {
		log.Fatal("Run:", err)
	}

	log.Println("end of main")

}
