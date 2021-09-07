package api

import (
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
)

type SeederHandler struct {
	store courses.Store
}

func (h *SeederHandler) Seed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ss []courses.Subject
		s1 := courses.Subject{Title: "Software Fundamentals", Slug: "software-fundamentals", ImgURL: "software-fundamentals.png"}
		s2 := courses.Subject{Title: "Python", Slug: "python", ImgURL: "python.png"}
		s3 := courses.Subject{Title: "VueJS", Slug: "vuejs", ImgURL: "vue.png"}
		ss = append(ss, s1, s2, s3)

		for i, s := range ss {
			if err := h.store.CreateSubject(&s); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ss[i] = s
		}

		var cc []courses.Course
		c1 := courses.Course{
			SubjectID:   ss[0].ID,
			Title:       "Design Patterns",
			Description: "Design patterns are typical solutions to common problems in software design. Each pattern is like a blueprint that you can customize to solve a particular design problem in your code.",
			Slug:        "design-patterns",
			ImgURL:      "design-patterns.png",
		}
		c2 := courses.Course{
			SubjectID:   ss[0].ID,
			Title:       "Algorithms",
			Description: "In mathematics and computer science, an algorithm is a finite sequence of well-defined, computer-implementable instructions, typically to solve a class of specific problems or to perform a computation. Algorithms are always unambiguous and are used as specifications for performing calculations, data processing, automated reasoning, and other tasks.",
			Slug:        "algorithms",
			ImgURL:      "algorithms.png",
		}
		c3 := courses.Course{
			SubjectID:   ss[1].ID,
			Title:       "Python Practical Guide",
			Description: "Python is an interpreted high-level general-purpose programming language. Its design philosophy emphasizes code readability with its use of significant indentation. Its language constructs as well as its object-oriented approach aim to help programmers write clear, logical code for small and large-scale projects.",
			Slug:        "python-practical-guide",
			ImgURL:      "python.png",
		}
		c4 := courses.Course{
			SubjectID:   ss[2].ID,
			Title:       "Vue Composition Api",
			Description: "What makes the Vue 3 Composition API so much better than the Options API is code sharing. Inside the setup hook of a component, we can group parts of our code by logical concern. We then can extract pieces of reactive logic and share the code with other components.",
			Slug:        "vue-composition-api",
			ImgURL:      "python.png",
		}
		cc = append(cc, c1, c2, c3, c4)

		for i, c := range cc {
			if err := h.store.CreateCourse(&c); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			cc[i] = c
		}

		var ll []courses.Lesson
		l1 := courses.Lesson{
			CourseID: cc[0].ID,
			Title:    "Factory Method",
			Content:  "Here goes some content",
			Slug:     "factory-method",
			Category: "creational-patterns",
			Order:    1,
		}
		ll = append(ll, l1)

		for i, l := range ll {
			if err := h.store.CreateLesson(&l); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ll[i] = l
		}

		res := make(map[string]interface{})
		res["subjects"] = ss
		res["courses"] = cc
		res["lessons"] = ll

		http_utils.RespondJson(w, http.StatusCreated, res)
	}
}
