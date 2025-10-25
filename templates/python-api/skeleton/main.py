"""
${{values.name}} - ${{values.description}}
"""

import os
from datetime import datetime
from fastapi import FastAPI
from fastapi.responses import JSONResponse
import uvicorn

app = FastAPI(
    title="${{values.name}}",
    description="${{values.description}}",
    version="1.0.0"
)

@app.get("/health")
async def health():
    """Health check endpoint"""
    return {
        "status": "healthy",
        "service": "${{values.name}}",
        "version": "1.0.0",
        "timestamp": datetime.now().isoformat()
    }

@app.get("/ready")
async def ready():
    """Readiness check endpoint"""
    return {"status": "ready"}

@app.get("/info")
async def info():
    """Service information"""
    return {
        "name": "${{values.name}}",
        "description": "${{values.description}}",
        "version": "1.0.0",
        "owner": "${{values.owner}}"
    }

@app.get("/api/v1/hello")
async def hello():
    """Example API endpoint"""
    return {
        "message": f"Hello from ${{values.name}}",
        "version": "1.0.0"
    }

if __name__ == "__main__":
    port = int(os.getenv("PORT", "${{values.port}}"))
    uvicorn.run(app, host="0.0.0.0", port=port)

