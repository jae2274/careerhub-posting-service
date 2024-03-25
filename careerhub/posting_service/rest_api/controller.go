package restapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/rest_api/dto"
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

func (restApiCtrler *RestApiController) RegisterRoutes(rootPath string) {
	restApiCtrler.router.HandleFunc(rootPath+"/job_postings", restApiCtrler.GetJobPostings).Methods("GET")
	restApiCtrler.router.HandleFunc(rootPath+"/job_postings/{site}/{postingId}", restApiCtrler.GetJobPostingDetail).Methods("GET")
	restApiCtrler.router.HandleFunc(rootPath+"/categories", restApiCtrler.GetCategories).Methods("GET")
	restApiCtrler.router.HandleFunc(rootPath+"/skills", restApiCtrler.GetSkills).Methods("GET")
}

func (restApiCtrler *RestApiController) GetJobPostings(w http.ResponseWriter, r *http.Request) {

	reqCtx := r.Context()

	var req dto.GetJobPostingsRequest
	err := req.Set(r)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobPostings, err := restApiCtrler.service.GetJobPostings(reqCtx, &req)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Failed to get job postings", http.StatusInternalServerError)
		return
	}

	// jobPostings를 JSON으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") //TODO: 이후 세부적으로 설정 필요
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"job_postings": jobPostings,
	})
}

func (restApiCtrler *RestApiController) GetJobPostingDetail(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	vars := mux.Vars(r)
	site, ok := vars["site"]
	if !ok {
		llog.Msg("Invalid site value").Level(llog.ERROR).Data("site", site).Log(reqCtx)
		http.Error(w, "Invalid site value", http.StatusBadRequest)
		return
	}

	postingId, ok := vars["postingId"]
	if !ok {
		llog.Msg("Invalid postingId value").Level(llog.ERROR).Data("postingId", postingId).Log(reqCtx)
		http.Error(w, "Invalid postingId value", http.StatusBadRequest)
		return
	}

	jobPostingDetail, err := restApiCtrler.service.GetJobPostingDetail(reqCtx, site, postingId)

	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Failed to get job posting detail", http.StatusInternalServerError)
		return
	}

	// jobPostingDetail을 JSON으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") //TODO: 이후 세부적으로 설정 필요
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobPostingDetail)
}

func (restApiCtrler *RestApiController) GetCategories(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	categories, err := restApiCtrler.service.GetAllCategories(reqCtx)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}

	// categories를 JSON으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") //TODO: 이후 세부적으로 설정 필요
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func (restApiCtrler *RestApiController) GetSkills(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()

	skills, err := restApiCtrler.service.GetAllSkills(reqCtx)
	if err != nil {
		llog.LogErr(reqCtx, err)
		http.Error(w, "Failed to get skills", http.StatusInternalServerError)
		return
	}

	// skills를 JSON으로 변환하여 응답
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") //TODO: 이후 세부적으로 설정 필요
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"skills": skills,
	})
}
