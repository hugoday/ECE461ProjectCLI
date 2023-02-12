#aubrey gatewood 2/9/23

import os
import math
from os import path
import sys
if(os.path.exists('src/metric_scores/rampuptime/output2.txt')):
    os.system('rm src/metric_scores/rampuptime/output2.txt')
os.system('cloc src/metric_scores/repos >> src/metric_scores/rampuptime/output2.txt')
try:
    with open('src/metric_scores/rampuptime/output2.txt') as f:
        lastline = str(f.readlines()[-2])
        os.system('rm src/metric_scores/rampuptime/output2.txt')

    while '  ' in lastline:
        lastline = lastline.replace('  ', ' ')
    while ' ' in lastline:
        lastline = lastline.replace(' ', ',')
    lastline = lastline.split(',')
    commentLines = lastline[3] # number of comment lines
    codeLines = lastline[4] # number of code lines
    commentLines = int(commentLines)
    codeLines = int(codeLines)
    result = min(3 * commentLines / codeLines, 1)
    result = str((result)) #final ratio
    with open('src/metric_scores/rampuptime/RU_Result.txt', 'w') as f:
        f.write(result)
except: 
    print('Failed because repo is empty.')
