#aubrey gatewood 2/9/23

import os
from os import path
import sys

os.system('cloc src/repos/rnd >> output.txt')
list = []
try:
    with open('output.txt') as f:
        lastline = str(f.readlines()[-2])
except:
    print('Failed because repo is empty.') 
# print(lastline)
while '  ' in lastline:
    lastline = lastline.replace('  ', ' ')
while ' ' in lastline:
    lastline = lastline.replace(' ', ',')
lastline = lastline.split(',')
# print(lastline)
commentLines = lastline[3] # number of comment lines
codeLines = lastline[4] # number of code lines
# print(commentLines)
try: 
    commentLines = int(commentLines)
    codeLines = int(codeLines)
    result = str((commentLines / codeLines)) #final ratio
    f = open('result.txt', 'w')
    f.write(result)
    f.close()
except: 
    print('Failed because repo is empty.')
# print(result)
# result = str(result)
