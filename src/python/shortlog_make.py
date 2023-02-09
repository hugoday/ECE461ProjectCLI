import os, sys, re

def run_shortlog(url):
# /    print("making dir " + url)
    # repo_dir = 'src/repos' + '/' + url
    # # url = url.split('/')[-1]
    # print(url)
    print("running shortlog for " + url)
    
    if os.path.exists("shortlog.txt"):
        os.system("rm shortlog.txt")
    os.chdir(url)
    # print("in dir" + os.getcwd())
    os.system("git shortlog -se | sort -n")
    # os.system("mv shortlog.txt " + "src/go")
    # print("curr dir is " + os.getcwd())



if __name__ == "__main__":
    run_shortlog(url = sys.argv[1])
