package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Abc struct {
	Unique_platform_id      string `json:"unique_platform_id"`
	Is_sandbox              bool   `json:"is_sandbox"`
	Oauth_app_installed     bool   `json:"oauth_app_installed"`
	Is_active               bool   `json:"is_active"`
	Auth_type               string `json:"auth_type"`
	Sync_source_instance_id int    `json:"sync_source_instance_id"`
	Auth_status             string `json:"auth_status"`
	Id                      int    `json:"id"`
}

type Response_1 struct {
	Sync_source_instances []Abc `json:"sync_source_instances"`
}

type Response_2 struct {
	Is_valid             bool   `json:"is_valid"`
	Username             string `json:"username"`
	Platform_username    string `json:"platform_username"`
	Is_oauth_active      bool   `json:"is_oauth_active"`
	Authorization_status string `json:"authorization_status"`
	Updated              string `json:"updated"`
}

func MyNewURI(clientID string, IsSandbox string) ([]byte, int) {
	myURL := "https://organization-external.prod1.6si.com/api/v1/oauth/org/" + clientID + "/authorize/get_sync_source_instances/?sync_source_type=Salesforce&is_sandbox=" + IsSandbox
	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	first_response, _ := ioutil.ReadAll(res.Body)
	return first_response, res.StatusCode
}

func MySecondURI(clientID string, SyncSrcID int) ([]byte, int) {
	myURL := "https://organization-external.prod1.6si.com/api/v1/oauth/org/" + clientID + "/authorize/get_oauth_integration_details/?sync_source_instance_id=" + strconv.Itoa(SyncSrcID)
	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	response, _ := ioutil.ReadAll(res.Body)
	return response, res.StatusCode
}

func DecodeFirstJSON(JSON []byte) Response_1 {
	var resp1out Response_1
	if json.Valid(JSON) {
		err := json.Unmarshal(JSON, &resp1out)
		if err != nil {
			panic(err)
		}
		fmt.Println("Unique_platform_id: ", resp1out.Sync_source_instances[0].Unique_platform_id)
		fmt.Println("is_sandbox: ", resp1out.Sync_source_instances[0].Is_sandbox)
		fmt.Println("oauth_app_installed: ", resp1out.Sync_source_instances[0].Oauth_app_installed)
		fmt.Println("is_active: ", resp1out.Sync_source_instances[0].Is_active)
		fmt.Println("auth_type: ", resp1out.Sync_source_instances[0].Auth_type)
		fmt.Println("sync_source_instance_id: ", resp1out.Sync_source_instances[0].Sync_source_instance_id)
		fmt.Println("auth_status: ", resp1out.Sync_source_instances[0].Auth_status)
	}
	return resp1out
}

func DecodeSecondJSON(JSON []byte) Response_2 {
	var resp2out Response_2
	if json.Valid(JSON) {
		err := json.Unmarshal(JSON, &resp2out)
		if err != nil {
			panic(err)
		}
		fmt.Println("Is_valid: ", resp2out.Is_valid)
		fmt.Println("Username: ", resp2out.Username)
		fmt.Println("Platform_username: ", resp2out.Platform_username)
		fmt.Println("Is_oauth_active: ", resp2out.Is_oauth_active)
		fmt.Println("Authorization_status: ", resp2out.Authorization_status)
		fmt.Println("Updated: ", resp2out.Updated)

	}
	return resp2out
}

func myFunc(clientID string, IsSandbox string) Response_2 {
	responseFronFirst, stts := MyNewURI(clientID, IsSandbox)
	/* if stts != 200 {
		panic("ERR while reading the JSON, error code: " + strconv.Itoa(stts))
	} */
	response1 := DecodeFirstJSON(responseFronFirst)
	_ = response1
	/*
		unique_platform_id := response1.Sync_source_instances[0].Unique_platform_id
		is_sandbox := response1.Sync_source_instances[0].Is_sandbox
		oauth_app_installed := response1.Sync_source_instances[0].Oauth_app_installed
		is_active := response1.Sync_source_instances[0].Is_active
		auth_type := response1.Sync_source_instances[0].Auth_type
		auth_status := response1.Sync_source_instances[0].Auth_status
	*/
	sync_source_instance_id := response1.Sync_source_instances[0].Sync_source_instance_id
	responseFromSecond, stts := MySecondURI(clientID, sync_source_instance_id)
	/* if stts != 200 {
		panic("ERR while reading the JSON, error code: " + strconv.Itoa(stts))
	} */
	response2 := DecodeSecondJSON(responseFromSecond)
	_ = stts

	return response2
	/*
		is_valid := response2.Is_valid
		username := response2.Username
		platform_username := response2.Platform_username
		is_oauth_active := response2.Is_oauth_active
		authorization_status := response2.Authorization_status
	*/

}

func main() {
	http.HandleFunc("/", myFunc2)
	http.ListenAndServe(":9090", nil)

	//myFunc("4", "false")
}

func myFunc2(res http.ResponseWriter, req *http.Request) {
	clientID := req.URL.Query().Get("clientID")
	fmt.Println(clientID)
	FinalRes := myFunc(clientID, "false")
	_ = FinalRes
}
