package api

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/errors"
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
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		errs := make(map[string]interface{})
		for _, file := range files {
			b, err := ioutil.ReadFile(filePath + "/" + file)
			if err != nil {
				restErr := errors.NewInternatServerError("internal server error")
				http_utils.RespondJson(w, restErr.Status, restErr)
				return
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
