using System;
using ThunkMetrc.Client;
using ThunkMetrc.Wrapper.Services;
using Microsoft.Extensions.Logging;

namespace ThunkMetrc.Wrapper
{
    public class MetrcWrapper
    {
        public MetrcClient Client { get; }
        public IMetrcRateLimiter RateLimiter { get; }

        public MetrcWrapper(MetrcClient client, IMetrcRateLimiter? rateLimiter = null, RateLimiterConfig? config = null, ILogger? logger = null)
        {
            Client = client;
            RateLimiter = rateLimiter ?? new MetrcRateLimiter(config, logger);
            AdditivesTemplates = new AdditivesTemplatesService(Client, RateLimiter);
            CaregiversStatus = new CaregiversStatusService(Client, RateLimiter);
            Employees = new EmployeesService(Client, RateLimiter);
            Facilities = new FacilitiesService(Client, RateLimiter);
            Harvests = new HarvestsService(Client, RateLimiter);
            Items = new ItemsService(Client, RateLimiter);
            LabTests = new LabTestsService(Client, RateLimiter);
            Locations = new LocationsService(Client, RateLimiter);
            Packages = new PackagesService(Client, RateLimiter);
            PatientCheckIns = new PatientCheckInsService(Client, RateLimiter);
            Patients = new PatientsService(Client, RateLimiter);
            PatientsStatus = new PatientsStatusService(Client, RateLimiter);
            PlantBatches = new PlantBatchesService(Client, RateLimiter);
            Plants = new PlantsService(Client, RateLimiter);
            ProcessingJob = new ProcessingJobService(Client, RateLimiter);
            RetailId = new RetailIdService(Client, RateLimiter);
            Sales = new SalesService(Client, RateLimiter);
            Sandbox = new SandboxService(Client, RateLimiter);
            Strains = new StrainsService(Client, RateLimiter);
            Sublocations = new SublocationsService(Client, RateLimiter);
            Tags = new TagsService(Client, RateLimiter);
            Transfers = new TransfersService(Client, RateLimiter);
            Transporters = new TransportersService(Client, RateLimiter);
            UnitsOfMeasure = new UnitsOfMeasureService(Client, RateLimiter);
            WasteMethods = new WasteMethodsService(Client, RateLimiter);
            Webhooks = new WebhooksService(Client, RateLimiter);
        }
        public AdditivesTemplatesService AdditivesTemplates { get; }
        public CaregiversStatusService CaregiversStatus { get; }
        public EmployeesService Employees { get; }
        public FacilitiesService Facilities { get; }
        public HarvestsService Harvests { get; }
        public ItemsService Items { get; }
        public LabTestsService LabTests { get; }
        public LocationsService Locations { get; }
        public PackagesService Packages { get; }
        public PatientCheckInsService PatientCheckIns { get; }
        public PatientsService Patients { get; }
        public PatientsStatusService PatientsStatus { get; }
        public PlantBatchesService PlantBatches { get; }
        public PlantsService Plants { get; }
        public ProcessingJobService ProcessingJob { get; }
        public RetailIdService RetailId { get; }
        public SalesService Sales { get; }
        public SandboxService Sandbox { get; }
        public StrainsService Strains { get; }
        public SublocationsService Sublocations { get; }
        public TagsService Tags { get; }
        public TransfersService Transfers { get; }
        public TransportersService Transporters { get; }
        public UnitsOfMeasureService UnitsOfMeasure { get; }
        public WasteMethodsService WasteMethods { get; }
        public WebhooksService Webhooks { get; }
    }
}
