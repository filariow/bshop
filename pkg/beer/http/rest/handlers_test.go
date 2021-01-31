package rest_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/filariow/bshop/pkg/beer/http/rest"
	"github.com/filariow/bshop/pkg/beer/mocks"
	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
)

func toJSON(t *testing.T, data interface{}) io.Reader {
	p, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Error marshaling data to JSON")
	}

	return bytes.NewReader(p)
}

func Test_CreateBeer(t *testing.T) {
	type request struct {
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}
	type response struct {
		status int
	}
	type storageResponse struct {
		times int
		id    uint64
		err   error
	}
	type testCase struct {
		name string
		q    request
		a    response
		st   storageResponse
	}

	ttc := []testCase{
		{
			name: "Valid",
			q: request{
				Name:   "First",
				Cost:   1.0,
				Price:  2.0,
				Brewer: "Brewer",
				Size:   0.5,
				Vol:    4.5,
			},
			a: response{
				status: http.StatusOK,
			},
			st: storageResponse{
				times: 1,
				id:    1,
				err:   nil,
			},
		},
		{
			name: "Minimum valid beer details",
			q: request{
				Name:  "Name",
				Price: 1.0,
			},
			a: response{
				status: http.StatusOK,
			},
			st: storageResponse{
				times: 1,
			},
		},
		{
			name: "No beer name",
			q: request{
				Name:   "",
				Cost:   1.0,
				Price:  2.0,
				Brewer: "Brewer",
				Size:   0.5,
				Vol:    4.5,
			},
			a: response{
				status: http.StatusBadRequest,
			},
			st: storageResponse{
				times: 0,
			},
		},
		{
			name: "Invalid beer price",
			q: request{
				Name:   "Name",
				Cost:   1.0,
				Price:  -1.0,
				Brewer: "Brewer",
				Size:   0.5,
				Vol:    4.5,
			},
			a: response{
				status: http.StatusBadRequest,
			},
			st: storageResponse{
				times: 0,
			},
		},
		{
			name: "Invalid beer cost",
			q: request{
				Name:   "Name",
				Cost:   -1.0,
				Price:  1.0,
				Brewer: "Brewer",
				Size:   0.5,
				Vol:    4.5,
			},
			a: response{
				status: http.StatusBadRequest,
			},
			st: storageResponse{
				times: 0,
			},
		},
	}

	tf := func(t *testing.T, tc *testCase) {
		is := is.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockBeerRepository(ctrl)
		s := rest.New(r, "/beers")

		r.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(tc.st.id, tc.st.err).
			Times(tc.st.times)

		w := httptest.NewRecorder()
		q := httptest.NewRequest(http.MethodPost, "/beers", toJSON(t, &tc.q))
		s.ServeHTTP(w, q)

		is.Equal(w.Result().StatusCode, tc.a.status)
	}

	for _, tc := range ttc {
		t.Run(tc.name, func(t *testing.T) { tf(t, &tc) })
	}
}
