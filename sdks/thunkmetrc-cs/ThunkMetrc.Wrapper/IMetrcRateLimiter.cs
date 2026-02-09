using System;
using System.Threading;
using System.Threading.Tasks;

namespace ThunkMetrc.Wrapper
{
    public interface IMetrcRateLimiter
    {
        Task<T> ExecuteAsync<T>(string? facility, bool isGet, Func<Task<T>> op, CancellationToken cancellationToken = default);
    }
}
