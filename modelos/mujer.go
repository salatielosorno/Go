package modelos

/* Mujer hereda todo de hombre */
type Mujer struct {
	Hombre
	EsMadre bool
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Sexo() string { return "mujer" }
