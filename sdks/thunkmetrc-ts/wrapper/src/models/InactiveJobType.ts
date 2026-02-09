
import type { ProcessingJobInactiveJobTypeAttributesItem } from './ProcessingJobInactiveJobTypeAttributesItem';
export interface InactiveJobType {
    Attributes?: ProcessingJobInactiveJobTypeAttributesItem[];
    CategoryName?: string;
    Description?: string;
    ForItems?: boolean;
    ForProcessingJobs?: boolean;
    Id?: number;
    Name?: string;
    ProcessingSteps?: string;
    RequiresProcessingJobAttributes?: boolean;
}
