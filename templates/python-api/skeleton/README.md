# ${{values.name}}

${{values.description}}

## Owner
${{values.owner}}

## Technology Stack
- **Language**: Python 3.11
- **Framework**: FastAPI
- **Deployment**: Kubernetes
- **CI/CD**: Tekton Pipelines
- **GitOps**: ArgoCD

## Local Development

```bash
# Install dependencies
pip install -r requirements.txt

# Run locally
uvicorn main:app --reload --port ${{values.port}}
```

Visit: http://localhost:${{values.port}}

## API Documentation

Once running, visit:
- Swagger UI: http://localhost:${{values.port}}/docs
- ReDoc: http://localhost:${{values.port}}/redoc

## Deployment

This service is automatically deployed to Kubernetes via ArgoCD when changes are pushed to the `main` branch.

### Kubernetes Resources
- **Namespace**: `${{values.namespace}}`
- **Service**: `${{values.name}}`
- **Port**: `${{values.port}}`

### CI/CD Pipeline
The Tekton pipeline automatically:
1. Clones the repository
2. Runs tests
3. Builds Docker image
4. Pushes to GitHub Container Registry
5. Deploys to Kubernetes

## Monitoring

This service is automatically monitored by:
- **Prometheus**: Metrics collection
- **Grafana**: Dashboards and visualization
- **Loki**: Log aggregation

## Links
- [ArgoCD Application](http://argocd.impelligent-idp.local/applications/${{values.name}})
- [Backstage Catalog](http://backstage.impelligent-idp.local/catalog/default/component/${{values.name}})

