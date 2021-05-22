package property

type Property interface {
	ID() ID
}

type Prop struct {
	id ID
}

func MakeProp(id ID) Prop {
	return Prop{
		id: id,
	}
}

func (p Prop) ID() ID {
	return p.id
}
