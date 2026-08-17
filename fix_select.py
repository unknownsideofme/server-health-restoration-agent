import re

with open('lucid/frontend/src/index.css', 'r') as f:
    css = f.read()

css = css.replace('background: rgba(0, 0, 0, 0.4);\n  border: 1px solid var(--panel-border);\n  color: var(--text-light);\n', 'background: var(--input-bg);\n  border: 1px solid var(--panel-border);\n  color: var(--text-main);\n')
css = css.replace('background: rgba(0, 0, 0, 0.6);\n}', 'background: var(--input-bg-focus);\n}')

with open('lucid/frontend/src/index.css', 'w') as f:
    f.write(css)

