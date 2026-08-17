const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const { execFile } = require('child_process');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 8085;

const pythonBin = path.join(__dirname, '../../.venv/bin/python3');
const bridgeScript = path.join(__dirname, 'bridge.py');

// Middleware
app.use(cors());
app.use(express.json());
app.use(morgan('dev'));

// Helper to execute Python bridge
function runBridge(action, args = []) {
  return new Promise((resolve, reject) => {
    const allArgs = [bridgeScript, action, ...args];
    execFile(pythonBin, allArgs, { maxBuffer: 10 * 1024 * 1024 }, (error, stdout, stderr) => {
      if (error) {
        console.error(`Bridge error for action ${action}:`, stderr || error.message);
        return reject(new Error(stderr || error.message));
      }
      resolve(stdout);
    });
  });
}

// REST API Routes
app.get('/api/state', async (req, res) => {
  try {
    const stdout = await runBridge('telemetry');
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: 'Failed to fetch state', details: err.message });
  }
});

app.get('/api/copilot', async (req, res) => {
  try {
    const target = req.query.target || 'org-a-edge-router';
    const stdout = await runBridge('copilot', [target]);
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: 'Failed to generate copilot analysis', details: err.message });
  }
});

app.post('/api/copilot/chat', async (req, res) => {
  try {
    const prompt = req.body.prompt || 'What is likely to fail next?';
    const stdout = await runBridge('copilot-chat', [prompt]);
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: 'Failed to query copilot chat', details: err.message });
  }
});

app.post('/api/fault/inject', async (req, res) => {
  try {
    const scenario = req.body.scenario || 'congestion';
    const target = req.body.target || 'org-a-edge-router';
    const stdout = await runBridge('inject-fault', [scenario, target]);
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: 'Failed to inject fault', details: err.message });
  }
});

app.post('/api/fault/clear', async (req, res) => {
  try {
    const target = req.body.target || '';
    const stdout = await runBridge('clear-faults', [target]);
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: 'Failed to clear faults', details: err.message });
  }
});

app.post('/api/mcp/call', async (req, res) => {
  try {
    const toolName = req.body.tool_name || 'execute_autonomous_repair';
    const argumentsObj = req.body.arguments || {};
    const stdout = await runBridge('mcp-call', [toolName, JSON.stringify(argumentsObj)]);
    res.json(JSON.parse(stdout));
  } catch (err) {
    res.status(500).json({ error: `Failed to call MCP tool ${req.body.tool_name}`, details: err.message });
  }
});

// Prometheus Metrics Route
app.get('/metrics', async (req, res) => {
  try {
    const stdout = await runBridge('metrics');
    res.set('Content-Type', 'text/plain; version=0.0.4');
    res.send(stdout);
  } catch (err) {
    res.status(500).send(`# ERROR: Failed to fetch Prometheus metrics: ${err.message}\n`);
  }
});

// Serve compiled static frontend assets in production mode
const frontendDistPath = path.join(__dirname, '../frontend/dist');
app.use(express.static(frontendDistPath));

// Fallback all unspecified routes to index.html for React Router
app.get('*', (req, res) => {
  res.sendFile(path.join(frontendDistPath, 'index.html'));
});

// Start Server
app.listen(PORT, '0.0.0.0', () => {
  console.log(`Node.js Express backend proxy listening on port ${PORT}`);
  console.log(`Proxying queries through Python binary at: ${pythonBin}`);
});
