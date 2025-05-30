a = math.sin(3) + math.cos(10)
print("a: " .. tostring(a))

print("current_data: " .. tostring(os.date()))
print("current_time: " .. tostring(os.time()))

-- ----------------------------------------------

a = "one string"
b = string.gsub(a, "one", "another")
print(a)
print(b)

print("type(a): " .. type(a))
print("type(b): " .. type(b))

-- ----------------------------------------------

a = "a line"
b = 'another line'
print(a)
print(b)

-- ----------------------------------------------

a = "[]"
b = "\\[\\]"
print(a)
print(b)

-- ----------------------------------------------

global_text = ""
function try_set_text(value)
    max_len = 10
    if value:len() > max_len then
        print("len value is more " .. tostring(max_len))
        return false
    end

    global_text = value
    return true
end

print("global_text: " .. global_text)

try_set_text("new text")
print("global_text: " .. global_text)

try_set_text("new text new text new text")
print("global_text: " .. global_text)

-- ----------------------------------------------

print("\n")
print("one line\nnext line\n\"in quotes\", 'in quotes'")

-- ----------------------------------------------

print("alo\123\049123")

-- ----------------------------------------------

page = [[

123

]]

io.write(page)

-- ----------------------------------------------