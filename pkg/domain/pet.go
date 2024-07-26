package domain

type Contact struct {
	Owner string `json:"owner" bson:"owner"`
	Phone string `json:"phone" bson:"phone"`
	City  string `json:"city" bson:"city"`
	State string `json:"state" bson:"state"`
}

type PetID interface{} // String / Int

type Pet struct {
	ID       PetID   `json:"id" bson:"_id,omitempty"`
	Category string  `json:"category" bson:"category"`
	Breed    string  `json:"bread" bson:"breed,omitempty"`
	Age      int     `json:"age" bson:"age"`
	Gender   string  `json:"gender" bson:"gender"`
	Colors   string  `json:"colors" bson:"colors,omitempty"`
	Contact  Contact `json:"contact" bson:"contact,omitempty"`
	Price    float64 `json:"price" bson:"price,omitempty"`
}

type PetService interface {
	Get(id PetID) (*Pet, error)
	List(category string) ([]*Pet, error)
	Create(p *Pet) (*Pet, error)
	Delete(id PetID) error
}

type PetDB interface {
	Get(id PetID) (*Pet, error)
	List(category string) ([]*Pet, error)
	Create(p *Pet) (*Pet, error)
	Delete(id PetID) error
}
