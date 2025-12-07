"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetrcRateLimiter = exports.RateLimiterConfig = void 0;
class RateLimiterConfig {
    constructor() {
        this.enabled = false;
        this.maxGetPerSecondPerFacility = 50;
        this.maxGetPerSecondIntegrator = 150;
        this.maxConcurrentGetPerFacility = 10;
        this.maxConcurrentGetIntegrator = 30;
    }
}
exports.RateLimiterConfig = RateLimiterConfig;
class Semaphore {
    constructor(max) {
        this.max = max;
        this.tasks = [];
        this.count = 0;
    }
    async acquire() {
        if (this.count < this.max) {
            this.count++;
            return;
        }
        return new Promise(resolve => this.tasks.push(resolve));
    }
    release() {
        if (this.tasks.length > 0) {
            const next = this.tasks.shift();
            next();
        }
        else {
            this.count--;
        }
    }
}
class TokenBucket {
    constructor(rate, capacity) {
        this.rate = rate;
        this.capacity = capacity;
        this.tokens = capacity;
        this.lastRefill = Date.now();
    }
    async wait() {
        this.refill();
        if (this.tokens >= 1) {
            this.tokens -= 1;
            return;
        }
        const missing = 1 - this.tokens;
        const waitMs = (missing / this.rate) * 1000;
        await new Promise(resolve => setTimeout(resolve, waitMs));
        return this.wait();
    }
    refill() {
        const now = Date.now();
        const elapsedSec = (now - this.lastRefill) / 1000;
        this.tokens = Math.min(this.capacity, this.tokens + (elapsedSec * this.rate));
        this.lastRefill = now;
    }
}
class MetrcRateLimiter {
    constructor(config) {
        this.facilityRates = new Map();
        this.facilitySems = new Map();
        this.config = config || new RateLimiterConfig();
        this.integratorRate = new TokenBucket(this.config.maxGetPerSecondIntegrator, this.config.maxGetPerSecondIntegrator);
        this.integratorSem = new Semaphore(this.config.maxConcurrentGetIntegrator);
    }
    async execute(facility, isGet, op) {
        if (!this.config.enabled || !isGet) {
            return op();
        }
        await this.integratorSem.acquire();
        try {
            if (facility) {
                const sem = this.getFacilitySem(facility);
                await sem.acquire();
                try {
                    return await this.executeRateLimited(facility, op);
                }
                finally {
                    sem.release();
                }
            }
            else {
                return await this.executeRateLimited(facility, op);
            }
        }
        finally {
            this.integratorSem.release();
        }
    }
    async executeRateLimited(facility, op) {
        await this.integratorRate.wait();
        if (facility) {
            await this.getFacilityRate(facility).wait();
        }
        // Retry loop
        while (true) {
            try {
                return await op();
            }
            catch (err) {
                // Check for 429
                if (err?.response?.status === 429 || (err?.message && err.message.includes('429'))) {
                    // Backoff
                    const retryAfterInfo = err?.response?.headers?.['retry-after'];
                    let waitMs = 1000;
                    if (retryAfterInfo) {
                        const seconds = parseInt(retryAfterInfo, 10);
                        if (!isNaN(seconds))
                            waitMs = seconds * 1000;
                    }
                    await new Promise(resolve => setTimeout(resolve, waitMs));
                    continue;
                }
                throw err;
            }
        }
    }
    getFacilityRate(facility) {
        let tb = this.facilityRates.get(facility);
        if (!tb) {
            tb = new TokenBucket(this.config.maxGetPerSecondPerFacility, this.config.maxGetPerSecondPerFacility);
            this.facilityRates.set(facility, tb);
        }
        return tb;
    }
    getFacilitySem(facility) {
        let sem = this.facilitySems.get(facility);
        if (!sem) {
            sem = new Semaphore(this.config.maxConcurrentGetPerFacility);
            this.facilitySems.set(facility, sem);
        }
        return sem;
    }
}
exports.MetrcRateLimiter = MetrcRateLimiter;
