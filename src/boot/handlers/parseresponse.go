package handler

import (
	//"bytes"
	//"encoding/json"
	//"encoding/xml"
	"fmt"
	"net/http"
	//"runtime/debug"
	"boot/log"
	//"strings"
)

func ProcessError(w http.ResponseWriter, r *http.Request, e error) {
	if e != nil {
		fmt.Fprintf(w, "Error: %s", e.Error())
		log.Error(e)
		//fmt.Printf("STACK: %s", string(debug.Stack()))
	}
}

func ProcessResponse(w http.ResponseWriter, r *http.Request, d map[string]interface{}) {
	if len(d) < 1 {
		ProcessError(w, r, fmt.Errorf("No Response found"))
	}
	//if aHdr := r.Header.Get(consts.HEADER_ACCEPT); strings.EqualFold(aHdr, consts.APPLICATION_XML) {
	//	parseXML(w, r, d)
	//} else {
	//	parseJSON(w, r, d)
	//}
}

//func parseXML(w http.ResponseWriter, r *http.Request, d []*db.Contact) {

//	w.Header().Set(consts.HEADER_CONTENTTYPE, consts.APPLICATION_XML)

//}
//func parseJSON(w http.ResponseWriter, r *http.Request, d []*db.Contact) {

//	w.Header().Set(consts.HEADER_CONTENTTYPE, consts.APPLICATION_JSON)
//	res, err := json.MarshalIndent(d, "", "  ")

//	fmt.Fprint(w, string(res))
//}
