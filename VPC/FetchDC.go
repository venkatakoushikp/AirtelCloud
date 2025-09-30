package VPC

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://www.cv-prod-india-1.arista.io/api/resources/studio/v1/Inputs"
	studioID = "studio-evpn-services"
	TenantCount = 4 // Currently they are 4 Tenants Chennai-AZ1 Chennai-AZ2 Manesar-AZ1 Manesar-AZ2
)

var DCMAP = make(map[string]string)
var VRFs = make(map[int]map[int]string)
func SetHeaders(req *http.Request, token string){
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	temp:= fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", temp)
}

func FetchDC(token string, workspace_id string, i int){

	/* Removing this block since it would apply them in alphabetical order
		params := url.Values{}
		params.Set("key.studio_id", "studio-evpn-services")
		params.Set("key.workspace_id", workspace_id)
		params.Add("key.path.values", "tenants")
		params.Add("key.path.values", "")
		params.Add("key.path.values", "name")
	*/
	urlValues := fmt.Sprintf("key.studioId=%s&key.workspaceId=%s&key.path.values=tenants&key.path.values=%d&key.path.values=name", studioID, workspace_id,i) 
	Fullurl_ := baseURL + "?" + urlValues
	req ,err := http.NewRequest("GET", Fullurl_, nil)
	
	if err!= nil {
		fmt.Printf("%v", err)
	}
	SetHeaders(req, token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err!= nil{
		fmt.Printf("%v", err)
	}	
	defer response.Body.Close()
 	body, _ := io.ReadAll(response.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	result_ := result["value"].(map[string]interface{})
	value_:= result_["key"].(map[string]interface{})
	tenant_ := value_["path"].(map[string]interface{})["values"].([]interface {})
	DCMAP[result_["inputs"].(string)] = tenant_[1].(string)

}

func FetchVLAN(token string, workspace_id string, i int){
	urlValues := fmt.Sprintf("key.studioId=%s&key.workspaceId=%s&key.path.values=tenants&key.path.values=%d&key.path.values=vlans", studioID, workspace_id,i) 
	Fullurl_ := baseURL + "?" + urlValues
	req, err := http.NewRequest("GET", Fullurl_, nil)
	if err!=nil {
		fmt.Printf("%v", err)
	}
	SetHeaders(req, token)

	client := &http.Client{}
	response, err := client.Do(req)
	if err!= nil{
		fmt.Printf("%v", err)
	}	
	defer response.Body.Close()
 	body, _ := io.ReadAll(response.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	result_ := (result["value"].(map[string]interface{})["inputs"]).(string)
	InputsBytes := []byte(result_)
	var finalInput []map[string]interface{}
	json.Unmarshal(InputsBytes, &finalInput)
	finalInput_node = 


}

func VPCmain(token string, workspaceID string){
	for DC_COUNTER:=range TenantCount{
		FetchDC(token, workspaceID,DC_COUNTER)
	}
	fmt.Println("--------- Fetch DC --------------")
	for key,value := range DCMAP{
		fmt.Println(key, value)
	}
	fmt.Println("---------------------------------")
	FetchVLAN(token, workspaceID, 0)
		fmt.Println("--------- Fetch DC --------------")
	for key,value := range VRFs{
		fmt.Println(key, value)
	}
	fmt.Println("---------------------------------")
	
}

