package handler

import (
	//"encoding/json"
	"boot/config"
	"boot/consts"
	"boot/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Param struct {
	Id       string
	ApiToken string
}

func GetServiceResponse(optName string, p *Param) (map[string]interface{}, error) {
	var data map[string]interface{}
	url := fmt.Sprintf("%s/%s?api_token=%s", config.Url.HostPort, optName, p.ApiToken)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error while fetching data. Url:%s, Error:%s", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Info("GOT RESPONSE-S-- ", resp.Status)
	log.Info("GOT RESPONSE-B-- ", string(body))
	return data, nil
}
func GetShippedProjects(w http.ResponseWriter, r *http.Request) {
	if p, e := parseGetReq(r); e != nil {
		ProcessError(w, r, e)
		return
	} else {
		if out, err := GetServiceResponse(consts.SHIPPEDPROJECTS, p); err != nil {
			ProcessError(w, r, e)
		} else {
			ProcessResponse(w, r, out)
		}
	}
}

func GetShippedProjectServices(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Operation Not implemented yet")
}
func GetShippedProjectEnvs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Operation Not implemented yet")
}
func GetShippedBuildPacks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Operation Not implemented yet")
}
func GetShippedDependencies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Operation Not implemented yet")
}
func GetShippedProjectBuilds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Operation Not implemented yet")
}

func parseGetReq(r *http.Request) (*Param, error) {
	p := &Param{}
	if r.Method != "GET" {
		return nil, fmt.Errorf("Onlgy Get Request is configured right now.")
	}

	u, err := url.Parse(r.RequestURI)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse url: %s", r.RequestURI)
	}
	//parse query
	q, err := url.ParseQuery(u.RawQuery)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse query in url: %s", r.RequestURI)
	}
	if value := q.Get("id"); len(value) > 0 {

		p.Id = value
	}
	if value := q.Get("api_token"); len(value) > 0 {
		p.ApiToken = value
	} else {
		return nil, fmt.Errorf("Please specify 'api_token' as input param. ")
	}
	return p, nil
}
