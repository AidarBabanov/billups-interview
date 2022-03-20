package game

import (
	"encoding/json"
	"github.com/AidarBabanov/billups-interview/rest"
	"github.com/go-chi/chi"
	"net/http"
)

//go:generate mockery --name "Game" --inpackage --structname "mockGame" --filename "game.mock.go"

type Game interface {
	GetChoices() []string
	GetChoice() (int, string)
	Play(choiceID int) (plChoiceID, comChoiceID int, res, description string)
}

type Resource struct {
	*chi.Mux
	service Game
}

func NewResource(gameService Game) *Resource {
	res := &Resource{Mux: chi.NewMux(), service: gameService}
	res.Get("/choices", rest.APIHandlerFunc(res.GetChoices))
	res.Get("/choice", rest.APIHandlerFunc(res.GetChoice))
	res.Post("/play", rest.APIHandlerFunc(res.PostPlay))
	return res
}

// swagger:model
type getChoiceResponseBody struct {
	// example: 1
	ID int `json:"id"`
	// example: rock
	Name string `json:"name"`
}

// swagger:response
type getChoicesResponse struct {
	Body []getChoiceResponseBody
}

// GetChoices returns all possible choices for the game.
//
//	swagger:route GET /choices Game GetChoices
//
//	Returns all possible choices for the game.
//
//	Returns all possible choices for the game.
//
//	responses:
//		200: getChoicesResponse
//		default: errorResponse
func (res *Resource) GetChoices(w http.ResponseWriter, _ *http.Request) error {
	choicesList := res.service.GetChoices()
	var resp []getChoiceResponseBody
	for id, name := range choicesList {
		resp = append(resp, getChoiceResponseBody{ID: id + 1, Name: name})
	}
	return rest.WriteJson(w, http.StatusOK, resp)
}

// swagger:response
type getChoiceResponse struct {
	Body getChoiceResponseBody
}

// GetChoice returns random choice ID.
//
//	swagger:route GET /choice Game GetChoice
//
//	Returns random choice ID.
//
//  Returns random choice ID (1-5).
//
//	responses:
//		200: getChoiceResponse
//		default: errorResponse
func (res *Resource) GetChoice(w http.ResponseWriter, _ *http.Request) error {
	id, name := res.service.GetChoice()
	resp := getChoiceResponseBody{ID: id, Name: name}
	return rest.WriteJson(w, http.StatusOK, resp)
}

// swagger:parameters PostPlay
type postPlayParams struct {
	// in:body
	Body struct {
		// Player choice ID
		// example: 1
		// enum: 1, 2, 3, 4, 5
		Player int `json:"player"`
	}
}

// swagger:model
type postPlayResponseBody struct {
	// example: tie
	Results string `json:"results"`
	// example: 1
	Player int `json:"player"`
	// example: 1
	Computer int `json:"computer"`
	// example: Two solid Dwaynes.
	Description string `json:"description"`
}

// swagger:response
type postPlayResponse struct {
	Body postPlayResponseBody
}

// PostPlay performs game for a player against computer.
//
//	swagger:route POST /play Game PostPlay
//
//	Performs game for a player against computer.
//
//  Performs game for a player against computer.
//
//	responses:
//		200: postPlayResponse
//		default: errorResponse
func (res *Resource) PostPlay(w http.ResponseWriter, r *http.Request) error {
	req := postPlayParams{}
	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return rest.BadRequestErrorf("wrong body format")
	}
	if req.Body.Player == 0 || req.Body.Player > 5 {
		return rest.BadRequestErrorf("player choice ID must be in range [1:5]")
	}
	plChoice, comChoice, result, description := res.service.Play(req.Body.Player)
	resp := postPlayResponseBody{Player: plChoice, Computer: comChoice, Results: result, Description: description}
	return rest.WriteJson(w, http.StatusOK, resp)
}
