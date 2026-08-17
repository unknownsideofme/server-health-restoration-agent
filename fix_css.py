import re

with open('lucid/frontend/src/index.css', 'r') as f:
    css = f.read()

# Add new variables to root
new_root_vars = """
  --cabinet-bg: rgba(10, 15, 30, 0.4);
  --device-bg: rgba(30, 41, 59, 0.25);
  --copilot-bg: rgba(15, 23, 42, 0.5);
  --alert-text-danger: #fda4af;
  --alert-text-success: #a7f3d0;
  --btn-primary-hover: #5046e4;
"""
css = css.replace('  --terminal-bg: #040609;\n', '  --terminal-bg: #040609;\n' + new_root_vars)

# Add new variables to theme-light
new_light_vars = """
  --cabinet-bg: rgba(255, 255, 255, 0.5);
  --device-bg: rgba(255, 255, 255, 0.8);
  --copilot-bg: rgba(255, 255, 255, 0.6);
  --alert-text-danger: #e11d48;
  --alert-text-success: #059669;
  --btn-primary-hover: #4f46e5;
"""
css = css.replace('  --terminal-bg: #1e293b;\n', '  --terminal-bg: #1e293b;\n' + new_light_vars)

# Replace values with variables
css = css.replace('background: rgba(10, 15, 30, 0.4);', 'background: var(--cabinet-bg);')
css = css.replace('background: rgba(30, 41, 59, 0.25);', 'background: var(--device-bg);')
css = css.replace('background: rgba(15, 23, 42, 0.5);', 'background: var(--copilot-bg);')
css = css.replace('color: #fda4af;', 'color: var(--alert-text-danger);')
css = css.replace('color: #a7f3d0;', 'color: var(--alert-text-success);')
css = css.replace('background: #5046e4;', 'background: var(--btn-primary-hover);')
css = css.replace('border-color: #5046e4;', 'border-color: var(--btn-primary-hover);')

# The buttons background is completely transparent if they are danger/success. We should fix btn text colors inside the alert feed as well if they were missed.
with open('lucid/frontend/src/index.css', 'w') as f:
    f.write(css)
