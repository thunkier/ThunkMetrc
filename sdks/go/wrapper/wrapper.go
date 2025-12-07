package wrapper

import (
	"context"
	"github.com/thunkmetrc/sdks/go/client"
)

type MetrcWrapper struct {
	Client *client.MetrcClient
	RateLimiter *RateLimiter
}

func New(client *client.MetrcClient) *MetrcWrapper {
	return &MetrcWrapper{
		Client: client,
		RateLimiter: NewRateLimiter(DefaultConfig()),
	}
}

func NewClient(baseUrl, vendorKey, userKey string) *MetrcWrapper {
    return New(client.New(baseUrl, vendorKey, userKey))
}

func NewWithConfig(client *client.MetrcClient, config RateLimiterConfig) *MetrcWrapper {
	return &MetrcWrapper{
		Client: client,
		RateLimiter: NewRateLimiter(config),
	}
}

// POST CreateAdjust V1
// Permissions Required:
//   - ManageProcessingJobs
func (w *MetrcWrapper) ProcessingJobsCreateAdjustV1(licensenumber string, body []ProcessingJobsCreateAdjustV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateAdjustV1(licensenumber, body)
	})
}

// POST CreateAdjust V2
// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsCreateAdjustV2(licensenumber string, body []ProcessingJobsCreateAdjustV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateAdjustV2(licensenumber, body)
	})
}

// POST CreateJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsCreateJobtypesV1(licensenumber string, body []ProcessingJobsCreateJobtypesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateJobtypesV1(licensenumber, body)
	})
}

// POST CreateJobtypes V2
// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsCreateJobtypesV2(licensenumber string, body []ProcessingJobsCreateJobtypesV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateJobtypesV2(licensenumber, body)
	})
}

// POST CreateStart V1
// Permissions Required:
//   - ManageProcessingJobs
func (w *MetrcWrapper) ProcessingJobsCreateStartV1(licensenumber string, body []ProcessingJobsCreateStartV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateStartV1(licensenumber, body)
	})
}

// POST CreateStart V2
// Initiates new processing jobs at a Facility, including job details and associated packages.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsCreateStartV2(licensenumber string, body []ProcessingJobsCreateStartV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreateStartV2(licensenumber, body)
	})
}

// POST Createpackages V1
// Permissions Required:
//   - ManageProcessingJobs
func (w *MetrcWrapper) ProcessingJobsCreatepackagesV1(licensenumber string, body []ProcessingJobsCreatepackagesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreatepackagesV1(licensenumber, body)
	})
}

// POST Createpackages V2
// Creates packages from processing jobs at a Facility, including optional location and note assignments.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsCreatepackagesV2(licensenumber string, body []ProcessingJobsCreatepackagesV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsCreatepackagesV2(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsDeleteV2(id, licensenumber)
	})
}

// DELETE DeleteJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsDeleteJobtypesV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsDeleteJobtypesV1(id, licensenumber)
	})
}

// DELETE DeleteJobtypes V2
// Archives a Processing Job Type at a Facility, making it inactive for future use.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsDeleteJobtypesV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsDeleteJobtypesV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a ProcessingJob by Id.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetActive V2
// Retrieves active processing jobs at a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetInactive V2
// Retrieves inactive processing jobs at a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetJobtypesActive V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetJobtypesActive V2
// Retrieves a list of all active processing job types defined within a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetJobtypesAttributes V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesAttributesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesAttributesV1(licensenumber)
	})
}

// GET GetJobtypesAttributes V2
// Retrieves all processing job attributes available for a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesAttributesV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesAttributesV2(licensenumber)
	})
}

// GET GetJobtypesCategories V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesCategoriesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesCategoriesV1(licensenumber)
	})
}

// GET GetJobtypesCategories V2
// Retrieves all processing job categories available for a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesCategoriesV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesCategoriesV2(licensenumber)
	})
}

// GET GetJobtypesInactive V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetJobtypesInactive V2
// Retrieves a list of all inactive processing job types defined within a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsGetJobtypesInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ProcessingJobsGetJobtypesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// PUT UpdateFinish V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateFinishV1(licensenumber string, body []ProcessingJobsUpdateFinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateFinishV1(licensenumber, body)
	})
}

// PUT UpdateFinish V2
// Completes processing jobs at a Facility by recording final notes and waste measurements.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateFinishV2(licensenumber string, body []ProcessingJobsUpdateFinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateFinishV2(licensenumber, body)
	})
}

// PUT UpdateJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateJobtypesV1(licensenumber string, body []ProcessingJobsUpdateJobtypesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateJobtypesV1(licensenumber, body)
	})
}

// PUT UpdateJobtypes V2
// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateJobtypesV2(licensenumber string, body []ProcessingJobsUpdateJobtypesV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateJobtypesV2(licensenumber, body)
	})
}

// PUT UpdateUnfinish V1
// Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateUnfinishV1(licensenumber string, body []ProcessingJobsUpdateUnfinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateUnfinishV1(licensenumber, body)
	})
}

// PUT UpdateUnfinish V2
// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
// 
//   Permissions Required:
//   - Manage Processing Job
func (w *MetrcWrapper) ProcessingJobsUpdateUnfinishV2(licensenumber string, body []ProcessingJobsUpdateUnfinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ProcessingJobsUpdateUnfinishV2(licensenumber, body)
	})
}

// POST Create V2
// Creates new sublocation records for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsCreateV2(licensenumber string, body []SublocationsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SublocationsCreateV2(licensenumber, body)
	})
}

// DELETE Delete V2
// Archives an existing Sublocation record for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SublocationsDeleteV2(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Sublocation by its Id, with an optional license number.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SublocationsGetV2(id, licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SublocationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive sublocations for the specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SublocationsGetInactiveV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT Update V2
// Updates existing sublocation records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) SublocationsUpdateV2(licensenumber string, body []SublocationsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SublocationsUpdateV2(licensenumber, body)
	})
}

// POST CreateDelivery V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesCreateDeliveryV1(licensenumber string, body []SalesCreateDeliveryV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryV1(licensenumber, body)
	})
}

// POST CreateDelivery V2
// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (w *MetrcWrapper) SalesCreateDeliveryV2(licensenumber string, body []SalesCreateDeliveryV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryV2(licensenumber, body)
	})
}

// POST CreateDeliveryRetailer V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerV1(licensenumber string, body []SalesCreateDeliveryRetailerV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerV1(licensenumber, body)
	})
}

// POST CreateDeliveryRetailer V2
// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerV2(licensenumber string, body []SalesCreateDeliveryRetailerV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerV2(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerDepart V1
// Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerDepartV1(licensenumber string, body []SalesCreateDeliveryRetailerDepartV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerDepartV1(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerDepart V2
// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerDepartV2(licensenumber string, body []SalesCreateDeliveryRetailerDepartV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerDepartV2(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerEnd V1
// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerEndV1(licensenumber string, body []SalesCreateDeliveryRetailerEndV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerEndV1(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerEnd V2
// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerEndV2(licensenumber string, body []SalesCreateDeliveryRetailerEndV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerEndV2(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerRestock V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerRestockV1(licensenumber string, body []SalesCreateDeliveryRetailerRestockV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerRestockV1(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerRestock V2
// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerRestockV2(licensenumber string, body []SalesCreateDeliveryRetailerRestockV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerRestockV2(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerSale V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesCreateDeliveryRetailerSaleV1(licensenumber string, body []SalesCreateDeliveryRetailerSaleV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerSaleV1(licensenumber, body)
	})
}

// POST CreateDeliveryRetailerSale V2
// Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (w *MetrcWrapper) SalesCreateDeliveryRetailerSaleV2(licensenumber string, body []SalesCreateDeliveryRetailerSaleV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateDeliveryRetailerSaleV2(licensenumber, body)
	})
}

// POST CreateReceipt V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesCreateReceiptV1(licensenumber string, body []SalesCreateReceiptV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateReceiptV1(licensenumber, body)
	})
}

// POST CreateReceipt V2
// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales (Write)
//   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
//   - Industry/Facility Type/Advanced Sales
//   - WebApi Sales Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (w *MetrcWrapper) SalesCreateReceiptV2(licensenumber string, body []SalesCreateReceiptV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateReceiptV2(licensenumber, body)
	})
}

// POST CreateTransactionByDate V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesCreateTransactionByDateV1(date string, licensenumber string, body []SalesCreateTransactionByDateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesCreateTransactionByDateV1(date, licensenumber, body)
	})
}

// DELETE DeleteDelivery V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesDeleteDeliveryV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteDeliveryV1(id, licensenumber)
	})
}

// DELETE DeleteDelivery V2
// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesDeleteDeliveryV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteDeliveryV2(id, licensenumber)
	})
}

