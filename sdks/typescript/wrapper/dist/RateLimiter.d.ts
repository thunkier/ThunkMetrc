export declare class RateLimiterConfig {
    enabled: boolean;
    maxGetPerSecondPerFacility: number;
    maxGetPerSecondIntegrator: number;
    maxConcurrentGetPerFacility: number;
    maxConcurrentGetIntegrator: number;
}
export declare class MetrcRateLimiter {
    private config;
    private integratorRate;
    private integratorSem;
    private facilityRates;
    private facilitySems;
    constructor(config?: RateLimiterConfig);
    execute<T>(facility: string | null, isGet: boolean, op: () => Promise<T>): Promise<T>;
    private executeRateLimited;
    private getFacilityRate;
    private getFacilitySem;
}
