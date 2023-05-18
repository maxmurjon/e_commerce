package postgres


type authorRepo struct{
	db *pgxpool.Pool
}

