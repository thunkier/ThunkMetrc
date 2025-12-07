package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/thunkmetrc/automation/pkg/bruno"
)

func generateRustWrapper(groups map[string][]bruno.Request, version string) {
	outDir := filepath.Join(outputDir, "rust", "wrapper")
	os.MkdirAll(filepath.Join(outDir, "src"), 0755)

	generateRustWrapperCargo(outDir, version)
	generateRustRateLimiter(outDir)
	generateRustWrapperCode(outDir, groups)
}

func generateRustRateLimiter(dir string) {
	content := `use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::{Mutex, Semaphore};
use std::time::{Duration, Instant};

#[derive(Clone)]
pub struct RateLimiterConfig {
    pub enabled: bool,
    pub max_get_per_second_per_facility: f64,
    pub max_get_per_second_integrator: f64,
    pub max_concurrent_get_per_facility: usize,
    pub max_concurrent_get_integrator: usize,
}

impl Default for RateLimiterConfig {
    fn default() -> Self {
        Self {
            enabled: false,
            max_get_per_second_per_facility: 50.0,
            max_get_per_second_integrator: 150.0,
            max_concurrent_get_per_facility: 10,
            max_concurrent_get_integrator: 30,
        }
    }
}

pub struct TokenBucket {
    rate: f64,
    capacity: f64,
    tokens: f64,
    last_refill: Instant,
}

impl TokenBucket {
    pub fn new(rate: f64, capacity: f64) -> Self {
        Self {
            rate,
            capacity,
            tokens: capacity,
            last_refill: Instant::now(),
        }
    }

    fn refill(&mut self) {
        let now = Instant::now();
        let elapsed = now.duration_since(self.last_refill).as_secs_f64();
        self.tokens = (self.tokens + elapsed * self.rate).min(self.capacity);
        self.last_refill = now;
    }

    pub fn try_consume(&mut self, amount: f64) -> Option<Duration> {
        self.refill();
        if self.tokens >= amount {
            self.tokens -= amount;
            None
        } else {
            let missing = amount - self.tokens;
            Some(Duration::from_secs_f64(missing / self.rate))
        }
    }
}

pub struct MetrcRateLimiter {
    config: RateLimiterConfig,
    integrator_rate: Mutex<TokenBucket>,
    integrator_sem: Arc<Semaphore>,
    facility_rates: Mutex<HashMap<String, Arc<Mutex<TokenBucket>>>>,
    facility_sems: Mutex<HashMap<String, Arc<Semaphore>>>,
}

impl MetrcRateLimiter {
    pub fn new(config: Option<RateLimiterConfig>) -> Self {
        let config = config.unwrap_or_default();
        Self {
            integrator_rate: Mutex::new(TokenBucket::new(config.max_get_per_second_integrator, config.max_get_per_second_integrator)),
            integrator_sem: Arc::new(Semaphore::new(config.max_concurrent_get_integrator)),
            facility_rates: Mutex::new(HashMap::new()),
            facility_sems: Mutex::new(HashMap::new()),
            config,
        }
    }

    pub async fn execute<F, Fut, T, E>(&self, facility: Option<&str>, is_get: bool, op: F) -> Result<T, E>
    where
        F: Fn() -> Fut,
        Fut: std::future::Future<Output = Result<T, E>>,
        E: std::fmt::Debug + std::fmt::Display, 
    {
        if !self.config.enabled || !is_get {
            return op().await;
        }

        // 1. Integrator Semaphore
        let _sem_permit = self.integrator_sem.acquire().await.map_err(|e| panic!("Semaphore closed: {:?}", e)).unwrap();

        // 2. Facility Semaphore
        let _fac_permit = if let Some(f) = facility {
             let sem = {
                 let mut sems = self.facility_sems.lock().await;
                 sems.entry(f.to_string())
                     .or_insert_with(|| Arc::new(Semaphore::new(self.config.max_concurrent_get_per_facility)))
                     .clone()
             };
             Some(sem.acquire_owned().await.map_err(|e| panic!("Semaphore closed: {:?}", e)).unwrap())
        } else {
             None
        };

        // 3. Global Rate
        loop {
            let mut bucket = self.integrator_rate.lock().await;
            if let Some(wait) = bucket.try_consume(1.0) {
                 drop(bucket);
                 tokio::time::sleep(wait).await;
            } else {
                 break;
            }
        }

        // 4. Facility Rate
        if let Some(f) = facility {
            let bucket_arc = {
                let mut rates = self.facility_rates.lock().await;
                rates.entry(f.to_string())
                    .or_insert_with(|| Arc::new(Mutex::new(TokenBucket::new(self.config.max_get_per_second_per_facility, self.config.max_get_per_second_per_facility))))
                    .clone()
            };
            loop {
                 let mut bucket = bucket_arc.lock().await;
                 if let Some(wait) = bucket.try_consume(1.0) {
                      drop(bucket);
                      tokio::time::sleep(wait).await;
                 } else {
                      break;
                 }
            }
        }

        // 5. Retry Loop
        loop {
             let res = op().await;
             match res {
                 Ok(v) => return Ok(v),
                 Err(e) => {
                      let msg = format!("{:?}", e);
                      if msg.contains("429") {
                           tokio::time::sleep(Duration::from_secs(1)).await;
                           continue;
                      }
                      return Err(e);
                 }
             }
        }
    }
}
`
	os.WriteFile(filepath.Join(dir, "src", "ratelimiter.rs"), []byte(content), 0644)
}

func generateRustWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	var modelsSb strings.Builder

	sb.WriteString("pub use thunkmetrc_client::MetrcClient;\n")
	sb.WriteString("use serde_json::Value;\n")
	sb.WriteString("use serde::{Deserialize, Serialize};\n")
	sb.WriteString("use std::error::Error;\n")
	sb.WriteString("use std::sync::Arc;\n")
	sb.WriteString("mod ratelimiter;\n")
	sb.WriteString("use ratelimiter::{MetrcRateLimiter, RateLimiterConfig};\n\n")

	sb.WriteString("pub struct MetrcWrapper {\n")
	sb.WriteString("    client: MetrcClient,\n")
	sb.WriteString("    rate_limiter: Arc<MetrcRateLimiter>,\n")
	sb.WriteString("}\n\n")

	sb.WriteString("impl MetrcWrapper {\n")
	sb.WriteString("    pub fn new(client: MetrcClient) -> Self {\n")
	sb.WriteString("        MetrcWrapper {\n")
	sb.WriteString("            client,\n")
	sb.WriteString("            rate_limiter: Arc::new(MetrcRateLimiter::new(Some(RateLimiterConfig::default()))),\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n\n")

	sb.WriteString("    pub fn new_with_config(client: MetrcClient, config: RateLimiterConfig) -> Self {\n")
	sb.WriteString("        MetrcWrapper {\n")
	sb.WriteString("            client,\n")
	sb.WriteString("            rate_limiter: Arc::new(MetrcRateLimiter::new(Some(config))),\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n\n")

	for group, reqs := range groups {
		safeGroup := ToSnakeCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToSnakeCase(cleanName(req.Name))
			rustMethodName := fmt.Sprintf("%s_%s", safeGroup, methodName)

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s: &str", ToSnakeCase(p)))
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				rustParamName := ToSnakeCase(paramName)
				argsList = append(argsList, fmt.Sprintf("%s: Option<String>", rustParamName))
			}

			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					pascalGroup := ToPascalCase(cleanName(group))
					structName := fmt.Sprintf("%s%sRequest", pascalGroup, ToPascalCase(cleanName(req.Name)))
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						structName = fmt.Sprintf("%sItem", structName)
						generateRustStruct(structName, targetSchema, &modelsSb)

						argsList = append(argsList, fmt.Sprintf("body: Option<&Vec<%s>>", structName))
					} else {
						generateRustStruct(structName, targetSchema, &modelsSb)
						argsList = append(argsList, fmt.Sprintf("body: Option<&%s>", structName))
					}
				} else {
					argsList = append(argsList, "body: Option<&Value>")
				}
			} else {
				argsList = append(argsList, "body: Option<&Value>")
			}

			sigArgs := strings.Join(argsList, ", ")
			if len(sigArgs) > 0 {
				sigArgs = "&self, " + sigArgs
			} else {
				sigArgs = "&self"
			}

			// Build call args
			var ownedCallArgs []string
			for _, p := range pathParams {
				ownedCallArgs = append(ownedCallArgs, fmt.Sprintf("&%s", ToSnakeCase(p)))
			}
			for _, k := range sortedKeys {
				// Query params are passed as owned Option<String>
				paramName := strings.TrimSuffix(k, "optional")
				ownedCallArgs = append(ownedCallArgs, ToSnakeCase(paramName))
			}

			valConversion := ""
			if req.Method == "POST" || req.Method == "PUT" && req.BodySchema != nil {
				valConversion = fmt.Sprintf("\n        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };")
				ownedCallArgs = append(ownedCallArgs, "body_val.as_ref()")
			} else {
				ownedCallArgs = append(ownedCallArgs, "body.as_ref()")
			}
			ownedCallArgsStr := strings.Join(ownedCallArgs, ", ")

			// Determine Facility
			facilityArg := "None"
			for _, p := range pathParams {
				if strings.ToLower(cleanName(p)) == "licensenumber" {
					facilityArg = fmt.Sprintf("Some(%s)", ToSnakeCase(p))
					break
				}
			}

			isGet := "false"
			if req.Method == "GET" {
				isGet = "true"
			}

			// Client call closure construction
			// We need to clone arguments for the async block or move them?
			// self.client needs to be cloned? MetrcClient is lightweight (Arc internal usually? or struct with fields).
			// Checking rust_gen: MetrcClient struct fields are Strings. Clone needed?
			// MetrcClient should be Clone.

			// To be safe, we might need to clone self.client into the closure.
			// And clone string args.

			cloneBlock := ""
			var cloneArgs []string
			for _, p := range pathParams {
				pName := ToSnakeCase(p)
				cloneBlock += fmt.Sprintf("        let %s = %s.to_string();\n", pName, pName)
				cloneArgs = append(cloneArgs, pName) // Use the owned String
			}
			if req.Method == "POST" || req.Method == "PUT" && req.BodySchema != nil {
				// body_val needs to be cloned? Value is Clone.
			}

			// Simplified: The client methods take &str? No, previously I saw client generated with String.
			// Wait, step 1198 lines 104: callArgs append ToSnakeCase(p).
			// If client takes &str, passing String is mismatch (need &). If client takes String, need to_string().
			// I added `.to_string()` in line 242 above.

			// To pass into async closure, we need owned data.
			// So we convert to String OUTSIDE closure, then move into closure.

			sb.WriteString(fmt.Sprintf("    /// %s %s\n", req.Method, req.Name))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("    /// %s\n", line))
				}
				sb.WriteString("    ///\n")
			}
			sb.WriteString(fmt.Sprintf("    pub async fn %s(%s) -> Result<Option<Value>, Box<dyn Error>> {%s\n", rustMethodName, sigArgs, valConversion))
			sb.WriteString(cloneBlock) // Creates owned strings

			sb.WriteString(fmt.Sprintf("        let client = self.client.clone();\n"))
			// body_val ownership?
			if req.Method == "POST" || req.Method == "PUT" && req.BodySchema != nil {
				sb.WriteString("        let body_val = body_val.clone();\n") // Clone Value
			} else {
				sb.WriteString("        let body = body.cloned();\n") // Clone Option<&Value> -> Option<Value> ?
				// Wait, body is Option<&Value>. We need owned Value for async move?
				// Or we keep it as reference if scope allows?
				// Since execute takes future, future must live long. References are hard.
				// Better to Clone inputs to owned versions.
				// So body should be Option<Value>.
			}

			// Fix call args to use owned

			sb.WriteString(fmt.Sprintf("        self.rate_limiter.execute(%s, %s, move || {\n", facilityArg, isGet))
			sb.WriteString("            let client = client.clone();\n")
			// Re-clone vars if Fn called multiple times?
			for _, p := range pathParams {
				sb.WriteString(fmt.Sprintf("            let %s = %s.clone();\n", ToSnakeCase(p), ToSnakeCase(p)))
			}
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				rustParamName := ToSnakeCase(paramName)
				sb.WriteString(fmt.Sprintf("            let %s = %s.clone();\n", rustParamName, rustParamName))
			}
			if req.Method == "POST" || req.Method == "PUT" && req.BodySchema != nil {
				sb.WriteString("            let body_val = body_val.clone();\n")
			} else {
				sb.WriteString("            let body = body.clone();\n")
			}

			sb.WriteString(fmt.Sprintf("            async move { client.%s(%s).await }\n", rustMethodName, ownedCallArgsStr))
			sb.WriteString("        }).await\n")
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString("}\n\n")
	sb.WriteString(modelsSb.String())

	os.WriteFile(filepath.Join(dir, "src", "lib.rs"), []byte(sb.String()), 0644)
}
func generateRustWrapperCargo(dir string, version string) {
	var cargoToml = fmt.Sprintf(`[package]
name = "thunkmetrc-wrapper"
version = "%s"
edition = "2021"

[dependencies]
thunkmetrc-client = { path = "../client", version = "%s" }
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
tokio = { version = "1", features = ["full"] }
`, version, version)
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(cargoToml), 0644)
}

