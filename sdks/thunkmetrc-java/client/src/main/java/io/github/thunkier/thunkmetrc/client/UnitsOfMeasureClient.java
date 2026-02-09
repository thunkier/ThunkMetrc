package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class UnitsOfMeasureClient {
    private final MetrcClient client;

    public UnitsOfMeasureClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Retrieves all active units of measure.
     * GET GetActiveUnitsOfMeasure
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getActiveUnitsOfMeasure(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/unitsofmeasure/v2/active");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves all inactive units of measure.
     * GET GetInactiveUnitsOfMeasure
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getInactiveUnitsOfMeasure(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/unitsofmeasure/v2/inactive");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

