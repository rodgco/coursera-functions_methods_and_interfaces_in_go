package data

type Data struct {
	name string
	value int
}

func New(name string, value int) *Data {
	return &Data{name, value}
}

func (d *Data) Name() string {
	return d.name
}

func (d *Data) Value() int {
	return d.value
}

func (d *Data) Double() {
	d.value *= 2
}
