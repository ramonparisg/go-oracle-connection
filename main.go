package main

import (
	"database/sql"
	slog "github.com/go-eden/slf4go"
	_ "github.com/godror/godror"
	"os"
	"strconv"
	"strings"
)

type COUpdateWrapperDTO struct {
	orderNumber int64
	source      string
	xml         string
	success     bool
	referenceId string
}

func getUrl() string {
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbUrl := os.Getenv("DB_URL")

	return DbUser + "/" + DbPass + "@" + DbUrl
}

func main() {
	slog.Info("===== EMPEZÓ EL PROCESO =====")

	coUpdates, err := findAllCoUpdates()
	if err != nil {
		slog.Errorf("No pudo buscar coupdates. Err: %s", err)
		return
	}

	slog.Infof("Se encontraron %s coupdates", strconv.Itoa(len(coUpdates)))
}

func findAllCoUpdates() ([]*COUpdateWrapperDTO, error) {
	db, err := sql.Open("godror", getUrl())

	if err != nil {
		slog.Error("No conectó")
		return nil, err
	}
	defer db.Close()

	//noinspection SqlNoDataSourceInspection
	SQL := os.Getenv("QUERY")

	rows, err := db.Query(SQL)
	if err != nil {
		slog.Error("Error running query")
		return nil, err
	}
	defer rows.Close()
	slog.Info("Query correcta")
	var coUpdates []*COUpdateWrapperDTO
	for rows.Next() {
		var referenceId, origin, txml string
		var hasErrors bool
		rows.Scan(&referenceId, &origin, &txml, &hasErrors)
		splicedReferenceId := strings.Split(referenceId, "_")
		orderNumber, _ := strconv.ParseInt(splicedReferenceId[len(splicedReferenceId)-1], 0, 64)
		wrapper := &COUpdateWrapperDTO{
			orderNumber: orderNumber,
			source:      origin,
			xml:         txml,
			success:     !hasErrors,
			referenceId: referenceId,
		}

		coUpdates = append(coUpdates, wrapper)
	}

	return coUpdates, nil
}
