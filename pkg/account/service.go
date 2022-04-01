package account

import (
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"

	"github.com/pkedy/megabank/pkg/api"
)

type (
	Service struct {
		log        logr.Logger
		daprClient dapr.Client
	}
)

func New(log logr.Logger, daprClient dapr.Client) *Service {
	return &Service{
		log:        log,
		daprClient: daprClient,
	}
}

func (s *Service) RegisterService(app *mux.Router) {
	app.HandleFunc("/accounts/v1/{id}", s.GetBalance).Methods("GET")
	app.HandleFunc("/accounts/v1/{id}/withdraw", s.Withdraw).Methods("POST")
	app.HandleFunc("/accounts/v1/{id}/deposit", s.Deposit).Methods("POST")
}

// GetBalance godoc
// @Summary      Get the account balance
// @Description  Get the account balance by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "Account ID"
// @Success      200  {object} integer
// @Router       /accounts/v1/{id} [get]
func (s *Service) GetBalance(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	s.log.Info("GetBalance", "id", id)
	actor := api.NewAccountActor(id)
	s.daprClient.ImplActorClientStub(actor)
	result, err := actor.Balance(req.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// Withdraw godoc
// @Summary      Withdraw from an account
// @Description  Withdraw from an account by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Account ID"
// @Param        body body      api.AmountRequest true  "Amount request"
// @Success      204
// @Router       /accounts/v1/{id}/withdraw [post]
func (s *Service) Withdraw(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	s.log.Info("Withdraw", "id", id)

	var request api.AmountRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		handleError(w, err)
		return
	}

	actor := api.NewAccountActor(id)
	s.daprClient.ImplActorClientStub(actor)
	err := actor.Withdraw(req.Context(), request.Amount)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Deposit godoc
// @Summary      Deposit into an account
// @Description  Deposit into an account by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param        body body      api.AmountRequest true  "Amount request"
// @Success      204
// @Router       /accounts/v1/{id}/deposit [post]
func (s *Service) Deposit(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	s.log.Info("Deposit", "id", id)

	var request api.AmountRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		handleError(w, err)
		return
	}

	actor := api.NewAccountActor(id)
	s.daprClient.ImplActorClientStub(actor)
	err := actor.Deposit(req.Context(), request.Amount)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleError(w http.ResponseWriter, err error) {
	// TODO: This needs a proper API error implementation
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
