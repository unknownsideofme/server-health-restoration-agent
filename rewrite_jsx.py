import re

with open('lucid/frontend/src/App.jsx', 'r') as f:
    jsx = f.read()

icons_code = """
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
"""

jsx = jsx.replace('const ComponentIcon = ({ kind, size = 14 }) => {', icons_code)

# Add isDarkMode state to App
app_start = """export default function App() {
  const [telemetry, setTelemetry] = useState(null);
"""
app_start_new = """export default function App() {
  const [isDarkMode, setIsDarkMode] = useState(true);
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
"""
jsx = jsx.replace(app_start, app_start_new)

# Add toggle button to header
header_start = """          <div className="live-badge">
            <div className="live-dot"></div>
            SYSTEM LIVE
          </div>
        </div>
      </header>"""

header_start_new = """          <div className="live-badge">
            <div className="live-dot"></div>
            SYSTEM LIVE
          </div>
          <button 
            onClick={() => setIsDarkMode(!isDarkMode)} 
            className="btn" 
            style={{ marginLeft: '12px', padding: '6px 12px', background: 'var(--glass-dark)' }}
            title="Toggle Light/Dark Mode"
          >
            {isDarkMode ? <IconSun size={16} /> : <IconMoon size={16} />}
          </button>
        </div>
      </header>"""

jsx = jsx.replace(header_start, header_start_new)

with open('lucid/frontend/src/App.jsx', 'w') as f:
    f.write(jsx)

