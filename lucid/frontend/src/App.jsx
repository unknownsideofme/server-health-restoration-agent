/* eslint-disable react/prop-types */
import { useState, useEffect } from 'react';

// Custom SVG Icons for LUCID NOC UI
const LucidLogo = () => (
  <svg className="lucid-logo-svg" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="lucid-grad" x1="0%" y1="0%" x2="100%" y2="100%">
        <stop offset="0%" stopColor="#0ea5e9" />
        <stop offset="100%" stopColor="#6366f1" />
      </linearGradient>
    </defs>
    <path d="M12 2L2 22h20L12 2z" stroke="url(#lucid-grad)" strokeWidth="2" strokeLinejoin="round" />
    <circle cx="12" cy="11" r="3" fill="url(#lucid-grad)" />
    <line x1="12" y1="14" x2="7" y2="21" stroke="#0ea5e9" strokeWidth="1.5" />
    <line x1="12" y1="14" x2="17" y2="21" stroke="#6366f1" strokeWidth="1.5" />
  </svg>
);

const IconShield = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
  </svg>
);

const IconCpu = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <rect x="4" y="4" width="16" height="16" rx="2" />
    <path d="M9 9h6v6H9zm0-5V2m6 2V2M9 22v-2m6 2v-2M20 9h2m-2 6h2M2 9h2m-2 6h2" />
  </svg>
);

const IconRouter = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <circle cx="12" cy="12" r="3" />
    <circle cx="4" cy="12" r="2" />
    <circle cx="20" cy="12" r="2" />
    <circle cx="12" cy="4" r="2" />
    <circle cx="12" cy="20" r="2" />
    <path d="M6 12h3M15 12h3M12 6v3M12 15v3" />
  </svg>
);

const IconSwitch = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <rect x="2" y="5" width="20" height="14" rx="2" />
    <path d="M7 10h10M17 14H7" />
    <path d="M14 8l3 2-3 2M10 12l-3 2 3 2" />
  </svg>
);

const IconRack = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <rect x="2" y="3" width="20" height="5" rx="1" />
    <rect x="2" y="10" width="20" height="5" rx="1" />
    <rect x="2" y="17" width="20" height="5" rx="1" />
    <line x1="6" y1="5.5" x2="6.01" y2="5.5" strokeLinecap="round" strokeWidth="3" />
    <line x1="6" y1="12.5" x2="6.01" y2="12.5" strokeLinecap="round" strokeWidth="3" />
    <line x1="6" y1="19.5" x2="6.01" y2="19.5" strokeLinecap="round" strokeWidth="3" />
  </svg>
);

const IconSubnet = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M4 9h16M4 15h16M9 4v16M15 4v16" />
  </svg>
);

const IconService = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M21 16V8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16z" />
    <path d="M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12" />
  </svg>
);

const IconOrg = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M3 21h18M3 7V5a2 2 0 012-2h14a2 2 0 012 2v2M5 21V7m14 11V7M9 7h6M9 11h6M9 15h6" />
  </svg>
);

const IconAlert = ({ size = 14, className = "" }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className={`svg-icon ${className}`}>
    <path d="M10.29 3.86L1.82 18a2 2 0 001.71 3h16.94a2 2 0 001.71-3L13.71 3.86a2 2 0 00-3.42 0zM12 9v4M12 17h.01" />
  </svg>
);

const IconWrench = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M14.7 6.3a1 1 0 000 1.4l1.6 1.6a1 1 0 001.4 0l3.77-3.77a6 6 0 01-7.94 7.94l-6.91 6.91a2.12 2.12 0 01-3-3l6.91-6.91a6 6 0 017.94-7.94l-3.76 3.76z" />
  </svg>
);

const IconCheck = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <polyline points="20 6 9 17 4 12" />
  </svg>
);

const IconGlobe = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <circle cx="12" cy="12" r="10" />
    <line x1="2" y1="12" x2="22" y2="12" />
    <path d="M12 2a15.3 15.3 0 014 10 15.3 15.3 0 01-4 10 15.3 15.3 0 01-4-10 15.3 15.3 0 014-10z" />
  </svg>
);

const IconActivity = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <polyline points="22 12 18 12 15 21 9 3 6 12 2 12" />
  </svg>
);

const IconTerminal = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <polyline points="4 17 10 11 4 5" />
    <line x1="12" y1="19" x2="20" y2="19" />
  </svg>
);

const IconSend = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <line x1="22" y1="2" x2="11" y2="13" />
    <polygon points="22 2 15 22 11 13 2 9 22 2" />
  </svg>
);

