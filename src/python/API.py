import requests
from dateutil import parser
import datetime
import sys
import os

# Used for authentication
headers = {"Authorization":  "token " + os.getenv("GITHUB_TOKEN")}

# Call this function in Go, pass in a URL, returns a score for responsiveness
def graphQLMetric(url):
    type(url)
    # Number of commits to check for in Repository
    maxCommits = 100

    # Variables to pass into query
    variables = {
        "url": url,
        "commits": maxCommits
      }
    
    # Query, pass in variables for url and number of commits to get
    # Looksup URl, sees if its a repository
    # Gets name, date created, then searches the repository for commits
    # Then grabs first x commits (x being whatever maxCommits is set to)
    # Pulls the commit date of each of those commits
    # Query was formed using Github GraphQL Explorer to see which data could be pulled
    query = """ 
    query MyQuery($url: URI!, $commits: Int) {
      resource(url: $url) {
        ... on Repository {
          name
          createdAt
          defaultBranchRef {
            target {
              ... on Commit {
                history(first: $commits) {
                  edges {
                    node {
                      ... on Commit {
                        committedDate
                      }
                    }
                  }
                }
              } 
            }
          }
        }
      }
    } 
    """

    # Performs the Query, passing in query and variables
    result = performQuery(query, variables) # Perform the query
    if result == 0: # If Query request has an expection
      return 0
    # Takes results from Query and returns repo origin, commits this year, total commits (capped at maxCommits)
    repoDate, commitsThisYear, totalCommits = parseGraphQLresult(result, maxCommits)

    # Takes Parsed Query results and calculates the score
    score = calculateResponsiveness(repoDate, commitsThisYear, totalCommits, maxCommits)

    # Return score
    return score 

# Performs the Query, takes in query and variables, returns request.json
def performQuery(query, variables): 
    # uses requests to perform query with query, variables, and headers (containing github token)
    request = requests.post('https://api.github.com/graphql', json={'query': query, 'variables': variables}, headers=headers)
    if request.status_code == 200: # checks if request was successful
        return request.json() # returns results
    else: # in case of expection
      return 0 # error case

# Parses GraphQL results for scoring, recieves results, and maxCommits, returns score
def parseGraphQLresult(result, maxCommits):
    # Declare Variables
    totalCommits = 0
    commitsThisYear = 0
      
    # Get date a year ago
    lastYear = datetime.date.today() - datetime.timedelta(days=365)

    # Get Repo origin date
    repoDate, z  = result['data']['resource']['createdAt'].split('T')# Removes time from datetime
    repoDate = datetime.date.fromisoformat(repoDate) # Converts repo origin date to desired format

    # Checks each commit to see if it was within the last year
    for x in result['data']['resource']['defaultBranchRef']['target']['history']['edges']:

      # Gets commit date
      date, time = x['node']['committedDate'].split('T') # Removes time from datetime
      commitDate = datetime.date.fromisoformat(date) # Converts date to desired format

      if commitDate > lastYear: # Checks if commit is within last year
        commitsThisYear += 1 # Increments commits this year
      totalCommits += 1 # Increments total # of commits

    return repoDate, commitsThisYear, totalCommits # Returns repo origin date, commits this year, and total commits

# Calculates responsiveness of repository based on creation date, number of commits within last year,
# and if there are at least the number of desired commits in history
def calculateResponsiveness(repoDate, commitsThisYear, totalCommits, maxCommits):
    # Scoring - Look at plan document for insight into how scores are produced

    # Variable Declarations
    score = 0

    # Date Times for origin comparison
    lastYear = datetime.date.today() - datetime.timedelta(days=365) # Date 1 year ago
    lastTwoYears = datetime.date.today() - datetime.timedelta(days=730) # Date 2 years ago

    # Calculates score
    if totalCommits == maxCommits: # number of commits in repo history >= maxCommits
      if repoDate > lastYear: # check if repo is less than a year old
        score = 20
      else:
        if repoDate > lastTwoYears: # checks if repo is less than two years old
          score = 20 * commitsThisYear / maxCommits  
        else:
          score = min(20, (30 * commitsThisYear/maxCommits))
    else: # There are less than maxCommits, so scores are lowered
      if repoDate > lastYear: # checks if repo is less than a year old
        if commitsThisYear < 20:
          score = 0
        else:
          score = 5 * commitsThisYear / maxCommits
      else:
        score = 10 * commitsThisYear / maxCommits
    
    return score / 20 # returns Responsiveness Score on scale of [0,1]

if __name__ == "__main__":
  score = graphQLMetric(sys.argv[1])
  print(score)
  
