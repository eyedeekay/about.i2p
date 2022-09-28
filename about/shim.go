package about

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"cerca/database"
	"cerca/server"

	"github.com/eyedeekay/about.i2p/about/html"
)

func SafeDirectory(u *server.CercaForum) string {
	if u.Directory == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		u.Directory = filepath.Join(dir, "CercaForum")
	}
	os.MkdirAll(u.Directory, 0755)
	return u.Directory
}

func generateTemplates() (*template.Template, error) {
	views := []string{
		"about",
		"footer",
		"generic-message",
		"head",
		"index",
		"login",
		"login-component",
		"new-thread",
		"register",
		"register-success",
		"thread",
		"password-reset",
		"change-password",
		"change-password-success",
	}

	rootTemplate := template.New("root")

	for _, view := range views {
		newTemplate, err := rootTemplate.Funcs(server.TemplateFuncs).ParseFS(html.Templates, fmt.Sprintf("%s.html", view))
		if err != nil {
			return nil, fmt.Errorf("could not get files: %w", err)
		}
		rootTemplate = newTemplate
	}

	return rootTemplate, nil
}

// NewServer sets up a new CercaForum object. Always use this to initialize
// new CercaForum objects. Pass the result to http.Serve() with your choice
// of net.Listener.
func NewServer(allowlist []string, sessionKey, dir string) (*server.CercaForum, error) {
	server.Templates = template.Must(generateTemplates())
	s := &server.CercaForum{
		ServeMux:  http.ServeMux{},
		Directory: dir,
	}

	dbpath := filepath.Join(SafeDirectory(s), "forum.db")
	db := database.InitDB(dbpath)

	/* note: be careful with trailing slashes; go's default handler is a bit sensitive */
	// TODO (2022-01-10): introduce middleware to make sure there is never an issue with trailing slashes
	handler := server.NewRequestHandler(db, sessionKey, allowlist)
	s.ServeMux.HandleFunc("/reset/", handler.ResetPasswordRoute)
	s.ServeMux.HandleFunc("/about", handler.AboutRoute)
	s.ServeMux.HandleFunc("/logout", handler.LogoutRoute)
	s.ServeMux.HandleFunc("/login", handler.LoginRoute)
	s.ServeMux.HandleFunc("/register", handler.RegisterRoute)
	s.ServeMux.HandleFunc("/post/delete/", handler.DeletePostRoute)
	s.ServeMux.HandleFunc("/thread/new/", handler.NewThreadRoute)
	s.ServeMux.HandleFunc("/thread/", handler.ThreadRoute)
	s.ServeMux.HandleFunc("/robots.txt", handler.RobotsRoute)
	s.ServeMux.HandleFunc("/", handler.IndexRoute)

	fileserver := http.FileServer(http.FS(html.Templates))
	s.ServeMux.Handle("/assets/", fileserver) //http.StripPrefix("/assets/", fileserver))
	return s, nil
}
