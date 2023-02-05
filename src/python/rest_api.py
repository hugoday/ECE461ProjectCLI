import os

# --header ''Accept: application/vnd.github+json''
def getIssues(url=""):
  _, _, url = url.partition(".com/")
  if os.path.exists('src/python/issues/closed.txt'):
      os.remove('src/python/issues/closed.txt')
  if os.path.exists('src/python/issues/open.txt'):
      os.remove('src/python/issues/open.txt')
  #get token from environment variable
  # os.system('curl -i -H https://api.github.com/search/issues?q=repo:'+url+'+type:issue+state:closed >> src/python/issues/closed.txt')
  # os.system('curl -i -H https://api.github.com/search/issues?q=repo:'+url+'+type:issue+state:open >> src/python/issues/open.txt')
  # os.system('curl -i -H "Authorization: token [TOKEN]" https://api.github.com/search/issues?q=repo:'+url+'+type:issue+state:closed >> src/python/issues/closed.txt')
  # os.system('curl -i -H "Authorization: token [TOKEN]" https://api.github.com/search/issues?q=repo:'+url+'+type:issue+state:open >> src/python/issues/open.txt')
  os.system('curl -i -H "Authorization: token [TOKEN]" https://api.github.com/search/issues?q=repo:hugoday/resume+type:issue+state:closed >> src/python/issues/closed.txt')
  os.system('curl -i -H "Authorization: token [TOKEN]" https://api.github.com/search/issues?q=repo:hugoday/resume+type:issue+state:open >> src/python/issues/open.txt')

def deleteIssues():
  if os.path.exists('src/python/issues/closed.txt'):
      os.remove('src/python/issues/closed.txt')
  if os.path.exists('src/python/issues/open.txt'):
      os.remove('src/python/issues/open.txt')