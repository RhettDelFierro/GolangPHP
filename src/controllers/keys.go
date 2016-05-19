package controllers

import(
	"io/ioutil"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"io"
	"fmt"
)

const (
	privKey = "keys/app.rsa"
	pubKey = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys(){
	var err error

	signKey, err = ioutil.ReadFile(privKey)
	if err != nil {
		fmt.Println("privKey not reading")
		panic(err)
	}

	verifyKey, err = ioutil.ReadFile(pubKey)
	if err != nil {
		fmt.Println("pubKey not reading")
		panic(err)
	}
}

//generating the token to
func GenerateToken(name, role string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RSA256"))


	//Setting the claims. This info will be used through the app.
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	t.Claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//Validate the tokens on each route that needs them
func AuthorizeToken(w http.ResponseWriter, req *http.ResponseWriter, next http.HandlerFunc){
	//checking the token from the request.
	//make sure to put Authorization: Bearer <token info> in header on front end.
	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error){
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		//write the error during validation
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			//time from Claims expired.
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(401)
				w.Write("jwt has expired")
				return

			default:
				w.WriteHeader(500)
				w.Write("error parsing access Token")
				return
			}

		default:
			w.WriteHeader(500)
			return
		}
	}
	if token.Valid {
		//call back on the HandlerFunc because this is a wrapping function on middleware.
		//will use with negroni:
		next(w, req)
	} else {
		w.WriteHeader(401)
		w.Write("Invalid Access Token")
	}
}