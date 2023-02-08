
import os

def run_shortlog(url):
    orig_dir = os.getcwd()
    repo_dir = orig_dir + '/' + url

    print("running shortlog for " + url)
    if os.path.exists("shortlog.txt"):
        os.system("rm shortlog.txt")

    os.chdir(repo_dir)
    os.system("git shortlog -se | sort -n >> shortlog.txt")
    os.system("mv shortlog.txt " + "../")



if __name__ == "__main__":
    run_shortlog("kubernetes")
