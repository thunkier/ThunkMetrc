using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/additives", body, queryParams);
        }

        /// <summary>
        /// Records additive usage details applied to specific plants at a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/additives/bylocation", body, queryParams);
        }

        /// <summary>
        /// Records additive usage for plants based on their location within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/bylocation", body, queryParams);
        }

        /// <summary>
        /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationUsingtemplateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/bylocation/usingtemplate", body, queryParams);
        }

        /// <summary>
        /// Records additive usage for plants using predefined additive templates at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesUsingtemplateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/usingtemplate", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreateChangegrowthphasesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/changegrowthphases", body, queryParams);
        }

        /// <summary>
        /// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateHarvestplantsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/harvestplants", body, queryParams);
        }

        /// <summary>
        /// Creates harvest product records from plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateManicureV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/manicure", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateManicureplantsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/manicureplants", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreateMoveplantsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/moveplants", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantsCreatePlantbatchPackageV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/create/plantbatch/packages", body, queryParams);
        }

        /// <summary>
        /// Creates packages from plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantsCreatePlantbatchPackageV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/plantbatch/packages", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreatePlantingsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/create/plantings", body, queryParams);
        }

        /// <summary>
        /// Creates new plant batches at a specified Facility from existing plant data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreatePlantingsV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/plantings", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantsCreateWasteV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v1/waste", body, queryParams);
        }

        /// <summary>
        /// Records waste events for plants at a Facility, including method, reason, and location details.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantsCreateWasteV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/waste", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Destroy Veg/Flower Plants
        /// </summary>
        public async Task PlantsDeleteV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plants/v1", null, queryParams);
        }

        /// <summary>
        /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Destroy Veg/Flower Plants
        /// </summary>
        public async Task PlantsDeleteV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plants/v2", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Plant by Id.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View/Manage Plants Additives
        /// </summary>
        public async Task<object> PlantsGetAdditivesV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/additives", null, queryParams);
        }

        /// <summary>
        /// Retrieves additive records applied to plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View/Manage Plants Additives
        /// </summary>
        public async Task<object> PlantsGetAdditivesV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/additives", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> PlantsGetAdditivesTypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/additives/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of all plant additive types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetAdditivesTypesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/additives/types", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetByLabelV1(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/{Uri.EscapeDataString(label)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Plant by label.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetByLabelV2(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/{Uri.EscapeDataString(label)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetFloweringV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/flowering", null, queryParams);
        }

        /// <summary>
        /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetFloweringV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/flowering", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetGrowthPhasesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/growthphases", null, queryParams);
        }

        /// <summary>
        /// Retrieves the list of growth phases supported by a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetGrowthPhasesV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/growthphases", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves inactive plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves mother-phase plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother", null, queryParams);
        }

        /// <summary>
        /// Retrieves inactive mother-phase plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother/onhold", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/onhold", null, queryParams);
        }

        /// <summary>
        /// Retrieves plants that are currently on hold at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/onhold", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetVegetativeV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/vegetative", null, queryParams);
        }

        /// <summary>
        /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetVegetativeV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/vegetative", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of recorded plant waste events for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWasteV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteMethodsAllV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/waste/methods/all", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of all available plant waste methods for use within a Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteMethodsAllV2(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/methods/all", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWastePackageV2(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/{Uri.EscapeDataString(id)}/package", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWastePlantV2(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/{Uri.EscapeDataString(id)}/plant", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteReasonsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v1/waste/reasons", null, queryParams);
        }

        /// <summary>
        /// Retriveves available reasons for recording mature plant waste at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteReasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/reasons", null, queryParams);
        }

        /// <summary>
        /// Adjusts the recorded count of plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateAdjustV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/adjust", body, queryParams);
        }

        /// <summary>
        /// Changes the growth phases of plants within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateGrowthphaseV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/growthphase", body, queryParams);
        }

        /// <summary>
        /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsUpdateHarvestV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/harvest", body, queryParams);
        }

        /// <summary>
        /// Moves plant batches to new locations within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateLocationV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/location", body, queryParams);
        }

        /// <summary>
        /// Merges multiple plant groups into a single group within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsUpdateMergeV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/merge", body, queryParams);
        }

        /// <summary>
        /// Splits an existing plant group into multiple groups within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plant
        /// </summary>
        public async Task PlantsUpdateSplitV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/split", body, queryParams);
        }

        /// <summary>
        /// Updates the strain information for plants within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateStrainV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/strain", body, queryParams);
        }

        /// <summary>
        /// Replaces existing plant tags with new tags for plants within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateTagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/tag", body, queryParams);
        }

    }
}
