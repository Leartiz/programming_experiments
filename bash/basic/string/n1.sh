lvl="INFO"
echo "${lvl,}"
echo "${lvl,,}"

echo file_{A..C}.txt
echo file_{1..5}.txt

lvl=1
echo "${lvl,}"
echo "${lvl,,}"

lvl=ON
echo "${lvl,}"
echo "${lvl,,}"

lvl="info"
echo "${lvl^^}"
echo "hello world" | tr '[:lower:]' '[:upper:]'