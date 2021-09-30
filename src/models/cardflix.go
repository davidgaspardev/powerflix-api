package model

type Cardflix struct {
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Cover       string `json:"cover,omitempty" bson:"cover,omitempty"`
	Modules     []struct {
		Level       string `json:"level,omitempty" bson:"level,omitempty"`
		Description string `json:"description,omitempty" bson:"description,omitempty"`
		Frequency   struct {
			Repetition int `json:"repetition,omitempty" bson:"repetition,omitempty"`
			Series     int `json:"series,omitempty" bson:"series,omitempty"`
		} `json:"frequency,omitempty" bson:"frequency,omitempty"`
		Exercises []struct {
			Order    int    `json:"order,omitempty" bson:"order,omitempty"`
			Name     string `json:"name,omitempty" bson:"name,omitempty"`
			Metadata string `json:"metadata,omitempty" bson:"metadata,omitempty"`
		} `json:"exercises,omitempty" bson:"exercises,omitempty"`
	} `json:"modules,omitempty" bson:"modules,omitempty"`
}
