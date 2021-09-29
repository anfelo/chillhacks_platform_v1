package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
)

type JobHandler struct {
	store    courses.Store
	sessions *scs.SessionManager
}

func (h *JobHandler) RunMigrations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath, _ := filepath.Abs("./migrations")
		files, err := getUpMigrationsFileNames(filePath)
		if err != nil {
			log.Fatal(err)
		}

		errs := make(map[string]interface{})
		for _, file := range files {
			b, err := ioutil.ReadFile(filePath + "/" + file)
			if err != nil {
				log.Fatal(err)
			}
			s := string(b)
			if err := h.store.RunMigration(s); err != nil {
				errs[file] = err.Error()
			} else {
				errs[file] = "migration created"
			}
		}

		http_utils.RespondJson(w, http.StatusOK, errs)
	}
}

func getUpMigrationsFileNames(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), ".up.sql") {
			files = append(files, file.Name())
		}
	}
	return files, nil
}
