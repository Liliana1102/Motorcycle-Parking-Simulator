package models

type Entrada struct {
	states      []string
	actualState string
}

func NewEntrada() *Entrada  {
	states := []string{"Entrando", "Saliendo", "Parado"}
	return &Entrada {
		states:      states,
		actualState: "Parado",
	}
}

func (e *Entrada ) GetState() string {
	return e.actualState
}

func (e *Entrada ) SetState(n int) {
	if n >= 0 && n < len(e.states) {
		e.actualState = e.states[n]
	}
}
