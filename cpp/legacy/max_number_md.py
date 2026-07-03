import re

pattern = re.compile(r"^- (n([0-9]+)).*$")
f = open("NOTE.md")

max_n = 0
for line in f.readlines():
    match_result = re.match(pattern, line)
    if match_result:
        n = int(match_result.group(2))
        if max_n < n:
            max_n = n

print(max_n)