// DELETE DeleteDeliveryRetailer V1
// Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesDeleteDeliveryRetailerV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteDeliveryRetailerV1(id, licensenumber)
	})
}

// DELETE DeleteDeliveryRetailer V2
// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesDeleteDeliveryRetailerV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteDeliveryRetailerV2(id, licensenumber)
	})
}

// DELETE DeleteReceipt V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesDeleteReceiptV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteReceiptV1(id, licensenumber)
	})
}

// DELETE DeleteReceipt V2
// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
// 
//   Permissions Required:
//   - Manage Sales
func (w *MetrcWrapper) SalesDeleteReceiptV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesDeleteReceiptV2(id, licensenumber)
	})
}

// GET GetCounties V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) SalesGetCountiesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetCountiesV1(no)
	})
}

// GET GetCounties V2
// Returns a list of counties available for sales deliveries.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) SalesGetCountiesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetCountiesV2(no)
	})
}

// GET GetCustomertypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) SalesGetCustomertypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetCustomertypesV1(no)
	})
}

// GET GetCustomertypes V2
// Returns a list of customer types.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) SalesGetCustomertypesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetCustomertypesV2(no)
	})
}

// GET GetDeliveriesActive V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveriesActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart)
	})
}

// GET GetDeliveriesActive V2
// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveriesActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart)
	})
}

// GET GetDeliveriesInactive V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveriesInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart)
	})
}

// GET GetDeliveriesInactive V2
// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveriesInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart)
	})
}

// GET GetDeliveriesRetailerActive V1
// Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveriesRetailerActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesRetailerActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetDeliveriesRetailerActive V2
// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveriesRetailerActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesRetailerActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetDeliveriesRetailerInactive V1
// Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveriesRetailerInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesRetailerInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetDeliveriesRetailerInactive V2
// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveriesRetailerInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesRetailerInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetDeliveriesReturnreasons V1
// Permissions Required:
//   -
func (w *MetrcWrapper) SalesGetDeliveriesReturnreasonsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesReturnreasonsV1(licensenumber)
	})
}

// GET GetDeliveriesReturnreasons V2
// Returns a list of return reasons for sales deliveries based on the provided License Number.
// 
//   Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveriesReturnreasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveriesReturnreasonsV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetDelivery V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveryV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveryV1(id, licensenumber)
	})
}

// GET GetDelivery V2
// Retrieves a sales delivery record by its Id, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesGetDeliveryV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveryV2(id, licensenumber)
	})
}

// GET GetDeliveryRetailer V1
// Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveryRetailerV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveryRetailerV1(id, licensenumber)
	})
}

// GET GetDeliveryRetailer V2
// Retrieves a retailer delivery record by its ID, with an optional License Number.
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesGetDeliveryRetailerV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetDeliveryRetailerV2(id, licensenumber)
	})
}

// GET GetPatientRegistrationsLocations V1
// Permissions Required:
//   -
func (w *MetrcWrapper) SalesGetPatientRegistrationsLocationsV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetPatientRegistrationsLocationsV1(no)
	})
}

// GET GetPatientRegistrationsLocations V2
// Returns a list of valid Patient registration locations for sales.
// 
//   Permissions Required:
//   -
func (w *MetrcWrapper) SalesGetPatientRegistrationsLocationsV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetPatientRegistrationsLocationsV2(no)
	})
}

// GET GetPaymenttypes V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesGetPaymenttypesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetPaymenttypesV1(licensenumber)
	})
}

// GET GetPaymenttypes V2
// Returns a list of available payment types for the specified License Number.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesGetPaymenttypesV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetPaymenttypesV2(licensenumber)
	})
}

// GET GetReceipt V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesGetReceiptV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptV1(id, licensenumber)
	})
}

// GET GetReceipt V2
// Retrieves a sales receipt by its Id, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (w *MetrcWrapper) SalesGetReceiptV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptV2(id, licensenumber)
	})
}

// GET GetReceiptsActive V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesGetReceiptsActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptsActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart)
	})
}

// GET GetReceiptsActive V2
// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (w *MetrcWrapper) SalesGetReceiptsActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptsActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart)
	})
}

// GET GetReceiptsExternalByExternalNumber V2
// Retrieves a Sales Receipt by its external number, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (w *MetrcWrapper) SalesGetReceiptsExternalByExternalNumberV2(externalnumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptsExternalByExternalNumberV2(externalnumber, licensenumber)
	})
}

// GET GetReceiptsInactive V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesGetReceiptsInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptsInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart)
	})
}

// GET GetReceiptsInactive V2
// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (w *MetrcWrapper) SalesGetReceiptsInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetReceiptsInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart)
	})
}

// GET GetTransactions V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesGetTransactionsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetTransactionsV1(licensenumber)
	})
}

// GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart string, salesdateend string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart, salesdateend, licensenumber)
	})
}

// PUT UpdateDelivery V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryV1(licensenumber string, body []SalesUpdateDeliveryV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryV1(licensenumber, body)
	})
}

// PUT UpdateDelivery V2
// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryV2(licensenumber string, body []SalesUpdateDeliveryV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryComplete V1
// Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryCompleteV1(licensenumber string, body []SalesUpdateDeliveryCompleteV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryCompleteV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryComplete V2
// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryCompleteV2(licensenumber string, body []SalesUpdateDeliveryCompleteV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryCompleteV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryHub V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryHubV1(licensenumber string, body []SalesUpdateDeliveryHubV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryHub V2
// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales Delivery, Manage Sales Delivery Hub
func (w *MetrcWrapper) SalesUpdateDeliveryHubV2(licensenumber string, body []SalesUpdateDeliveryHubV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubAccept V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesUpdateDeliveryHubAcceptV1(licensenumber string, body []SalesUpdateDeliveryHubAcceptV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubAcceptV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubAccept V2
// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (w *MetrcWrapper) SalesUpdateDeliveryHubAcceptV2(licensenumber string, body []SalesUpdateDeliveryHubAcceptV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubAcceptV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubDepart V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesUpdateDeliveryHubDepartV1(licensenumber string, body []SalesUpdateDeliveryHubDepartV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubDepartV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubDepart V2
// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (w *MetrcWrapper) SalesUpdateDeliveryHubDepartV2(licensenumber string, body []SalesUpdateDeliveryHubDepartV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubDepartV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubVerifyID V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesUpdateDeliveryHubVerifyIdV1(licensenumber string, body []SalesUpdateDeliveryHubVerifyIdV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubVerifyIdV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryHubVerifyID V2
// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (w *MetrcWrapper) SalesUpdateDeliveryHubVerifyIdV2(licensenumber string, body []SalesUpdateDeliveryHubVerifyIdV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryHubVerifyIdV2(licensenumber, body)
	})
}

// PUT UpdateDeliveryRetailer V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryRetailerV1(licensenumber string, body []SalesUpdateDeliveryRetailerV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryRetailerV1(licensenumber, body)
	})
}

// PUT UpdateDeliveryRetailer V2
// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
//   - Manage Retailer Delivery
func (w *MetrcWrapper) SalesUpdateDeliveryRetailerV2(licensenumber string, body []SalesUpdateDeliveryRetailerV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateDeliveryRetailerV2(licensenumber, body)
	})
}

// PUT UpdateReceipt V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesUpdateReceiptV1(licensenumber string, body []SalesUpdateReceiptV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateReceiptV1(licensenumber, body)
	})
}

// PUT UpdateReceipt V2
// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales
func (w *MetrcWrapper) SalesUpdateReceiptV2(licensenumber string, body []SalesUpdateReceiptV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateReceiptV2(licensenumber, body)
	})
}

// PUT UpdateReceiptFinalize V2
// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// 
//   Permissions Required:
//   - Manage Sales
func (w *MetrcWrapper) SalesUpdateReceiptFinalizeV2(licensenumber string, body []SalesUpdateReceiptFinalizeV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateReceiptFinalizeV2(licensenumber, body)
	})
}

// PUT UpdateReceiptUnfinalize V2
// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// 
//   Permissions Required:
//   - Manage Sales
func (w *MetrcWrapper) SalesUpdateReceiptUnfinalizeV2(licensenumber string, body []SalesUpdateReceiptUnfinalizeV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateReceiptUnfinalizeV2(licensenumber, body)
	})
}

// PUT UpdateTransactionByDate V1
// Permissions Required:
//   - Sales
func (w *MetrcWrapper) SalesUpdateTransactionByDateV1(date string, licensenumber string, body []SalesUpdateTransactionByDateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SalesUpdateTransactionByDateV1(date, licensenumber, body)
	})
}

// POST Create V1
// Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsCreateV1(licensenumber string, body []StrainsCreateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsCreateV1(licensenumber, body)
	})
}

