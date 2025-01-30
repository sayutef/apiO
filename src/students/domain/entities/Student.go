package entities

type Student struct {
	ID        int    `json:"id_student" db:"id_student"`
	Name      string `json:"name_student" db:"name_student"`
	LastName  string `json:"last_name_student" db:"last_name_student"`
	Matricule string `json:"matricule_student" db:"matricule_student"`
	Age       int    `json:"age_student" db:"age_student"`
}
