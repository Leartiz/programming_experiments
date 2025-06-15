import re
import os

dir_names = os.listdir(".")

pattern = re.compile(r"^(n([0-9]+)).*$")
max_n = 0
for dir_name in dir_names:
    match_result = re.match(pattern, dir_name)
    if match_result:
        n = int(match_result.group(2))
        if max_n < n:
            max_n = n

print(max_n)


