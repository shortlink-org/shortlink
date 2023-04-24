package UoW

type UnitOfWork interface {
	RegisterNew(interface{})
}
