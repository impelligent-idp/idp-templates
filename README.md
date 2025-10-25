# IDP Templates

> **Software templates for scaffolding new services with the Impelligent IDP**

## Overview

This repository contains Backstage software templates and golden path definitions for creating new services, applications, and infrastructure components.

## Templates Available (v1.0.0)

### Application Templates ✅

1. **Go Microservice** (`templates/go-service/`) ✅ **READY**
   - REST API with Gin framework
   - Dockerfile multi-stage build
   - Kubernetes manifests
   - Tekton pipeline
   - ArgoCD application definition
   - Full catalog-info.yaml

2. **Python API** (`templates/python-api/`) ✅ **READY**
   - FastAPI framework
   - Requirements.txt dependency management
   - Docker deployment
   - Kubernetes manifests
   - Backstage catalog integration

### Planned Templates (v1.1.0+) 🔜

3. **React Frontend** (`templates/react-frontend/`)
   - TypeScript + Vite
   - Component library setup
   - Docker nginx deployment
   - CI/CD pipeline
   - Environment configuration

4. **Database Setup** (`templates/database/`)
   - PostgreSQL configuration
   - Migration scripts
   - Backup setup

5. **Message Queue** (`templates/message-queue/`)
   - Kafka/RabbitMQ setup
   - Consumer/Producer patterns

## Repository Structure

```
idp-templates/
├── README.md
├── catalog-info.yaml           # Backstage catalog definition
├── templates/
│   ├── go-service/
│   │   ├── template.yaml       # Backstage template definition
│   │   ├── skeleton/           # Template files
│   │   │   ├── src/
│   │   │   ├── Dockerfile
│   │   │   ├── k8s/
│   │   │   │   ├── deployment.yaml
│   │   │   │   └── service.yaml
│   │   │   └── tekton/
│   │   │       └── pipeline.yaml
│   │   └── README.md
│   ├── react-frontend/
│   ├── python-api/
│   ├── database/
│   └── message-queue/
└── docs/
    ├── creating-templates.md
    └── template-variables.md
```

## Template Variables

All templates support these common variables:

```yaml
${{values.name}}              # Service name
${{values.owner}}             # Team/owner
${{values.description}}       # Service description
${{values.repoUrl}}           # GitHub repository URL
${{values.namespace}}         # Kubernetes namespace
```

## Usage

### v1.0.0: Manual Template Usage

For v1.0.0, use templates manually by copying skeleton directories:

```bash
# Clone templates repo
git clone https://github.com/impelligent-idp/idp-templates.git
cd idp-templates

# Option 1: Go Microservice
cp -r templates/go-service/skeleton ~/git/my-new-service
cd ~/git/my-new-service

# Customize:
# 1. Replace service name in all files
# 2. Update go.mod with your module path
# 3. Customize main.go, Dockerfile, k8s/, tekton/, argocd/
# 4. Update catalog-info.yaml with your info
# 5. git init && git add . && git commit -m "Initial commit"
# 6. git push to your repo
# 7. Deploy via ArgoCD

# Option 2: Python API
cp -r templates/python-api/skeleton ~/git/my-new-api
# Follow similar steps
```

### v1.1.0+: Via Backstage UI (Planned) 🔜

Once Backstage UI is fully operational:
1. Open Backstage at http://backstage.impelligent-idp.local
2. Click "Create" → "Choose a template"
3. Fill in the form (name, owner, description)
4. Click "Create" - template scaffolds everything automatically
5. Repo, pipeline, and ArgoCD app created automatically

## Creating New Templates

See [docs/creating-templates.md](docs/creating-templates.md) for guide.

## Integration with IDP

These templates are automatically connected to:
- ✅ ArgoCD (for GitOps deployment)
- ✅ Tekton (for CI/CD pipelines)
- 🔜 Backstage (for UI-based scaffolding)

## Contributing

Add new templates by:
1. Creating a new directory in `templates/`
2. Adding `template.yaml` (Backstage definition)
3. Creating `skeleton/` with template files
4. Documenting variables and usage
5. Testing template generation

---

**Part of**: [Impelligent IDP](https://github.com/impelligent-idp/impelligent-idp)
