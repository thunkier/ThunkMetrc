"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetrcClient = void 0;
const axios_1 = __importDefault(require("axios"));
class MetrcClient {
    constructor(baseUrl, vendorKey, userKey) {
        this.client = axios_1.default.create({
            baseURL: baseUrl,
            auth: {
                username: vendorKey,
                password: userKey
            }
        });
    }
    async locationsCreateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/create?${queryStr}` : `/locations/v1/create`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async locationsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async locationsCreateUpdateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/update?${queryStr}` : `/locations/v1/update`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async locationsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/${encodeURIComponent(id)}?${queryStr}` : `/locations/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async locationsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async locationsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/${encodeURIComponent(id)}?${queryStr}` : `/locations/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetActiveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/active?${queryStr}` : `/locations/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2/active?${queryStr}` : `/locations/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2/inactive?${queryStr}` : `/locations/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetTypesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v1/types?${queryStr}` : `/locations/v1/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsGetTypesV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2/types?${queryStr}` : `/locations/v2/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async locationsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async patientsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async patientsCreateAddV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/add?${queryStr}` : `/patients/v1/add`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async patientsCreateUpdateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/update?${queryStr}` : `/patients/v1/update`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async patientsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/${encodeURIComponent(id)}?${queryStr}` : `/patients/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async patientsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async patientsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/${encodeURIComponent(id)}?${queryStr}` : `/patients/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientsGetActiveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/active?${queryStr}` : `/patients/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientsGetActiveV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2/active?${queryStr}` : `/patients/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries?${queryStr}` : `/sales/v1/deliveries`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer?${queryStr}` : `/sales/v1/deliveries/retailer`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerDepartV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/depart?${queryStr}` : `/sales/v1/deliveries/retailer/depart`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerDepartV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/depart?${queryStr}` : `/sales/v2/deliveries/retailer/depart`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerEndV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/end?${queryStr}` : `/sales/v1/deliveries/retailer/end`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerEndV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/end?${queryStr}` : `/sales/v2/deliveries/retailer/end`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerRestockV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/restock?${queryStr}` : `/sales/v1/deliveries/retailer/restock`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerRestockV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/restock?${queryStr}` : `/sales/v2/deliveries/retailer/restock`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerSaleV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/sale?${queryStr}` : `/sales/v1/deliveries/retailer/sale`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateDeliveryRetailerSaleV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/sale?${queryStr}` : `/sales/v2/deliveries/retailer/sale`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateReceiptV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts?${queryStr}` : `/sales/v1/receipts`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateReceiptV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesCreateTransactionByDateV1(date, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(date)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(date)}`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async salesDeleteDeliveryV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesDeleteDeliveryV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesDeleteDeliveryRetailerV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesDeleteDeliveryRetailerV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesDeleteReceiptV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/receipts/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesDeleteReceiptV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async salesGetCountiesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/counties?${queryStr}` : `/sales/v1/counties`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetCountiesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/counties?${queryStr}` : `/sales/v2/counties`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetCustomertypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/customertypes?${queryStr}` : `/sales/v1/customertypes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetCustomertypesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/customertypes?${queryStr}` : `/sales/v2/customertypes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/active?${queryStr}` : `/sales/v1/deliveries/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/active?${queryStr}` : `/sales/v2/deliveries/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/inactive?${queryStr}` : `/sales/v1/deliveries/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/inactive?${queryStr}` : `/sales/v2/deliveries/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesRetailerActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/active?${queryStr}` : `/sales/v1/deliveries/retailer/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesRetailerActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/active?${queryStr}` : `/sales/v2/deliveries/retailer/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesRetailerInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/inactive?${queryStr}` : `/sales/v1/deliveries/retailer/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesRetailerInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/inactive?${queryStr}` : `/sales/v2/deliveries/retailer/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesReturnreasonsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/returnreasons?${queryStr}` : `/sales/v1/deliveries/returnreasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveriesReturnreasonsV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/returnreasons?${queryStr}` : `/sales/v2/deliveries/returnreasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveryV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveryV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveryRetailerV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetDeliveryRetailerV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetPatientRegistrationsLocationsV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/patientregistration/locations?${queryStr}` : `/sales/v1/patientregistration/locations`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetPatientRegistrationsLocationsV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/patientregistration/locations?${queryStr}` : `/sales/v2/patientregistration/locations`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetPaymenttypesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/paymenttypes?${queryStr}` : `/sales/v1/paymenttypes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetPaymenttypesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/paymenttypes?${queryStr}` : `/sales/v2/paymenttypes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/receipts/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptsActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts/active?${queryStr}` : `/sales/v1/receipts/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptsActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/active?${queryStr}` : `/sales/v2/receipts/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptsExternalByExternalNumberV2(externalNumber, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}?${queryStr}` : `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptsInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts/inactive?${queryStr}` : `/sales/v1/receipts/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetReceiptsInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        if (salesDateEnd)
            query.append('salesDateEnd', salesDateEnd);
        if (salesDateStart)
            query.append('salesDateStart', salesDateStart);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/inactive?${queryStr}` : `/sales/v2/receipts/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetTransactionsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/transactions?${queryStr}` : `/sales/v1/transactions`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart, salesDateEnd, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(salesDateStart)}/${encodeURIComponent(salesDateEnd)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(salesDateStart)}/${encodeURIComponent(salesDateEnd)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries?${queryStr}` : `/sales/v1/deliveries`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryCompleteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/complete?${queryStr}` : `/sales/v1/deliveries/complete`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryCompleteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/complete?${queryStr}` : `/sales/v2/deliveries/complete`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/hub?${queryStr}` : `/sales/v1/deliveries/hub`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/hub?${queryStr}` : `/sales/v2/deliveries/hub`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubAcceptV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/hub/accept?${queryStr}` : `/sales/v1/deliveries/hub/accept`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubAcceptV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/hub/accept?${queryStr}` : `/sales/v2/deliveries/hub/accept`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubDepartV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/hub/depart?${queryStr}` : `/sales/v1/deliveries/hub/depart`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubDepartV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/hub/depart?${queryStr}` : `/sales/v2/deliveries/hub/depart`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubVerifyIdV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/hub/verifyID?${queryStr}` : `/sales/v1/deliveries/hub/verifyID`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryHubVerifyIdV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/hub/verifyID?${queryStr}` : `/sales/v2/deliveries/hub/verifyID`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryRetailerV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/deliveries/retailer?${queryStr}` : `/sales/v1/deliveries/retailer`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateDeliveryRetailerV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateReceiptV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/receipts?${queryStr}` : `/sales/v1/receipts`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateReceiptV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateReceiptFinalizeV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/finalize?${queryStr}` : `/sales/v2/receipts/finalize`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateReceiptUnfinalizeV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v2/receipts/unfinalize?${queryStr}` : `/sales/v2/receipts/unfinalize`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async salesUpdateTransactionByDateV1(date, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(date)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(date)}`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async tagsGetPackageAvailableV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/tags/v2/package/available?${queryStr}` : `/tags/v2/package/available`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async tagsGetPlantAvailableV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/tags/v2/plant/available?${queryStr}` : `/tags/v2/plant/available`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async tagsGetStagedV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/tags/v2/staged?${queryStr}` : `/tags/v2/staged`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersCreateExternalIncomingV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/external/incoming?${queryStr}` : `/transfers/v1/external/incoming`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transfersCreateExternalIncomingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transfersCreateTemplatesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transfersCreateTemplatesOutgoingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transfersDeleteExternalIncomingV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/external/incoming/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v1/external/incoming/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transfersDeleteExternalIncomingV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/external/incoming/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/external/incoming/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transfersDeleteTemplatesV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v1/templates/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transfersDeleteTemplatesOutgoingV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transfersGetDeliveriesPackagesStatesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/packages/states?${queryStr}` : `/transfers/v1/deliveries/packages/states`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveriesPackagesStatesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/packages/states?${queryStr}` : `/transfers/v2/deliveries/packages/states`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v1/${encodeURIComponent(id)}/deliveries`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/${encodeURIComponent(id)}/deliveries`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches?${queryStr}` : `/transfers/v1/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches?${queryStr}` : `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageWholesaleV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages/wholesale?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages/wholesale`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryPackageWholesaleV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryTransportersV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryTransportersV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryTransportersDetailsV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters/details`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetDeliveryTransportersDetailsV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetHubV2(body, estimatedArrivalEnd, estimatedArrivalStart, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (estimatedArrivalEnd)
            query.append('estimatedArrivalEnd', estimatedArrivalEnd);
        if (estimatedArrivalStart)
            query.append('estimatedArrivalStart', estimatedArrivalStart);
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/hub?${queryStr}` : `/transfers/v2/hub`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetIncomingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/incoming?${queryStr}` : `/transfers/v1/incoming`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetIncomingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/incoming?${queryStr}` : `/transfers/v2/incoming`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetOutgoingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/outgoing?${queryStr}` : `/transfers/v1/outgoing`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/outgoing?${queryStr}` : `/transfers/v2/outgoing`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetRejectedV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/rejected?${queryStr}` : `/transfers/v1/rejected`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetRejectedV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/rejected?${queryStr}` : `/transfers/v2/rejected`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesDeliveryV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v1/templates/${encodeURIComponent(id)}/deliveries`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesDeliveryPackageV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/packages`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesDeliveryTransportersV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesDeliveryTransportersDetailsV1(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters/details`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesOutgoingDeliveryV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesOutgoingDeliveryPackageV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesOutgoingDeliveryTransportersV2(id, body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTypesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/types?${queryStr}` : `/transfers/v1/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersGetTypesV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/types?${queryStr}` : `/transfers/v2/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transfersUpdateExternalIncomingV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/external/incoming?${queryStr}` : `/transfers/v1/external/incoming`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async transfersUpdateExternalIncomingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async transfersUpdateTemplatesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async transfersUpdateTemplatesOutgoingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async wasteMethodsGetAllV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/wastemethods/v2?${queryStr}` : `/wastemethods/v2`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async additivesTemplatesCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async additivesTemplatesGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/additivestemplates/v2/${encodeURIComponent(id)}?${queryStr}` : `/additivestemplates/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async additivesTemplatesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/additivestemplates/v2/active?${queryStr}` : `/additivestemplates/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async additivesTemplatesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/additivestemplates/v2/inactive?${queryStr}` : `/additivestemplates/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async additivesTemplatesUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async facilitiesGetAllV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/facilities/v1?${queryStr}` : `/facilities/v1`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async facilitiesGetAllV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/facilities/v2?${queryStr}` : `/facilities/v2`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsCreateFinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/finish?${queryStr}` : `/harvests/v1/finish`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreatePackageV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/create/packages?${queryStr}` : `/harvests/v1/create/packages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreatePackageV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/packages?${queryStr}` : `/harvests/v2/packages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreatePackageTestingV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/create/packages/testing?${queryStr}` : `/harvests/v1/create/packages/testing`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreatePackageTestingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/packages/testing?${queryStr}` : `/harvests/v2/packages/testing`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreateRemoveWasteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/removewaste?${queryStr}` : `/harvests/v1/removewaste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreateUnfinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/unfinish?${queryStr}` : `/harvests/v1/unfinish`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsCreateWasteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async harvestsDeleteWasteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/waste/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/waste/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async harvestsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/active?${queryStr}` : `/harvests/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/active?${queryStr}` : `/harvests/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/inactive?${queryStr}` : `/harvests/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/inactive?${queryStr}` : `/harvests/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/onhold?${queryStr}` : `/harvests/v1/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/onhold?${queryStr}` : `/harvests/v2/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetWasteV2(body, harvestId, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (harvestId)
            query.append('harvestId', harvestId);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetWasteTypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/waste/types?${queryStr}` : `/harvests/v1/waste/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsGetWasteTypesV2(body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/waste/types?${queryStr}` : `/harvests/v2/waste/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async harvestsUpdateFinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/finish?${queryStr}` : `/harvests/v2/finish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateLocationV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/location?${queryStr}` : `/harvests/v2/location`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateMoveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/move?${queryStr}` : `/harvests/v1/move`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateRenameV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v1/rename?${queryStr}` : `/harvests/v1/rename`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateRenameV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/rename?${queryStr}` : `/harvests/v2/rename`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateRestoreHarvestedPlantsV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/restore/harvestedplants?${queryStr}` : `/harvests/v2/restore/harvestedplants`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async harvestsUpdateUnfinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/harvests/v2/unfinish?${queryStr}` : `/harvests/v2/unfinish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async itemsCreateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/create?${queryStr}` : `/items/v1/create`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreateBrandV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreateFileV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/file?${queryStr}` : `/items/v2/file`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreatePhotoV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/photo?${queryStr}` : `/items/v1/photo`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreatePhotoV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/photo?${queryStr}` : `/items/v2/photo`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsCreateUpdateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/update?${queryStr}` : `/items/v1/update`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async itemsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async itemsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async itemsDeleteBrandV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/brand/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/brand/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async itemsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetActiveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/active?${queryStr}` : `/items/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/active?${queryStr}` : `/items/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetBrandsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/brands?${queryStr}` : `/items/v1/brands`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetBrandsV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/brands?${queryStr}` : `/items/v2/brands`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetCategoriesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/categories?${queryStr}` : `/items/v1/categories`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetCategoriesV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/categories?${queryStr}` : `/items/v2/categories`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetFileV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/file/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/file/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetInactiveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/inactive?${queryStr}` : `/items/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/inactive?${queryStr}` : `/items/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetPhotoV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v1/photo/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/photo/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsGetPhotoV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/photo/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/photo/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async itemsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async itemsUpdateBrandV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async labTestsCreateRecordV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/record?${queryStr}` : `/labtests/v1/record`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async labTestsCreateRecordV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/record?${queryStr}` : `/labtests/v2/record`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async labTestsGetBatchesV2(body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/batches?${queryStr}` : `/labtests/v2/batches`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetLabtestdocumentV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/labtestdocument/${encodeURIComponent(id)}?${queryStr}` : `/labtests/v1/labtestdocument/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetLabtestdocumentV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/labtestdocument/${encodeURIComponent(id)}?${queryStr}` : `/labtests/v2/labtestdocument/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetResultsV1(body, licenseNumber, packageId) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (packageId)
            query.append('packageId', packageId);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/results?${queryStr}` : `/labtests/v1/results`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetResultsV2(body, licenseNumber, packageId, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (packageId)
            query.append('packageId', packageId);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/results?${queryStr}` : `/labtests/v2/results`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetStatesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/states?${queryStr}` : `/labtests/v1/states`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetStatesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/states?${queryStr}` : `/labtests/v2/states`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetTypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/types?${queryStr}` : `/labtests/v1/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsGetTypesV2(body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/types?${queryStr}` : `/labtests/v2/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async labTestsUpdateLabtestdocumentV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/labtestdocument?${queryStr}` : `/labtests/v1/labtestdocument`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async labTestsUpdateLabtestdocumentV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/labtestdocument?${queryStr}` : `/labtests/v2/labtestdocument`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async labTestsUpdateResultReleaseV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v1/results/release?${queryStr}` : `/labtests/v1/results/release`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async labTestsUpdateResultReleaseV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/labtests/v2/results/release?${queryStr}` : `/labtests/v2/results/release`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async strainsCreateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v1/create?${queryStr}` : `/strains/v1/create`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async strainsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async strainsCreateUpdateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v1/update?${queryStr}` : `/strains/v1/update`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async strainsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v1/${encodeURIComponent(id)}?${queryStr}` : `/strains/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async strainsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async strainsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v1/${encodeURIComponent(id)}?${queryStr}` : `/strains/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async strainsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async strainsGetActiveV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v1/active?${queryStr}` : `/strains/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async strainsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2/active?${queryStr}` : `/strains/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async strainsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2/inactive?${queryStr}` : `/strains/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async strainsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async transportersCreateDriverV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transportersCreateVehicleV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async transportersDeleteDriverV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transportersDeleteVehicleV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async transportersGetDriverV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transportersGetDriversV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transportersGetVehicleV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transportersGetVehiclesV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async transportersUpdateDriverV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async transportersUpdateVehicleV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesCreateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/create?${queryStr}` : `/packages/v1/create`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2?${queryStr}` : `/packages/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateAdjustV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/adjust?${queryStr}` : `/packages/v1/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateAdjustV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateChangeItemV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/change/item?${queryStr}` : `/packages/v1/change/item`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateChangeLocationV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/change/locations?${queryStr}` : `/packages/v1/change/locations`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateFinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/finish?${queryStr}` : `/packages/v1/finish`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreatePlantingsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/create/plantings?${queryStr}` : `/packages/v1/create/plantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreatePlantingsV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/plantings?${queryStr}` : `/packages/v2/plantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateRemediateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/remediate?${queryStr}` : `/packages/v1/remediate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateTestingV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/create/testing?${queryStr}` : `/packages/v1/create/testing`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateTestingV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/testing?${queryStr}` : `/packages/v2/testing`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesCreateUnfinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/unfinish?${queryStr}` : `/packages/v1/unfinish`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async packagesDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async packagesGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/${encodeURIComponent(id)}?${queryStr}` : `/packages/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/active?${queryStr}` : `/packages/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/active?${queryStr}` : `/packages/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetAdjustReasonsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/adjust/reasons?${queryStr}` : `/packages/v1/adjust/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetAdjustReasonsV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/adjust/reasons?${queryStr}` : `/packages/v2/adjust/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetByLabelV1(label, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/${encodeURIComponent(label)}?${queryStr}` : `/packages/v1/${encodeURIComponent(label)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetByLabelV2(label, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(label)}?${queryStr}` : `/packages/v2/${encodeURIComponent(label)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/inactive?${queryStr}` : `/packages/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/inactive?${queryStr}` : `/packages/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetIntransitV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/intransit?${queryStr}` : `/packages/v2/intransit`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetLabsamplesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/labsamples?${queryStr}` : `/packages/v2/labsamples`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/onhold?${queryStr}` : `/packages/v1/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/onhold?${queryStr}` : `/packages/v2/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetSourceHarvestV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}/source/harvests?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}/source/harvests`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetTransferredV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/transferred?${queryStr}` : `/packages/v2/transferred`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetTypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/types?${queryStr}` : `/packages/v1/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesGetTypesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/types?${queryStr}` : `/packages/v2/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async packagesUpdateAdjustV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateChangeNoteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v1/change/note?${queryStr}` : `/packages/v1/change/note`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateDecontaminateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/decontaminate?${queryStr}` : `/packages/v2/decontaminate`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateDonationFlagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/donation/flag?${queryStr}` : `/packages/v2/donation/flag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateDonationUnflagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/donation/unflag?${queryStr}` : `/packages/v2/donation/unflag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateExternalidV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/externalid?${queryStr}` : `/packages/v2/externalid`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateFinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/finish?${queryStr}` : `/packages/v2/finish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateItemV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/item?${queryStr}` : `/packages/v2/item`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateLabTestRequiredV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/labtests/required?${queryStr}` : `/packages/v2/labtests/required`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateLocationV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/location?${queryStr}` : `/packages/v2/location`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateNoteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/note?${queryStr}` : `/packages/v2/note`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateRemediateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/remediate?${queryStr}` : `/packages/v2/remediate`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateTradesampleFlagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/tradesample/flag?${queryStr}` : `/packages/v2/tradesample/flag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateTradesampleUnflagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/tradesample/unflag?${queryStr}` : `/packages/v2/tradesample/unflag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateUnfinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/unfinish?${queryStr}` : `/packages/v2/unfinish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async packagesUpdateUsebydateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/packages/v2/usebydate?${queryStr}` : `/packages/v2/usebydate`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async sublocationsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async sublocationsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async sublocationsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async sublocationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2/active?${queryStr}` : `/sublocations/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async sublocationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2/inactive?${queryStr}` : `/sublocations/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async sublocationsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async unitsOfMeasureGetActiveV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/unitsofmeasure/v1/active?${queryStr}` : `/unitsofmeasure/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async unitsOfMeasureGetActiveV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/unitsofmeasure/v2/active?${queryStr}` : `/unitsofmeasure/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async unitsOfMeasureGetInactiveV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/unitsofmeasure/v2/inactive?${queryStr}` : `/unitsofmeasure/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesCreateAdditivesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/additives?${queryStr}` : `/plantbatches/v1/additives`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateAdditivesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/additives?${queryStr}` : `/plantbatches/v2/additives`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateAdditivesUsingtemplateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/additives/usingtemplate?${queryStr}` : `/plantbatches/v2/additives/usingtemplate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateAdjustV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/adjust?${queryStr}` : `/plantbatches/v1/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateAdjustV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/adjust?${queryStr}` : `/plantbatches/v2/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateChangegrowthphaseV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/changegrowthphase?${queryStr}` : `/plantbatches/v1/changegrowthphase`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateGrowthphaseV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/growthphase?${queryStr}` : `/plantbatches/v2/growthphase`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreatePackageV2(body, isFromMotherPlant, licenseNumber) {
        const query = new URLSearchParams();
        if (isFromMotherPlant)
            query.append('isFromMotherPlant', isFromMotherPlant);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/packages?${queryStr}` : `/plantbatches/v2/packages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreatePackageFrommotherplantV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/create/packages/frommotherplant?${queryStr}` : `/plantbatches/v1/create/packages/frommotherplant`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreatePackageFrommotherplantV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/packages/frommotherplant?${queryStr}` : `/plantbatches/v2/packages/frommotherplant`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreatePlantingsV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/plantings?${queryStr}` : `/plantbatches/v2/plantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateSplitV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/split?${queryStr}` : `/plantbatches/v1/split`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateSplitV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/split?${queryStr}` : `/plantbatches/v2/split`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateWasteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/waste?${queryStr}` : `/plantbatches/v1/waste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateWasteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreatepackagesV1(body, isFromMotherPlant, licenseNumber) {
        const query = new URLSearchParams();
        if (isFromMotherPlant)
            query.append('isFromMotherPlant', isFromMotherPlant);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/createpackages?${queryStr}` : `/plantbatches/v1/createpackages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesCreateplantingsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/createplantings?${queryStr}` : `/plantbatches/v1/createplantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantBatchesDeleteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1?${queryStr}` : `/plantbatches/v1`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async plantBatchesDeleteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2?${queryStr}` : `/plantbatches/v2`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async plantBatchesGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/${encodeURIComponent(id)}?${queryStr}` : `/plantbatches/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/${encodeURIComponent(id)}?${queryStr}` : `/plantbatches/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/active?${queryStr}` : `/plantbatches/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/active?${queryStr}` : `/plantbatches/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/inactive?${queryStr}` : `/plantbatches/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/inactive?${queryStr}` : `/plantbatches/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetTypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/types?${queryStr}` : `/plantbatches/v1/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetTypesV2(body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/types?${queryStr}` : `/plantbatches/v2/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetWasteV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetWasteReasonsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/waste/reasons?${queryStr}` : `/plantbatches/v1/waste/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesGetWasteReasonsV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/waste/reasons?${queryStr}` : `/plantbatches/v2/waste/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantBatchesUpdateLocationV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/location?${queryStr}` : `/plantbatches/v2/location`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantBatchesUpdateMoveplantbatchesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v1/moveplantbatches?${queryStr}` : `/plantbatches/v1/moveplantbatches`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantBatchesUpdateNameV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/name?${queryStr}` : `/plantbatches/v2/name`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantBatchesUpdateStrainV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/strain?${queryStr}` : `/plantbatches/v2/strain`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantBatchesUpdateTagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plantbatches/v2/tag?${queryStr}` : `/plantbatches/v2/tag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/additives?${queryStr}` : `/plants/v1/additives`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesBylocationV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/additives/bylocation?${queryStr}` : `/plants/v1/additives/bylocation`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesBylocationV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives/bylocation?${queryStr}` : `/plants/v2/additives/bylocation`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesBylocationUsingtemplateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives/bylocation/usingtemplate?${queryStr}` : `/plants/v2/additives/bylocation/usingtemplate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateAdditivesUsingtemplateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives/usingtemplate?${queryStr}` : `/plants/v2/additives/usingtemplate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateChangegrowthphasesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/changegrowthphases?${queryStr}` : `/plants/v1/changegrowthphases`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateHarvestplantsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/harvestplants?${queryStr}` : `/plants/v1/harvestplants`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateManicureV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/manicure?${queryStr}` : `/plants/v2/manicure`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateManicureplantsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/manicureplants?${queryStr}` : `/plants/v1/manicureplants`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateMoveplantsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/moveplants?${queryStr}` : `/plants/v1/moveplants`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreatePlantbatchPackageV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/create/plantbatch/packages?${queryStr}` : `/plants/v1/create/plantbatch/packages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreatePlantbatchPackageV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/plantbatch/packages?${queryStr}` : `/plants/v2/plantbatch/packages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreatePlantingsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/create/plantings?${queryStr}` : `/plants/v1/create/plantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreatePlantingsV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/plantings?${queryStr}` : `/plants/v2/plantings`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateWasteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/waste?${queryStr}` : `/plants/v1/waste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsCreateWasteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async plantsDeleteV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1?${queryStr}` : `/plants/v1`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async plantsDeleteV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2?${queryStr}` : `/plants/v2`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async plantsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/${encodeURIComponent(id)}?${queryStr}` : `/plants/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(id)}?${queryStr}` : `/plants/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetAdditivesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/additives?${queryStr}` : `/plants/v1/additives`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetAdditivesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetAdditivesTypesV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/additives/types?${queryStr}` : `/plants/v1/additives/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetAdditivesTypesV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/additives/types?${queryStr}` : `/plants/v2/additives/types`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetByLabelV1(label, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/${encodeURIComponent(label)}?${queryStr}` : `/plants/v1/${encodeURIComponent(label)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetByLabelV2(label, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(label)}?${queryStr}` : `/plants/v2/${encodeURIComponent(label)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetFloweringV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/flowering?${queryStr}` : `/plants/v1/flowering`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetFloweringV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/flowering?${queryStr}` : `/plants/v2/flowering`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetGrowthPhasesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/growthphases?${queryStr}` : `/plants/v1/growthphases`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetGrowthPhasesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/growthphases?${queryStr}` : `/plants/v2/growthphases`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/inactive?${queryStr}` : `/plants/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/inactive?${queryStr}` : `/plants/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetMotherV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/mother?${queryStr}` : `/plants/v2/mother`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetMotherInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/mother/inactive?${queryStr}` : `/plants/v2/mother/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetMotherOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/mother/onhold?${queryStr}` : `/plants/v2/mother/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/onhold?${queryStr}` : `/plants/v1/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/onhold?${queryStr}` : `/plants/v2/onhold`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetVegetativeV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/vegetative?${queryStr}` : `/plants/v1/vegetative`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetVegetativeV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/vegetative?${queryStr}` : `/plants/v2/vegetative`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWasteV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWasteMethodsAllV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/waste/methods/all?${queryStr}` : `/plants/v1/waste/methods/all`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWasteMethodsAllV2(body, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste/methods/all?${queryStr}` : `/plants/v2/waste/methods/all`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWastePackageV2(id, body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/package?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/package`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWastePlantV2(id, body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/plant?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/plant`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWasteReasonsV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v1/waste/reasons?${queryStr}` : `/plants/v1/waste/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsGetWasteReasonsV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/waste/reasons?${queryStr}` : `/plants/v2/waste/reasons`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async plantsUpdateAdjustV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/adjust?${queryStr}` : `/plants/v2/adjust`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateGrowthphaseV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/growthphase?${queryStr}` : `/plants/v2/growthphase`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateHarvestV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/harvest?${queryStr}` : `/plants/v2/harvest`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateLocationV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/location?${queryStr}` : `/plants/v2/location`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateMergeV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/merge?${queryStr}` : `/plants/v2/merge`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateSplitV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/split?${queryStr}` : `/plants/v2/split`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateStrainV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/strain?${queryStr}` : `/plants/v2/strain`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async plantsUpdateTagV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/plants/v2/tag?${queryStr}` : `/plants/v2/tag`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async employeesGetAllV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/employees/v1?${queryStr}` : `/employees/v1`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async employeesGetAllV2(body, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/employees/v2?${queryStr}` : `/employees/v2`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async employeesGetPermissionsV2(body, employeeLicenseNumber, licenseNumber) {
        const query = new URLSearchParams();
        if (employeeLicenseNumber)
            query.append('employeeLicenseNumber', employeeLicenseNumber);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/employees/v2/permissions?${queryStr}` : `/employees/v2/permissions`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientCheckInsCreateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async patientCheckInsCreateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async patientCheckInsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v1/${encodeURIComponent(id)}?${queryStr}` : `/patient-checkins/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async patientCheckInsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v2/${encodeURIComponent(id)}?${queryStr}` : `/patient-checkins/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async patientCheckInsGetAllV1(body, checkinDateEnd, checkinDateStart, licenseNumber) {
        const query = new URLSearchParams();
        if (checkinDateEnd)
            query.append('checkinDateEnd', checkinDateEnd);
        if (checkinDateStart)
            query.append('checkinDateStart', checkinDateStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientCheckInsGetAllV2(body, checkinDateEnd, checkinDateStart, licenseNumber) {
        const query = new URLSearchParams();
        if (checkinDateEnd)
            query.append('checkinDateEnd', checkinDateEnd);
        if (checkinDateStart)
            query.append('checkinDateStart', checkinDateStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientCheckInsGetLocationsV1(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v1/locations?${queryStr}` : `/patient-checkins/v1/locations`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientCheckInsGetLocationsV2(body, No) {
        const query = new URLSearchParams();
        if (No)
            query.append('No', No);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v2/locations?${queryStr}` : `/patient-checkins/v2/locations`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientCheckInsUpdateV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async patientCheckInsUpdateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v1/statuses/${encodeURIComponent(patientLicenseNumber)}?${queryStr}` : `/patients/v1/statuses/${encodeURIComponent(patientLicenseNumber)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}?${queryStr}` : `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsCreateAdjustV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/adjust?${queryStr}` : `/processing/v1/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreateAdjustV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/adjust?${queryStr}` : `/processing/v2/adjust`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreateJobtypesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes?${queryStr}` : `/processing/v1/jobtypes`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreateJobtypesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreateStartV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/start?${queryStr}` : `/processing/v1/start`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreateStartV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/start?${queryStr}` : `/processing/v2/start`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreatepackagesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/createpackages?${queryStr}` : `/processing/v1/createpackages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsCreatepackagesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/createpackages?${queryStr}` : `/processing/v2/createpackages`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async processingJobsDeleteV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async processingJobsDeleteV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async processingJobsDeleteJobtypesV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/jobtypes/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async processingJobsDeleteJobtypesV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/jobtypes/${encodeURIComponent(id)}`;
        const { data } = await this.client.delete(fullUrl, body);
        return data;
    }
    async processingJobsGetV1(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetV2(id, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/active?${queryStr}` : `/processing/v1/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/active?${queryStr}` : `/processing/v2/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/inactive?${queryStr}` : `/processing/v1/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/inactive?${queryStr}` : `/processing/v2/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes/active?${queryStr}` : `/processing/v1/jobtypes/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes/active?${queryStr}` : `/processing/v2/jobtypes/active`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesAttributesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes/attributes?${queryStr}` : `/processing/v1/jobtypes/attributes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesAttributesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes/attributes?${queryStr}` : `/processing/v2/jobtypes/attributes`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesCategoriesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes/categories?${queryStr}` : `/processing/v1/jobtypes/categories`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesCategoriesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes/categories?${queryStr}` : `/processing/v2/jobtypes/categories`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes/inactive?${queryStr}` : `/processing/v1/jobtypes/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsGetJobtypesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        const query = new URLSearchParams();
        if (lastModifiedEnd)
            query.append('lastModifiedEnd', lastModifiedEnd);
        if (lastModifiedStart)
            query.append('lastModifiedStart', lastModifiedStart);
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        if (pageNumber)
            query.append('pageNumber', pageNumber);
        if (pageSize)
            query.append('pageSize', pageSize);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes/inactive?${queryStr}` : `/processing/v2/jobtypes/inactive`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async processingJobsUpdateFinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/finish?${queryStr}` : `/processing/v1/finish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async processingJobsUpdateFinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/finish?${queryStr}` : `/processing/v2/finish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async processingJobsUpdateJobtypesV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/jobtypes?${queryStr}` : `/processing/v1/jobtypes`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async processingJobsUpdateJobtypesV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async processingJobsUpdateUnfinishV1(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v1/unfinish?${queryStr}` : `/processing/v1/unfinish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async processingJobsUpdateUnfinishV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/processing/v2/unfinish?${queryStr}` : `/processing/v2/unfinish`;
        const { data } = await this.client.put(fullUrl, body);
        return data;
    }
    async retailIdCreateAssociateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/associate?${queryStr}` : `/retailid/v2/associate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async retailIdCreateGenerateV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/generate?${queryStr}` : `/retailid/v2/generate`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async retailIdCreateMergeV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/merge?${queryStr}` : `/retailid/v2/merge`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async retailIdCreatePackageInfoV2(body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/packages/info?${queryStr}` : `/retailid/v2/packages/info`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async retailIdGetReceiveByLabelV2(label, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/receive/${encodeURIComponent(label)}?${queryStr}` : `/retailid/v2/receive/${encodeURIComponent(label)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async retailIdGetReceiveQrByShortCodeV2(shortCode, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}?${queryStr}` : `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async sandboxCreateIntegratorSetupV2(body, userKey) {
        const query = new URLSearchParams();
        if (userKey)
            query.append('userKey', userKey);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/sandbox/v2/integrator/setup?${queryStr}` : `/sandbox/v2/integrator/setup`;
        const { data } = await this.client.post(fullUrl, body);
        return data;
    }
    async caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/caregivers/v1/status/${encodeURIComponent(caregiverLicenseNumber)}?${queryStr}` : `/caregivers/v1/status/${encodeURIComponent(caregiverLicenseNumber)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
    async caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber, body, licenseNumber) {
        const query = new URLSearchParams();
        if (licenseNumber)
            query.append('licenseNumber', licenseNumber);
        const queryStr = query.toString();
        const fullUrl = queryStr ? `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}?${queryStr}` : `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}`;
        const { data } = await this.client.get(fullUrl, body);
        return data;
    }
}
exports.MetrcClient = MetrcClient;
// --- Request Models ---
