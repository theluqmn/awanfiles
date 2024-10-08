package user

// import (
// 	"main/utils"
// 	"net/http"
// 	"github.com/labstack/echo"
// )

// func Account(c echo.Context) error {
//     db := DatabaseGet()

// 	// Check if username and password are provided
// 	if c.QueryParam("username") == "" || c.QueryParam("password") == "" {
//         return c.String(http.StatusBadRequest, "Missing username or password")
//     }

//     stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
//     if err != nil {
//         return c.String(http.StatusInternalServerError, "Failed to prepare statement")
//     }
//     defer stmt.Close()

//     _, err = stmt.Exec(c.QueryParam("username"), c.QueryParam("password"))
//     if err != nil {
//         return c.String(http.StatusInternalServerError, "Failed to create account")
//     }

// 	utils.Log("Created new account, with username " + c.QueryParam("username"))
//     return c.String(http.StatusOK, "Created account")
// }