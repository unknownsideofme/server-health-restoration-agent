import re

with open('lucid/frontend/src/index.css', 'r') as f:
    css = f.read()

css = css.replace('--alert-text-success: #a7f3d0;', '--alert-text-success: #a7f3d0;\n  --alert-bg-danger: rgba(244, 63, 94, 0.04);\n  --alert-border-danger: rgba(244, 63, 94, 0.15);\n  --alert-border-danger-hover: rgba(244, 63, 94, 0.35);')
css = css.replace('--alert-text-success: #059669;', '--alert-text-success: #059669;\n  --alert-bg-danger: rgba(225, 29, 72, 0.08);\n  --alert-border-danger: rgba(225, 29, 72, 0.25);\n  --alert-border-danger-hover: rgba(225, 29, 72, 0.5);')

css = css.replace('background: rgba(244, 63, 94, 0.04);', 'background: var(--alert-bg-danger);')
css = css.replace('border: 1px solid rgba(244, 63, 94, 0.15);', 'border: 1px solid var(--alert-border-danger);')
css = css.replace('border-color: rgba(244, 63, 94, 0.35);', 'border-color: var(--alert-border-danger-hover);')

# Animation keyframes
css = css.replace('from { border-color: rgba(244, 63, 94, 0.15); }', 'from { border-color: var(--alert-border-danger); }')
css = css.replace('to { border-color: rgba(244, 63, 94, 0.3); }', 'to { border-color: var(--alert-border-danger-hover); }')

with open('lucid/frontend/src/index.css', 'w') as f:
    f.write(css)

