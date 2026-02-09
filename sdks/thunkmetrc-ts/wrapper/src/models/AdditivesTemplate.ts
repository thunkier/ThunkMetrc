
import type { AdditivesTemplateActiveIngredientsItem } from './AdditivesTemplateActiveIngredientsItem';
export interface AdditivesTemplate {
    ActiveIngredients?: AdditivesTemplateActiveIngredientsItem[];
    AdditiveType?: string;
    AdditiveTypeName?: string;
    ApplicationDevice?: string;
    EpaRegistrationNumber?: string;
    FacilityId?: number;
    Id?: number;
    IsActive?: boolean;
    LastModified?: string;
    Name?: string;
    Note?: string;
    ProductSupplier?: string;
    ProductTradeName?: string;
    RestrictiveEntryIntervalQuantityDescription?: string;
    RestrictiveEntryIntervalTimeDescription?: string;
}
