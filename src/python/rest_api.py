import os

# --header ''Accept: application/vnd.github+json''

if os.path.exists('closed.txt'):
    os.remove('closed.txt')
if os.path.exists('open.txt'):
    os.remove('open.txt')
os.system('curl -i -H "Authorization: token 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" https://api.github.com/search/issues?q=repo:octocat/Spoon-Knife+type:issue+state:closed >> closed.txt')
os.system('curl -i -H "Authorization: token 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" https://api.github.com/search/issues?q=repo:octocat/Spoon-Knife+type:issue+state:open >> open.txt')