package authentication

import (
	"SpecForge_api_backend/utilities/globalUtility"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAuthenticationData struct {
	secretKey string
	timestamp int
}

const CONST_SERVER_SECRET_KEY = "|xh&?LLr]|OA`_B?zQI-er;#?W_'?zfK"

func AuthenticateUser(ctx *gin.Context) {
	// Extract Bearer token from Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing my friend"})
		ctx.Abort()
		return
	}

	// Check for Bearer token format and split to get 'ak'
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		ctx.Abort()
		return
	}
	ak := tokenParts[1]

	decodedAK := globalUtility.DecodeBase64(ak)
	if decodedAK == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization key"})
		ctx.Abort()
		return
	}

	var UserAKData UserAuthenticationData
	AKObject := globalUtility.JsonDecode(decodedAK)
	if payload, ok := AKObject.(map[string]interface{}); ok {
		if secretKey, ok := payload["secret_key"]; ok {
			UserAKData.secretKey = globalUtility.ConvertValueToString(secretKey)
		}
		if timestamp, ok := payload["time"]; ok {
			UserAKData.timestamp = globalUtility.ConvertValueToInt(timestamp)
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization key"})
		ctx.Abort()
		return
	}

	if UserAKData.secretKey != CONST_SERVER_SECRET_KEY {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization key"})
		ctx.Abort()
		return
	}
	
	current_date := time.Now().Format("2006-01-02")
	time := globalUtility.ConvertUnixToTime(int64(globalUtility.ConvertValueToInt(UserAKData.timestamp))).Format("2006-01-02")
	if time != current_date {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization key"})
		ctx.Abort()
		return
	}
}