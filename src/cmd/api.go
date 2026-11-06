package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlcRepo "virtual.economy/src/internal/adapters/postgresql/sqlc/gen"
	"virtual.economy/src/internal/env"
	"virtual.economy/src/internal/exchange"
	"virtual.economy/src/internal/item"
	"virtual.economy/src/internal/json"
	"virtual.economy/src/internal/market"
	"virtual.economy/src/internal/profile"
	"virtual.economy/src/internal/wallet"
)

func main() {
	ctx := context.Background()

	env.Load()
	conn, err := pgxpool.New(ctx, env.Get("DATABASE_URL", "postgresql://superadmin:nooneknows@localhost:5432/virtualeconomy"))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) { // simple ping pong
		json.Write(w, http.StatusOK, "pong")
	})

	queries := sqlcRepo.New(conn)

	profileRepo := profile.NewSqlcRepository(queries)
	itemRepo := item.NewSqlcRepository(conn)
	walletRepo := wallet.NewSqlcRepository(queries)
	exchangeRepo := exchange.NewSqlcRepository(queries)

	itemService := item.NewService(itemRepo)
	itemHandler := item.NewHandler(itemService)
	item.Router(itemHandler, mux)

	walletService := wallet.NewService(walletRepo)
	walletHandler := wallet.NewHandler(walletService)
	wallet.Router(walletHandler, mux)

	exchangePolicy, err := exchange.NewPolicy(5)

	if err != nil {
		log.Fatal(err)
	}

	exchangeService := exchange.NewService(exchangeRepo, walletService, exchangePolicy)
	exchangeHandler := exchange.NewHandler(exchangeService)
	exchange.Router(exchangeHandler, mux)

	profileService := profile.NewService(profileRepo, exchangeService)
	profileHandler := profile.NewHandler(profileService)
	profile.Router(profileHandler, mux)

	marketService := market.NewService(itemService, exchangeService, profileService)
	marketHandler := market.NewHandler(marketService)
	market.Router(marketHandler, mux)

	http.ListenAndServe(env.Get("PORT", ":8080"), mux)
}
