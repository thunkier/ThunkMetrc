import * as dotenv from "dotenv";
import * as path from "path";

// Load .env from repo root
dotenv.config({ path: path.resolve(__dirname, "../../../../.env") });
