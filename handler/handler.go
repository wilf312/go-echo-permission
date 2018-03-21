package handler

import (
  "net/http"
	_"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
  return func(c echo.Context) error { 
    return c.String(http.StatusOK, "Hello World")
  }
}

// func PathParameter() echo.HandlerFunc {
//   return func(c echo.Context) error { 
//     pathParam := c.Param("pathParameter")
//     return c.String(http.StatusOK, "Hello World " + pathParam)
//   }
// }