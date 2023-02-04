package service

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// later -get it from env varaible
var secretKey = []byte("SecretKey")

// func generateJwt(userId string) (string, error) {
// 	//token := jwt.New(jwt.SigningMethodEdDSA)
// 	// claims := token.Claims(jwt.MapClaims)
// 	// claims = token.Claims(jwt.MapClaims)
// 	// claims["exp"] = time.Now().Add(time.Hour * 72)
// 	// claims["authorized"] = true
// 	// claims["user"] = userId
// }
