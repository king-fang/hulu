package form

type GeneralGetId struct {
	ID int `alias:"请求ID" valid:"Required;Min(1)"`
}
