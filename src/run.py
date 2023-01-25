import sys, getopt, os

#convert into executable

def main(argv):
    arg = sys.argv[1]
    if arg == "install":
        #install dependencies
        return 0
    elif arg == "build":
        #complete any compilation
        return 0
    elif arg == "test":
        #run tests
        return 0
    else:
        #URL file maybe
        return 0
    return 1

    