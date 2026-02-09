package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class WasteMethodsClient {
    private final MetrcClient client;

    public WasteMethodsClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Retrieves all available waste methods.
     * GET GetWasteMethodsWasteMethods
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getWasteMethodsWasteMethods(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/wastemethods/v2");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

