package model

type Action struct {
	ActionName string `json:"action_name"`
	Status     string `json:"status"`
	StatusDesc string `json:"status_desc"`
}

type Actions struct {
	Type    string   `json:"type"`
	Payload []Action `json:"payload"`
}

func (a *Actions) DeleteEntity(param string) error {
	return nil
}

func (a *Actions) GetEntity(param string) (interface{}, error) {
	return nil, nil
}

func (a *Actions) UpdateData(payload interface{}) error {
	return nil
}

func (a *Actions) InsertData(payload interface{}) error {
	return nil
}

func (a *Actions) SetElement(typ string, value interface{}) error {
	return nil
}

func (a *Actions) GetElement(msg string) (*string, error) {
	return nil, nil
}

func (a *Actions) FindDocument(key string, val string) (interface{}, error) {
	return nil, nil
}
