
import type { LabTestsBatchItemCategoriesItem } from './LabTestsBatchItemCategoriesItem';
import type { LabTestsBatchLabTestTypesItem } from './LabTestsBatchLabTestTypesItem';
export interface Batch {
    Id?: number;
    ItemCategories?: LabTestsBatchItemCategoriesItem[];
    ItemCategoryCount?: number;
    LabTestTypeCount?: number;
    LabTestTypes?: LabTestsBatchLabTestTypesItem[];
    Name?: string;
    RequiresAllFromLabTestBatch?: boolean;
}
