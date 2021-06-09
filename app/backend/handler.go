package backend

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"policyService/app"
	"policyService/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Database interface{
	 Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	 CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error
}

type Handler struct {
    db Database
}

func New(d Database) *Handler {
    return &Handler{
        db: d,
    }
}


// only size check
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.RequestPayload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	users:= h.db.Collection("users")
	user_rules := h.db.Collection("user_rules")

	var user models.User
	err = users.FindOne(r.Context(), bson.M{"email": t.Sasl_username}).Decode(&user)
	// log.Println(user)
	if err != nil {
		res, _ := app.Action("reject_not_authorized")
		json.NewEncoder(w).Encode(res)
		return
	}

	// size check
	// var data bson.M
	var data models.User_rule
	err = user_rules.FindOne(r.Context(), bson.M{"user_id":user.Id, "rule_id":13}).Decode(&data)
	// log.Println(data)
	if err != nil {
        w.Write([]byte(`{"This user have not rules 13"}`))
		log.Printf("findone rules error %+v", err)
		return
	}
	if app.ToInt(t.Smtp_session_data.Size) > data.Value {
		res, code := app.Action(models.Rules["reject_exceed_msg_size"])
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(res)
		// json.NewEncoder(w).Encode(t)
		return
		} else {
			w.Header().Set("Content-Type", "application/json")
			res, _ := app.Action(models.Rules["default"])
			// w.WriteHeader(code)
			json.NewEncoder(w).Encode(res)
			return
	 }
}
