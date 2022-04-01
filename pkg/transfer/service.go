package transfer

import (
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"

	"github.com/pkedy/megabank/pkg/api"
)

type Service struct {
	log        logr.Logger
	daprClient dapr.Client
}

func New(log logr.Logger, daprClient dapr.Client) *Service {
	return &Service{
		log:        log,
		daprClient: daprClient,
	}
}

func (s *Service) RegisterService(app *mux.Router) {
	app.HandleFunc("/transfers/v1/{id}", s.TransferFunds).Methods("POST")
}

// TransferFunds godoc
// @Summary      Transfer a specific amount from one account to another
// @Description  Transfer a specific amount of money from one account to another
// @Tags         transfers
// @Accept       json
// @Produce      json
// @Param        id   path     string  true  "Reference ID"
// @Param        body body     api.TransferRequest true  "Transfer request"
// @Success      200  {object} api.TransactionResult
// @Router       /transfers/v1/{id} [post]
func (s *Service) TransferFunds(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	s.log.Info("Withdraw", "id", id)

	var request api.TransferRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		handleError(w, err)
		return
	}

	actor := api.NewTransferActor(id)
	s.daprClient.ImplActorClientStub(actor)
	resp, err := actor.TransferFunds(req.Context(), &request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(resp)
}

func handleError(w http.ResponseWriter, err error) {
	// TODO: This needs a proper API error implementation
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
