export class RateLimiterConfig {
  enabled: boolean = false;
  maxGetPerSecondPerFacility: number = 50;
  maxGetPerSecondIntegrator: number = 150;
  maxConcurrentGetPerFacility: number = 10;
  maxConcurrentGetIntegrator: number = 30;
}

class Semaphore {
  private tasks: (() => void)[] = [];
  private count = 0;

  constructor(private max: number) {}

  async acquire(): Promise<void> {
    if (this.count < this.max) {
      this.count++;
      return;
    }
    return new Promise<void>(resolve => this.tasks.push(resolve));
  }

  release(): void {
    if (this.tasks.length > 0) {
      const next = this.tasks.shift()!;
      next();
    } else {
      this.count--;
    }
  }
}

class TokenBucket {
  private tokens: number;
  private lastRefill: number;

  constructor(private rate: number, private capacity: number) {
    this.tokens = capacity;
    this.lastRefill = Date.now();
  }

  async wait(): Promise<void> {
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

  private refill() {
    const now = Date.now();
    const elapsedSec = (now - this.lastRefill) / 1000;
    this.tokens = Math.min(this.capacity, this.tokens + (elapsedSec * this.rate));
    this.lastRefill = now;
  }
}

export class MetrcRateLimiter {
  private config: RateLimiterConfig;
  
  private integratorRate: TokenBucket;
  private integratorSem: Semaphore;
  
  private facilityRates: Map<string, TokenBucket> = new Map();
  private facilitySems: Map<string, Semaphore> = new Map();

  constructor(config?: RateLimiterConfig) {
    this.config = config || new RateLimiterConfig();
    this.integratorRate = new TokenBucket(this.config.maxGetPerSecondIntegrator, this.config.maxGetPerSecondIntegrator);
    this.integratorSem = new Semaphore(this.config.maxConcurrentGetIntegrator);
  }

  public async execute<T>(facility: string | null, isGet: boolean, op: () => Promise<T>): Promise<T> {
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
        } finally {
            sem.release();
        }
      } else {
        return await this.executeRateLimited(facility, op);
      }
    } finally {
      this.integratorSem.release();
    }
  }

  private async executeRateLimited<T>(facility: string | null, op: () => Promise<T>): Promise<T> {
    await this.integratorRate.wait();
    if (facility) {
      await this.getFacilityRate(facility).wait();
    }

    // Retry loop
    while (true) {
      try {
        return await op();
      } catch (err: any) {
        // Check for 429
        if (err?.response?.status === 429 || (err?.message && err.message.includes('429'))) {
           // Backoff
           const retryAfterInfo = err?.response?.headers?.['retry-after'];
           let waitMs = 1000;
           if (retryAfterInfo) {
             const seconds = parseInt(retryAfterInfo, 10);
             if (!isNaN(seconds)) waitMs = seconds * 1000;
           }
           await new Promise(resolve => setTimeout(resolve, waitMs));
           continue;
        }
        throw err;
      }
    }
  }

  private getFacilityRate(facility: string): TokenBucket {
    let tb = this.facilityRates.get(facility);
    if (!tb) {
      tb = new TokenBucket(this.config.maxGetPerSecondPerFacility, this.config.maxGetPerSecondPerFacility);
      this.facilityRates.set(facility, tb);
    }
    return tb;
  }

  private getFacilitySem(facility: string): Semaphore {
    let sem = this.facilitySems.get(facility);
    if (!sem) {
      sem = new Semaphore(this.config.maxConcurrentGetPerFacility);
      this.facilitySems.set(facility, sem);
    }
    return sem;
  }
}
