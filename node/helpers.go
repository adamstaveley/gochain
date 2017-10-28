package node

// import (
// 	"encoding/json"
// 	"io"
// )

// func decode(body io.ReadCloser, v struct{}) struct{} {
// 	d := json.NewDecoder(body)
// 	err := d.Decode(&v)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return v
// }