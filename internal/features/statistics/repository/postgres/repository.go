package statistics_postgres_repository

import core_postgres_pool "github.com/simonkefir/golang-todoapp/internal/core/repository/postgres/pool"

type StatisticsRepostiory struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(
	pool core_postgres_pool.Pool,
) *StatisticsRepostiory {
	return &StatisticsRepostiory{
		pool: pool,
	}
}