// POST Create V2
// Creates new strain records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsCreateV2(licensenumber string, body []StrainsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsCreateV2(licensenumber, body)
	})
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsCreateUpdateV1(licensenumber string, body []StrainsCreateUpdateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsCreateUpdateV1(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Archives an existing strain record for a Facility
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsDeleteV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.StrainsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Strain record by its Id, with an optional license number.
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.StrainsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsGetActiveV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.StrainsGetActiveV1(licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.StrainsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.StrainsGetInactiveV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT Update V2
// Updates existing strain records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Strains
func (w *MetrcWrapper) StrainsUpdateV2(licensenumber string, body []StrainsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.StrainsUpdateV2(licensenumber, body)
	})
}

// GET GetPackageAvailable V2
// Returns a list of available package tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
func (w *MetrcWrapper) TagsGetPackageAvailableV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TagsGetPackageAvailableV2(licensenumber)
	})
}

// GET GetPlantAvailable V2
// Returns a list of available plant tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
func (w *MetrcWrapper) TagsGetPlantAvailableV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TagsGetPlantAvailableV2(licensenumber)
	})
}

// GET GetStaged V2
// Returns a list of staged tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
//   - RetailId.AllowPackageStaging Key Value enabled
func (w *MetrcWrapper) TagsGetStagedV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TagsGetStagedV2(licensenumber)
	})
}

// POST Create V2
// Creates new additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (w *MetrcWrapper) AdditivesTemplatesCreateV2(licensenumber string, body []AdditivesTemplatesCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.AdditivesTemplatesCreateV2(licensenumber, body)
	})
}

// GET Get V2
// Retrieves an Additive Template by its Id.
// 
//   Permissions Required:
//   - Manage Additives
func (w *MetrcWrapper) AdditivesTemplatesGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.AdditivesTemplatesGetV2(id, licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (w *MetrcWrapper) AdditivesTemplatesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.AdditivesTemplatesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (w *MetrcWrapper) AdditivesTemplatesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.AdditivesTemplatesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// PUT Update V2
// Updates existing additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (w *MetrcWrapper) AdditivesTemplatesUpdateV2(licensenumber string, body []AdditivesTemplatesUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.AdditivesTemplatesUpdateV2(licensenumber, body)
	})
}

// POST CreateFinish V1
// Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (w *MetrcWrapper) HarvestsCreateFinishV1(licensenumber string, body []HarvestsCreateFinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreateFinishV1(licensenumber, body)
	})
}

// POST CreatePackage V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) HarvestsCreatePackageV1(licensenumber string, body []HarvestsCreatePackageV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreatePackageV1(licensenumber, body)
	})
}

// POST CreatePackage V2
// Creates packages from harvested products for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) HarvestsCreatePackageV2(licensenumber string, body []HarvestsCreatePackageV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreatePackageV2(licensenumber, body)
	})
}

// POST CreatePackageTesting V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) HarvestsCreatePackageTestingV1(licensenumber string, body []HarvestsCreatePackageTestingV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreatePackageTestingV1(licensenumber, body)
	})
}

// POST CreatePackageTesting V2
// Creates packages for testing from harvested products for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) HarvestsCreatePackageTestingV2(licensenumber string, body []HarvestsCreatePackageTestingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreatePackageTestingV2(licensenumber, body)
	})
}

// POST CreateRemoveWaste V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsCreateRemoveWasteV1(licensenumber string, body []HarvestsCreateRemoveWasteV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreateRemoveWasteV1(licensenumber, body)
	})
}

// POST CreateUnfinish V1
// Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (w *MetrcWrapper) HarvestsCreateUnfinishV1(licensenumber string, body []HarvestsCreateUnfinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreateUnfinishV1(licensenumber, body)
	})
}

// POST CreateWaste V2
// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsCreateWasteV2(licensenumber string, body []HarvestsCreateWasteV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsCreateWasteV2(licensenumber, body)
	})
}

// DELETE DeleteWaste V2
// Discontinues a specific harvest waste record by Id for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Discontinue Harvest Waste
func (w *MetrcWrapper) HarvestsDeleteWasteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsDeleteWasteV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
// 
//   Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active harvests for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive harvests for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetOnhold V1
// Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetOnhold V2
// Retrieves a list of harvests on hold for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetWaste V2
// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
// 
//   Permissions Required:
//   - View Harvests
func (w *MetrcWrapper) HarvestsGetWasteV2(harvestid string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetWasteV2(harvestid, licensenumber, pagenumber, pagesize)
	})
}

// GET GetWasteTypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) HarvestsGetWasteTypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetWasteTypesV1(no)
	})
}

// GET GetWasteTypes V2
// Retrieves a list of Waste types for harvests.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) HarvestsGetWasteTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.HarvestsGetWasteTypesV2(pagenumber, pagesize)
	})
}

// PUT UpdateFinish V2
// Marks one or more harvests as finished for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (w *MetrcWrapper) HarvestsUpdateFinishV2(licensenumber string, body []HarvestsUpdateFinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateFinishV2(licensenumber, body)
	})
}

// PUT UpdateLocation V2
// Updates the Location of Harvest for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsUpdateLocationV2(licensenumber string, body []HarvestsUpdateLocationV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateLocationV2(licensenumber, body)
	})
}

// PUT UpdateMove V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsUpdateMoveV1(licensenumber string, body []HarvestsUpdateMoveV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateMoveV1(licensenumber, body)
	})
}

// PUT UpdateRename V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsUpdateRenameV1(licensenumber string, body []HarvestsUpdateRenameV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateRenameV1(licensenumber, body)
	})
}

// PUT UpdateRename V2
// Renames one or more harvests for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (w *MetrcWrapper) HarvestsUpdateRenameV2(licensenumber string, body []HarvestsUpdateRenameV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateRenameV2(licensenumber, body)
	})
}

// PUT UpdateRestoreHarvestedPlants V2
// Restores previously harvested plants to their original state for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (w *MetrcWrapper) HarvestsUpdateRestoreHarvestedPlantsV2(licensenumber string, body []HarvestsUpdateRestoreHarvestedPlantsV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateRestoreHarvestedPlantsV2(licensenumber, body)
	})
}

// PUT UpdateUnfinish V2
// Reopens one or more previously finished harvests for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (w *MetrcWrapper) HarvestsUpdateUnfinishV2(licensenumber string, body []HarvestsUpdateUnfinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.HarvestsUpdateUnfinishV2(licensenumber, body)
	})
}

// POST CreateRecord V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsCreateRecordV1(licensenumber string, body []LabTestsCreateRecordV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsCreateRecordV1(licensenumber, body)
	})
}

// POST CreateRecord V2
// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsCreateRecordV2(licensenumber string, body []LabTestsCreateRecordV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsCreateRecordV2(licensenumber, body)
	})
}

// GET GetBatches V2
// Retrieves a list of Lab Test batches.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) LabTestsGetBatchesV2(pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetBatchesV2(pagenumber, pagesize)
	})
}

// GET GetLabtestdocument V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsGetLabtestdocumentV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetLabtestdocumentV1(id, licensenumber)
	})
}

// GET GetLabtestdocument V2
// Retrieves a specific Lab Test result document by its Id for a given Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsGetLabtestdocumentV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetLabtestdocumentV2(id, licensenumber)
	})
}

// GET GetResults V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) LabTestsGetResultsV1(licensenumber string, packageid string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetResultsV1(licensenumber, packageid)
	})
}

// GET GetResults V2
// Retrieves Lab Test results for a specified Package.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsGetResultsV2(licensenumber string, packageid string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetResultsV2(licensenumber, packageid, pagenumber, pagesize)
	})
}

// GET GetStates V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) LabTestsGetStatesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetStatesV1(no)
	})
}

// GET GetStates V2
// Returns a list of all lab testing states.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) LabTestsGetStatesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetStatesV2(no)
	})
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) LabTestsGetTypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetTypesV1(no)
	})
}

// GET GetTypes V2
// Returns a list of Lab Test types.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) LabTestsGetTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LabTestsGetTypesV2(pagenumber, pagesize)
	})
}

// PUT UpdateLabtestdocument V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsUpdateLabtestdocumentV1(licensenumber string, body []LabTestsUpdateLabtestdocumentV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsUpdateLabtestdocumentV1(licensenumber, body)
	})
}

// PUT UpdateLabtestdocument V2
// Updates one or more documents for previously submitted lab tests.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsUpdateLabtestdocumentV2(licensenumber string, body []LabTestsUpdateLabtestdocumentV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsUpdateLabtestdocumentV2(licensenumber, body)
	})
}

// PUT UpdateResultRelease V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsUpdateResultReleaseV1(licensenumber string, body []LabTestsUpdateResultReleaseV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsUpdateResultReleaseV1(licensenumber, body)
	})
}

