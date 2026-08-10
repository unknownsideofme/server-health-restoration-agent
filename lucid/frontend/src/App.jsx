import React, { useState, useEffect } from 'react';

export default function App() {
  const [orgFilter, setOrgFilter] = useState('ALL');
  const [categoryFilter, setCategoryFilter] = useState('ALL');
  const [telemetry, setTelemetry] = useState(null);
  const [copilotText, setCopilotText] = useState('Select any component card or prompt local LLM daemon...');
  const [userPrompt, setUserPrompt] = useState('');
  const [lastSync, setLastSync] = useState('Syncing...');
  const [orgOrder, setOrgOrder] = useState(['org-a', 'org-b', 'org-c', 'org-d']);

  useEffect(() => {
    fetchTelemetryData();
    const interval = setInterval(fetchTelemetryData, 2500);
    return () => clearInterval(interval);
  }, []);

  const fetchTelemetryData = async () => {
    try {
      const res = await fetch('/api/telemetry');
      const data = await res.json();
      setTelemetry(data);
      setLastSync(new Date().toLocaleTimeString());
    } catch (err) {
      console.error('Failed to fetch telemetry data:', err);
    }
  };

  const handleSelectComponent = async (name) => {
    setCopilotText(`Querying Local LLM Daemon for component '${name}'...`);
    try {
      const res = await fetch(`/api/copilot?target=${name}`);
      const data = await res.json();
      const out = `=== AIRGAP AI NOC COPILOT DECISION SUPPORT (44 Components) ===
Target Component: ${data.component} (Status: ${data.status} | Risk: ${data.risk_score}%)

Q1: WHAT IS LIKELY TO FAIL NEXT — AND WHEN?
${data.q1_forecast}

Q2: WHY IS RISK ASSESSED AS ELEVATED — WHICH SIGNALS CONTRIBUTED?
${data.q2_reasoning}

Q3: WHAT CORRECTIVE ACTION SHOULD BE TAKEN BEFORE SLA IMPACT?
${data.q3_corrective_action}

LLM INSIGHT (qwen2.5:0.5b Daemon Port 11435):
${data.llm_insight}`;
      setCopilotText(out);
    } catch (e) {
      setCopilotText("Error fetching Copilot decision support.");
    }
  };

  const handleSendPrompt = async () => {
    if (!userPrompt) return;
    const promptCopy = userPrompt;
    setUserPrompt('');
    setCopilotText('Querying Dedicated Local LLM Daemon (Port 11435)...');
    try {
      const res = await fetch('/api/copilot/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ prompt: promptCopy })
      });
      const data = await res.json();
      setCopilotText(data.answer);
    } catch (e) {
      setCopilotText("Error querying local LLM daemon.");
    }
  };

  const handleInjectFault = async (scenario) => {
    await fetch('/api/fault/inject', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ scenario: scenario, target: 'org-a-edge-router' })
    });
    fetchTelemetryData();
    handleSelectComponent('org-a-edge-router');
  };

  const handleMCPFix = async (target = 'org-a-edge-router') => {
    setCopilotText(`Executing MCP Skill Autonomous Repair on ${target}...`);
    try {
      const res = await fetch('/api/mcp/call', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ tool_name: 'execute_autonomous_repair', arguments: { component_name: target } })
      });
      const data = await res.json();
      const resultObj = data.result ? data.result.result : {};
      const actions = resultObj.actions ? resultObj.actions.join('\n- ') : 'Autonomous repair completed.';
      setCopilotText(`=== MCP AGENT AUTONOMOUS REPAIR RESULT ===\nTarget: ${target}\nStatus: COMPLETED\n\nActions Taken:\n- ` + actions);
      fetchTelemetryData();
    } catch (e) {
      setCopilotText("Error executing MCP repair skill.");
    }
  };

  const handleClearFaults = async () => {
    try {
      await fetch('/api/fault/clear', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({})
      });
      fetchTelemetryData();
    } catch (e) {
      console.error('Failed to clear faults:', e);
    }
  };

  // HTML5 Drag and Drop Handlers
  const handleDragStart = (e, index) => {
    e.dataTransfer.setData('draggedIndex', index);
  };

  const handleDrop = (e, dropIndex) => {
    const dragIndex = e.dataTransfer.getData('draggedIndex');
    if (dragIndex === null || dragIndex === undefined || dragIndex === "") return;
    const dragIdxInt = parseInt(dragIndex, 10);
    if (isNaN(dragIdxInt)) return;
    const newOrder = [...orgOrder];
    const [movedItem] = newOrder.splice(dragIdxInt, 1);
    newOrder.splice(dropIndex, 0, movedItem);
    setOrgOrder(newOrder);
  };

  return (
    <div>
      <header>
        <div className="brand">
          <div className="brand-icon">React</div>
          <div>
            <div className="brand-title">Autonomous Air-Gapped NOC Copilot (44 Components)</div>
            <div style={{ fontSize: '11px', color: 'var(--text-muted)' }}>
              Live Traffic Metrics (RPS, MB/s, TCP) & Interactive HTML5 Drag-and-Drop Custom Layout
            </div>
          </div>
        </div>
        <div className="live-badge">
          <div className="live-dot"></div>
          <span>Live Telemetry ● {lastSync}</span>
        </div>
      </header>

      {/* Navigation & Quick Toggle Filters */}
      <div className="nav-bar">
        <div className="nav-row">
          <div style={{ fontSize: '12px', fontWeight: '700', color: 'var(--accent-cyan)' }}>
            ⚡ QUICK CATEGORY FILTERS
          </div>
          <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
            <span style={{ fontSize: '12px', color: 'var(--text-muted)' }}>Scope Filter:</span>
            <select value={orgFilter} onChange={(e) => setOrgFilter(e.target.value)}>
              <option value="ALL">All Organizations</option>
              <option value="org-a">Org A (10.10.0.0/24)</option>
              <option value="org-b">Org B (10.11.0.0/24)</option>
              <option value="org-c">Org C (10.12.0.0/24)</option>
              <option value="org-d">Org D (10.13.0.0/24)</option>
            </select>
          </div>
        </div>

        <div className="quick-filters">
          <div
            className={`filter-btn ${categoryFilter === 'ALL' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('ALL')}
          >
            🌐 ALL (44 Components)
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'admin' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('admin')}
          >
            🛡️ ADMIN CONTROL PLANE
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'worker' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('worker')}
          >
            ⚙️ WORKLOAD COMPUTE NODES
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'router' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('router')}
          >
            🌐 EDGE ROUTERS
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'tor' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('tor')}
          >
            🔀 TOR SWITCHES
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'rack' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('rack')}
          >
            🗄️ HARDWARE RACKS
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'subnet' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('subnet')}
          >
            📡 CIDR SUBNETS
          </div>
          <div
            className={`filter-btn ${categoryFilter === 'service' ? 'active' : ''}`}
            onClick={() => setCategoryFilter('service')}
          >
            ⚡ MICROSERVICES & APPS
          </div>
        </div>
      </div>

      <div className="main-grid">
        <div>
          <div className="card">
            <div className="card-title">
              <span>Interactive Topology Grid (Drag cards to reorder layout)</span>
              <span style={{ fontSize: '11px', color: 'var(--text-muted)' }}>Showing 44 live components</span>
            </div>

            <InteractiveOrgGrid
              orgOrder={orgOrder}
              orgFilter={orgFilter}
              categoryFilter={categoryFilter}
              telemetry={telemetry}
              onSelect={handleSelectComponent}
              onDragStart={handleDragStart}
              onDrop={handleDrop}
            />
          </div>

          <div className="card">
            <div className="card-title">
              <span>AI Copilot Decision Support (Q1, Q2, Q3)</span>
              <span
                style={{
                  fontSize: '11px',
                  padding: '2px 8px',
                  borderRadius: '10px',
                  background: 'rgba(6,182,212,0.2)',
                  color: 'var(--accent-cyan)',
                }}
              >
                Local LLM Daemon (Port 11435)
              </span>
            </div>
            <div className="copilot-output">{copilotText}</div>
            <div className="chat-input-row">
              <input
                type="text"
                value={userPrompt}
                onChange={(e) => setUserPrompt(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && handleSendPrompt()}
                placeholder="Ask NOC Copilot (e.g. 'What is likely to fail next?')"
              />
              <button className="btn" onClick={handleSendPrompt}>
                Ask Copilot
              </button>
            </div>
          </div>
        </div>

        <div>
          <div className="card">
            <div className="card-title">Fault Injection Control Panel</div>
            <div className="btn-group">
              <button className="btn btn-danger" onClick={() => handleInjectFault('congestion')}>
                <span>⚡ Progressive Congestion</span>
                <span style={{ fontSize: '10px' }}>(org-a-edge-router)</span>
              </button>
              <button className="btn btn-danger" onClick={() => handleInjectFault('route_flap')}>
                <span>⚡ BGP Route Flap Cascade</span>
                <span style={{ fontSize: '10px' }}>(tor-a1)</span>
              </button>
              <button className="btn btn-danger" onClick={() => handleInjectFault('tunnel_deg')}>
                <span>⚡ Tunnel Degradation</span>
                <span style={{ fontSize: '10px' }}>(org-a-edge-router)</span>
              </button>
              <button className="btn btn-danger" onClick={() => handleInjectFault('policy_drift')}>
                <span>⚡ Policy Drift</span>
                <span style={{ fontSize: '10px' }}>(org-a-edge-router)</span>
              </button>
              <button
                className="btn btn-clear"
                style={{ borderColor: 'rgba(6,182,212,0.5)', background: 'rgba(6,182,212,0.15)' }}
                onClick={() => handleMCPFix()}
              >
                <span>🛠️ Execute Autonomous Repair (MCP Skill)</span>
              </button>
              <button className="btn btn-clear" onClick={handleClearFaults}>
                <span>✓ Clear All Active Faults</span>
              </button>
            </div>
          </div>

          <div className="card">
            <div className="card-title">Active Predictive Alerts</div>
            <div style={{ fontSize: '12px' }}>
              {telemetry && telemetry.alerts && telemetry.alerts.length > 0 ? (
                telemetry.alerts.map((al, idx) => (
                  <div
                    key={idx}
                    style={{
                      padding: '8px',
                      marginBottom: '6px',
                      borderRadius: '6px',
                      background: 'rgba(239, 68, 68, 0.15)',
                      border: '1px solid rgba(239, 68, 68, 0.3)',
                    }}
                  >
                    <div style={{ fontWeight: 600, color: '#ef4444' }}>🚨 {al.issue_type}</div>
                    <div style={{ fontSize: '11px', marginTop: '2px' }}>{al.summary}</div>
                    <div style={{ fontSize: '10px', color: '#9ca3af', marginTop: '2px' }}>
                      TTI: {al.tti_minutes} mins | Conf: {Math.round(al.confidence * 100)}%
                    </div>
                  </div>
                ))
              ) : (
                <div style={{ color: 'var(--text-muted)', fontStyle: 'italic' }}>
                  No critical predictive alerts active.
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

function InteractiveOrgGrid({
  orgOrder,
  orgFilter,
  categoryFilter,
  telemetry,
  onSelect,
  onDragStart,
  onDrop,
}) {
  if (!telemetry || !telemetry.by_org) return <div>Loading live 44-component telemetry matrix...</div>;

  const by_org = telemetry.by_org;
  const filteredKeys = orgOrder.filter((k) => orgFilter === 'ALL' || k === orgFilter);

  return (
    <div>
      {filteredKeys.map((orgKey, idx) => {
        const orgData = by_org[orgKey];
        if (!orgData) return null;

        return (
          <div
            key={orgKey}
            className="org-section-card"
            draggable={true}
            onDragStart={(e) => onDragStart(e, idx)}
            onDragOver={(e) => e.preventDefault()}
            onDrop={(e) => onDrop(e, idx)}
          >
            <div className="org-header">
              <div className="org-title">
                <span className="drag-handle">⋮⋮</span>
                <span>🏢 {orgData.name}</span>
              </div>
              <div className="org-badge">
                Subnet: 10.{10 + ['org-a', 'org-b', 'org-c', 'org-d'].indexOf(orgKey)}.0.0/24
              </div>
            </div>

            <div className="sub-grid">
              {(categoryFilter === 'ALL' || categoryFilter === 'router' || categoryFilter === 'subnet') && (
                <div className="sub-section">
                  <div className="sub-title">🌐 Edge Router & Subnets</div>
                  <RichComponentList items={orgData.routers.concat(orgData.subnets)} onSelect={onSelect} />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'tor' || categoryFilter === 'rack') && (
                <div className="sub-section">
                  <div className="sub-title">🔀 TOR Switches & Racks</div>
                  <RichComponentList items={orgData.tors.concat(orgData.racks)} onSelect={onSelect} />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'admin') && (
                <div className="sub-section">
                  <div className="sub-title">🛡️ Admin Control Plane</div>
                  <RichComponentList items={orgData.admin_servers} onSelect={onSelect} />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'worker') && (
                <div className="sub-section">
                  <div className="sub-title">⚙️ Workload Compute</div>
                  <RichComponentList items={orgData.worker_servers} onSelect={onSelect} />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'service') &&
                orgData.services &&
                orgData.services.length > 0 && (
                  <div className="sub-section">
                    <div className="sub-title">⚡ Microservices & Apps</div>
                    <RichComponentList items={orgData.services} onSelect={onSelect} />
                  </div>
                )}
            </div>
          </div>
        );
      })}
    </div>
  );
}

function RichComponentList({ items, onSelect }) {
  if (!items || items.length === 0) {
    return <div style={{ fontSize: '11px', color: 'var(--text-muted)', fontStyle: 'italic' }}>None</div>;
  }

  return (
    <div>
      {items.map((item) => {
        const tti_html = item.tti_minutes ? (
          <span style={{ color: '#ef4444', fontWeight: 'bold' }}>TTI: {item.tti_minutes}m</span>
        ) : (
          'Nominal'
        );
        const m = item.current_metrics || {};
        const util = m.interface_utilization_pct || 25.0;
        const rps = m.http_requests_per_sec || 15.0;
        const mbps = m.network_traffic_bytes_sec
          ? (m.network_traffic_bytes_sec / (1024 * 1024)).toFixed(1)
          : 0.5;
        const tcp = m.active_tcp_connections || 45;
        const err = m.error_rate_pct || 0.0;

        return (
          <div
            key={item.name}
            className={`topo-item status-${item.status}`}
            onClick={() => onSelect(item.name)}
          >
            <div className="item-header">
              <span className="item-name">{item.name}</span>
              <span className={`status-pill ${item.status}`}>{item.status}</span>
            </div>
            <div className="metric-row">
              <span>Risk / Util:</span>
              <span className="metric-val">
                {item.risk_score || 15}% / {util}%
              </span>
            </div>
            <div className="metric-row">
              <span>Traffic RPS / MBps:</span>
              <span className="metric-val">
                {rps} rps / {mbps} MB/s
              </span>
            </div>
            <div className="metric-row">
              <span>Active TCP / Err:</span>
              <span className="metric-val">
                {tcp} / {err}%
              </span>
            </div>
            <div className="metric-row">
              <span>Lead Time (TTI):</span>
              <span className="metric-val">{tti_html}</span>
            </div>
          </div>
        );
      })}
    </div>
  );
}
