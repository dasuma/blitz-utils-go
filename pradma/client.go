package pradma

import (
	"fmt"
	"net/http"

	generic "github.com/dasuma/blitz-utils-go/models"
	"github.com/dasuma/blitz-utils-go/rest"
)

var (
	cityServices = "https://blitz-city.azurewebsites.net"
)

//PradmaService service
var PradmaService pradmaService = &pradmaImpl{}

type pradmaImpl struct {
}

type pradmaService interface {
	GetCity(name string, year int) (interface{}, error)
}

func (s *pradmaImpl) GetCity(name string, year int) (interface{}, error) {
	url := fmt.Sprintf("%s/ms-city/city/%s?year=%v", cityServices, name, year)
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	var src interface{}
	res, err := rest.HTTPRequest(http.MethodGet, url, nil, &src, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}