// PUT UpdateResultRelease V2
// Releases Lab Test results for one or more packages.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) LabTestsUpdateResultReleaseV2(licensenumber string, body []LabTestsUpdateResultReleaseV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LabTestsUpdateResultReleaseV2(licensenumber, body)
	})
}

// POST Create V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsCreateV1(licensenumber string, body []LocationsCreateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsCreateV1(licensenumber, body)
	})
}

// POST Create V2
// Creates new locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsCreateV2(licensenumber string, body []LocationsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsCreateV2(licensenumber, body)
	})
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsCreateUpdateV1(licensenumber string, body []LocationsCreateUpdateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsCreateUpdateV1(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Archives a specified Location, identified by its Id, for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsDeleteV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Location by its Id.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetActiveV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetActiveV1(licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetInactiveV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetTypes V1
// Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetTypesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetTypesV1(licensenumber)
	})
}

// GET GetTypes V2
// Retrieves a list of active location types for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsGetTypesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.LocationsGetTypesV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT Update V2
// Updates existing locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (w *MetrcWrapper) LocationsUpdateV2(licensenumber string, body []LocationsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.LocationsUpdateV2(licensenumber, body)
	})
}

// GET GetStatusesByPatientLicenseNumber V1
// Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Patients
func (w *MetrcWrapper) PatientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber, licensenumber)
	})
}

// GET GetStatusesByPatientLicenseNumber V2
// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Patients
func (w *MetrcWrapper) PatientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber, licensenumber)
	})
}

// POST CreateAdditives V1
// Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesV1(licensenumber string, body []PlantsCreateAdditivesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesV1(licensenumber, body)
	})
}

// POST CreateAdditives V2
// Records additive usage details applied to specific plants at a Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesV2(licensenumber string, body []PlantsCreateAdditivesV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesV2(licensenumber, body)
	})
}

// POST CreateAdditivesBylocation V1
// Permissions Required:
//   - Manage Plants
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesBylocationV1(licensenumber string, body []PlantsCreateAdditivesBylocationV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesBylocationV1(licensenumber, body)
	})
}

// POST CreateAdditivesBylocation V2
// Records additive usage for plants based on their location within a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesBylocationV2(licensenumber string, body []PlantsCreateAdditivesBylocationV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesBylocationV2(licensenumber, body)
	})
}

// POST CreateAdditivesBylocationUsingtemplate V2
// Records additive usage for plants by location using a predefined additive template at a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesBylocationUsingtemplateV2(licensenumber string, body []PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesBylocationUsingtemplateV2(licensenumber, body)
	})
}

// POST CreateAdditivesUsingtemplate V2
// Records additive usage for plants using predefined additive templates at a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantsCreateAdditivesUsingtemplateV2(licensenumber string, body []PlantsCreateAdditivesUsingtemplateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateAdditivesUsingtemplateV2(licensenumber, body)
	})
}

// POST CreateChangegrowthphases V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsCreateChangegrowthphasesV1(licensenumber string, body []PlantsCreateChangegrowthphasesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateChangegrowthphasesV1(licensenumber, body)
	})
}

// POST CreateHarvestplants V1
// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (w *MetrcWrapper) PlantsCreateHarvestplantsV1(licensenumber string, body []PlantsCreateHarvestplantsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateHarvestplantsV1(licensenumber, body)
	})
}

// POST CreateManicure V2
// Creates harvest product records from plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (w *MetrcWrapper) PlantsCreateManicureV2(licensenumber string, body []PlantsCreateManicureV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateManicureV2(licensenumber, body)
	})
}

// POST CreateManicureplants V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (w *MetrcWrapper) PlantsCreateManicureplantsV1(licensenumber string, body []PlantsCreateManicureplantsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateManicureplantsV1(licensenumber, body)
	})
}

// POST CreateMoveplants V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsCreateMoveplantsV1(licensenumber string, body []PlantsCreateMoveplantsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateMoveplantsV1(licensenumber, body)
	})
}

// POST CreatePlantbatchPackage V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantsCreatePlantbatchPackageV1(licensenumber string, body []PlantsCreatePlantbatchPackageV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreatePlantbatchPackageV1(licensenumber, body)
	})
}

// POST CreatePlantbatchPackage V2
// Creates packages from plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantsCreatePlantbatchPackageV2(licensenumber string, body []PlantsCreatePlantbatchPackageV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreatePlantbatchPackageV2(licensenumber, body)
	})
}

// POST CreatePlantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsCreatePlantingsV1(licensenumber string, body []PlantsCreatePlantingsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreatePlantingsV1(licensenumber, body)
	})
}

// POST CreatePlantings V2
// Creates new plant batches at a specified Facility from existing plant data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsCreatePlantingsV2(licensenumber string, body []PlantsCreatePlantingsV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreatePlantingsV2(licensenumber, body)
	})
}

// POST CreateWaste V1
// Permissions Required:
//   - Manage Plants Waste
func (w *MetrcWrapper) PlantsCreateWasteV1(licensenumber string, body []PlantsCreateWasteV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateWasteV1(licensenumber, body)
	})
}

// POST CreateWaste V2
// Records waste events for plants at a Facility, including method, reason, and location details.
// 
//   Permissions Required:
//   - Manage Plants Waste
func (w *MetrcWrapper) PlantsCreateWasteV2(licensenumber string, body []PlantsCreateWasteV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsCreateWasteV2(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Destroy Veg/Flower Plants
func (w *MetrcWrapper) PlantsDeleteV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsDeleteV1(licensenumber)
	})
}

// DELETE Delete V2
// Removes plants from a Facility’s inventory while recording the reason for their disposal.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Destroy Veg/Flower Plants
func (w *MetrcWrapper) PlantsDeleteV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsDeleteV2(licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Plant by Id.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetV2(id, licensenumber)
	})
}

// GET GetAdditives V1
// Permissions Required:
//   - View/Manage Plants Additives
func (w *MetrcWrapper) PlantsGetAdditivesV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetAdditivesV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetAdditives V2
// Retrieves additive records applied to plants at a specified Facility.
// 
//   Permissions Required:
//   - View/Manage Plants Additives
func (w *MetrcWrapper) PlantsGetAdditivesV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetAdditivesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetAdditivesTypes V1
// Permissions Required:
//   -
func (w *MetrcWrapper) PlantsGetAdditivesTypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetAdditivesTypesV1(no)
	})
}

// GET GetAdditivesTypes V2
// Retrieves a list of all plant additive types defined within a Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetAdditivesTypesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetAdditivesTypesV2(no)
	})
}

// GET GetByLabel V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetByLabelV1(label string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetByLabelV1(label, licensenumber)
	})
}

// GET GetByLabel V2
// Retrieves a Plant by label.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetByLabelV2(label string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetByLabelV2(label, licensenumber)
	})
}

// GET GetFlowering V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetFloweringV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetFloweringV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetFlowering V2
// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetFloweringV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetFloweringV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetGrowthPhases V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetGrowthPhasesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetGrowthPhasesV1(licensenumber)
	})
}

// GET GetGrowthPhases V2
// Retrieves the list of growth phases supported by a specified Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetGrowthPhasesV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetGrowthPhasesV2(licensenumber)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetInactive V2
// Retrieves inactive plants at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetMother V2
// Retrieves mother-phase plants at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (w *MetrcWrapper) PlantsGetMotherV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetMotherV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetMotherInactive V2
// Retrieves inactive mother-phase plants at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (w *MetrcWrapper) PlantsGetMotherInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetMotherInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetMotherOnhold V2
// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (w *MetrcWrapper) PlantsGetMotherOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetMotherOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetOnhold V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetOnhold V2
// Retrieves plants that are currently on hold at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetVegetative V1
// Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetVegetativeV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetVegetativeV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetVegetative V2
// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (w *MetrcWrapper) PlantsGetVegetativeV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetVegetativeV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetWaste V2
// Retrieves a list of recorded plant waste events for a specific Facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (w *MetrcWrapper) PlantsGetWasteV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWasteV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetWasteMethodsAll V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetWasteMethodsAllV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWasteMethodsAllV1(no)
	})
}

// GET GetWasteMethodsAll V2
// Retrieves a list of all available plant waste methods for use within a Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetWasteMethodsAllV2(pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWasteMethodsAllV2(pagenumber, pagesize)
	})
}

// GET GetWastePackage V2
// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (w *MetrcWrapper) PlantsGetWastePackageV2(id string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWastePackageV2(id, licensenumber, pagenumber, pagesize)
	})
}

// GET GetWastePlant V2
// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (w *MetrcWrapper) PlantsGetWastePlantV2(id string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWastePlantV2(id, licensenumber, pagenumber, pagesize)
	})
}

// GET GetWasteReasons V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetWasteReasonsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWasteReasonsV1(licensenumber)
	})
}

