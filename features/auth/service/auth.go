package service

import (
	repo "consolidated/features/auth/repository"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func BasicAuth(c *gin.Context) (userAuth repo.Auth, errs []ErrosMsg) {

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		// opt.SkipOnEmpty = false
	})

	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		// err := errors.New("Unauthorized")
		// return userAuth, err
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		// return userAuth, errors.New("Unauthorized")
	}
	userAuth.Username = pair[0]
	userAuth.Password = pair[1]
	// userAuth.Age = "adr"

	// m := map[string]map[string]string{}

	// m["key1"] = map[string]string{}
	// m["key1"]["1"] = "dgggg"
	// m["key1"]["2"] = "wwwwwww"
	// m["key1"]["3"] = "dddddd"

	// m["key2"] = map[string]string{}
	// m["key2"]["3"] = "kkkkk"

	// // m["key2"]["3"] = "dddddd"

	// fmt.Println(m["key2"])
	// b, _ := json.Marshal(m)
	// fmt.Println(string(b))

	v := validate.Struct(userAuth)

	if !v.Validate() { // validate ok
		errs := []ErrosMsg{}

		for fieldName, validations := range v.Errors {
			er := ErrosMsg{
				Field:       fieldName,
				Description: make([]map[string]string, 0),
			}

			fmt.Println(fieldName)
			for validateType, errMessage := range validations {
				des := map[string]string{}
				des[validateType] = errMessage
				er.Description = append(er.Description, des)
			}
			errs = append(errs, er)
		}

		// b, _ := json.Marshal(errs)
		// fmt.Println(string(b))

		return userAuth, errs
	}

	return userAuth, nil
}

type ErrosMsg struct {
	Field       string
	Description []map[string]string
}

// [{
//     "Field": "Username",
//     "Description": [
//         {"required": "Username is required and not empty"},
// 	{"required": "Username is required and not empty"},
//     ]
// },
// {
//     "Field": "Password",
//     "Description": {
//         "required": "Password is required and not empty"
//     }
// }
// ]
