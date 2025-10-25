# IDP Templates

> **Software templates for scaffolding new services with the Impelligent IDP**

## Overview

This repository contains Backstage software templates and golden path definitions for creating new services, applications, and infrastructure components.

## Templates Available

### Application Templates

1. **Go Microservice** (`templates/go-service/`)
   - REST API with Gin framework
   - Dockerfile multi-stage build
   - Kubernetes manifests
   - Tekton pipeline
   - ArgoCD application definition

2. **React Frontend** (`templates/react-frontend/`)
   - TypeScript + Vite
   - Component library setup
   - Docker nginx deployment
   - CI/CD pipeline
   - Environment configuration

3. **Python API** (`templates/python-api/`)
   - FastAPI framework
   - Poetry dependency management
   - Docker deployment
   - Unit test setup
   - Pipeline configuration

### Infrastructure Templates

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

### Via Backstage (When Available)
1. Open Backstage UI
2. Click "Create" → "Choose a template"
3. Fill in the form
4. Template creates repo, pipeline, and deployment

### Manual Usage
```bash
# Clone template
git clone https://github.com/impelligent-idp/idp-templates.git
cd idp-templates/templates/go-service/skeleton

# Copy to new project
cp -r . /path/to/new-service

# Customize
# - Replace ${{values.name}} with your service name
# - Update dependencies
# - Customize configuration
```

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
