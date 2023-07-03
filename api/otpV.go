package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pquerna/otp/totp"
)

func OtpV() {

	a, _ := totp.ValidateCustom("869876", "K5F2K6UEHYK4JZ5MBERXGOSI", time.Now(), totp.ValidateOpts{Digits: 6})
	fmt.Println("===> verified", a)

	router := mux.NewRouter()
	router.HandleFunc("/", OtpVHandler).Methods("POST")
	http.ListenAndServe(":6060", router)

}

type OtpReq struct {
	Otp string
}

func OtpVHandler(w http.ResponseWriter, r *http.Request) {
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
