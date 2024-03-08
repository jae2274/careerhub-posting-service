package dto

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jae2274/goutils/llog"
)

type GetJobPostingsRequest struct {
	Page     int
	Size     int
	QueryReq QueryReq
}

type QueryReq struct {
	Categories []cateogoryQuery `json:"categories"`
	// skillNames: []
	// tagIds: []
	// minCareer: null
	// maxCareer: null
}

type cateogoryQuery struct {
	Site         string
	CategoryName string
}

const initPage = 0

func (req *GetJobPostingsRequest) Set(r *http.Request) error {

	reqCtx := r.Context()

	queryValues := r.URL.Query()

	// "page" 값 추출
	pageStr := queryValues.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return fmt.Errorf("invalid page value. %s", pageStr)
	} else if page < initPage {
		llog.Msg("Invalid page value").Level(llog.ERROR).Data("page", page).Log(reqCtx)
		return fmt.Errorf("invalid page value. %d", page)
	}

	// "size" 값 추출
	sizeStr := queryValues.Get("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size value. %s", sizeStr)
	} else if size < 1 || size > 100 {
		return fmt.Errorf("invalid size value. %d", size)
	}

	queryReq, err := GetQuery(queryValues.Get("encoded_query"))
	if err != nil {
		return err
	}

	req.Page = page
	req.Size = size
	req.QueryReq = queryReq

	return nil
}

func GetQuery(encodedQuery string) (QueryReq, error) {

	bytes, err := base64.StdEncoding.DecodeString(encodedQuery)
	if err != nil {
		query := string(bytes)
		log.Println(query)
		return QueryReq{}, fmt.Errorf("invalid encoded_query value. failed to decode. %s", encodedQuery)
	}

	var queryReq QueryReq
	err = json.Unmarshal(bytes, &queryReq)
	if err != nil {
		return QueryReq{}, fmt.Errorf("invalid encoded_query value. failed to unmarshal. %s", string(bytes))
	}

	return queryReq, nil // TODO
}
