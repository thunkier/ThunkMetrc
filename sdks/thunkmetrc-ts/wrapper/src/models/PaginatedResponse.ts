export interface PaginatedResponse<T> {
    Total?: number;
    PageSize?: number;
    Current?: number;
    TotalPages?: number;
    Next?: number;
    Previous?: number;
    Data?: T[];
}
