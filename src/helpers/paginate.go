package helpers

const ITEMS_PER_PAGE = 100

type pagination struct {
	TotalItems   int
	TotalPages   int
	CurrentPage  int
	ItemsPerPage int
	LastPage     int
	FirstPage    int
}

func NewPagination(totalItems int, currentPage int) *pagination {
	totalPages := totalItems / ITEMS_PER_PAGE
	if totalItems%ITEMS_PER_PAGE > 0 {
		totalPages++
	}

	return &pagination{
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		CurrentPage:  currentPage,
		ItemsPerPage: ITEMS_PER_PAGE,
		LastPage:     totalPages,
		FirstPage:    1,
	}
}

func (p *pagination) NextPage() int {
	if p.CurrentPage < p.LastPage {
		return p.CurrentPage + 1
	}
	return p.LastPage
}

func (p *pagination) PreviousPage() int {
	if p.CurrentPage > p.FirstPage {
		return p.CurrentPage - 1
	}
	return 1
}

func (p *pagination) SetNextPage() {
	if p.CurrentPage < p.LastPage {
		p.CurrentPage++
	}
}

func (p *pagination) SetPreviousPage() {
	if p.CurrentPage > p.FirstPage {
		p.CurrentPage--
	}
}
