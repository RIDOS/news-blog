package pagination

type Paginator interface {
	GetAll(limit, offset int) ([]interface{}, error)
	Count() (int, error)
}
