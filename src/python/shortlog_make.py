import os, sys, re

def run_shortlog(url):
    url = re.findall("https://github.com/([\w-]+)/([\w-]+)", url)[0][1]
    orig_dir = os.getcwd()
    repo_dir = 'src/repos' + '/' + url

    print("running shortlog for " + url)
    if os.path.exists("shortlog.txt"):
        os.system("rm shortlog.txt")

    os.chdir(repo_dir)
    os.system("git shortlog -se | sort -n >> shortlog.txt")
    # os.system("mv shortlog.txt " + "src/go")



if __name__ == "__main__":
    run_shortlog(url = sys.argv[1])
