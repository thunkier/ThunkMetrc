from typing import List, Optional, Any, Dict
from datetime import datetime, timedelta, timezone
from thunkmetrc.wrapper import MetrcWrapper

async def active_packages_inventory_sync(
    wrapper: MetrcWrapper,
    license_number: str,
    last_known_sync: Optional[datetime] = None,
    buffer_minutes: int = 5
) -> List[Dict[str, Any]]:
    """
    Syncs all active packages modified within the time window.
    
    :param wrapper: MetrcWrapper instance
    :param license_number: Facility license number
    :param last_known_sync: Last successful sync time (datetime). If None, defaults to 24h ago.
    :param buffer_minutes: Buffer to subtract from start time (default 5).
    :return: List of active packages
    """
    end_time = datetime.now(timezone.utc)
    
    if last_known_sync:
        start_time = last_known_sync - timedelta(minutes=buffer_minutes)
    else:
        # Default to 24 hours ago if no sync state
        start_time = end_time - timedelta(hours=24)

    # Ensure UTC if naive (assume UTC if naive, or convert)
    if start_time.tzinfo is None:
        start_time = start_time.replace(tzinfo=timezone.utc)
    if end_time.tzinfo is None:
        end_time = end_time.replace(tzinfo=timezone.utc)

    # Format as ISO 8601 strings (e.g. 2023-10-26T12:00:00+00:00)
    # isoformat() produces valid ISO 8601
    start_str = start_time.isoformat()
    end_str = end_time.isoformat()

    all_packages: List[Dict[str, Any]] = []
    page = 1
    page_size = 20
    previous_first_id = None

    while True:
        # Use the standardized method name
        response = await wrapper.get_packages(
            last_modified_end=end_str,
            last_modified_start=start_str,
            license_number=license_number,
            page_number=str(page),
            page_size=str(page_size)
        )

        if response:
            data = None
            # Response should be a dict if parsed from JSON
            # e.g. { "Data": [...], "TotalPages": 1 }
            if isinstance(response, dict):
                data = response.get("Data")
                if isinstance(data, list):
                    # ID-based loop detection
                    if len(data) > 0:
                        first_id = data[0].get("Id") if isinstance(data[0], dict) else None
                        if first_id is not None and first_id == previous_first_id:
                            break
                        previous_first_id = first_id
                    elif len(data) == 0:
                        break
                    all_packages.extend(data)
                
                total_pages = response.get("TotalPages")
                if isinstance(total_pages, int):
                    if page >= total_pages:
                        break
                else:
                    # No pagination info or single page?
                    break
            elif isinstance(response, list):
                # Flat array response
                if len(response) == 0:
                    break
                first_id = response[0].get("Id") if isinstance(response[0], dict) else None
                if first_id is not None and first_id == previous_first_id:
                    break
                previous_first_id = first_id
                all_packages.extend(response)
                # Flat arrays without TotalPages - break after first
                break
            else:
                break
        else:
            break

        page += 1

    return all_packages
