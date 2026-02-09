
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    ActiveJobType,
    CreateAdjustProcessingJobRequestItem,
    CreateJobTypesRequestItem,
    CreateProcessingJobPackagesRequestItem,
    FinishProcessingJobRequestItem,
    InactiveJobType,
    JobTypesAttribute,
    JobTypesCategory,
    ProcessingJob,
    StartProcessingJobRequestItem,
    UnfinishProcessingJobRequestItem,
    UpdateJobTypesRequestItem,
    WriteResponse,
} from '../models';

export class ProcessingJobService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateAdjustProcessingJob
   */
  public async createAdjustProcessingJob(body: CreateAdjustProcessingJobRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createAdjustProcessingJob(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateJobTypes
   */
  public async createProcessingJobJobTypes(body: CreateJobTypesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createProcessingJobJobTypes(body,licenseNumber,
        )
    );
  }

  /**
   * Creates packages from processing jobs at a Facility, including optional location and note assignments.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateProcessingJobPackages
   */
  public async createProcessingJobPackages(body: CreateProcessingJobPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createProcessingJobPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Archives a Processing Job Type at a Facility, making it inactive for future use.
   * Permissions Required:
   * - Manage Processing Job
   * DELETE DeleteJobTypeById
   */
  public async deleteProcessingJobJobTypeById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteProcessingJobJobTypeById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
   * Permissions Required:
   * - Manage Processing Job
   * DELETE DeleteProcessingJobById
   */
  public async deleteProcessingJobById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteProcessingJobById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Completes processing jobs at a Facility by recording final notes and waste measurements.
   * Permissions Required:
   * - Manage Processing Job
   * PUT FinishProcessingJob
   */
  public async finishProcessingJobProcessingJob(body: FinishProcessingJobRequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.finishProcessingJobProcessingJob(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of all active processing job types defined within a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetActiveJobTypes
   */
  public async getProcessingJobActiveJobTypes(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ActiveJobType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getProcessingJobActiveJobTypes(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves active processing jobs at a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetActiveProcessingJob
   */
  public async getActiveProcessingJob(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ProcessingJob>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveProcessingJob(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of all inactive processing job types defined within a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetInactiveJobTypes
   */
  public async getProcessingJobInactiveJobTypes(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<InactiveJobType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getProcessingJobInactiveJobTypes(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves inactive processing jobs at a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetInactiveProcessingJob
   */
  public async getInactiveProcessingJob(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ProcessingJob>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveProcessingJob(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves all processing job attributes available for a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetJobTypesAttributes
   */
  public async getProcessingJobJobTypesAttributes(body?: any, licenseNumber?: string): Promise<PaginatedResponse<JobTypesAttribute>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getProcessingJobJobTypesAttributes(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves all processing job categories available for a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetJobTypesCategories
   */
  public async getProcessingJobJobTypesCategories(body?: any, licenseNumber?: string): Promise<PaginatedResponse<JobTypesCategory>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getProcessingJobJobTypesCategories(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a ProcessingJob by Id.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetProcessingJobById
   */
  public async getProcessingJobById(id: string, body?: any, licenseNumber?: string): Promise<ProcessingJob> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getProcessingJobById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Initiates new processing jobs at a Facility, including job details and associated packages.
   * Permissions Required:
   * - Manage Processing Job
   * POST StartProcessingJob
   */
  public async startProcessingJobProcessingJob(body: StartProcessingJobRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.startProcessingJobProcessingJob(body,licenseNumber,
        )
    );
  }

  /**
   * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
   * Permissions Required:
   * - Manage Processing Job
   * PUT UnfinishProcessingJob
   */
  public async unfinishProcessingJobProcessingJob(body: UnfinishProcessingJobRequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.unfinishProcessingJobProcessingJob(body,licenseNumber,
        )
    );
  }

  /**
   * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
   * Permissions Required:
   * - Manage Processing Job
   * PUT UpdateJobTypes
   */
  public async updateProcessingJobJobTypes(body: UpdateJobTypesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateProcessingJobJobTypes(body,licenseNumber,
        )
    );
  }
}

