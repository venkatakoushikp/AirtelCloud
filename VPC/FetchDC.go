package VPC

import (
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://www.cv-prod-india-1.arista.io/api/resources/studio/v1/Inputs"
	studioID = "studio-evpn-services"
	TenantCount = 4 // Currently they are 4 Tenants Chennai-AZ1 Chennai-AZ2 Manesar-AZ1 Manesar-AZ2
	
)

/*

{
  "value": {
    "key": {
      "studioId": "studio-evpn-services",
      "workspaceId": "",
      "path": {
        "values": [
          "tenants",
          "0",
          "name"
        ]
      }
    },
    "createdAt": "2025-04-01T05:13:37.233414941Z",
    "createdBy": "Siddharth-Arista",
    "lastModifiedAt": "2025-09-29T10:37:25.988347618Z",
    "lastModifiedBy": "Coredge",
    "inputs": "\"Chennai-AZ1\""
  },
  "time": "2025-09-30T04:50:40.465841712Z"
}


*/

type Response struct {
	Value `json:"value"`
	Time `json:"time`
}

var DCMAP = make(map[int]string)

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
	fmt.Println(Fullurl_)
	req ,err := http.NewRequest("GET", Fullurl_, nil)
	
	if err!= nil {
		fmt.Printf("%v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	temp:= fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", temp)
	client := &http.Client{}
	response, err := client.Do(req)

	if err!= nil{
		fmt.Printf("%v", err)
	}
	
	defer response.Body.Close()
	fmt.Println("Status Code:", response.Status)
 	body, _ := io.ReadAll(response.Body)
	body1 := string(body)
	fmt.Print(body1)


}

func VPCmain(token string, workspaceID string){
	for DC_COUNTER:=range TenantCount{
		FetchDC(token, workspaceID,DC_COUNTER)
	}
	
}

