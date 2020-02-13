package types

// Vertex represents a player position on the game board
type Vertex int

// Transportation represents a means of transportation tying
// two positions on the board together
type Transportation string

// Edge is represents a connection between two positions or
// vertices on the board
type Edge struct {
	Transportation Transportation `json:"ex:transportation"`
	Destination    Vertex         `json:"ex:destination"`
}
