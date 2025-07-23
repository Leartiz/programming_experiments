function try_set_text(text, err)
    random_int = math.random(1, 2)
    print("random_int: " .. tostring(random_int))

    if random_int % 2 == 0 then
       return true 
    end    

    if err == ni
    err.message = "err"
    return false
end

-- ----------------------------------------------

local err = {}
print("err.message: " .. tostring(err.message))

try_set_text("123", err)
print("err.message: " .. tostring(err.message))

try_set_text("123", err)
print("err.message: " .. tostring(err.message))

-- ----------------------------------------------

