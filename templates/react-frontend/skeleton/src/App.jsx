import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const [appInfo, setAppInfo] = useState(null)

  useEffect(() => {
    setAppInfo({
      name: '${{values.name}}',
      description: '${{values.description}}',
      owner: '${{values.owner}}',
      platform: 'Impelligent IDP',
      version: '1.0.0'
    })
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        <h1>🚀 ${{values.name}}</h1>
        <p className="description">${{values.description}}</p>
        
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            Count is {count}
          </button>
        </div>

        {appInfo && (
          <div className="info-card">
            <h2>Application Info</h2>
            <dl>
              <dt>Name:</dt>
              <dd>{appInfo.name}</dd>
              
              <dt>Description:</dt>
              <dd>{appInfo.description}</dd>
              
              <dt>Owner:</dt>
              <dd>{appInfo.owner}</dd>
              
              <dt>Platform:</dt>
              <dd>{appInfo.platform}</dd>
              
              <dt>Version:</dt>
              <dd>{appInfo.version}</dd>
            </dl>
          </div>
        )}

        <div className="links">
          <a href="http://backstage.impelligent-idp.local:3001" target="_blank" rel="noopener noreferrer">
            📊 Backstage
          </a>
          <a href="http://argocd.impelligent-idp.local" target="_blank" rel="noopener noreferrer">
            🔄 ArgoCD
          </a>
          <a href="http://grafana.impelligent-idp.local" target="_blank" rel="noopener noreferrer">
            📈 Grafana
          </a>
        </div>
      </header>
    </div>
  )
}

export default App

