package workspaceCreation

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"math/rand"

)

const url_ = "https://www.cv-prod-india-1.arista.io/api/resources/workspace/v1/WorkspaceConfig"
func CreateWorkspace(token string, workspaceName string) bool{

	payload := fmt.Sprintf(`
{
"key": {
	"workspaceId": "%s"
},
"displayName": "TAC BLR",
"description": "Test Script for Airtel Cloud",
"requestParams": {
	"requestId": "%d"
}
}
	`, workspaceName, rand.Intn(100000))
	fmt.Print(payload)
	req ,err := http.NewRequest("POST", url_, strings.NewReader(payload))
	
	if err!= nil {
		fmt.Printf("%v", err)
		return false 
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	temp:= fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", temp)

	client := &http.Client{}
	response, err := client.Do(req)

	if err!= nil{
		fmt.Printf("%v", err)
		return false
	}
	
	defer response.Body.Close()
	fmt.Println("Status Code:", response.Status)
 	body, err := io.ReadAll(response.Body)
 	if err != nil {
 		
 	}
 	fmt.Println("Response Body:", string(body))
	return true


}
