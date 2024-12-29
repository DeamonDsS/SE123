package users


import (

   "errors"

   "net/http"

   "time"


   "github.com/gin-gonic/gin"

   "golang.org/x/crypto/bcrypt"

   "gorm.io/gorm"


   "github.com/SE67/config"

    "github.com/SE67/entity"

    "github.com/SE67/services"

)


type (

   Authen struct {

       Email    string `json:"email"`

       Password string `json:"password"`

   }


   signUp struct {

       FirstName string    `json:"first_name"`

       LastName  string    `json:"last_name"`

       Email     string    `json:"email"`

       Age       uint8     `json:"age"`

       Password  string    `json:"password"`

       BirthDay  time.Time `json:"birthday"`

       GenderID  uint      `json:"gender_id"`

   }

)


func SignUp(c *gin.Context) {

   var payload signUp


   // Bind JSON payload to the struct

   if err := c.ShouldBindJSON(&payload); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

       return

   }


   db := config.DB()

   var userCheck entity.Users


   // Check if the user with the provided email already exists

   result := db.Where("email = ?", payload.Email).First(&userCheck)

   if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {

       // If there's a database error other than "record not found"

       c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})

       return

   }


   if userCheck.ID != 0 {

       // If the user with the provided email already exists

       c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})

       return

   }


   // Hash the user's password

   hashedPassword, _ := config.HashPassword(payload.Password)


   // Create a new user

   user := entity.Users{

       FirstName: payload.FirstName,

       LastName:  payload.LastName,

       Email:     payload.Email,

       Age:       payload.Age,

       Password:  hashedPassword,

       BirthDay:  payload.BirthDay,

       GenderID:  payload.GenderID,
       
       IsAdmin:   false,
       

   }


   // Save the user to the database

   if err := db.Create(&user).Error; err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

       return

   }


   c.JSON(http.StatusCreated, gin.H{"message": "Sign-up successful"})

}


func SignIn(c *gin.Context) {
    var payload Authen
    var user entity.Users

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB().Where("email = ?", payload.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password is incorrect"})
        return
    }

    jwtWrapper := services.JwtWrapper{
        SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
        Issuer:          "AuthService",
        ExpirationHours: 24,
    }

    signedToken, err := jwtWrapper.GenerateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
        return
    }

    // Include IsAdmin in the response
    c.JSON(http.StatusOK, gin.H{
        "token_type": "Bearer",
        "token":      signedToken,
        "id":         user.ID,
        "is_admin":   user.IsAdmin,
    })
}
