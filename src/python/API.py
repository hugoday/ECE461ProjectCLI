import requests

# Used for authentication
headers = {"Authorization":  "token <insert token here>"}

def performQuery(query): 
    # calls request with passed in query and authentication token to Github GraphQl
    request = requests.post('https://api.github.com/graphql', json={'query': query}, headers=headers)
    #check for http post to indicate success
    if request.status_code == 200:
        # return json from query
        return request.json()
    # handle exception


def graphQLMetric(url):
    # fill in query for desired returns from github url (need to figure out metric and info to pull)
    # currently trying to figure out how to get total repo contributions for a year compared to alltime for repo
    query = """{

    } """
    result = performQuery(query) # Perform the query
    print(result) # Print query results
    score = calculateMetricScore(result)
    return score


def calculateMetricScore(result):
    # Calculate metric score based off of query results
    return 0
