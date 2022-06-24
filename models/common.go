package models

type Response struct {
	Error   string
	Message string
	Data    interface{}
}

// // getting time
// dt := time.Now()
// fmt.Println(dt.Format("01-02-2006 15:04:05"))
