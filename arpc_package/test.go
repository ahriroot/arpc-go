//119459-110-11 20:00:00

package api

import "encoding/json"

type TestRequestV1 struct {
    UserId int
}

func (b *TestRequestV1) New(user_id int) {
	b.UserId = user_id
}

func (b *TestRequestV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TestRequestV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

type TestResponseV1 struct {
    UserId   int
    Username string
}

func (b *TestResponseV1) New(user_id int, username string) {
	b.UserId = user_id
    b.Username = username
}

func (b *TestResponseV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *TestResponseV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}
