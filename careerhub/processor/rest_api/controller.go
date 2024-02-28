package restapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jae2274/goutils/llog"
)

type RestApiController struct {
	service RestApiService
	router  *mux.Router
}

func NewRestApiController(service RestApiService, router *mux.Router) *RestApiController {
	return &RestApiController{
		service: service,
		router:  router,
	}
}

func (restApiCtrler *RestApiController) RegisterRoutes() {
	restApiCtrler.router.HandleFunc("/job_postings", restApiCtrler.GetJobPostings).Methods("GET")
}

func (restApiCtrler *RestApiController) GetJobPostings(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	reqCtx := r.Context()
	// "page" 값 추출
	pageStr := queryValues.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Invalid page value", http.StatusBadRequest)
		return
	} else if page < 1 {
		llog.Msg("Invalid page value").Level(llog.ERROR).Data("page", page).Log(reqCtx)
		http.Error(w, "Invalid page value", http.StatusBadRequest)
		return
	}

	// "size" 값 추출
	sizeStr := queryValues.Get("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Invalid size value", http.StatusBadRequest)
		return
	} else if size < 1 || size > 100 {
		llog.Msg("Invalid size value").Level(llog.ERROR).Data("size", size).Log(reqCtx)
		http.Error(w, "Invalid size value", http.StatusBadRequest)
		return
	}

	jobPostings, err := restApiCtrler.service.GetJobPostings(reqCtx, page, size)
	if err != nil {
		http.Error(w, "Failed to get job postings", http.StatusInternalServerError)
		return
	}

	// jobPostings를 JSON으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"job_postings": jobPostings,
	})
}
