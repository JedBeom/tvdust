package main

import (
	"github.com/JedBeom/airq"
	"github.com/kpango/glg"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Grade string
	Bg string
}

func index() http.HandlerFunc {
	var (
		// init sync.Once
		t *template.Template
	)

	return func(w http.ResponseWriter, r *http.Request) {
		/*init.Do(func(){
			t = template.Must(template.ParseFiles("tmpl/main.html"))
		})
		*/

		t = template.Must(template.ParseFiles("tmpl/main.html"))

		q, err := airq.NowByStation("연향동")
		if err != nil {
			_ = glg.Error(err)
		}

		err = t.Execute(w, setGrade(q.Pm10GradeWHO, q.Pm25GradeWHO))
		if err != nil {
			log.Println(err)
		}

		_ = glg.Info("Pm10:", q.Pm10GradeWHO, "Pm25:", q.Pm25GradeWHO)
	}
}

