#aubrey gatewood 2/9/23

import os
from os import path
import sys
if(os.path.exists('output2.txt')):
    os.system('rm output2.txt')
os.system('cloc src/repos/rnd >> output2.txt')
try:
    with open('output2.txt') as f:
        lastline = str(f.readlines()[-2])
except:
    print('Failed because repo is empty.') 
while '  ' in lastline:
    lastline = lastline.replace('  ', ' ')
while ' ' in lastline:
    lastline = lastline.replace(' ', ',')
lastline = lastline.split(',')
commentLines = lastline[3] # number of comment lines
codeLines = lastline[4] # number of code lines
try: 
    commentLines = int(commentLines)
    codeLines = int(codeLines)
    result = str((commentLines / codeLines)) #final ratio
    with open('RU_Result.txt', 'w') as f:
        f.write(result)
except: 
    print('Failed because repo is empty.')
