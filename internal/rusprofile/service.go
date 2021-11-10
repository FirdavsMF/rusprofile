package rusprofile

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	grpc "rusprofile/pkg/grpc"
	"strconv"
)

type service struct {
	host string
	get  func(url string) (resp *http.Response, err error)
}

func NewService(host string) Service {
	return &service{host: host, get: http.Get}
}

func (s *service) GetCompany(inn int64) (*grpc.Company, error) {
	resp, err := s.get(s.buildURL(fmt.Sprint(inn)))
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var newCompany = &grpc.Company{}

	document.Find(".copy_target").Each(func(i int, selection *goquery.Selection) {
		switch i {
		case 1:
			inn, _ := strconv.Atoi(selection.Text())
			newCompany.Inn = int64(inn)
		case 2:
			kpp, _ := strconv.Atoi(selection.Text())
			newCompany.Kpp = int64(kpp)
		}
	})

	document.Find(".company-name").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			newCompany.Name = selection.Text()
		}
	})

	document.Find(".gtm_main_fl").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			newCompany.OwnerName = selection.Text()
		}
	})

	return newCompany, nil
}

func (s *service) buildURL(query string) string {
	return fmt.Sprintf("%s?query=%s", s.host, query)
}
