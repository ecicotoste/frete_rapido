package repository

import (
	"database/sql"
	"log/slog"
	"strings"

	"api.frete.rapido/internal/entity"
	_ "github.com/go-sql-driver/mysql"
)

type MetricsRepositoryMysql struct {
	DB *sql.DB
}

func NewMetricsRepositoryMysql(db *sql.DB) *MetricsRepositoryMysql {
	return &MetricsRepositoryMysql{DB: db}
}

func (r *MetricsRepositoryMysql) Insert(metrics *entity.Metrics) error {
	_, err := r.DB.Exec("Insert into metrics (id, regtransp, company, price) values (?,?,?,?)",
		metrics.ID, metrics.IdTransp, metrics.CompName, metrics.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *MetricsRepositoryMysql) Query(limQuote string) ([]*entity.RMetrics, error) {

	query := " with ids as (select distinct id, dthr from metrics order by dthr desc " + limQuote + ") "
	query += " select count(*) tot_transp, company, ROUND(SUM(price), 2) tot_price, ROUND(avg(price), 2) as media_price,"
	query += " (select min(price) from metrics) as min_all_price, "
	query += " (select max(price) from metrics) as max_all_price"
	query += " from metrics where id in (select id from ids)"
	query += " group by company;"

	slog.Info("=====> StrQuery: " + query)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ListMetrics []*entity.RMetrics
	for rows.Next() {
		Metrics := &entity.RMetrics{}
		err = rows.Scan(&Metrics.TotResTransp, &Metrics.Carrier, &Metrics.TotalFinalPrice, &Metrics.AverageFinalPrice, &Metrics.MinAllPrice, &Metrics.MaxAllPrice)
		if err != nil {
			if !strings.Contains(err.Error(), "no rows in result set") {
				return nil, err
			}
		}
		ListMetrics = append(ListMetrics, Metrics)
	}
	return ListMetrics, nil
}
