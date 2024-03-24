package main

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
)

//
//func handleFunctionality() {
//
//	jsonChainData, err := os.ReadFile("bam-api/chain.json")
//	if err != nil {
//		panic(err)
//	}
//
//	var readBlockchain blockchain.Blockchain
//	err = json.Unmarshal(jsonChainData, &readBlockchain)
//	if err != nil {
//		panic(err)
//	}
//	chainlen := len(readBlockchain.Chain)
//	latestBlockData := readBlockchain.Chain[chainlen-1].Data
//	//fmt.Println(latestBlockData)
//
//	var messageData dal.MessagingStruct
//	var UUIDsData dal.UUIDsConfig
//
//	err = json.Unmarshal([]byte(latestBlockData), &messageData)
//
//	jsonconfigData, err := os.ReadFile("config.json")
//
//	err = json.Unmarshal(jsonconfigData, &UUIDsData)
//	if err != nil {
//		panic(err)
//	}
//	messageChange, _ := strconv.Atoi(messageData.Change)
//	if messageData.Name == "lighting" {
//		for i := 0; i < len(UUIDsData.LightingUUIDs); i++ {
//			if UUIDsData.LightingUUIDs[i] == messageData.UUID {
//				if messageData.Function == "status" {
//					lighting.UpdateStatus(messageData.StatusChange)
//				}
//				if messageData.Function == "brightness" {
//					lighting.SetBrightness(messageChange)
//				}
//			}
//		}
//	}
//	if messageData.Name == "hvac" {
//		for i := 0; i < len(UUIDsData.HvacUUIDs); i++ {
//			if UUIDsData.HvacUUIDs[i] == messageData.UUID {
//				if messageData.Function == "status" {
//					hvac.UpdateStatus(messageData.StatusChange)
//				}
//				if messageData.Function == "fan" {
//					hvac.UpdateFanSpeed(messageChange)
//				}
//				if messageData.Function == "temerature" {
//					hvac.UpdateTemperature(messageChange)
//				}
//				if messageData.Function == "mode" {
//					hvac.UpdateMode(messageData.Change)
//				}
//			}
//		}
//	}
//	if messageData.Name == "security" {
//		for i := 0; i < len(UUIDsData.SecurityUUIDs); i++ {
//			if UUIDsData.SecurityUUIDs[i] == messageData.UUID {
//				if messageData.Function == "status" {
//					security.UpdateAlarmStatus(messageData.StatusChange)
//				}
//			}
//		}
//	}
//	if messageData.Name == "appliances" {
//		for i := 0; i < len(UUIDsData.AppliancesUUIDs); i++ {
//			if UUIDsData.AppliancesUUIDs[i] == messageData.UUID {
//				if messageData.Function == "status" {
//					appliances.UpdateStatus(messageData.StatusChange)
//				}
//				if messageData.Function == "temperature" {
//					appliances.UpdateTemperature(messageChange)
//				}
//				if messageData.Function == "timerstoptime" {
//					appliances.UpdateTimeStopTime(messageChange)
//				}
//				if messageData.Function == "power" {
//					appliances.UpdatePower(messageChange)
//				}
//				if messageData.Function == "energysavingmode" {
//					appliances.UpdateEnergySavingMode(messageData.StatusChange)
//				}
//				if messageData.Function == "washtime" {
//					appliances.UpdateWashTime(messageChange)
//				}
//			}
//		}
//	}
//	if messageData.Name == "energy" {
//		for i := 0; i < len(UUIDsData.EnergyUUIDs); i++ {
//			if UUIDsData.EnergyUUIDs[i] == messageData.UUID {
//				if messageData.Function == "status" {
//					//energy.UpdateAlarmStatus(messageData.StatusChange)
//				}
//			}
//		}
//	}
//
//}

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

	messaging.BlockReceiver()
	//go handleFunctionality()
}
