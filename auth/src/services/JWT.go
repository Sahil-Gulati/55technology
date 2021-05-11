package services;

import(
	"fmt"
	utils "../utils"
	stubs "../stubs"
)

type JWT struct {
	
}
func (jwt JWT) GetToken() string {
	return jwt.GetSignatureHash(
		jwt.GetHeadersHash(),
		jwt.GetBodyHash(),
	)
}
func (jwt JWT) GetHeadersHash() string {
	headers := stubs.Headers{Algorithm: "SHA256", TokenType: "JWT"}
	jsonBody := utils.JsonEncode(headers)
	return utils.Base64encode(jsonBody)
}
func (jwt JWT) GetBodyHash() string {
	/**
	 * Returning headers hash instead of body hash
	 */
	headers := stubs.Headers{Algorithm: "SHA256", TokenType: "JWT"}
	jsonBody := utils.JsonEncode(headers)
	return utils.Base64encode(jsonBody)
}
func (jwt JWT) GetSignatureHash(headersHash string, bodyHash string) string {
	jsonBody := utils.JsonEncode(headersHash + bodyHash)
	return fmt.Sprintf(
		"%s.%s.%s",
		headersHash,
		bodyHash,
		utils.Base64encode(jsonBody),
	)
}