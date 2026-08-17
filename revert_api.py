import re

with open('lucid/backend/server.js', 'r') as f:
    server_js = f.read()

# Revert /noc/ to /api/
server_js = server_js.replace('/noc/', '/api/')

# Revert Base64 back to JSON
server_js = re.sub(
    r"res.type\('text/plain'\);\n\s*res\.send\(Buffer\.from\((.*?)\)\.toString\('base64'\)\);",
    r"res.json(JSON.parse(\1));",
    server_js
)

with open('lucid/backend/server.js', 'w') as f:
    f.write(server_js)

with open('lucid/frontend/src/App.jsx', 'r') as f:
    app_jsx = f.read()

# Revert /noc/ to /api/
app_jsx = app_jsx.replace('/noc/', '/api/')

# Revert Base64 decoding back to JSON
app_jsx = re.sub(
    r"const base64Text = await res\.text\(\);\n\s*const data = JSON\.parse\(atob\(base64Text\)\);",
    r"const data = await res.json();",
    app_jsx
)

with open('lucid/frontend/src/App.jsx', 'w') as f:
    f.write(app_jsx)
