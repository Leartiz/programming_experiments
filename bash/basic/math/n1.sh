result=$((1 + 3 + 10))
echo "result = $result"

result=$(expr 5 + 2)
echo "result = $result"

duration=101230

sec=$(( duration / 1000 ))
rem=$(( duration % 1000 ))

frac=$(printf "%03d" "$rem")
frac2="${frac:0:2}"

echo "Duration=${duration}ms (${sec}.${frac2}s)"