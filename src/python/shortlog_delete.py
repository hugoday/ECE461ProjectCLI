import os, sys

def delete_shortlog():
    if(os.path.exists("src/go/shortlog.txt")):
        os.remove("src/go/shortlog.txt")
    


if __name__ == "__main__":
    delete_shortlog()
