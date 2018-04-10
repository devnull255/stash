// forms.go
package main

import (
	"html/template"
        "log"
	"net/http"
	"os"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func addContactDetails(cd ContactDetails) error {
	sess, err := session.NewSession()
	if err != nil {
	    return fmt.Errorf("Error creating session: %s", err)
        }
	svc := s3.New(sess)
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	f, _ := os.Create("/var/log/golang/golang-server.log")
	defer f.Close()
	log.SetOutput(f)
	tmpl := template.Must(template.ParseFiles("public/forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":"+port, nil)
}
