my_string="hello world"
read var1 var2 <<< "$my_string"

echo "[$IFS]"
echo $var1 "|" $var2
echo

# NOTE:
#   не будет сброшен, нужно сбрасывать в ручную
oldIFS=$IFS
IFS=':'
my_string="hello:world"
read var1 var2 <<< "$my_string"

echo "[$IFS]"
echo $var1 "|" $var2
echo

my_string="hello world"
read var1 var2 <<< "$my_string"

echo "[$IFS]"
echo $var1 "|" $var2
echo

IFS=$oldIFS
my_string="hello world"
read var1 var2 <<< "$my_string"

echo "[$IFS]"
echo $var1 "|" $var2
echo

