package middleware

import (
	"fmt"
	"net/http"
	"strings"

	// "strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
type CustomClaims struct {
    UserID int `json:"user_id"`
    Roles  []int `json:"roles"`
    jwt.StandardClaims
}

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		stringToken := strings.Replace(tokenString, "Bearer ", "", 1)
		valid, claims := VerifyToken(stringToken, "my-secret-key")
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"errorVerify": "error Verify"})
			c.Abort()
			return
		}

		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		// 	c.Abort()
		// 	return
		// }

		// Check if the token contains a "roles" claim
		// role, ok := claims["RoleId"]
		// fmt.Println(role)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
		// 	c.Abort()
		// 	return
		// }

		// roleNum := role
		// // fmt.Println("ini role num",roleNum)
		// // if roleNum != 1 && roleNum != 2 && roleNum != 3 {
		// // 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		// // 	c.Abort()
		// // 	return
		// // }
		// fmt.Println(roleNum)

		// c.Set("role", roleNum)
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["Username"] = claims.UserId
		c.Keys["Roles"] = claims.RoleId
		c.Next()
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.Keys) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in array key"})
			c.Abort()
			return
		}

		rolesVal := c.Keys["Roles"]
		fmt.Println("roles", rolesVal)
		if rolesVal == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you cant access this API"})
			c.Abort()
			return
		}

		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "you cant access this API"})
			c.Abort()
			return
			}
		}

		c.Next()
	}
}
// func RBACMiddleware(allowedRoles ...int) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
// 			return
// 		}

// 		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
// 		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}

// 		claims, ok := token.Claims.(*CustomClaims)
// 		if !ok || !token.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			return
// 		}

// 		if len(allowedRoles) > 0 {
// 			roles, ok := claims.Roles.([]interface{})
// 			if !ok {
// 				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid roles claim in token"})
// 				return
// 			}

// 			allowed := false
// 			for _, r := range roles {
// 				role, ok := r.(float64)
// 				if !ok {
// 					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid roles claim in token"})
// 					return
// 				}

// 				for _, allowedRole := range allowedRoles {
// 					if int(role) == allowedRole {
// 						allowed = true
// 						break
// 					}
// 				}
// 			}

// 			if !allowed {
// 				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
// 				return
// 			}
// 		}

// 		c.Set("user", claims.UserID)
// 		c.Next()
// 	}
// }