// GET GetWasteReasons V2
// Retriveves available reasons for recording mature plant waste at a specified Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantsGetWasteReasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantsGetWasteReasonsV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT UpdateAdjust V2
// Adjusts the recorded count of plants at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsUpdateAdjustV2(licensenumber string, body []PlantsUpdateAdjustV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateAdjustV2(licensenumber, body)
	})
}

// PUT UpdateGrowthphase V2
// Changes the growth phases of plants within a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsUpdateGrowthphaseV2(licensenumber string, body []PlantsUpdateGrowthphaseV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateGrowthphaseV2(licensenumber, body)
	})
}

// PUT UpdateHarvest V2
// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (w *MetrcWrapper) PlantsUpdateHarvestV2(licensenumber string, body []PlantsUpdateHarvestV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateHarvestV2(licensenumber, body)
	})
}

// PUT UpdateLocation V2
// Moves plant batches to new locations within a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsUpdateLocationV2(licensenumber string, body []PlantsUpdateLocationV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateLocationV2(licensenumber, body)
	})
}

// PUT UpdateMerge V2
// Merges multiple plant groups into a single group within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (w *MetrcWrapper) PlantsUpdateMergeV2(licensenumber string, body []PlantsUpdateMergeV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateMergeV2(licensenumber, body)
	})
}

// PUT UpdateSplit V2
// Splits an existing plant group into multiple groups within a Facility.
// 
//   Permissions Required:
//   - View Plant
func (w *MetrcWrapper) PlantsUpdateSplitV2(licensenumber string, body []PlantsUpdateSplitV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateSplitV2(licensenumber, body)
	})
}

// PUT UpdateStrain V2
// Updates the strain information for plants within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsUpdateStrainV2(licensenumber string, body []PlantsUpdateStrainV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateStrainV2(licensenumber, body)
	})
}

// PUT UpdateTag V2
// Replaces existing plant tags with new tags for plants within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantsUpdateTagV2(licensenumber string, body []PlantsUpdateTagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantsUpdateTagV2(licensenumber, body)
	})
}

// POST CreateAssociate V2
// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
func (w *MetrcWrapper) RetailIdCreateAssociateV2(licensenumber string, body []RetailIdCreateAssociateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.RetailIdCreateAssociateV2(licensenumber, body)
	})
}

// POST CreateGenerate V2
// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
func (w *MetrcWrapper) RetailIdCreateGenerateV2(licensenumber string, body RetailIdCreateGenerateV2Request) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.RetailIdCreateGenerateV2(licensenumber, body)
	})
}

// POST CreateMerge V2
// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Key Value Settings/Retail ID Merge Packages Enabled
func (w *MetrcWrapper) RetailIdCreateMergeV2(licensenumber string, body RetailIdCreateMergeV2Request) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.RetailIdCreateMergeV2(licensenumber, body)
	})
}

// POST CreatePackageInfo V2
// Retrieves Package information for given list of Package labels.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - Admin/Employees/Packages Page/Product Labels(Manage)
func (w *MetrcWrapper) RetailIdCreatePackageInfoV2(licensenumber string, body RetailIdCreatePackageInfoV2Request) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.RetailIdCreatePackageInfoV2(licensenumber, body)
	})
}

// GET GetReceiveByLabel V2
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Manage RetailId
//   - WebApi Retail ID Read Write State (All or ReadOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
func (w *MetrcWrapper) RetailIdGetReceiveByLabelV2(label string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.RetailIdGetReceiveByLabelV2(label, licensenumber)
	})
}

// GET GetReceiveQrByShortCode V2
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Manage RetailId
//   - WebApi Retail ID Read Write State (All or ReadOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
func (w *MetrcWrapper) RetailIdGetReceiveQrByShortCodeV2(shortcode string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.RetailIdGetReceiveQrByShortCodeV2(shortcode, licensenumber)
	})
}

// POST CreateIntegratorSetup V2
// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) SandboxCreateIntegratorSetupV2(userkey string, body interface{}) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.SandboxCreateIntegratorSetupV2(userkey, body)
	})
}

// GET GetAll V1
// This endpoint provides a list of facilities for which the authenticated user has access.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) FacilitiesGetAllV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.FacilitiesGetAllV1(no)
	})
}

// GET GetAll V2
// This endpoint provides a list of facilities for which the authenticated user has access.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) FacilitiesGetAllV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.FacilitiesGetAllV2(no)
	})
}

// POST Create V2
// Adds new patients to a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsCreateV2(licensenumber string, body []PatientsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsCreateV2(licensenumber, body)
	})
}

// POST CreateAdd V1
// Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsCreateAddV1(licensenumber string, body []PatientsCreateAddV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsCreateAddV1(licensenumber, body)
	})
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsCreateUpdateV1(licensenumber string, body []PatientsCreateUpdateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsCreateUpdateV1(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Removes a Patient, identified by an Id, from a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsDeleteV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Patient by Id.
// 
//   Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsGetActiveV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsGetActiveV1(licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active patients for a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsGetActiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientsGetActiveV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT Update V2
// Updates Patient information for a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (w *MetrcWrapper) PatientsUpdateV2(licensenumber string, body []PatientsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientsUpdateV2(licensenumber, body)
	})
}

// POST CreateExternalIncoming V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersCreateExternalIncomingV1(licensenumber string, body []TransfersCreateExternalIncomingV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersCreateExternalIncomingV1(licensenumber, body)
	})
}

// POST CreateExternalIncoming V2
// Creates external incoming shipment plans for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (w *MetrcWrapper) TransfersCreateExternalIncomingV2(licensenumber string, body []TransfersCreateExternalIncomingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersCreateExternalIncomingV2(licensenumber, body)
	})
}

// POST CreateTemplates V1
// Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersCreateTemplatesV1(licensenumber string, body []TransfersCreateTemplatesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersCreateTemplatesV1(licensenumber, body)
	})
}

// POST CreateTemplatesOutgoing V2
// Creates new transfer templates for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (w *MetrcWrapper) TransfersCreateTemplatesOutgoingV2(licensenumber string, body []TransfersCreateTemplatesOutgoingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersCreateTemplatesOutgoingV2(licensenumber, body)
	})
}

// DELETE DeleteExternalIncoming V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersDeleteExternalIncomingV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersDeleteExternalIncomingV1(id, licensenumber)
	})
}

// DELETE DeleteExternalIncoming V2
// Voids an external incoming shipment plan for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (w *MetrcWrapper) TransfersDeleteExternalIncomingV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersDeleteExternalIncomingV2(id, licensenumber)
	})
}

// DELETE DeleteTemplates V1
// Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersDeleteTemplatesV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersDeleteTemplatesV1(id, licensenumber)
	})
}

// DELETE DeleteTemplatesOutgoing V2
// Archives a transfer template for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (w *MetrcWrapper) TransfersDeleteTemplatesOutgoingV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersDeleteTemplatesOutgoingV2(id, licensenumber)
	})
}

// GET GetDeliveriesPackagesStates V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) TransfersGetDeliveriesPackagesStatesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveriesPackagesStatesV1(no)
	})
}

// GET GetDeliveriesPackagesStates V2
// Returns a list of available shipment Package states.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) TransfersGetDeliveriesPackagesStatesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveriesPackagesStatesV2(no)
	})
}

// GET GetDelivery V1
// Please note: that the {id} parameter above represents a Shipment Plan ID.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryV1(id, no)
	})
}

// GET GetDelivery V2
// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryV2(id, pagenumber, pagesize)
	})
}

// GET GetDeliveryPackage V1
// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageV1(id, no)
	})
}

// GET GetDeliveryPackage V2
// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageV2(id, pagenumber, pagesize)
	})
}

// GET GetDeliveryPackageRequiredlabtestbatches V1
// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageRequiredlabtestbatchesV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageRequiredlabtestbatchesV1(id, no)
	})
}

// GET GetDeliveryPackageRequiredlabtestbatches V2
// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageRequiredlabtestbatchesV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageRequiredlabtestbatchesV2(id, pagenumber, pagesize)
	})
}

// GET GetDeliveryPackageWholesale V1
// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageWholesaleV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageWholesaleV1(id, no)
	})
}

// GET GetDeliveryPackageWholesale V2
// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryPackageWholesaleV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryPackageWholesaleV2(id, pagenumber, pagesize)
	})
}

// GET GetDeliveryTransporters V1
// Please note: that the {id} parameter above represents a Shipment Delivery ID.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryTransportersV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryTransportersV1(id, no)
	})
}

// GET GetDeliveryTransporters V2
// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryTransportersV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryTransportersV2(id, pagenumber, pagesize)
	})
}

// GET GetDeliveryTransportersDetails V1
// Please note: The {id} parameter above represents a Shipment Delivery ID.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetDeliveryTransportersDetailsV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryTransportersDetailsV1(id, no)
	})
}

