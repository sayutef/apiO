package entities

type Asignature struct {
	ID          int    `json:"id_asignature" db:"id_asignature"`
	Name        string `json:"name_asignature" db:"name_asignature"`
	Description string `json:"description_asignature" db:"description_asignature"`
}