func generateRustStruct(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		sb.WriteString(fmt.Sprintf("#[derive(Serialize, Deserialize, Debug, Clone)]\n"))
		sb.WriteString(fmt.Sprintf("pub struct %s {\n", name))

		nested := make(map[string]*bruno.Schema)

		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, propName := range keys {
			propSchema := schema.Properties[propName]
			rustField := ToSnakeCase(propName)
			if rustField == "type" {
				rustField = "type_"
			}

			rsType, nestedSchema, nestedName := resolveRustType(propSchema, name, ToPascalCase(cleanName(propName)))
			if nestedSchema != nil {
				nested[nestedName] = nestedSchema
			}

			sb.WriteString(fmt.Sprintf("    #[serde(rename = \"%s\")]\n", propName))
			sb.WriteString(fmt.Sprintf("    pub %s: Option<%s>,\n", rustField, rsType))
		}
		sb.WriteString("}\n\n")

		var nestedKeys []string
		for k := range nested {
			nestedKeys = append(nestedKeys, k)
		}
		sort.Strings(nestedKeys)

		for _, subName := range nestedKeys {
			generateRustStruct(subName, nested[subName], sb)
		}
	}
}

func resolveRustType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "String", nil, ""
	case bruno.TypeInt:
		return "i64", nil, ""
	case bruno.TypeFloat:
		return "f64", nil, ""
	case bruno.TypeBool:
		return "bool", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "Value", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolveRustType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("Vec<%s>", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subName := fmt.Sprintf("%s%s", parentName, propName)
		return subName, s, subName
	default:
		return "Value", nil, ""
	}
}
