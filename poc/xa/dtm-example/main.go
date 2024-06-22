package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dtm-labs/dtmcli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// Service address of the transaction participant
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8082

var qsBusi = fmt.Sprintf("http://localhost:%d%s", qsBusiPort, qsBusiAPI)

func main() {
	QsStartSvr()
	_ = QsFireRequest()
	select {}
}

// QsStartSvr quick start: start server
func QsStartSvr() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	qsAddRoute(r)
	log.Printf("quick start examples listening at %d", qsBusiPort)
	go func() {
		_ = http.ListenAndServe(fmt.Sprintf(":%d", qsBusiPort), r)
	}()
	time.Sleep(100 * time.Millisecond)
}

func qsAddRoute(r chi.Router) {
	r.Post(qsBusiAPI+"/TransIn", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("TransIn")
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "")
		// render.Status(r, http.StatusConflict) // Status 409 for Failure. Won't be retried
	})
	r.Post(qsBusiAPI+"/TransInCompensate", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("TransInCompensate")
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "")
	})
	r.Post(qsBusiAPI+"/TransOut", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("TransOut")
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "")
	})
	r.Post(qsBusiAPI+"/TransOutCompensate", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("TransOutCompensate")
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "")
	})
}

const dtmServer = "http://localhost:36789/api/dtmsvr"

// QsFireRequest quick start: fire request
func QsFireRequest() string {
	req := map[string]interface{}{"amount": 30} // Payload for the microservice
	// DtmServer is the address of the DTM service
	saga := dtmcli.NewSaga(dtmServer, dtmcli.MustGenGid(dtmServer)).
		// Add a TransOut sub-transaction, with the forward operation as url: qsBusi+"/TransOut", and the reverse operation as url: qsBusi+"/TransOutCompensate"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCompensate", req).
		// Add a TransIn sub-transaction, with the forward operation as url: qsBusi+"/TransIn", and the reverse operation as url: qsBusi+"/TransInCompensate"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCompensate", req)
	// Submit the saga transaction, DTM will complete all sub-transactions/rollback all sub-transactions
	err := saga.Submit()

	if err != nil {
		panic(err)
	}
	log.Printf("transaction: %s submitted", saga.Gid)
	return saga.Gid
}