const IconInfo = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <circle cx="12" cy="12" r="10" />
    <line x1="12" y1="16" x2="12" y2="12" />
    <line x1="12" y1="8" x2="12.01" y2="8" />
  </svg>
);


const IconSun = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <circle cx="12" cy="12" r="5" />
    <line x1="12" y1="1" x2="12" y2="3" />
    <line x1="12" y1="21" x2="12" y2="23" />
    <line x1="4.22" y1="4.22" x2="5.64" y2="5.64" />
    <line x1="18.36" y1="18.36" x2="19.78" y2="19.78" />
    <line x1="1" y1="12" x2="3" y2="12" />
    <line x1="21" y1="12" x2="23" y2="12" />
    <line x1="4.22" y1="19.78" x2="5.64" y2="18.36" />
    <line x1="18.36" y1="5.64" x2="19.78" y2="4.22" />
  </svg>
);

const IconMoon = ({ size = 14 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="svg-icon">
    <path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z" />
  </svg>
);

const ComponentIcon = ({ kind, size = 14 }) => {

  if (kind === 'router') return <IconRouter size={size} />;
  if (kind === 'tor') return <IconSwitch size={size} />;
  if (kind === 'subnet') return <IconSubnet size={size} />;
  if (kind === 'server') return <IconCpu size={size} />;
  return <IconService size={size} />;
};

// Parser to turn raw Copilot console dump into beautiful UI elements
function parseCopilotOutput(text) {
  if (!text) return null;

  if (text.includes('=== AIRGAP AI NOC COPILOT DECISION SUPPORT')) {
    const q1Index = text.indexOf('Q1:');
    const q2Index = text.indexOf('Q2:');
    const q3Index = text.indexOf('Q3:');
    const llmIndex = text.indexOf('LLM INSIGHT');

    const targetMatch = text.match(/Target Component:\s*([^\n]+)/);
    const targetLine = targetMatch ? targetMatch[1] : 'Unknown Target';

    let q1 = '';
    let q2 = '';
    let q3 = '';
    let llm = '';

    if (q1Index !== -1 && q2Index !== -1) {
      q1 = text.substring(q1Index, q2Index).replace(/Q1:[^\n]*/, '').trim();
    }
    if (q2Index !== -1 && q3Index !== -1) {
      q2 = text.substring(q2Index, q3Index).replace(/Q2:[^\n]*/, '').trim();
    }
    if (q3Index !== -1) {
      const endOfQ3 = llmIndex !== -1 ? llmIndex : text.length;
      q3 = text.substring(q3Index, endOfQ3).replace(/Q3:[^\n]*/, '').trim();
    }
    if (llmIndex !== -1) {
      llm = text.substring(llmIndex).replace(/LLM INSIGHT[^\n]*/, '').trim();
    }

    return {
      type: 'copilot_support',
      target: targetLine,
      q1,
      q2,
      q3,
      llm
    };
  }

  if (text.includes('=== MCP AGENT AUTONOMOUS REPAIR RESULT')) {
    const targetMatch = text.match(/Target:\s*([^\n]+)/);
    const target = targetMatch ? targetMatch[1] : '';
    const actionsIndex = text.indexOf('Actions Taken:');
    const actions = actionsIndex !== -1 ? text.substring(actionsIndex + 14).trim() : '';

    return {
      type: 'mcp_repair',
      target,
      actions
    };
  }

  return {
    type: 'raw',
    content: text
  };
}

export default function App() {
  const [isDarkMode, setIsDarkMode] = useState(true);
  const [orgFilter, setOrgFilter] = useState('ALL');
  const [categoryFilter, setCategoryFilter] = useState('ALL');
  const [telemetry, setTelemetry] = useState(null);

  useEffect(() => {
    if (isDarkMode) {
      document.body.classList.remove('theme-light');
      document.body.classList.add('theme-dark');
    } else {
      document.body.classList.remove('theme-dark');
      document.body.classList.add('theme-light');
    }
  }, [isDarkMode]);
  const [copilotText, setCopilotText] = useState('Select any component card or prompt local LLM daemon...');
  const [userPrompt, setUserPrompt] = useState('');
  const [lastSync, setLastSync] = useState('Syncing...');
  const [orgOrder, setOrgOrder] = useState(['org-a', 'org-b', 'org-c', 'org-d']);
  const [selectedComponent, setSelectedComponent] = useState(null);
  const [selectedFaultScenario, setSelectedFaultScenario] = useState('congestion');
  const [selectedFaultTarget, setSelectedFaultTarget] = useState('');

  useEffect(() => {
    fetchTelemetryData();
    const interval = setInterval(fetchTelemetryData, 2500);
    return () => clearInterval(interval);
  }, []);

  const fetchTelemetryData = async () => {
    try {
      const res = await fetch('/api/state');
      const data = await res.json();
      setTelemetry(data);
      setLastSync(new Date().toLocaleTimeString());
    } catch (err) {
      console.error('Failed to fetch telemetry data:', err);
    }
  };

  const handleSelectComponent = async (name) => {
    setCopilotText(`Querying Local LLM Daemon for component '${name}'...`);
    
    // Find component details in telemetry to display in details sidebar
    if (telemetry?.topology) {
      const matched = telemetry.topology.find(c => c.name === name);
      if (matched) {
        setSelectedComponent(matched);
      }
    }
    
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
    } catch {
      setCopilotText("Error fetching Copilot decision support.");
    }
  };

  const handleSendPrompt = async (forcedPrompt = null) => {
    const promptToSend = forcedPrompt || userPrompt;
    if (!promptToSend) return;
    
    setUserPrompt('');
    setCopilotText('Querying Dedicated Local LLM Daemon (Port 11435)...');
    try {
      const res = await fetch('/api/copilot/chat', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ prompt: promptToSend })
      });
      const data = await res.json();
      setCopilotText(data.answer);
    } catch {
      setCopilotText("Error querying local LLM daemon.");
    }
  };

  const handleInjectFault = async (scenario, target = 'org-a-edge-router') => {
    setCopilotText(`Injecting progressive fault '${scenario}' into '${target}'...`);
    try {
      await fetch('/api/fault/inject', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ scenario: scenario, target: target })
      });
      fetchTelemetryData();
      handleSelectComponent(target);
    } catch (e) {
      console.error('Failed to inject fault:', e);
    }
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
    } catch {
      setCopilotText("Error executing MCP repair skill.");
    }
  };

  const handleClearFaults = async () => {
    setCopilotText('Clearing all active infrastructure faults...');
    try {
      await fetch('/api/fault/clear', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({})
      });
      fetchTelemetryData();
      setCopilotText('All infrastructure faults successfully cleared.');
      setSelectedComponent(null);
    } catch (e) {
      console.error('Failed to clear faults:', e);
    }
  };

  // Drag and Drop
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

  // Calculate live global stats for KPI metrics
  const componentsList = telemetry?.topology || [];
  const avgHealthScore = componentsList.length > 0 
    ? Math.round(componentsList.reduce((acc, c) => acc + (c.current_metrics?.health_score ?? 100), 0) / componentsList.length)
    : 100;
  const activeAlertsCount = telemetry?.alerts?.length || 0;
  const activeFaultsCount = telemetry?.active_faults ? Object.keys(telemetry.active_faults).length : 0;
  const totalRPS = componentsList.reduce((acc, c) => acc + (c.current_metrics?.http_requests_per_sec || 0), 0).toFixed(1);

  // Link selected component to the active live metrics dictionary
  const activeComponent = componentsList.find(c => c.name === selectedComponent?.name) || selectedComponent;

  // Prompt suggestions tags
  const promptSuggestions = [
    "What is likely to fail next?",
    "Check Org A router status",
    "Explain BGP route flap logs",
    "Run self-healing optimization"
  ];

  return (
    <div>
      <header>
        <div className="brand">
          <LucidLogo />
          <div>
            <div className="brand-title">LUCID NOC</div>
            <div style={{ fontSize: '12px', color: 'var(--text-muted)' }}>
              Autonomous Air-Gapped Network Operations Dashboard (44 Components)
            </div>
          </div>
        </div>
        <div style={{ display: 'flex', alignItems: 'center', gap: '12px' }}>
          <div className="live-badge">
            <div className="live-dot"></div>
            <span>Live Telemetry ● {lastSync}</span>
          </div>
          <button 
            onClick={() => setIsDarkMode(!isDarkMode)} 
            className="btn" 
            style={{ padding: '6px 12px', background: 'var(--glass-dark)' }}
            title="Toggle Light/Dark Mode"
          >
            {isDarkMode ? <IconSun size={16} /> : <IconMoon size={16} />}
          </button>
        </div>
      </header>

      {/* KPI Stats bar */}
      <div className="kpi-grid">
        <div className="kpi-card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span className="kpi-label">Avg Infrastructure Health</span>
            <IconActivity size={16} style={{ color: avgHealthScore > 85 ? 'var(--status-healthy)' : 'var(--status-degraded)' }} />
          </div>
          <span className="kpi-value" style={{ color: avgHealthScore > 85 ? 'var(--status-healthy)' : avgHealthScore > 60 ? 'var(--status-degraded)' : 'var(--status-critical)' }}>
            {avgHealthScore}%
          </span>
        </div>
        <div className="kpi-card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span className="kpi-label">Predictive Alerts</span>
            <IconAlert size={16} style={{ color: activeAlertsCount > 0 ? 'var(--status-critical)' : 'var(--text-muted)' }} />
          </div>
          <span className="kpi-value" style={{ color: activeAlertsCount > 0 ? 'var(--status-critical)' : 'var(--text-main)' }}>
            {activeAlertsCount}
          </span>
        </div>
        <div className="kpi-card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span className="kpi-label">Active Faults</span>
            <IconAlert size={16} style={{ color: activeFaultsCount > 0 ? 'var(--status-degraded)' : 'var(--text-muted)' }} />
          </div>
          <span className="kpi-value" style={{ color: activeFaultsCount > 0 ? 'var(--status-degraded)' : 'var(--text-main)' }}>
            {activeFaultsCount}
          </span>
        </div>
        <div className="kpi-card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span className="kpi-label">Aggregated Request Rate</span>
            <IconActivity size={16} style={{ color: 'var(--accent-cyan)' }} />
          </div>
          <span className="kpi-value">{totalRPS} rps</span>
        </div>
      </div>

      {/* Navigation & Controls */}
      <div className="controls-row">
        <div className="filter-section">
          <span className="filter-section-title" style={{ display: 'inline-flex', alignItems: 'center', gap: '5px' }}>
            <IconOrg size={12} />
            <span>Tenant Scope</span>
          </span>
          <select value={orgFilter} onChange={(e) => setOrgFilter(e.target.value)}>
            <option value="ALL">All Organizations</option>
            <option value="org-a">Org A (10.10.0.0/24)</option>
            <option value="org-b">Org B (10.11.0.0/24)</option>
            <option value="org-c">Org C (10.12.0.0/24)</option>
            <option value="org-d">Org D (10.13.0.0/24)</option>
          </select>
        </div>

        <div className="segmented-control">
          {[
            { id: 'ALL', label: 'All Components', icon: <IconGlobe size={12} /> },
            { id: 'admin', label: 'Admin Control', icon: <IconShield size={12} /> },
            { id: 'worker', label: 'Workloads', icon: <IconCpu size={12} /> },
            { id: 'router', label: 'Edge Routers', icon: <IconRouter size={12} /> },
            { id: 'tor', label: 'TOR Switches', icon: <IconSwitch size={12} /> },
            { id: 'rack', label: 'Racks', icon: <IconRack size={12} /> },
            { id: 'subnet', label: 'Subnets', icon: <IconSubnet size={12} /> },
            { id: 'service', label: 'Services', icon: <IconService size={12} /> }
          ].map(btn => (
            <button
              key={btn.id}
              className={`segmented-btn ${categoryFilter === btn.id ? 'active' : ''}`}
              onClick={() => setCategoryFilter(btn.id)}
              style={{ display: 'inline-flex', alignItems: 'center', gap: '5px' }}
            >
              {btn.icon}
              <span>{btn.label}</span>
            </button>
          ))}
        </div>
      </div>

      <div className="noc-layout-grid">
        {/* Left Column: Topology Grid */}
        <div>
          <div className="card">
            <div className="card-title">
              <span>Interactive Data Centers Grid</span>
              <span className="card-title-subtitle">Drag cabinets to customize topology layout</span>
            </div>

            <InteractiveOrgGrid
              orgOrder={orgOrder}
              orgFilter={orgFilter}
              categoryFilter={categoryFilter}
              telemetry={telemetry}
              selectedName={activeComponent?.name}
              onSelect={handleSelectComponent}
              onDragStart={handleDragStart}
              onDrop={handleDrop}
            />
          </div>
        </div>

        {/* Right Column: AI Diagnostics & Details Sidebar */}
        <div>
          {/* Contextual Detail Panel */}
          {activeComponent && (
            <div className="detail-panel">
              <div className="detail-header">
                <div className="detail-meta">
                  <div className="detail-title">{activeComponent.name}</div>
                  <div className="detail-subtitle" style={{ display: 'flex', alignItems: 'center', gap: '4px', marginTop: '3px' }}>
                    <ComponentIcon kind={activeComponent.kind} size={11} />
                    <span>{activeComponent.kind} {activeComponent.role ? `(${activeComponent.role})` : ''}</span>
                  </div>
                </div>
                <div className="detail-health-score">
                  <div className={`detail-score-circle ${
                    (activeComponent.current_metrics?.health_score ?? 100) > 80 ? 'healthy' :
                    (activeComponent.current_metrics?.health_score ?? 100) > 50 ? 'degraded' : 'critical'
                  }`}>
                    {activeComponent.current_metrics?.health_score ?? 100}%
                  </div>
                </div>
              </div>

              {activeComponent.tti_minutes && (
                <div style={{
                  background: 'rgba(244, 63, 94, 0.08)',
                  border: '1px solid rgba(244, 63, 94, 0.2)',
                  borderRadius: '8px',
                  padding: '10px 12px',
                  marginBottom: '14px',
                  fontSize: '12px',
                  color: '#fda4af',
                  fontWeight: 600,
                  display: 'flex',
                  alignItems: 'center',
                  gap: '6px'
                }}>
                  <IconAlert size={14} className="shake-anim" />
                  <span>Projected Failure (TTI): {activeComponent.tti_minutes} mins</span>
                </div>
              )}

              <div className="detail-metrics-grid">
                {[
                  { label: 'Utilization', val: `${activeComponent.current_metrics?.interface_utilization_pct ?? 0}%`, fill: activeComponent.current_metrics?.interface_utilization_pct },
                  { label: 'Throughput', val: `${activeComponent.current_metrics?.http_requests_per_sec ?? 0} RPS`, fill: Math.min(100, (activeComponent.current_metrics?.http_requests_per_sec ?? 0) * 4) },
                  { label: 'Latency', val: `${activeComponent.current_metrics?.latency_ms ?? 0} ms`, fill: Math.min(100, (activeComponent.current_metrics?.latency_ms ?? 0) * 1.5) },
                  { label: 'Packet Loss', val: `${activeComponent.current_metrics?.packet_loss_pct ?? 0}%`, fill: Math.min(100, (activeComponent.current_metrics?.packet_loss_pct ?? 0) * 20) },
                  { label: 'Active TCP', val: `${activeComponent.current_metrics?.active_tcp_connections ?? 0} conn`, fill: Math.min(100, (activeComponent.current_metrics?.active_tcp_connections ?? 0) / 2) },
                  { label: 'Error Rate', val: `${activeComponent.current_metrics?.error_rate_pct ?? 0}%`, fill: Math.min(100, (activeComponent.current_metrics?.error_rate_pct ?? 0) * 10) }
                ].map((item, idx) => (
                  <div key={idx} className="detail-metric-card">
                    <span className="detail-metric-label">{item.label}</span>
                    <span className="detail-metric-val">{item.val}</span>
                    <div className="metric-bar-container">
                      <div 
                        className="metric-bar-fill" 
                        style={{ 
                          width: `${item.fill ?? 0}%`, 
                          background: item.label === 'Packet Loss' || item.label === 'Error Rate' 
                            ? 'var(--status-critical)' 
                            : 'var(--accent-cyan)'
                        }}
                      />
                    </div>
                  </div>
                ))}
              </div>

              <div className="detail-actions">
                <button className="btn btn-primary" onClick={() => handleMCPFix(activeComponent.name)}>
                  <IconWrench size={13} />
                  <span>Run Autonomous Repair (MCP Skill)</span>
                </button>
                
                <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '6px' }}>
                  <button className="btn btn-danger" onClick={() => handleInjectFault('congestion', activeComponent.name)}>
                    <IconAlert size={12} />
                    <span>Congestion</span>
                  </button>
                  <button className="btn btn-danger" onClick={() => handleInjectFault('route_flap', activeComponent.name)}>
                    <IconAlert size={12} />
                    <span>Route Flap</span>
                  </button>
                </div>
              </div>
            </div>
          )}

          {/* AI Terminal Diagnostic Console */}
          <div className="card">
            <div className="card-title">
              <span>LUCID AI Diagnostics Copilot</span>
              <span className="terminal-badge" style={{ display: 'inline-flex', alignItems: 'center', gap: '4px' }}>
                <IconCpu size={10} />
                <span>Offline LLM</span>
              </span>
            </div>

            <div className="terminal-box">
              <div className="terminal-header">
                <span style={{ display: 'inline-flex', alignItems: 'center', gap: '5px' }}>
                  <IconTerminal size={11} />
                  <span>Console stdout / LLM Link</span>
                </span>
                <span>Port 11435</span>
              </div>
              
              {(() => {
                const parsed = parseCopilotOutput(copilotText);
                if (!parsed) return null;

                if (parsed.type === 'copilot_support') {
                  return (
                    <div>
                      <div style={{ borderBottom: '1px solid rgba(255,255,255,0.06)', paddingBottom: '8px', marginBottom: '10px' }}>
                        <div style={{ color: 'var(--accent-cyan)', fontSize: '11px', fontWeight: 600, display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconInfo size={11} />
                          <span>DIAGNOSTIC TARGET</span>
                        </div>
                        <div style={{ fontFamily: 'JetBrains Mono', fontSize: '12px', marginTop: '2px' }}>{parsed.target}</div>
                      </div>

                      <div className="copilot-card" style={{ borderLeft: '3px solid var(--status-degraded)', marginBottom: '8px' }}>
                        <div className="copilot-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconAlert size={11} style={{ color: 'var(--status-degraded)' }} />
                          <span>Q1: Failure Forecast</span>
                        </div>
                        <div className="copilot-section-content">{parsed.q1}</div>
                      </div>

                      <div className="copilot-card" style={{ borderLeft: '3px solid var(--accent-blue)', marginBottom: '8px' }}>
                        <div className="copilot-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconActivity size={11} style={{ color: 'var(--accent-blue)' }} />
                          <span>Q2: Root Cause Signals</span>
                        </div>
                        <div className="copilot-section-content">{parsed.q2}</div>
                      </div>

                      <div className="copilot-card" style={{ borderLeft: '3px solid var(--status-healthy)', marginBottom: '8px' }}>
                        <div className="copilot-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconWrench size={11} style={{ color: 'var(--status-healthy)' }} />
                          <span>Q3: Corrective Actions</span>
                        </div>
                        <div className="copilot-section-content">{parsed.q3}</div>
                      </div>

                      <div className="copilot-card" style={{ background: 'rgba(99, 102, 241, 0.04)', border: '1px solid rgba(99, 102, 241, 0.15)' }}>
                        <div className="copilot-section-title" style={{ color: 'var(--accent-purple)', display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconCpu size={11} />
                          <span>Local LLM Reasoning</span>
                        </div>
                        <div className="copilot-section-content" style={{ margin: 0, fontStyle: 'italic' }}>{parsed.llm}</div>
                      </div>
                    </div>
                  );
                }

                if (parsed.type === 'mcp_repair') {
                  return (
                    <div>
                      <div className="copilot-card" style={{ borderLeft: '3px solid var(--status-healthy)', background: 'rgba(16, 185, 129, 0.03)' }}>
                        <div className="copilot-section-title" style={{ color: 'var(--status-healthy)', fontSize: '12px', display: 'flex', alignItems: 'center', gap: '5px' }}>
                          <IconCheck size={14} />
                          <span>Autonomous Healing Executed</span>
                        </div>
                        <div style={{ fontSize: '12px', color: 'var(--text-muted)', marginBottom: '8px' }}>Target device: <strong>{parsed.target}</strong></div>
                        <div style={{ fontSize: '12px', color: 'var(--text-light)', whiteSpace: 'pre-wrap' }}>
                          <strong>Actions taken:</strong>
                          <div style={{ marginTop: '4px', paddingLeft: '8px', borderLeft: '1px solid rgba(255,255,255,0.08)' }}>
                            {parsed.actions}
                          </div>
                        </div>
                      </div>
                    </div>
                  );
                }

                return <div className="copilot-raw">{parsed.content}</div>;
              })()}
            </div>

            <div className="prompt-suggestions">
              {promptSuggestions.map((s, idx) => (
                <div key={idx} className="suggestion-tag" onClick={() => handleSendPrompt(s)}>
                  {s}
                </div>
              ))}
            </div>

            <div className="chat-input-row">
              <input
                type="text"
                value={userPrompt}
                onChange={(e) => setUserPrompt(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && handleSendPrompt()}
                placeholder="Query AI copilot (e.g. 'What triggers BGP flaps?')..."
              />
              <button className="btn btn-primary" onClick={() => handleSendPrompt()}>
                <IconSend size={12} />
                <span>Send Query</span>
              </button>
            </div>
          </div>

          {/* Fault injection Control Center */}
          <div className="card">
            <div className="card-title">Fault Injection Control Center</div>
            
            <div className="btn-group">
              <div style={{ display: 'flex', flexDirection: 'column', gap: '10px', marginBottom: '10px' }}>
                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <label style={{ fontSize: '12px', color: 'var(--text-muted)' }}>Target Component</label>
                  <select 
                    style={{ background: 'rgba(0,0,0,0.2)', border: '1px solid var(--panel-border)', color: 'var(--text-main)', padding: '8px', borderRadius: '6px', outline: 'none' }}
                    value={selectedFaultTarget}
                    onChange={(e) => setSelectedFaultTarget(e.target.value)}
                  >
                    <option value="">-- Select Target --</option>
                    {telemetry?.topology?.map(comp => (
                      <option key={comp.name} value={comp.name}>{comp.name} ({comp.kind})</option>
                    ))}
                  </select>
                </div>
                <div style={{ display: 'flex', flexDirection: 'column', gap: '4px' }}>
                  <label style={{ fontSize: '12px', color: 'var(--text-muted)' }}>Fault Scenario</label>
                  <select 
                    style={{ background: 'rgba(0,0,0,0.2)', border: '1px solid var(--panel-border)', color: 'var(--text-main)', padding: '8px', borderRadius: '6px', outline: 'none' }}
                    value={selectedFaultScenario}
                    onChange={(e) => setSelectedFaultScenario(e.target.value)}
                  >
                    <option value="congestion">Progressive Congestion</option>
                    <option value="route_flap">Route Flap Cascade</option>
                    <option value="tunnel_deg">Tunnel Degradation</option>
                    <option value="policy_drift">Policy Drift</option>
                  </select>
                </div>
                <button 
                  className="btn btn-danger" 
                  onClick={() => handleInjectFault(selectedFaultScenario, selectedFaultTarget || telemetry?.topology?.[0]?.name)}
                  disabled={!telemetry?.topology?.length}
                >
                  <IconAlert size={12} />
                  <span>Inject Fault into Live Cluster</span>
                </button>
              </div>

              <button className="btn btn-success" onClick={() => handleMCPFix(selectedFaultTarget || 'org-a-edge-router')}>
                <IconWrench size={12} />
                <span>Execute Default Repair Skill (MCP Agent)</span>
              </button>
              
              <button className="btn" style={{ background: 'rgba(255,255,255,0.02)', borderColor: 'rgba(255,255,255,0.08)' }} onClick={handleClearFaults}>
                <IconCheck size={12} />
                <span>Reset Active Infrastructure Faults</span>
              </button>
            </div>
          </div>

          {/* Predictive Alerts Board */}
          <div className="card">
            <div className="card-title">Live Predictive Incident Feed</div>
            
            <div className="alerts-feed">
              {telemetry?.alerts && telemetry.alerts.length > 0 ? (
                telemetry.alerts.map((al, idx) => (
                  <div
                    key={idx}
                    className="alert-item"
                    onClick={() => handleSelectComponent(al.target)}
                  >
                    <div className="alert-title">
                      <IconAlert size={12} />
                      <span>{al.issue_type}</span>
                    </div>
                    <div className="alert-desc">{al.summary}</div>
                    <div className="alert-footer">
                      <span>TTI: {al.tti_minutes}m</span>
                      <span>Conf: {Math.round(al.confidence * 100)}%</span>
                    </div>
                  </div>
                ))
              ) : (
                <div style={{ color: 'var(--text-muted)', fontStyle: 'italic', fontSize: '12px', textAlign: 'center', padding: '10px 0' }}>
                  No active predictive anomalies detected.
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
  selectedName,
  onSelect,
  onDragStart,
  onDrop,
}) {
  if (!telemetry || !telemetry.by_org) {
    return <div style={{ color: 'var(--text-muted)', fontStyle: 'italic', fontSize: '13px' }}>Booting live 44-component telemetry matrix...</div>;
  }

  const by_org = telemetry.by_org;
  const filteredKeys = orgOrder.filter((k) => orgFilter === 'ALL' || k === orgFilter);

  return (
    <div>
      {filteredKeys.map((orgKey, idx) => {
        const orgData = by_org[orgKey];
        if (!orgData) return null;

        const subnetIndex = ['org-a', 'org-b', 'org-c', 'org-d'].indexOf(orgKey);

        return (
          <div
            key={orgKey}
            className="cabinet-card"
            draggable={true}
            onDragStart={(e) => onDragStart(e, idx)}
            onDragOver={(e) => e.preventDefault()}
            onDrop={(e) => onDrop(e, idx)}
          >
            <div className="cabinet-header">
              <div className="cabinet-title">
                <span className="cabinet-drag-handle">⋮⋮</span>
                <IconOrg size={15} />
                <span>{orgData.name}</span>
              </div>
              <div className="cabinet-badge">
                CIDR: 10.{10 + subnetIndex}.0.0/24
              </div>
            </div>

            <div className="cabinet-sections">
              {(categoryFilter === 'ALL' || categoryFilter === 'router' || categoryFilter === 'subnet') && (
                <div className="cabinet-section">
                  <div className="cabinet-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                    <IconRouter size={11} />
                    <span>Edge & Subnets</span>
                  </div>
                  <RichComponentList 
                    items={orgData.routers.concat(orgData.subnets)} 
                    selectedName={selectedName}
                    onSelect={onSelect} 
                  />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'tor' || categoryFilter === 'rack') && (
                <div className="cabinet-section">
                  <div className="cabinet-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                    <IconSwitch size={11} />
                    <span>TOR & Racks</span>
                  </div>
                  <RichComponentList 
                    items={orgData.tors.concat(orgData.racks)} 
                    selectedName={selectedName}
                    onSelect={onSelect} 
                  />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'admin') && (
                <div className="cabinet-section">
                  <div className="cabinet-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                    <IconShield size={11} />
                    <span>Admin Control</span>
                  </div>
                  <RichComponentList 
                    items={orgData.admin_servers} 
                    selectedName={selectedName}
                    onSelect={onSelect} 
                  />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'worker') && (
                <div className="cabinet-section">
                  <div className="cabinet-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                    <IconCpu size={11} />
                    <span>Workload Nodes</span>
                  </div>
                  <RichComponentList 
                    items={orgData.worker_servers} 
                    selectedName={selectedName}
                    onSelect={onSelect} 
                  />
                </div>
              )}

              {(categoryFilter === 'ALL' || categoryFilter === 'service') &&
                orgData.services &&
                orgData.services.length > 0 && (
                  <div className="cabinet-section">
                    <div className="cabinet-section-title" style={{ display: 'flex', alignItems: 'center', gap: '5px' }}>
                      <IconService size={11} />
                      <span>Microservices</span>
                    </div>
                    <RichComponentList 
                      items={orgData.services} 
                      selectedName={selectedName}
                      onSelect={onSelect} 
                    />
                  </div>
                )}
            </div>
          </div>
        );
      })}
    </div>
  );
}

function RichComponentList({ items, selectedName, onSelect }) {
  if (!items || items.length === 0) {
    return <div style={{ fontSize: '11px', color: 'var(--text-muted)', fontStyle: 'italic', padding: '4px 0' }}>None</div>;
  }

  return (
    <div>
      {items.map((item) => {
        const isSelected = item.name === selectedName;
        const m = item.current_metrics || {};
        const health = m.health_score ?? 100;
        const util = m.interface_utilization_pct || 0;
        const rps = m.http_requests_per_sec || 0;
        const tcp = m.active_tcp_connections || 0;
        const latency = m.latency_ms || 0;
        const tti = item.tti_minutes;

        return (
          <div
            key={item.name}
            className={`device-item status-${item.status} ${isSelected ? 'selected' : ''}`}
            onClick={() => onSelect(item.name)}
          >
            <div className="device-header">
              <span className="device-name" title={item.name}>{item.name}</span>
              <span className={`device-badge ${item.status}`}>{item.status}</span>
            </div>
            
            {/* Health Bar Indicator */}
            <div className="device-health-bar-bg">
              <div 
                className={`device-health-bar-fill ${
                  health > 80 ? 'healthy' : health > 50 ? 'degraded' : 'critical'
                }`}
                style={{ width: `${health}%` }}
              />
            </div>

            {/* Metrics Dashboard Inside the Grid Card */}
            <div className="device-metrics-grid">
              <div className="device-metric-cell">
                <span className="device-metric-lbl">Util:</span>
                <span className="device-metric-num">{util}%</span>
              </div>
              <div className="device-metric-cell">
                <span className="device-metric-lbl">RPS:</span>
                <span className="device-metric-num">{rps}</span>
              </div>
              <div className="device-metric-cell">
                <span className="device-metric-lbl">TCP:</span>
                <span className="device-metric-num">{tcp}</span>
              </div>
              <div className="device-metric-cell">
                {tti ? (
                  <span className="device-metric-num tti-warning">
                    <IconAlert size={9} /> {tti}m
                  </span>
                ) : (
                  <>
                    <span className="device-metric-lbl">Lat:</span>
                    <span className="device-metric-num">{latency}ms</span>
                  </>
                )}
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
