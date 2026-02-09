package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class SandboxClient {
    private final MetrcClient client;

    public SandboxClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * POST CreateIntegratorSetup
     * @param userKey Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSandboxIntegratorSetup(
        String userKey, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sandbox/v2/integrator/setup");
        StringBuilder query = new StringBuilder();
        if (userKey != null) {
            if (query.length() > 0) query.append("&");
            query.append("userKey=").append(URLEncoder.encode(userKey, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

}

