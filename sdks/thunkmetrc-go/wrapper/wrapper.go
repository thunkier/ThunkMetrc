package wrapper

import (
	"log/slog"

	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
	"github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)

type MetrcWrapper struct {
	Client      *client.MetrcClient
	RateLimiter services.RateLimiter
	AdditivesTemplates *services.AdditivesTemplatesService
	CaregiversStatus *services.CaregiversStatusService
	Employees *services.EmployeesService
	Facilities *services.FacilitiesService
	Harvests *services.HarvestsService
	Items *services.ItemsService
	LabTests *services.LabTestsService
	Locations *services.LocationsService
	Packages *services.PackagesService
	PatientCheckIns *services.PatientCheckInsService
	Patients *services.PatientsService
	PatientsStatus *services.PatientsStatusService
	PlantBatches *services.PlantBatchesService
	Plants *services.PlantsService
	ProcessingJob *services.ProcessingJobService
	RetailId *services.RetailIdService
	Sales *services.SalesService
	Sandbox *services.SandboxService
	Strains *services.StrainsService
	Sublocations *services.SublocationsService
	Tags *services.TagsService
	Transfers *services.TransfersService
	Transporters *services.TransportersService
	UnitsOfMeasure *services.UnitsOfMeasureService
	WasteMethods *services.WasteMethodsService
	Webhooks *services.WebhooksService
}

func New(c *client.MetrcClient) *MetrcWrapper {
	rl := services.NewRateLimiter(services.DefaultConfig(), c.Logger)
	return NewWithLimiter(c, rl)
}

func NewMetrcWrapper(baseUrl string, accessKey string, logger *slog.Logger) *MetrcWrapper {
	c := client.New(baseUrl, accessKey, "")
	rl := services.NewRateLimiter(services.DefaultConfig(), logger)

	return NewWithLimiter(c, rl)
}

func NewWithLimiter(client *client.MetrcClient, rateLimiter services.RateLimiter) *MetrcWrapper {
	return &MetrcWrapper{
		Client:      client,
		RateLimiter: rateLimiter,
			AdditivesTemplates: &services.AdditivesTemplatesService{Client: client, RateLimiter: rateLimiter},
			CaregiversStatus: &services.CaregiversStatusService{Client: client, RateLimiter: rateLimiter},
			Employees: &services.EmployeesService{Client: client, RateLimiter: rateLimiter},
			Facilities: &services.FacilitiesService{Client: client, RateLimiter: rateLimiter},
			Harvests: &services.HarvestsService{Client: client, RateLimiter: rateLimiter},
			Items: &services.ItemsService{Client: client, RateLimiter: rateLimiter},
			LabTests: &services.LabTestsService{Client: client, RateLimiter: rateLimiter},
			Locations: &services.LocationsService{Client: client, RateLimiter: rateLimiter},
			Packages: &services.PackagesService{Client: client, RateLimiter: rateLimiter},
			PatientCheckIns: &services.PatientCheckInsService{Client: client, RateLimiter: rateLimiter},
			Patients: &services.PatientsService{Client: client, RateLimiter: rateLimiter},
			PatientsStatus: &services.PatientsStatusService{Client: client, RateLimiter: rateLimiter},
			PlantBatches: &services.PlantBatchesService{Client: client, RateLimiter: rateLimiter},
			Plants: &services.PlantsService{Client: client, RateLimiter: rateLimiter},
			ProcessingJob: &services.ProcessingJobService{Client: client, RateLimiter: rateLimiter},
			RetailId: &services.RetailIdService{Client: client, RateLimiter: rateLimiter},
			Sales: &services.SalesService{Client: client, RateLimiter: rateLimiter},
			Sandbox: &services.SandboxService{Client: client, RateLimiter: rateLimiter},
			Strains: &services.StrainsService{Client: client, RateLimiter: rateLimiter},
			Sublocations: &services.SublocationsService{Client: client, RateLimiter: rateLimiter},
			Tags: &services.TagsService{Client: client, RateLimiter: rateLimiter},
			Transfers: &services.TransfersService{Client: client, RateLimiter: rateLimiter},
			Transporters: &services.TransportersService{Client: client, RateLimiter: rateLimiter},
			UnitsOfMeasure: &services.UnitsOfMeasureService{Client: client, RateLimiter: rateLimiter},
			WasteMethods: &services.WasteMethodsService{Client: client, RateLimiter: rateLimiter},
			Webhooks: &services.WebhooksService{Client: client, RateLimiter: rateLimiter},
	}
}



