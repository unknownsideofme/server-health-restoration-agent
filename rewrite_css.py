import re

with open('lucid/frontend/src/index.css', 'r') as f:
    css = f.read()

# Add theme-light and new vars to root
root_vars = """
  --glass-dark: rgba(0, 0, 0, 0.4);
  --glass-dark-hover: rgba(0, 0, 0, 0.6);
  --input-bg: rgba(0, 0, 0, 0.3);
  --input-bg-focus: rgba(0, 0, 0, 0.5);
  
  --brand-gradient-start: #ffffff;
  --brand-gradient-end: var(--text-light);

  --btn-bg: rgba(30, 41, 59, 0.4);
  --btn-hover: rgba(14, 165, 233, 0.08);

  --card-shadow: 0 10px 30px rgba(0, 0, 0, 0.25);
  
  --bg-trans-10: rgba(0, 0, 0, 0.1);
  --bg-trans-20: rgba(0, 0, 0, 0.2);
  --bg-trans-30: rgba(0, 0, 0, 0.3);
  --bg-trans-40: rgba(0, 0, 0, 0.4);
  
  --terminal-bg: #040609;
"""

light_theme = """
body.theme-light {
  --bg-dark: #f0f4f8;
  --panel-bg: rgba(255, 255, 255, 0.7);
  --panel-border: rgba(15, 23, 42, 0.08);
  --panel-border-hover: rgba(14, 165, 233, 0.4);
  
  --text-main: #0f172a;
  --text-muted: #64748b;
  --text-light: #334155;

  --glass-dark: rgba(255, 255, 255, 0.6);
  --glass-dark-hover: rgba(255, 255, 255, 0.9);
  --input-bg: rgba(255, 255, 255, 0.5);
  --input-bg-focus: rgba(255, 255, 255, 0.8);

  --brand-gradient-start: #0f172a;
  --brand-gradient-end: #334155;

  --btn-bg: rgba(255, 255, 255, 0.8);
  --btn-hover: rgba(14, 165, 233, 0.15);

  --card-shadow: 0 10px 30px rgba(0, 0, 0, 0.04);
  
  --bg-trans-10: rgba(255, 255, 255, 0.3);
  --bg-trans-20: rgba(255, 255, 255, 0.5);
  --bg-trans-30: rgba(255, 255, 255, 0.6);
  --bg-trans-40: rgba(255, 255, 255, 0.7);
  
  --terminal-bg: #1e293b;
}
"""

css = css.replace('  --text-light: #cbd5e1;\n}', '  --text-light: #cbd5e1;\n' + root_vars + '}\n' + light_theme)

css = css.replace('rgba(0, 0, 0, 0.4)', 'var(--glass-dark)')
css = css.replace('rgba(0, 0, 0, 0.6)', 'var(--glass-dark-hover)')
css = css.replace('rgba(0, 0, 0, 0.3)', 'var(--input-bg)')
css = css.replace('rgba(0, 0, 0, 0.5)', 'var(--input-bg-focus)')
css = css.replace('rgba(30, 41, 59, 0.4)', 'var(--btn-bg)')
css = css.replace('rgba(14, 165, 233, 0.08)', 'var(--btn-hover)')
css = css.replace('0 10px 30px rgba(0, 0, 0, 0.25)', 'var(--card-shadow)')
css = css.replace('rgba(0, 0, 0, 0.2)', 'var(--bg-trans-20)')
css = css.replace('#040609', 'var(--terminal-bg)')
css = css.replace('linear-gradient(135deg, #ffffff 40%, var(--text-light) 90%)', 'linear-gradient(135deg, var(--brand-gradient-start) 40%, var(--brand-gradient-end) 90%)')
css = css.replace('linear-gradient(135deg, #ffffff, var(--text-light))', 'linear-gradient(135deg, var(--brand-gradient-start), var(--brand-gradient-end))')

# Colors that should flip
css = css.replace('color: #ffffff;', 'color: var(--text-main);')
css = css.replace('rgba(255, 255, 255, 0.05)', 'var(--panel-border)')
css = css.replace('rgba(255, 255, 255, 0.08)', 'var(--panel-border-hover)')
css = css.replace('rgba(255, 255, 255, 0.1)', 'var(--panel-border-hover)')
css = css.replace('rgba(255, 255, 255, 0.02)', 'var(--panel-border)')
css = css.replace('rgba(255, 255, 255, 0.04)', 'var(--panel-border)')

# Adjust layouts
css = css.replace('grid-template-columns: 1.85fr 1fr;', 'grid-template-columns: 2.2fr 1fr;')
css = css.replace('grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));', 'grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));')

with open('lucid/frontend/src/index.css', 'w') as f:
    f.write(css)

