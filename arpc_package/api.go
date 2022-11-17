//119459-110-11 20:00:00

package api

import "encoding/json"

type ApiRequestV1 struct {
	UserId int
}

func (b *ApiRequestV1) New(user_id int) {
	b.UserId = user_id
}

func (b *ApiRequestV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ApiRequestV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

type ApiResponseV1 struct {
	UserId   int
	Username string
}

func (b *ApiResponseV1) New(user_id int, username string) {
	b.UserId = user_id
	b.Username = username
}

func (b *ApiResponseV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ApiResponseV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

func GetUserV1(request *ApiRequestV1) (*ApiResponseV1, error) {
	return &ApiResponseV1{UserId: request.UserId, Username: "username"}, nil
}
