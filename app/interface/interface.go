package interfaces

type IEntity interface {
	DeleteEntity(key string , value string) error
	GetEntity(param string) (interface{}, error)
	UpdateData(key string, payload interface{}) error
	InsertData(payload interface{}) error
	SetElement(typ string, value interface{}) error
	GetElement(msg string) (*string, error)
	FindDocument(key string, val string) (interface{}, error)
}