// GET GetDeliveryTransportersDetails V2
// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetDeliveryTransportersDetailsV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetDeliveryTransportersDetailsV2(id, pagenumber, pagesize)
	})
}

// GET GetHub V2
// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetHubV2(estimatedarrivalend string, estimatedarrivalstart string, lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetHubV2(estimatedarrivalend, estimatedarrivalstart, lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetIncoming V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetIncomingV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetIncomingV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetIncoming V2
// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetIncomingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetIncomingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetOutgoing V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetOutgoingV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetOutgoingV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetOutgoing V2
// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetOutgoingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetRejected V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetRejectedV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetRejectedV1(licensenumber)
	})
}

// GET GetRejected V2
// Retrieves a list of shipments with rejected packages for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (w *MetrcWrapper) TransfersGetRejectedV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetRejectedV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetTemplates V1
// Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetTemplatesDelivery V1
// Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesDeliveryV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesDeliveryV1(id, no)
	})
}

// GET GetTemplatesDeliveryPackage V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersGetTemplatesDeliveryPackageV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesDeliveryPackageV1(id, no)
	})
}

// GET GetTemplatesDeliveryTransporters V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesDeliveryTransportersV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesDeliveryTransportersV1(id, no)
	})
}

// GET GetTemplatesDeliveryTransportersDetails V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesDeliveryTransportersDetailsV1(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesDeliveryTransportersDetailsV1(id, no)
	})
}

// GET GetTemplatesOutgoing V2
// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesOutgoingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetTemplatesOutgoingDelivery V2
// Retrieves a list of deliveries associated with a specific transfer template.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesOutgoingDeliveryV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesOutgoingDeliveryV2(id, pagenumber, pagesize)
	})
}

// GET GetTemplatesOutgoingDeliveryPackage V2
// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesOutgoingDeliveryPackageV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesOutgoingDeliveryPackageV2(id, pagenumber, pagesize)
	})
}

// GET GetTemplatesOutgoingDeliveryTransporters V2
// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesOutgoingDeliveryTransportersV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesOutgoingDeliveryTransportersV2(id, pagenumber, pagesize)
	})
}

// GET GetTemplatesOutgoingDeliveryTransportersDetails V2
// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (w *MetrcWrapper) TransfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id string, no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, no)
	})
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) TransfersGetTypesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTypesV1(licensenumber)
	})
}

// GET GetTypes V2
// Retrieves a list of available transfer types for a Facility based on its license number.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) TransfersGetTypesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransfersGetTypesV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT UpdateExternalIncoming V1
// Permissions Required:
//   - Transfers
func (w *MetrcWrapper) TransfersUpdateExternalIncomingV1(licensenumber string, body []TransfersUpdateExternalIncomingV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersUpdateExternalIncomingV1(licensenumber, body)
	})
}

// PUT UpdateExternalIncoming V2
// Updates external incoming shipment plans for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (w *MetrcWrapper) TransfersUpdateExternalIncomingV2(licensenumber string, body []TransfersUpdateExternalIncomingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersUpdateExternalIncomingV2(licensenumber, body)
	})
}

// PUT UpdateTemplates V1
// Permissions Required:
//   - Transfer Templates
func (w *MetrcWrapper) TransfersUpdateTemplatesV1(licensenumber string, body []TransfersUpdateTemplatesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersUpdateTemplatesV1(licensenumber, body)
	})
}

// PUT UpdateTemplatesOutgoing V2
// Updates existing transfer templates for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (w *MetrcWrapper) TransfersUpdateTemplatesOutgoingV2(licensenumber string, body []TransfersUpdateTemplatesOutgoingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransfersUpdateTemplatesOutgoingV2(licensenumber, body)
	})
}

// POST CreateDriver V2
// Creates new driver records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersCreateDriverV2(licensenumber string, body []TransportersCreateDriverV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersCreateDriverV2(licensenumber, body)
	})
}

// POST CreateVehicle V2
// Creates new vehicle records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersCreateVehicleV2(licensenumber string, body []TransportersCreateVehicleV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersCreateVehicleV2(licensenumber, body)
	})
}

// DELETE DeleteDriver V2
// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersDeleteDriverV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersDeleteDriverV2(id, licensenumber)
	})
}

// DELETE DeleteVehicle V2
// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersDeleteVehicleV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersDeleteVehicleV2(id, licensenumber)
	})
}

// GET GetDriver V2
// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
// 
//   Permissions Required:
//   - Transporters
func (w *MetrcWrapper) TransportersGetDriverV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransportersGetDriverV2(id, licensenumber)
	})
}

// GET GetDrivers V2
// Retrieves a list of drivers for a Facility.
// 
//   Permissions Required:
//   - Transporters
func (w *MetrcWrapper) TransportersGetDriversV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransportersGetDriversV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetVehicle V2
// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
// 
//   Permissions Required:
//   - Transporters
func (w *MetrcWrapper) TransportersGetVehicleV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransportersGetVehicleV2(id, licensenumber)
	})
}

// GET GetVehicles V2
// Retrieves a list of vehicles for a Facility.
// 
//   Permissions Required:
//   - Transporters
func (w *MetrcWrapper) TransportersGetVehiclesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.TransportersGetVehiclesV2(licensenumber, pagenumber, pagesize)
	})
}

// PUT UpdateDriver V2
// Updates existing driver records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersUpdateDriverV2(licensenumber string, body []TransportersUpdateDriverV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersUpdateDriverV2(licensenumber, body)
	})
}

// PUT UpdateVehicle V2
// Updates existing vehicle records for a facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (w *MetrcWrapper) TransportersUpdateVehicleV2(licensenumber string, body []TransportersUpdateVehicleV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.TransportersUpdateVehicleV2(licensenumber, body)
	})
}

// GET GetActive V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) UnitsOfMeasureGetActiveV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.UnitsOfMeasureGetActiveV1(no)
	})
}

// GET GetActive V2
// Retrieves all active units of measure.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) UnitsOfMeasureGetActiveV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.UnitsOfMeasureGetActiveV2(no)
	})
}

// GET GetInactive V2
// Retrieves all inactive units of measure.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) UnitsOfMeasureGetInactiveV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.UnitsOfMeasureGetInactiveV2(no)
	})
}

// GET GetAll V2
// Retrieves all available waste methods.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) WasteMethodsGetAllV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.WasteMethodsGetAllV2(no)
	})
}

// GET GetByCaregiverLicenseNumber V1
// Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Caregivers
func (w *MetrcWrapper) CaregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.CaregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber, licensenumber)
	})
}

// GET GetByCaregiverLicenseNumber V2
// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Caregivers
func (w *MetrcWrapper) CaregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.CaregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber, licensenumber)
	})
}

// GET GetAll V1
// Permissions Required:
//   - Manage Employees
func (w *MetrcWrapper) EmployeesGetAllV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.EmployeesGetAllV1(licensenumber)
	})
}

// GET GetAll V2
// Retrieves a list of employees for a specified Facility.
// 
//   Permissions Required:
//   - Manage Employees
//   - View Employees
func (w *MetrcWrapper) EmployeesGetAllV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.EmployeesGetAllV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetPermissions V2
// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
// 
//   Permissions Required:
//   - Manage Employees
func (w *MetrcWrapper) EmployeesGetPermissionsV2(employeelicensenumber string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.EmployeesGetPermissionsV2(employeelicensenumber, licensenumber)
	})
}

// POST Create V1
// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreateV1(licensenumber string, body []ItemsCreateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreateV1(licensenumber, body)
	})
}

// POST Create V2
// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreateV2(licensenumber string, body []ItemsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreateV2(licensenumber, body)
	})
}

// POST CreateBrand V2
// Creates one or more new item brands for the specified Facility identified by the License Number.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreateBrandV2(licensenumber string, body []ItemsCreateBrandV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreateBrandV2(licensenumber, body)
	})
}

// POST CreateFile V2
// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreateFileV2(licensenumber string, body []ItemsCreateFileV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreateFileV2(licensenumber, body)
	})
}

// POST CreatePhoto V1
// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreatePhotoV1(licensenumber string, body []ItemsCreatePhotoV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreatePhotoV1(licensenumber, body)
	})
}

// POST CreatePhoto V2
// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreatePhotoV2(licensenumber string, body []ItemsCreatePhotoV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreatePhotoV2(licensenumber, body)
	})
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsCreateUpdateV1(licensenumber string, body []ItemsCreateUpdateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsCreateUpdateV1(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Archives the specified Product by Id for the given Facility License Number.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsDeleteV2(id, licensenumber)
	})
}

// DELETE DeleteBrand V2
// Archives the specified Item Brand by Id for the given Facility License Number.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsDeleteBrandV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsDeleteBrandV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves detailed information about a specific Item by Id.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetActiveV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetActiveV1(licensenumber)
	})
}

