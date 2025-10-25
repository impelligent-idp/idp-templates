# ${{values.name}}

${{values.description}}

## Owner
${{values.owner}}

## Technology Stack
- **Framework**: React 18 with Vite
- **Server**: Nginx
- **Deployment**: Kubernetes with Ingress
- **CI/CD**: GitHub Actions
- **GitOps**: ArgoCD

## Local Development

```bash
# Install dependencies
npm install

# Run development server
npm run dev
```

Visit: http://localhost:3000

## Build

```bash
# Build for production
npm run build

# Preview production build
npm run preview
```

## Deployment

This application is automatically deployed via GitHub Actions and ArgoCD.

### Access Points
- **Application**: http://${{values.subdomain}}.impelligent-idp.local
- **Backstage**: [View in Catalog](http://backstage.impelligent-idp.local:3001/catalog/default/component/${{values.name}})
- **ArgoCD**: [View Deployment](http://argocd.impelligent-idp.local/applications/${{values.name}})

### Kubernetes Resources
- **Namespace**: `${{values.namespace}}`
- **Service**: `${{values.name}}`
- **Ingress**: `${{values.subdomain}}.impelligent-idp.local`

### CI/CD Pipeline
GitHub Actions automatically:
1. Builds the React application
2. Creates Docker image with Nginx
3. Pushes to GitHub Container Registry
4. Updates deployment manifest
5. ArgoCD syncs to Kubernetes

## Project Structure

```
.
├── src/
│   ├── App.jsx          # Main application component
│   ├── App.css          # Application styles
│   ├── main.jsx         # Entry point
│   └── index.css        # Global styles
├── k8s/
│   ├── deployment.yaml  # Kubernetes deployment
│   ├── service.yaml     # Kubernetes service
│   └── ingress.yaml     # Ingress configuration
├── argocd/
│   └── application.yaml # ArgoCD application
├── Dockerfile           # Multi-stage Docker build
├── nginx.conf           # Nginx configuration
└── vite.config.js       # Vite configuration
```

## Monitoring

This application is monitored by:
- **Prometheus**: Container metrics
- **Grafana**: Visualization
- **Loki**: Log aggregation

