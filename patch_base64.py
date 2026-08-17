import re

with open('lucid/backend/server.js', 'r') as f:
    server_js = f.read()

# Replace res.json(...) with base64 sending
server_js = re.sub(
    r"res\.json\(JSON\.parse\((.*?)\)\);",
    r"res.type('text/plain');\n    res.send(Buffer.from(\1).toString('base64'));",
    server_js
)

with open('lucid/backend/server.js', 'w') as f:
    f.write(server_js)

with open('lucid/frontend/src/App.jsx', 'r') as f:
    app_jsx = f.read()

# Replace const data = await res.json(); with base64 decoding
app_jsx = re.sub(
    r"const data = await res\.json\(\);",
    r"const base64Text = await res.text();\n      const data = JSON.parse(atob(base64Text));",
    app_jsx
)

with open('lucid/frontend/src/App.jsx', 'w') as f:
    f.write(app_jsx)
