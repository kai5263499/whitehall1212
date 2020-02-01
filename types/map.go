package types

type Vertex int

type Transportation string

type Edge struct {
	Transportation Transportation `json:"ex:transportation"`
	Destination    Vertex         `json:"ex:destination"`
}

type Edges []Edge
