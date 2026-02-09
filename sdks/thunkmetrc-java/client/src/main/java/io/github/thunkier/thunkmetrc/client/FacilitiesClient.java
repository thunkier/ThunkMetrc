package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class FacilitiesClient {
    private final MetrcClient client;

    public FacilitiesClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * GET GetFacilities
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getFacilities(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/facilities/v2");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

