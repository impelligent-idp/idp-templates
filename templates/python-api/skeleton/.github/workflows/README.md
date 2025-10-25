# CI/CD Workflows

This directory contains GitHub Actions workflows for automated build and deployment.

## deploy.yaml

Automatically triggers on every push to `main` branch:

1. **Build Docker Image**: Builds the Python FastAPI application
2. **Push to GHCR**: Pushes to GitHub Container Registry (`ghcr.io`)
3. **Update Deployment**: Updates the Kubernetes manifest with the new image tag
4. **Deploy to Kubernetes**: Applies the ArgoCD application (if KUBECONFIG is configured)

### Image Tags

The workflow creates multiple tags:
- `latest` - always points to the most recent main branch build
- `main-<sha>` - specific commit SHA for traceability
- `main` - branch tag

### Setting up Auto-Deployment

To enable automatic Kubernetes deployment, add a `KUBECONFIG` secret to your repository:

```bash
# Export your kubeconfig as base64
cat ~/.kube/config | base64 | pbcopy

# Then add it as a GitHub Secret:
# Settings → Secrets and variables → Actions → New repository secret
# Name: KUBECONFIG
# Value: <paste the base64 content>
```

### Manual Deployment

If KUBECONFIG is not configured, the image will still be built and pushed. To deploy manually:

```bash
kubectl apply -f argocd/application.yaml
```

ArgoCD will then automatically sync and deploy your application.

## Permissions

The workflow requires:
- `contents: read` - to checkout code
- `packages: write` - to push to GitHub Container Registry

These are automatically provided by GitHub Actions.