// GET GetActive V2
// Returns a list of active items for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetBrands V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetBrandsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetBrandsV1(licensenumber)
	})
}

// GET GetBrands V2
// Retrieves a list of active item brands for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetBrandsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetBrandsV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetCategories V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) ItemsGetCategoriesV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetCategoriesV1(licensenumber)
	})
}

// GET GetCategories V2
// Retrieves a list of item categories.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) ItemsGetCategoriesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetCategoriesV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetFile V2
// Retrieves a file by its Id for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetFileV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetFileV2(id, licensenumber)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetInactiveV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetInactiveV1(licensenumber)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive items for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetInactiveV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetPhoto V1
// Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetPhotoV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetPhotoV1(id, licensenumber)
	})
}

// GET GetPhoto V2
// Retrieves an image by its Id for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsGetPhotoV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.ItemsGetPhotoV2(id, licensenumber)
	})
}

// PUT Update V2
// Updates one or more existing products for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsUpdateV2(licensenumber string, body []ItemsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsUpdateV2(licensenumber, body)
	})
}

// PUT UpdateBrand V2
// Updates one or more existing item brands for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (w *MetrcWrapper) ItemsUpdateBrandV2(licensenumber string, body []ItemsUpdateBrandV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.ItemsUpdateBrandV2(licensenumber, body)
	})
}

// POST Create V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateV1(licensenumber string, body []PackagesCreateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateV1(licensenumber, body)
	})
}

// POST Create V2
// Creates new packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateV2(licensenumber string, body []PackagesCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateV2(licensenumber, body)
	})
}

// POST CreateAdjust V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreateAdjustV1(licensenumber string, body []PackagesCreateAdjustV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateAdjustV1(licensenumber, body)
	})
}

// POST CreateAdjust V2
// Records a list of adjustments for packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreateAdjustV2(licensenumber string, body []PackagesCreateAdjustV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateAdjustV2(licensenumber, body)
	})
}

// POST CreateChangeItem V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateChangeItemV1(licensenumber string, body []PackagesCreateChangeItemV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateChangeItemV1(licensenumber, body)
	})
}

// POST CreateChangeLocation V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateChangeLocationV1(licensenumber string, body []PackagesCreateChangeLocationV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateChangeLocationV1(licensenumber, body)
	})
}

// POST CreateFinish V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreateFinishV1(licensenumber string, body []PackagesCreateFinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateFinishV1(licensenumber, body)
	})
}

// POST CreatePlantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreatePlantingsV1(licensenumber string, body []PackagesCreatePlantingsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreatePlantingsV1(licensenumber, body)
	})
}

// POST CreatePlantings V2
// Creates new plantings from packages for a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreatePlantingsV2(licensenumber string, body []PackagesCreatePlantingsV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreatePlantingsV2(licensenumber, body)
	})
}

// POST CreateRemediate V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreateRemediateV1(licensenumber string, body []PackagesCreateRemediateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateRemediateV1(licensenumber, body)
	})
}

// POST CreateTesting V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateTestingV1(licensenumber string, body []PackagesCreateTestingV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateTestingV1(licensenumber, body)
	})
}

// POST CreateTesting V2
// Creates new packages for testing for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesCreateTestingV2(licensenumber string, body []PackagesCreateTestingV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateTestingV2(licensenumber, body)
	})
}

// POST CreateUnfinish V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesCreateUnfinishV1(licensenumber string, body []PackagesCreateUnfinishV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesCreateUnfinishV1(licensenumber, body)
	})
}

// DELETE Delete V2
// Discontinues a Package at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesDeleteV2(id, licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Package by its Id.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetAdjustReasons V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PackagesGetAdjustReasonsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetAdjustReasonsV1(licensenumber)
	})
}

// GET GetAdjustReasons V2
// Retrieves a list of adjustment reasons for packages at a specified Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PackagesGetAdjustReasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetAdjustReasonsV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetByLabel V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetByLabelV1(label string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetByLabelV1(label, licensenumber)
	})
}

// GET GetByLabel V2
// Retrieves a Package by its label.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetByLabelV2(label string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetByLabelV2(label, licensenumber)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetIntransit V2
// Retrieves a list of packages in transit for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetIntransitV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetIntransitV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetLabsamples V2
// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetLabsamplesV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetLabsamplesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetOnhold V1
// Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetOnhold V2
// Retrieves a list of packages on hold for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetSourceHarvest V2
// Retrieves the source harvests for a Package by its Id.
// 
//   Permissions Required:
//   - View Package Source Harvests
func (w *MetrcWrapper) PackagesGetSourceHarvestV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetSourceHarvestV2(id, licensenumber)
	})
}

// GET GetTransferred V2
// Retrieves a list of transferred packages for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
func (w *MetrcWrapper) PackagesGetTransferredV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetTransferredV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PackagesGetTypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetTypesV1(no)
	})
}

// GET GetTypes V2
// Retrieves a list of available Package types.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PackagesGetTypesV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PackagesGetTypesV2(no)
	})
}

// PUT UpdateAdjust V2
// Set the final quantity for a Package.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateAdjustV2(licensenumber string, body []PackagesUpdateAdjustV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateAdjustV2(licensenumber, body)
	})
}

// PUT UpdateChangeNote V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
//   - Manage Package Notes
func (w *MetrcWrapper) PackagesUpdateChangeNoteV1(licensenumber string, body []PackagesUpdateChangeNoteV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateChangeNoteV1(licensenumber, body)
	})
}

// PUT UpdateDecontaminate V2
// Updates the Product decontaminate information for a list of packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateDecontaminateV2(licensenumber string, body []PackagesUpdateDecontaminateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateDecontaminateV2(licensenumber, body)
	})
}

// PUT UpdateDonationFlag V2
// Flags one or more packages for donation at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateDonationFlagV2(licensenumber string, body []PackagesUpdateDonationFlagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateDonationFlagV2(licensenumber, body)
	})
}

// PUT UpdateDonationUnflag V2
// Removes the donation flag from one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateDonationUnflagV2(licensenumber string, body []PackagesUpdateDonationUnflagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateDonationUnflagV2(licensenumber, body)
	})
}

// PUT UpdateExternalid V2
// Updates the external identifiers for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Package Inventory
//   - External Id Enabled
func (w *MetrcWrapper) PackagesUpdateExternalidV2(licensenumber string, body []PackagesUpdateExternalidV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateExternalidV2(licensenumber, body)
	})
}

// PUT UpdateFinish V2
// Updates a list of packages as finished for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateFinishV2(licensenumber string, body []PackagesUpdateFinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateFinishV2(licensenumber, body)
	})
}

// PUT UpdateItem V2
// Updates the associated Item for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesUpdateItemV2(licensenumber string, body []PackagesUpdateItemV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateItemV2(licensenumber, body)
	})
}

// PUT UpdateLabTestRequired V2
// Updates the list of required lab test batches for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesUpdateLabTestRequiredV2(licensenumber string, body []PackagesUpdateLabTestRequiredV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateLabTestRequiredV2(licensenumber, body)
	})
}

// PUT UpdateLocation V2
// Updates the Location and Sublocation for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesUpdateLocationV2(licensenumber string, body []PackagesUpdateLocationV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateLocationV2(licensenumber, body)
	})
}

// PUT UpdateNote V2
// Updates notes associated with one or more packages for the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
//   - Manage Package Notes
func (w *MetrcWrapper) PackagesUpdateNoteV2(licensenumber string, body []PackagesUpdateNoteV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateNoteV2(licensenumber, body)
	})
}

// PUT UpdateRemediate V2
// Updates a list of Product remediations for packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateRemediateV2(licensenumber string, body []PackagesUpdateRemediateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateRemediateV2(licensenumber, body)
	})
}

// PUT UpdateTradesampleFlag V2
// Flags or unflags one or more packages at the specified Facility as trade samples.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateTradesampleFlagV2(licensenumber string, body []PackagesUpdateTradesampleFlagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateTradesampleFlagV2(licensenumber, body)
	})
}

// PUT UpdateTradesampleUnflag V2
// Removes the trade sample flag from one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateTradesampleUnflagV2(licensenumber string, body []PackagesUpdateTradesampleUnflagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateTradesampleUnflagV2(licensenumber, body)
	})
}

// PUT UpdateUnfinish V2
// Updates a list of packages as unfinished for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (w *MetrcWrapper) PackagesUpdateUnfinishV2(licensenumber string, body []PackagesUpdateUnfinishV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateUnfinishV2(licensenumber, body)
	})
}

// PUT UpdateUsebydate V2
// Updates the use-by date for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PackagesUpdateUsebydateV2(licensenumber string, body []PackagesUpdateUsebydateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PackagesUpdateUsebydateV2(licensenumber, body)
	})
}

