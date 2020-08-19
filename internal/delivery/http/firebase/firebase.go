package firebase

import (
	"context"
	"encoding/json"
	"errors"
	firebaseEntity "firebase/internal/entity/firebase"
	pEntity "firebase/internal/entity/returTnIn"
	"firebase/pkg/response"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// IfirebaseSvc is an interface to firebase Service
// Masukkan function dari service ke dalam interface ini
type IfirebaseSvc interface {
	InsertNewData(ctx context.Context, member firebaseEntity.Firebase) error
	TampilSemuaData(ctx context.Context) ([]firebaseEntity.Firebase, error)
	EditData(ctx context.Context, member firebaseEntity.Firebase) error
	GenerateMemberID(ctx context.Context) (string, error)
	TampilAllDataByNoReceive(ctx context.Context, NoReceive string) (pEntity.JSONRC, error)
}

type (
	// Handler ...
	Handler struct {
		firebaseSvc IfirebaseSvc
	}
)

// New for bridging product handler initialization
func New(is IfirebaseSvc) *Handler {
	return &Handler{
		firebaseSvc: is,
	}
}

// FirebaseHandler will receive request and return response
func (h *Handler) FirebaseHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      context.Context
		_token   string
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	if r.Header.Get("Authorization") == "" {
		_token = r.Header.Get("Authorization")
		if strings.Contains(_token, "Token ") {
			err = errors.New("403 - Forbidden")
		}

		type ctxKey string
		ctx = context.WithValue(ctx, ctxKey("_token"), _token)
	}

	if err == nil {
		switch r.Method {
		// Check if request method is GET
		case http.MethodGet:
			len := len(r.URL.Query())
			switch len {
			case 1:
				_, insertOK := r.URL.Query()["memberID"]
				_, NoReceiveOK := r.URL.Query()["NoReceive"]
				if insertOK {
					result, err = h.firebaseSvc.GenerateMemberID(context.Background())
				} else if NoReceiveOK {
					result, err = h.firebaseSvc.TampilAllDataByNoReceive(context.Background(), r.FormValue("NoReceive"))
				} else {
					err = errors.New("400")
				}
			case 0:
				result, err = h.firebaseSvc.TampilSemuaData(context.Background())
			default:
				err = errors.New("400")
			}
		// Check if request method is POST
		case http.MethodPost:
			len := len(r.URL.Query())
			switch len {
			case 1:
				var member firebaseEntity.Firebase
				_, insertOK := r.URL.Query()["insert"]
				body, _ := ioutil.ReadAll(r.Body)

				if insertOK {
					json.Unmarshal(body, &member)
					err = h.firebaseSvc.InsertNewData(context.Background(), member)
				} else {
					err = errors.New("400")
				}
			}
		// Check if request method is PUT
		case http.MethodPut:
			len := len(r.URL.Query())
			switch len {
			case 1:
				var member firebaseEntity.Firebase
				_, updateOK := r.URL.Query()["update"]
				body, _ := ioutil.ReadAll(r.Body)

				if updateOK {
					json.Unmarshal(body, &member)
					err = h.firebaseSvc.EditData(context.Background(), member)
				} else {
					err = errors.New("400")
				}
			}

		// Check if request method is DELETE
		case http.MethodDelete:

		default:
			err = errors.New("400")
		}
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "101 - Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   500,
				Msg:    "500 - Internal Server Error",
				Status: true,
			}
		}
		// If error 401
		if strings.Contains(err.Error(), "401") {
			// Replace error with server error
			errRes = response.Error{
				Code:   401,
				Msg:    "401 - Unauthorized",
				Status: true,
			}
		}
		// If error 403
		if strings.Contains(err.Error(), "403") {
			// Replace error with server error
			errRes = response.Error{
				Code:   403,
				Msg:    "403 - Forbidden",
				Status: true,
			}
		}
		// If error 400
		if strings.Contains(err.Error(), "400") {
			// Replace error with server error
			errRes = response.Error{
				Code:   400,
				Msg:    "400 - Bad Request",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = errRes.Code
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
