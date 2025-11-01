package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var stmtName = "get_imei_stmt"
var max_connection = 50

type GetImeiItemResponse struct {
	Msisdn       string `json:"msisdn"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	ImeiHash     string `json:"imei-hash"`
	Duplicated   string `json:"duplicated"`
	Timestmap    string `json:"timestamp"`
}

func main() {
	pgConf, _ := pgxpool.ParseConfig("postgres://postgres:1qazxsw23@192.168.17.98:5432/postgres?application_name=pgxpool_test")

	pgConf.MaxConnLifetime = 1 * time.Hour
	pgConf.MaxConns = int32(max_connection)
	pgConf.MinConns = int32(3)

	pgConf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// fmt.Print("Callback AfterConnect")
		_, err := conn.Prepare(ctx, stmtName,
			`select 
					m.msisdn,
					md5(substring(cl.imei,1,14)) imei,
					to_char(cl.event_dt,'YYYY-MM-DD"T"HH24:MI:SS') dt,
					edd.vendor_name,
					edd.model_name,
					exists(select from subs_ctrl.eir_imei_duplicates d where d.imei = cl.imei) is_duplicate
				from unnest($1::text[]) m(msisdn)
					left join subs_ctrl.eir_current_links cl on m.msisdn = cl.msisdn
					left join subs_ctrl.eir_dms_data edd on substr(cl.imei, 1, 8) = edd.tac`)
		if err != nil {
			fmt.Printf("prepare failed: %v", err)
			return err
		}
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, pgConf)
	if err != nil {
		fmt.Printf("Unable to connect to database: %s", err)
		return
	}

	// conn, err := pool.Acquire(ctx)
	// _, err = conn.Conn().Prepare(ctx, stmtName,
	// 	`select
	// 				m.msisdn,
	// 				md5(substring(cl.imei,1,14)) imei,
	// 				to_char(cl.event_dt,'YYYY-MM-DD"T"HH24:MI:SS') dt,
	// 				edd.vendor_name,
	// 				edd.model_name,
	// 				exists(select from subs_ctrl.eir_imei_duplicates d where d.imei = cl.imei) is_duplicate
	// 			from unnest($1::text[]) m(msisdn)
	// 				left join subs_ctrl.eir_current_links cl on m.msisdn = cl.msisdn
	// 				left join subs_ctrl.eir_dms_data edd on substr(cl.imei, 1, 8) = edd.tac`)
	// if err != nil {
	// 	fmt.Printf("prepare failed: %v", err)
	// 	return
	// }
	// conn.Release()

	// -------------------------------------------------------------------

	_, reg_err := pool.Exec(
		context.Background(),
		"call subs_ctrl.init_config_robot()")

	if reg_err != nil {
		fmt.Printf("Init SQL execution error: %v\n", reg_err)
		return
	}

	conns := make([]*pgxpool.Conn, 0, max_connection)
	for range max_connection {
		conn, err := pool.Acquire(context.Background())
		if err == nil {
			conns = append(conns, conn)
			// conn.Release()
		}
	}

	for i := 0; i < len(conns); i++ {
		conns[i].Release()
	}

	msisdns := []string{
		"998338964552", "998819887062", "998513498172", "998226895212", "998867497072", "998324428462", "998231190352", "998496839992", "998034479032", "998755889182",
		"998527024612", "998798640672", "998204306602", "998122059232", "998280554472", "998059234542", "998418663272", "998536052942", "998780444952", "998994533032",
		"998338964552", "998819887062", "998513498172", "998226895212", "998867497072", "998324428462", "998231190352", "998496839992", "998034479032", "998755889182",
		"998527024612", "998798640672", "998204306602", "998122059232", "998280554472", "998059234542", "998418663272", "998536052942", "998780444952", "998994533032",
		"998338964552", "998819887062", "998513498172", "998226895212", "998867497072", "998324428462", "998231190352", "998496839992", "998034479032", "998755889182",
		"998527024612", "998798640672", "998204306602", "998122059232", "998280554472", "998059234542", "998418663272", "998536052942", "998780444952", "998994533032",
		"998338964552", "998819887062", "998513498172", "998226895212", "998867497072", "998324428462", "998231190352", "998496839992", "998034479032", "998755889182",
		"998527024612", "998798640672", "998204306602", "998122059232", "998280554472", "998059234542", "998418663272", "998536052942", "998780444952", "998994533032",
	}

	concurrency := len(msisdns)
	var wg sync.WaitGroup
	wg.Add(concurrency)

	startTotal := time.Now()

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < concurrency; i++ {
		go func(idx int) {
			defer wg.Done()
			m := msisdns[idx]
			arr := []string{m}

			start := time.Now()
			rows, err := pool.Query(ctx, stmtName, arr)
			if err != nil {
				fmt.Printf("[%02d] Query error: %v\n", idx, err)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var r GetImeiItemResponse
				rows.Scan(&r.Msisdn, &r.ImeiHash, &r.Timestmap, &r.Manufacturer, &r.Model, &r.Duplicated)
				//fmt.Printf("%v", r)
			}

			fmt.Printf("[%02d] Query latency: %v\n", idx, time.Since(start))
		}(i)
	}

	wg.Wait()
	fmt.Printf("=== Total test time: %v ===\n", time.Since(startTotal))
	stat := pool.Stat()
	fmt.Printf("Pool stats: total=%d idle=%d acquired=%d\n",
		stat.TotalConns(), stat.IdleConns(), stat.AcquiredConns())

	/*
		for _, m := range msisdns {

			var arr []string
			arr = append(arr, m)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			start := time.Now()
			rows, _ := pool.Query(ctx, stmtName, arr)
			defer rows.Close()

			fmt.Printf("Query latency: %v\n", time.Since(start))
			for rows.Next() {
				var r GetImeiItemResponse
				rows.Scan(&r.Msisdn, &r.ImeiHash, &r.Timestmap, &r.Manufacturer, &r.Model, &r.Duplicated)
			}

		}
	*/

}
