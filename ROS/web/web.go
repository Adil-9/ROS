package web

import (
	"fmt"
	"net/http"
	"ros/dictionary"
	"ros/structure"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/home.html")
		if err != nil {
			Error_page(w, r)
			return
		}
		err = tmpl.Execute(w, "")
		if err != nil {
			Error_page(w, r)
			return
		}
	case "POST":
		r.ParseForm()
		var tosend structure.SENT
		tosend.Name = r.Form.Get("name")
		tosend.Surname = r.Form.Get("surname")
		// tosend.Email = r.Form.Get("email")
		// tosend.Login = r.Form.Get("login")
		tosend.Year = r.Form.Get("year")
		tosend.Password = r.Form.Get("password")
		tosend.Relation = r.Form["passwordlist"][0]

		if !dictionary.Validate(tosend.Password) {
			tmpl_exec(w, r, "Not allowed charachters in password")
			return
		} else if tosend.Password == "" {
			tmpl_exec(w, r, "Empty password not allowed")
			return
		}

		bool_rockyou, err := dictionary.RockYou(tosend.Password)

		var err_rockyou string
		var found_in_rockyou string

		if err != nil {
			err_rockyou = "Can not open rockyou file"
		}

		if bool_rockyou {
			found_in_rockyou = "The password is found in rockyou dictionary."
		}

		if tosend.Relation == "no relation" {
			strength := dictionary.RatePassword(tosend.Password)
			if strength <= 50 {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is very weak. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			} else if strength < 65 {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is weak. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			} else if strength < 75 {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is not weak but not realy reliable. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			} else if strength < 85 {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is good. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			} else if strength <= 100 {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is very Good. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			} else {
				tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d,\npassword is IMPOSSIBLY STRONG. %s%s", strength, found_in_rockyou, err_rockyou))
				return
			}
		} else if tosend.Relation == "relation" {
			strength := dictionary.RatePassword(tosend.Password)
			msg := dictionary.Depending(strength, tosend)

			tmpl_exec(w, r, fmt.Sprintf("Level of strength is : %d, \n%s", strength, msg))
			return
		}

		fmt.Printf("\nName: %s\nSurname: %s\nPassword: %s\nRelation: %s", tosend.Name, tosend.Surname, tosend.Password, tosend.Relation)

		tmpl_exec(w, r, "Some message")
	}
}
func Error_page(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error!"))
}

func tmpl_exec(w http.ResponseWriter, r *http.Request, msg string) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		Error_page(w, r)
		return
	}
	err = tmpl.Execute(w, msg)
	if err != nil {
		Error_page(w, r)
		return
	}
}
