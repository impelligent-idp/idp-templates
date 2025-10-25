# ${{values.name}}

> ${{values.description}}

**Owner**: ${{values.owner}}  
**Namespace**: `${{values.namespace}}`  
**Port**: `${{values.port}}`

## Quick Start

```bash
# Run locally
go run main.go

# Build
go build -o ${{values.name}} .

# Run tests
go test ./...

# Build Docker image
docker build -t ${{values.name}}:latest .

# Run in Docker
docker run -p ${{values.port}}:${{values.port}} ${{values.name}}:latest
```

## API Endpoints

- `GET /health` - Health check
- `GET /ready` - Readiness check
- `GET /info` - Service information
- `GET /api/v1/hello` - Example API endpoint

## Deployment

This service is automatically deployed via ArgoCD when changes are pushed to the `main` branch.

**ArgoCD Application**: https://localhost:8080/applications/${{values.name}}

## CI/CD Pipeline

Tekton pipeline runs on every push:
1. **fetch-source** - Clone repository
2. **run-tests** - Execute Go tests
3. **build-image** - Build and push Docker image
4. **deploy** - Apply Kubernetes manifests

View pipeline runs:
```bash
tkn pipelinerun list -n ${{values.namespace}}
tkn pipelinerun logs --last -f -n ${{values.namespace}}
```

## Monitoring

- **Metrics**: Prometheus scrapes `/metrics` endpoint
- **Logs**: Aggregated in Grafana via Loki
- **Dashboards**: Available in Grafana at http://localhost:3000

## Development

```bash
# Install dependencies
go mod download

# Run with hot reload (install air first: go install github.com/cosmtrek/air@latest)
air

# Format code
go fmt ./...

# Lint
golangci-lint run
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Service port | `${{values.port}}` |
| `SERVICE_NAME` | Service identifier | `${{values.name}}` |

## Links

- **Repository**: https://github.com/${{values.destination.owner}}/${{values.destination.repo}}
- **ArgoCD**: https://localhost:8080/applications/${{values.name}}
- **Grafana**: http://localhost:3000
- **Tekton**: http://localhost:9097

---

Generated from [Impelligent IDP Go Microservice Template](https://github.com/impelligent-idp/idp-templates)

