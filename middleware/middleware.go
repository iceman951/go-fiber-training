package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"ipanda.baac.tech/golib/cipherPayload"
)

func InitializeRequestIDFieldName() string {
	return viper.GetString("Middleware.RequestID.FieldName")
}

func InitializeCORS() (bool, cors.Config) {
	isEnable := viper.GetBool("Middleware.CORS.Enable")
	cfg := cors.ConfigDefault
	if !isEnable {
		return isEnable, cfg
	}

	cfg.AllowOrigins = viper.GetString("Middleware.CORS.AllowOrigin")
	cfg.AllowMethods = viper.GetString("Middleware.CORS.AllowMethod")

	return isEnable, cfg
}

func InitializeCipherPayload() (bool, cipherPayload.Config) {
	isEnable := viper.GetBool("Middleware.CipherPayload.Enable")
	cfg := cipherPayload.ConfigDefault
	if !isEnable {
		return isEnable, cfg
	}

	cfg.KeyPairs.AESKeyForEncrypt = []byte(viper.GetString("Middleware.CipherPayload.AESKeyForEncrypt"))
	cfg.KeyPairs.AESIVForEncrypt = []byte(viper.GetString("Middleware.CipherPayload.AESIVForEncrypt"))
	cfg.KeyPairs.AESKeyForDecrypt = []byte(viper.GetString("Middleware.CipherPayload.AESKeyForDecrypt"))
	cfg.KeyPairs.AESIVForDecrypt = []byte(viper.GetString("Middleware.CipherPayload.AESIVForDecrypt"))
	cfg.AllowMethod = viper.GetStringSlice("Middleware.CipherPayload.AllowMethod")
	cfg.DebugMode = viper.GetBool("Middleware.CipherPayload.Verbose")

	return isEnable, cfg
}
