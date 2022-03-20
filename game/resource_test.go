package game

import (
	"bytes"
	"encoding/json"
	"github.com/AidarBabanov/billups-interview/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChoicesHandler(t *testing.T) {
	srv := &mockGame{}
	resource := NewResource(srv)
	t.Run("OK", func(t *testing.T) {
		srv.On("GetChoices").Return(choices)
		r := httptest.NewRequest(http.MethodGet, "/choices", nil)
		w := httptest.NewRecorder()
		err := resource.GetChoices(w, r)
		require.NoError(t, err)
		expected, err := json.Marshal([]getChoiceResponseBody{
			{1, "rock"},
			{2, "paper"},
			{3, "scissors"},
			{4, "lizard"},
			{5, "spock"},
		})
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(expected), w.Body.String())
	})
}

func TestGetChoiceHandler(t *testing.T) {
	srv := &mockGame{}
	resource := NewResource(srv)
	t.Run("OK", func(t *testing.T) {
		srv.On("GetChoice").Return(1, "rock")
		r := httptest.NewRequest(http.MethodGet, "/choice", nil)
		w := httptest.NewRecorder()
		err := resource.GetChoice(w, r)
		require.NoError(t, err)
		expected, err := json.Marshal(getChoiceResponseBody{1, "rock"})
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(expected), w.Body.String())
	})
}

func TestPostPlayHandler(t *testing.T) {
	srv := &mockGame{}
	resource := NewResource(srv)
	t.Run("OK", func(t *testing.T) {
		srv.On("Play", 1).Return(1, 2, "lose", "Paper covers rock.")
		body, err := json.Marshal(struct {
			Player int `json:"player"`
		}{Player: 1})
		require.NoError(t, err)
		bodyReader := bytes.NewReader(body)
		r := httptest.NewRequest(http.MethodGet, "/play", bodyReader)
		w := httptest.NewRecorder()

		r.Header[rest.ContentType] = []string{rest.ContentTypeJson}
		err = resource.PostPlay(w, r)
		require.NoError(t, err)
		expected, err := json.Marshal(postPlayResponseBody{Results: "lose", Player: 1, Computer: 2, Description: "Paper covers rock."})
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(expected), w.Body.String())
	})

	t.Run("Error; wrong body format", func(t *testing.T) {
		srv.On("Play", 1).Return(1, 2, "lose", "Paper covers rock.")
		r := httptest.NewRequest(http.MethodGet, "/play", nil)
		w := httptest.NewRecorder()

		r.Header[rest.ContentType] = []string{rest.ContentTypeJson}
		err := resource.PostPlay(w, r)
		require.EqualError(t, err, "400: wrong body format")
	})

	t.Run("Error; wrong player id", func(t *testing.T) {
		srv.On("Play", 1).Return(1, 2, "lose", "Paper covers rock.")
		body, err := json.Marshal(struct {
			Player int `json:"player"`
		}{Player: 99})
		require.NoError(t, err)
		bodyReader := bytes.NewReader(body)
		r := httptest.NewRequest(http.MethodGet, "/play", bodyReader)
		w := httptest.NewRecorder()

		r.Header[rest.ContentType] = []string{rest.ContentTypeJson}
		err = resource.PostPlay(w, r)
		require.EqualError(t, err, "400: player choice ID must be in range [1:5]")
	})
}
