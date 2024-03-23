package main

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"CMPSC488SP24SecTuesday/blockchain"
	"CMPSC488SP24SecTuesday/dal"
	"CMPSC488SP24SecTuesday/lighting"
	"encoding/json"
	"os"
)

func handleFunctionality() {

	jsonChainData, err := os.ReadFile("bam-api/chain.json")
	if err != nil {
		panic(err)
	}

	var readBlockchain blockchain.Blockchain
	err = json.Unmarshal(jsonChainData, &readBlockchain)
	if err != nil {
		panic(err)
	}
	chainlen := len(readBlockchain.Chain)
	latestBlockData := readBlockchain.Chain[chainlen-1].Data
	//fmt.Println(latestBlockData)

	var messageData dal.MessagingStruct
	var UUIDsData dal.UUIDsConfig

	err = json.Unmarshal([]byte(latestBlockData), &messageData)

	jsonconfigData, err := os.ReadFile("config.json")

	err = json.Unmarshal(jsonconfigData, &UUIDsData)
	if err != nil {
		panic(err)
	}
	if messageData.Name == "lighting" {
		for i := 0; i < len(UUIDsData.LightingUUIDs); i++ {
			if UUIDsData.LightingUUIDs[i] == messageData.UUID {
				if messageData.Function == "status" {
					lighting.UpdateStatus(messageData.StatusChange)
				}
			}
		}
	}

}

//func findItemType(blockData []byte) []byte {
//	var lightingChange dal.UpdateLightingRequest
//	var HVACChange dal.UpdateHVACRequest
//	var securityChange dal.UpdateSecurityRequest
//	var applianceChange dal.UpdateAppliancesRequest
//	var energyChange dal.UpdateEnergyRequest
//
//	err := json.Unmarshal([]byte(blockData), &lightingChange)
//	if err != nil {
//		fmt.Printf("not a light update")
//	}
//	if err == nil {
//		return lightingChange
//	}
//	err = json.Unmarshal([]byte(blockData), &lightingChange)
//	if err != nil {
//		fmt.Printf("not a light update")
//	}
//}

func main() {

	go messaging.BlockReceiver()
	go handleFunctionality()
	//handleFunctionality()
}
