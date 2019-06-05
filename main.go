package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"	
    "strconv"
)

func main() {
    http.HandleFunc("/create", func (w http.ResponseWriter, r *http.Request) {
	    url1 := "https://www.activecampaign.com/api.php"
	    client := &http.Client{}
	    form := url.Values{}
        form.Add("account", "asdgaergoinqer4")
        form.Add("plan", "0")
        form.Add("name", "Erik")
        form.Add("notification", "etetland+demo@vendasta.com")
        form.Add("email", "etetland+demo@vendasta.com")
        form.Add("cname", "")
        req, _ := http.NewRequest("POST", url1, strings.NewReader(form.Encode()))
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
        q := req.URL.Query()
        q.Add("api_key", "KEYYYYYY")
        q.Add("api_action", "account_add")
        q.Add("api_output", "json")
        req.URL.RawQuery = q.Encode()
        fmt.Fprintf(w, req.URL.String())
        resp, err := client.Do(req)
        if err != nil {
        	fmt.Fprintf(w, "\n error: " + err.Error())
        	return
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Printf("%s", body)

	})
    http.HandleFunc("/apikey", func (w http.ResponseWriter, r *http.Request) {
	    url1 := "https://asdgaergoinqer4.activehosted.com/admin/api.php"
	    client := &http.Client{}
        req, _ := http.NewRequest("POST", url1, nil)
        q := req.URL.Query()
        q.Add("api_action", "user_me")
        q.Add("api_output", "json")
        q.Add("api_user", "admin")
        q.Add("api_pass", "PASSSSSS")
        req.URL.RawQuery = q.Encode()
        fmt.Fprintf(w, req.URL.String())
        resp, err := client.Do(req)
        if err != nil {
        	fmt.Fprintf(w, "\n error: " + err.Error())
        	return
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Printf("%s", body)

	})
	http.HandleFunc("/token", func (w http.ResponseWriter, r *http.Request) {
	    url1 := "http://asdgaergoinqer4.api-us1.com/api.php"
	    client := &http.Client{}
	    form := url.Values{}
        // form.Add("ln", c.ln)
        // form.Add("ip", c.ip)
        // form.Add("ua", c.ua)
        req, _ := http.NewRequest("POST", url1, strings.NewReader(form.Encode()))
        q := req.URL.Query()
        q.Add("api_key", "KEYYYYYY")
        q.Add("api_action", "singlesignon")
        q.Add("api_output", "json")
        q.Add("sso_addr", "71.17.243.75")
        q.Add("sso_user", "admin")
        q.Add("sso_duration", strconv.Itoa(7*60))
        req.URL.RawQuery = q.Encode()
        fmt.Fprintf(w, req.URL.String())
        resp, err := client.Do(req)
        if err != nil {
        	fmt.Fprintf(w, "err")
        	return
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        fmt.Printf("%s", body)
	})
	
	http.HandleFunc("/redirect", func (w http.ResponseWriter, r *http.Request) {
        url1 := "http://asdgaergoinqer4.api-us1.com/admin/main.php?_ssot=518f7fe616ed9daca16e82e79a736847"
        http.Redirect(w, r, url1, http.StatusSeeOther)
	})

	http.ListenAndServe(":8090", nil)
}
