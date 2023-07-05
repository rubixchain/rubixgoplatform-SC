package client

import (
	"fmt"

	"github.com/rubixchain/rubixgoplatform/core/model"
	"github.com/rubixchain/rubixgoplatform/server"
)

type SmartContractRequest struct {
	BinaryCode string
	RawCode    string
	SchemaCode string
	DID        string
	SCPath     string
}

type FetchSmartContractRequest struct {
	SmartContractToken     string
	SmartContractTokenPath string
}

func (c *Client) GenerateSmartContractToken(smartContractRequest *SmartContractRequest) (*model.BasicResponse, error) {

	fields := make(map[string]string)
	files := make(map[string]string)

	if smartContractRequest.BinaryCode != "" {
		files["binaryCodePath"] = smartContractRequest.BinaryCode
	}
	if smartContractRequest.RawCode != "" {
		files["rawCodePath"] = smartContractRequest.RawCode
	}
	if smartContractRequest.SchemaCode != "" {
		files["schemaFilePath"] = smartContractRequest.SchemaCode
	}
	if smartContractRequest.DID != "" {
		fields["did"] = smartContractRequest.DID
	}

	for key, value := range fields {
		fmt.Printf("Field: %s, Value: %s\n", key, value)
	}

	for key, value := range files {
		fmt.Printf("File: %s, Value: %s\n", key, value)
	}

	var basicResponse model.BasicResponse
	err := c.sendMutiFormRequest("POST", server.APIGenerateSmartContract, nil, fields, files, &basicResponse)
	if err != nil {
		return nil, err
	}

	return &basicResponse, nil

}

func (c *Client) FetchSmartContract(fetchSmartContractRequest *FetchSmartContractRequest) (*model.BasicResponse, error) {
	fields := make(map[string]string)
	if fetchSmartContractRequest.SmartContractToken != "" {
		fields["smartContractToken"] = fetchSmartContractRequest.SmartContractToken
	}

	var basicResponse model.BasicResponse
	err := c.sendMutiFormRequest("POST", server.APIFetchSmartContract, nil, fields, nil, &basicResponse)
	if err != nil {
		return nil, err
	}
	return &basicResponse, nil

}
func (c *Client) PublishNewEvent(contract string, did string, block string) (*model.BasicResponse, error) {
	var response model.BasicResponse
	newContract := model.NewContractEvent{
		Contract:          contract,
		Did:               did,
		ContractBlockHash: block,
	}
	err := c.sendJSONRequest("POST", server.APIPublishContract, nil, &newContract, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (c *Client) SubscribeContract(contract string) (*model.BasicResponse, error) {
	var response model.BasicResponse
	newSubscription := model.NewSubcription{
		Contract: contract,
	}
	err := c.sendJSONRequest("POST", server.APISubscribecontract, nil, &newSubscription, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