// POST Create V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsCreateV1(licensenumber string, body []PatientCheckInsCreateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsCreateV1(licensenumber, body)
	})
}

// POST Create V2
// Records patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsCreateV2(licensenumber string, body []PatientCheckInsCreateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsCreateV2(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsDeleteV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsDeleteV1(id, licensenumber)
	})
}

// DELETE Delete V2
// Archives a Patient Check-In, identified by its Id, for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsDeleteV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsDeleteV2(id, licensenumber)
	})
}

// GET GetAll V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsGetAllV1(checkindateend string, checkindatestart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientCheckInsGetAllV1(checkindateend, checkindatestart, licensenumber)
	})
}

// GET GetAll V2
// Retrieves a list of patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsGetAllV2(checkindateend string, checkindatestart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientCheckInsGetAllV2(checkindateend, checkindatestart, licensenumber)
	})
}

// GET GetLocations V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PatientCheckInsGetLocationsV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientCheckInsGetLocationsV1(no)
	})
}

// GET GetLocations V2
// Retrieves a list of Patient Check-In locations.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PatientCheckInsGetLocationsV2(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PatientCheckInsGetLocationsV2(no)
	})
}

// PUT Update V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsUpdateV1(licensenumber string, body []PatientCheckInsUpdateV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsUpdateV1(licensenumber, body)
	})
}

// PUT Update V2
// Updates patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (w *MetrcWrapper) PatientCheckInsUpdateV2(licensenumber string, body []PatientCheckInsUpdateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PatientCheckInsUpdateV2(licensenumber, body)
	})
}

// POST CreateAdditives V1
// Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantBatchesCreateAdditivesV1(licensenumber string, body []PlantBatchesCreateAdditivesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateAdditivesV1(licensenumber, body)
	})
}

// POST CreateAdditives V2
// Records Additive usage details for plant batches at a specific Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantBatchesCreateAdditivesV2(licensenumber string, body []PlantBatchesCreateAdditivesV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateAdditivesV2(licensenumber, body)
	})
}

// POST CreateAdditivesUsingtemplate V2
// Records Additive usage for plant batches at a Facility using predefined additive templates.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (w *MetrcWrapper) PlantBatchesCreateAdditivesUsingtemplateV2(licensenumber string, body []PlantBatchesCreateAdditivesUsingtemplateV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateAdditivesUsingtemplateV2(licensenumber, body)
	})
}

// POST CreateAdjust V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateAdjustV1(licensenumber string, body []PlantBatchesCreateAdjustV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateAdjustV1(licensenumber, body)
	})
}

// POST CreateAdjust V2
// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateAdjustV2(licensenumber string, body []PlantBatchesCreateAdjustV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateAdjustV2(licensenumber, body)
	})
}

// POST CreateChangegrowthphase V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateChangegrowthphaseV1(licensenumber string, body []PlantBatchesCreateChangegrowthphaseV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateChangegrowthphaseV1(licensenumber, body)
	})
}

// POST CreateGrowthphase V2
// Updates the growth phase of plants at a specified Facility based on tracking information.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateGrowthphaseV2(licensenumber string, body []PlantBatchesCreateGrowthphaseV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateGrowthphaseV2(licensenumber, body)
	})
}

// POST CreatePackage V2
// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantBatchesCreatePackageV2(isfrommotherplant string, licensenumber string, body []PlantBatchesCreatePackageV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreatePackageV2(isfrommotherplant, licensenumber, body)
	})
}

// POST CreatePackageFrommotherplant V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantBatchesCreatePackageFrommotherplantV1(licensenumber string, body []PlantBatchesCreatePackageFrommotherplantV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreatePackageFrommotherplantV1(licensenumber, body)
	})
}

// POST CreatePackageFrommotherplant V2
// Creates packages from mother plants at the specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantBatchesCreatePackageFrommotherplantV2(licensenumber string, body []PlantBatchesCreatePackageFrommotherplantV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreatePackageFrommotherplantV2(licensenumber, body)
	})
}

// POST CreatePlantings V2
// Creates new plantings for a Facility by generating plant batches based on provided planting details.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreatePlantingsV2(licensenumber string, body []PlantBatchesCreatePlantingsV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreatePlantingsV2(licensenumber, body)
	})
}

// POST CreateSplit V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateSplitV1(licensenumber string, body []PlantBatchesCreateSplitV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateSplitV1(licensenumber, body)
	})
}

// POST CreateSplit V2
// Splits an existing Plant Batch into multiple groups at the specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateSplitV2(licensenumber string, body []PlantBatchesCreateSplitV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateSplitV2(licensenumber, body)
	})
}

// POST CreateWaste V1
// Permissions Required:
//   - Manage Plants Waste
func (w *MetrcWrapper) PlantBatchesCreateWasteV1(licensenumber string, body []PlantBatchesCreateWasteV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateWasteV1(licensenumber, body)
	})
}

// POST CreateWaste V2
// Records waste information for plant batches based on the submitted data for the specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Waste
func (w *MetrcWrapper) PlantBatchesCreateWasteV2(licensenumber string, body []PlantBatchesCreateWasteV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateWasteV2(licensenumber, body)
	})
}

// POST Createpackages V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (w *MetrcWrapper) PlantBatchesCreatepackagesV1(isfrommotherplant string, licensenumber string, body []PlantBatchesCreatepackagesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreatepackagesV1(isfrommotherplant, licensenumber, body)
	})
}

// POST Createplantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (w *MetrcWrapper) PlantBatchesCreateplantingsV1(licensenumber string, body []PlantBatchesCreateplantingsV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesCreateplantingsV1(licensenumber, body)
	})
}

// DELETE Delete V1
// Permissions Required:
//   - View Immature Plants
//   - Destroy Immature Plants
func (w *MetrcWrapper) PlantBatchesDeleteV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesDeleteV1(licensenumber)
	})
}

// DELETE Delete V2
// Completes the destruction of plant batches based on the provided input data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Destroy Immature Plants
func (w *MetrcWrapper) PlantBatchesDeleteV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesDeleteV2(licensenumber)
	})
}

// GET Get V1
// Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetV1(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetV1(id, licensenumber)
	})
}

// GET Get V2
// Retrieves a Plant Batch by Id.
// 
//   Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetV2(id string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetV2(id, licensenumber)
	})
}

// GET GetActive V1
// Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetActive V2
// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetInactive V1
// Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber)
	})
}

// GET GetInactive V2
// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize)
	})
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PlantBatchesGetTypesV1(no string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetTypesV1(no)
	})
}

// GET GetTypes V2
// Retrieves a list of plant batch types.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantBatchesGetTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetTypesV2(pagenumber, pagesize)
	})
}

// GET GetWaste V2
// Retrieves waste details associated with plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (w *MetrcWrapper) PlantBatchesGetWasteV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetWasteV2(licensenumber, pagenumber, pagesize)
	})
}

// GET GetWasteReasons V1
// Permissions Required:
//   - None
func (w *MetrcWrapper) PlantBatchesGetWasteReasonsV1(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetWasteReasonsV1(licensenumber)
	})
}

// GET GetWasteReasons V2
// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
// 
//   Permissions Required:
//   - None
func (w *MetrcWrapper) PlantBatchesGetWasteReasonsV2(licensenumber string) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", true, func() (interface{}, error) {
		return w.Client.PlantBatchesGetWasteReasonsV2(licensenumber)
	})
}

// PUT UpdateLocation V2
// Moves one or more plant batches to new locations with in a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
func (w *MetrcWrapper) PlantBatchesUpdateLocationV2(licensenumber string, body []PlantBatchesUpdateLocationV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesUpdateLocationV2(licensenumber, body)
	})
}

// PUT UpdateMoveplantbatches V1
// Permissions Required:
//   - View Immature Plants
func (w *MetrcWrapper) PlantBatchesUpdateMoveplantbatchesV1(licensenumber string, body []PlantBatchesUpdateMoveplantbatchesV1RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesUpdateMoveplantbatchesV1(licensenumber, body)
	})
}

// PUT UpdateName V2
// Renames plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantBatchesUpdateNameV2(licensenumber string, body []PlantBatchesUpdateNameV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesUpdateNameV2(licensenumber, body)
	})
}

// PUT UpdateStrain V2
// Changes the strain of plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantBatchesUpdateStrainV2(licensenumber string, body []PlantBatchesUpdateStrainV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesUpdateStrainV2(licensenumber, body)
	})
}

// PUT UpdateTag V2
// Replaces tags for plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (w *MetrcWrapper) PlantBatchesUpdateTagV2(licensenumber string, body []PlantBatchesUpdateTagV2RequestItem) (interface{}, error) {
	return w.RateLimiter.Execute(context.TODO(), "", false, func() (interface{}, error) {
		return w.Client.PlantBatchesUpdateTagV2(licensenumber, body)
	})
}

