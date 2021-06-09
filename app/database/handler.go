package database

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

/**
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.RequestPayload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	users:= h.db.Collection("users")
	// rules:= h.db.Collection("rules")
	// policy_classes := h.db.Collection("policy_classes")
	policy_class_rules := h.db.Collection("policy_class_rules")
	// usages := h.db.Collection("usages")
	// user_rules := h.db.Collection("user_rules")
	user_policy_classes := h.db.Collection("user_policy_classes")

	var user bson.M
	err = users.FindOne(r.Context(), bson.M{"email": t.Sasl_username}).Decode(&user)
	if err != nil {
		res, _ := app.Action("reject_not_authorized")
		json.NewEncoder(w).Encode(res)
		return
	}

	// priority check
	user_id := user["id"]
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"priority": -1})
	cur, err := user_policy_classes.Find(r.Context(), bson.M{"user_id":user_id},findOptions)
	if err != nil {
		log.Println(err)
	}
	var priorities []bson.M
	defer cur.Close(r.Context())
	if err = cur.All(r.Context(), &priorities); err != nil {
		log.Print(err)
	}

	// log.Println(priorities)
	var rule_list []bson.M
	for _, v:= range priorities {
		policy_class_id := v["policy_class_id"]
		cur, err := policy_class_rules.Find(r.Context(), bson.M{"policy_class_id":policy_class_id})
		if err != nil {
			log.Println(err)
		}
		defer cur.Close(r.Context())
		if err = cur.All(r.Context(), &rule_list); err != nil {
			log.Print(err)
		}
		// log.Println(rule_list)
		for _, vv := range(rule_list) {
			// log.Printf("%v", vv["rule_id"])
			switch vv["rule_id"] {
			case 1.0:
				if  t.Smtp_session_data.Sasl_sender == "" {
					res, _ := app.Action("reject_null_sender")
					json.NewEncoder(w).Encode(res)
					return
				}
			case 2.0:
				if t.Sasl_username != t.Smtp_session_data.Sasl_username {
					res, _ := app.Action("reject_sender_login_mismatch")
					json.NewEncoder(w).Encode(res)
					return
				}
			case 3.0:
			case 4.0:
			case 5.0:
			case 6.0:
			case 7.0:
			case 8.0:
			case 9.0:
			case 10.0:
			case 11.0:
				// var data bson.M
				// err = policy_class_rules.FindOne(r.Context(), bson.M{"rule_id":11}).Decode(&data)
				// log.Println(data)
				// if err != nil { log.Printf("findone rules error %+v", err)}
				// size := app.ToInt(data["default_value"].(string))
				// if _size := app.ToInt(t.Smtp_session_data.Size); _size > size{
				// 	res, _:= app.Action("reject_message_size_exceeded")
				// 	// w.WriteHeader(code)
				// 	json.NewEncoder(w).Encode(res)
				// 	// json.NewEncoder(w).Encode(t)
				// 	return
				// 	}
			case 12.0:
			case 13.0:
				var data bson.M
				err = policy_class_rules.FindOne(r.Context(), bson.M{"rule_id":13}).Decode(&data)
				// log.Println(data)
				if err != nil { log.Printf("findone rules error %+v", err)}
				size := app.ToInt(data["default_value"].(string))
				if _size := app.ToInt(t.Smtp_session_data.Size); _size > size{
					res, _:= app.Action("reject_exceed_msg_size")
					// w.WriteHeader(code)
					json.NewEncoder(w).Encode(res)
					// json.NewEncoder(w).Encode(t)
					return
					}
			case 14.0:
				// var data bson.M
				// err = policy_class_rules.FindOne(r.Context(), bson.M{"rule_id":14}).Decode(&data)
				// log.Println(data)
				// if err != nil { log.Printf("findone rules error %+v", err)}
				// size := app.ToInt(data["default_value"].(string))
				// if _size := app.ToInt(t.Smtp_session_data.Size); _size > size{
				// 	res, _:= app.Action("reject_exceed_max_msgs")
				// 	// w.WriteHeader(code)
				// 	json.NewEncoder(w).Encode(res)
				// 	// json.NewEncoder(w).Encode(t)
				// 	return
				// 	}
			case 15.0:
				// var data bson.M
				// err = policy_class_rules.FindOne(r.Context(), bson.M{"rule_id":15}).Decode(&data)
				// log.Println(data)
				// if err != nil { log.Printf("findone rules error %+v", err)}
				// size := app.ToInt(data["default_value"].(string))
				// if _size := app.ToInt(t.Smtp_session_data.Size); _size > size{
				// 	res, _:= app.Action("reject_exceed_max_quota")
				// 	// w.WriteHeader(code)
				// 	json.NewEncoder(w).Encode(res)
				// 	// json.NewEncoder(w).Encode(t)
				// 	return
				// 	}
			case 16.0:
			case 17.0:
				if t.Sender != t.Smtp_session_data.Sasl_username {
					res, _ := app.Action("reject_sender_login_mismatch")
					json.NewEncoder(w).Encode(res)
					return
				}
			default:
				// w.Header().Set("Content-Type", "application/json")
				res, _ := app.Action(models.Rules["default"])
				// w.WriteHeader(code)
				json.NewEncoder(w).Encode(res)

				}
			}
		}
	}
**/



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











// type Collection interface{
//     InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
//     // InsertMany(ctx context.Context, document interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
//     FindOne(ctx context.Context, document interface{}, opts ...*options.FindOneOptions) (*mongo.SingleResult)
// 	Find(ctx context.Context, document interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
// }
// 
// type Handler struct {
//     coll Collection
// }
// 
// func New(c Collection) *Handler {
//     return &Handler{
//         coll: c,
//     }
// }

// func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var t models.RequestPayload
// 	err := decoder.Decode(&t)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// log.Printf("%T, %v", t, t)
// 	if  _, err := h.coll.InsertOne(r.Context(), t); err != nil {
//         fmt.Errorf("Insert question error: %+v", err)
//         return
//     }
// }
// 
// func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
// 	var data models.RequestPayload
// 	// var data bson.M
// 	if  err := h.coll.FindOne(r.Context(), r.Body).Decode(&data); err != nil {
//         fmt.Errorf("Insert question error: %+v", err)
//         return
//     }
// 	// log.Println(data)
// }
