package application

import (
	"ddd-go/domain"
	"ddd-go/infrastructure"
	"fmt"
	"net/http"
	"strings"
)

func Search(vendorCode string, request domain.SearchRequest) (int, domain.SearchResponse, domain.ErrorResponse) {
	vendorCode = strings.ToUpper(vendorCode)
	switch vendorCode {
	case "SJ":
		valid, response := request.Validate()
		if !valid {
			return http.StatusBadRequest, domain.SearchResponse{}, response
		} else {
			return http.StatusOK, infrastructure.SjSearch(request), domain.ErrorResponse{}
		}
	default:
		message := fmt.Sprintf("Vendor Code '%s' is not available", vendorCode)
		sr := domain.SearchResponse{domain.BaseResponse{false, message, ""}, domain.SearchResponseData{}}
		return http.StatusOK, sr, domain.ErrorResponse{}
	}
}

func Book() {

}

func BookInfo(request domain.BookInfoRequest) (int, domain.BookInfoResponse, domain.ErrorResponse) {
	valid, response := request.Validate()
	if !valid {
		return http.StatusBadRequest, domain.BookInfoResponse{}, response
	} else {
		return http.StatusOK, infrastructure.SjBookInfo(request), domain.ErrorResponse{}
	}
}

func SetPayment() {

}