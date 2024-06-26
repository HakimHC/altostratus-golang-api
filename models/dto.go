package models

type AsteroidsPostDTO struct {
	Name          string           `json:"name" validate:"required"`
	Diameter      float64          `json:"diameter" validate:"required,gte=0"`
	DiscoveryDate string           `json:"discovery_date" validate:"required"`
	Observations  string           `json:"observations,omitempty" validate:"omitempty,min=1"`
	Distances     []DistanceStruct `json:"distances,omitempty" validate:"omitempty,dive"`
}

type AsteroidsPatchDTO struct {
	Name          string           `json:"name"`
	Diameter      float64          `json:"diameter" validate:"gte=0"`
	DiscoveryDate string           `json:"discovery_date"`
	Observations  string           `json:"observations"`
	Distances     []DistanceStruct `json:"distances" validate:"dive"`
}
