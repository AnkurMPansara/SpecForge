package globalUtility

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func GetCurrentTimestamp() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05.000")
	return currentTime
}

func EncodeBase64(inputValue string) string {
	return base64.StdEncoding.EncodeToString([]byte(inputValue))
}

func DecodeBase64(inputValue string) string {
	decodedData, decodingErr := base64.StdEncoding.DecodeString(inputValue)
	if decodingErr != nil {
		return ""
	}
	return string(decodedData)
}

func JsonEncode(inputData interface{}) string {
	encoded, err := json.Marshal(inputData)
	if err != nil {
		return ""
	}
	return string(encoded)
}

func JsonDecode(inputData string) interface{} {
	var payload interface{}
	err := json.Unmarshal([]byte(inputData), &payload)
	if err != nil {
		return nil
	}
	return payload
}

func ConvertValueToString(inputData interface{}) string {
	return fmt.Sprintf("%v", inputData)
}

func ConvertValueToInt(inputData interface{}) int {
	switch val := inputData.(type) { //in go switch dont require break statement so no need to apply
	case int64:
		return int(val)
	case int32: //rune and int32 are same
		return int(val)
	case int16:
		return int(val)
	case int8:
		return int(val)
	case int:
		return val
	case uint:
		return int(val)
	case uint64:
		return int(val)
	case uint32:
		return int(val)
	case uint16:
		return int(val)
	case uint8: //byte and uint8 are same
		return int(val)
	case []byte: //byte and uint8 are same
		if strToFloatVal, err := strconv.Atoi(string(val)); err == nil {
			return strToFloatVal
		} else {
			return 0
		}
	case string:
		if strToFloatVal, err := strconv.Atoi(val); err == nil {
			return strToFloatVal
		} else {
			return 0
		}
	case float64:
		return int(val)
	case float32:
		return int(val)
	case bool:
		if val{
			return 1
		}else{
			return 0
		}
	default:
		return 0
	}
}

func ConvertUnixToTime(inputData int64) time.Time {
	return time.Unix(inputData, 0).UTC()
}