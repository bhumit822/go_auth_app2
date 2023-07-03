package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pquerna/otp/totp"
)

func Main() {

	// var a, err = totp.Generate(totp.GenerateOpts{Issuer: "app_auth.com", AccountName: "test@mail.com", SecretSize: 15})

	// fmt.Println(a.Secret())
	// fmt.Println(a.URL())
	// fmt.Println(err)

	a, _ := totp.ValidateCustom("869876", "K5F2K6UEHYK4JZ5MBERXGOSI", time.Now(), totp.ValidateOpts{Digits: 6})
	fmt.Println("===> verified", a)

	router := mux.NewRouter()
	// router.HandleFunc("/", HandleWell).Methods("GET")
	router.HandleFunc("/", Handler).Methods("POST")
	http.ListenAndServe(":6060", router)

}

type OtpReqIndex struct {
	Otp string
}

// func HandleWell(w http.ResponseWriter, r *http.Request) {

//		fmt.Fprintf(w, "WellCome to app Auth")
//	}
func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var otp OtpReqIndex
	err := decoder.Decode(&otp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("otp :==> ", otp.Otp)

	a, _ := totp.ValidateCustom(otp.Otp, "K5F2K6UEHYK4JZ5MBERXGOSI", time.Now(), totp.ValidateOpts{Digits: 6})

	if a {

		fmt.Fprintf(w, "===> verified")

	} else {

		fmt.Fprintf(w, "===> Not verified")
	}
}
