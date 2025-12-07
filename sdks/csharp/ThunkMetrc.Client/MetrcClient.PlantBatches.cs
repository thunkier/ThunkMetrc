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
        public async Task PlantBatchesCreateAdditivesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/additives", body, queryParams);
        }

        /// <summary>
        /// Records Additive usage details for plant batches at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantBatchesCreateAdditivesV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/additives", body, queryParams);
        }

        /// <summary>
        /// Records Additive usage for plant batches at a Facility using predefined additive templates.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantBatchesCreateAdditivesUsingtemplateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/additives/usingtemplate", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateAdjustV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/adjust", body, queryParams);
        }

        /// <summary>
        /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateAdjustV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/adjust", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateChangegrowthphaseV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/changegrowthphase", body, queryParams);
        }

        /// <summary>
        /// Updates the growth phase of plants at a specified Facility based on tracking information.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateGrowthphaseV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/growthphase", body, queryParams);
        }

        /// <summary>
        /// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageV2(object? body = null, string? isFromMotherPlant = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(isFromMotherPlant)) queryParams["isFromMotherPlant"] = isFromMotherPlant;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/packages", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageFrommotherplantV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/create/packages/frommotherplant", body, queryParams);
        }

        /// <summary>
        /// Creates packages from mother plants at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageFrommotherplantV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/packages/frommotherplant", body, queryParams);
        }

        /// <summary>
        /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreatePlantingsV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/plantings", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateSplitV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/split", body, queryParams);
        }

        /// <summary>
        /// Splits an existing Plant Batch into multiple groups at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateSplitV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/split", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantBatchesCreateWasteV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/waste", body, queryParams);
        }

        /// <summary>
        /// Records waste information for plant batches based on the submitted data for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantBatchesCreateWasteV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/waste", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatepackagesV1(object? body = null, string? isFromMotherPlant = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(isFromMotherPlant)) queryParams["isFromMotherPlant"] = isFromMotherPlant;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/createpackages", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateplantingsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v1/createplantings", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Destroy Immature Plants
        /// </summary>
        public async Task PlantBatchesDeleteV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plantbatches/v1", null, queryParams);
        }

        /// <summary>
        /// Completes the destruction of plant batches based on the provided input data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Destroy Immature Plants
        /// </summary>
        public async Task PlantBatchesDeleteV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plantbatches/v2", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Plant Batch by Id.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/active", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetTypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v1/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of plant batch types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves waste details associated with plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantBatchesGetWasteV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/waste", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetWasteReasonsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v1/waste/reasons", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetWasteReasonsV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/waste/reasons", null, queryParams);
        }

        /// <summary>
        /// Moves one or more plant batches to new locations with in a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants
        /// </summary>
        public async Task PlantBatchesUpdateLocationV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/location", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task PlantBatchesUpdateMoveplantbatchesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v1/moveplantbatches", body, queryParams);
        }

        /// <summary>
        /// Renames plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateNameV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/name", body, queryParams);
        }

        /// <summary>
        /// Changes the strain of plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateStrainV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/strain", body, queryParams);
        }

        /// <summary>
        /// Replaces tags for plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateTagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/tag", body, queryParams);
        }

    }
}
