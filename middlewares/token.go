package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// TokenAuthMiddleware JWT 토큰 유효성 검사를 한다
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Release Mode 일때만 토큰을 체크한다
		mode := os.Getenv("GIN_MODE")
		if mode == "release" {
			// Authorization header Check
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
				c.Abort()
				return
			}
			// Token Validation
			_, err := tokenValid(c.Request)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			//// Roles 목록을 String 슬라이스로 변환한다
			//roles := claims["roles"].([]interface{})
			//strRoles := make([]string, len(roles))
			//for i, v := range roles {
			//	strRoles[i] = fmt.Sprint(v)
			//}
			//c.Set("roles", strRoles)

			c.Next()
		}

		c.Next()
	}
}

//TokenValid Middleware에서 토큰의 유효성 검사를 한다
func tokenValid(r *http.Request) (jwt.MapClaims, error) {
	token, err := verifyJWTToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}

	return claims, nil
}

// VerifyJWTToken JWT 토큰의 유효성 검사를 한다
func verifyJWTToken(r *http.Request) (*jwt.Token, error){
	tokenString := extractToken(r)
	jwtAccSecret := os.Getenv("JWT_ACCESS_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtAccSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExtractToken HTTP Request 헤더에서 토큰을 가져온다
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")

	if len(strArr) == 2 {
		if strArr[0] == "Bearer" {
			return strArr[1]
		}
	}

	return ""
}

