
import type { ProcessingJobActiveJobTypeAttributesItem } from './ProcessingJobActiveJobTypeAttributesItem';
export interface ActiveJobType {
    Attributes?: ProcessingJobActiveJobTypeAttributesItem[];
    CategoryName?: string;
    Description?: string;
    ForItems?: boolean;
    ForProcessingJobs?: boolean;
    Id?: number;
    Name?: string;
    ProcessingSteps?: string;
    RequiresProcessingJobAttributes?: boolean;
}
