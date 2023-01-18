package domain

type Store struct {
	ID          string     `json:"id" bson:"_id"`
	Name        string     `json:"name" bson:"name"`
	Description string     `json:"description" bson:"description"`
	Address     string     `json:"address" bson:"address"`
	Phone       string     `json:"phone" bson:"phone"`
	Latitude    string     `json:"latitude" bson:"latitude"`
	Longitude   string     `json:"longitude" bson:"longitude"`
	MenuItems   []MenuItem `json:"menuItems" bson:"menuItems, omitempty"`
}

type MenuItem struct {
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Price       string   `json:"price" bson:"price"`
	Reviews     []Review `json:"reviews" bson:"reviews, omitempty"`
}

type Review struct {
	ID           string `json:"id" bson:"_id"`
	Review       string `json:"review" bson:"review"`
	Flavor       int    `json:"flavor" bson:"flavor"`
	Texture      int    `json:"texture" bson:"texture"`
	Presentation int    `json:"presentation" bson:"presentation"`
	Value        int    `json:"value" bson:"value"`
	Atmosphere   int    `json:"atmosphere" bson:"atmosphere"`
}
