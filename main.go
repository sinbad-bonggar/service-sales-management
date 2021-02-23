package main

import (
	"net/http"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/config"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/service"
	internalHttp "github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/handler"
	internalMiddleware "github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/middleware"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/storage/postgres"
	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	conf := config.NewConfig()

	db, _ := postgres.NewDBConnection()

	srv := internalHttp.New()

	permanentJourneyPlanPersist := postgres.NewPermanentJourneyPlanRepository(db)
	permanentJourneyPlanService := service.NewPermanentJourneyPlanService(permanentJourneyPlanPersist)
	permanentJourneyPlanHandler := handler.NewPermanentJourneyPlanHandler(permanentJourneyPlanService)

	srv.Router.With(internalMiddleware.SbpPassport).With(internalMiddleware.SearchPagination).Route("/v1", func(r chi.Router) {
		r.Route("/pjp", func(r chi.Router) {
			r.Get("/{id}", permanentJourneyPlanHandler.GetPermanentJourneyPlan)
			r.Get("/", permanentJourneyPlanHandler.GetAllPermanentJourneyPlan)
		})
	})

	http.ListenAndServe(":"+conf.AppPort, srv)
}
