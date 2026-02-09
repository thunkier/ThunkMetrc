use async_stream::stream;
use futures_core::stream::Stream;
use std::error::Error;
use crate::models::PaginatedResponse;

pub fn iterate_pages<T, F, Fut>(
    fetcher: F,
) -> impl Stream<Item = Result<T, Box<dyn Error + Send + Sync>>>
where
    T: Clone + Send + 'static,
    F: Fn(i32) -> Fut + Send + Sync + 'static,
    Fut: std::future::Future<Output = Result<PaginatedResponse<T>, Box<dyn Error + Send + Sync>>> + Send,
{
    stream! {
        let mut page = 1;
        loop {
            let response = match fetcher(page).await {
                Ok(res) => res,
                Err(e) => {
                    yield Err(e);
                    return;
                }
            };
            
            if let Some(data) = response.data {
                if data.is_empty() {
                    break;
                }
                for item in data {
                     yield Ok(item);
                }
            } else {
                 break;
            }

            if let Some(total_pages) = response.total_pages {
                 if (page as i64) >= total_pages {
                     break;
                 }
            }
            page += 1;
        }
    }
}
