```go
 	s, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Data: "+err.Error())
	}
	var body map[string]interface{}
	if err := json.Unmarshal(s, &body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Data: "+err.Error())
	}
	log.Printf("%s", body["email"])
	return c.JSON(http.StatusOK, body)
```
