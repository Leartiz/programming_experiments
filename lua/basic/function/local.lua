local function create_json_line(msg)
    local escape_str = function(s)
        s = tostring(s)
        s = s:gsub('\\', '\\\\')
        s = s:gsub('"', '\\"')
        s = s:gsub('\n', '\\n')
        s = s:gsub('\r', '\\r')
        s = s:gsub('\t', '\\t')
        return s
    end

    local json_line = '{' ..
        '"message_id":"' .. escape_str(msg.message_id) .. '",' ..
        '"addr_a":"' .. escape_str(msg.addr_a.addr) .. '",' ..
        '"addr_b":"' .. escape_str(msg.addr_b.addr) .. '",' ..
        '"parts":' .. tostring(msg.parts) ..
    '}\n'
    return json_line
end

local msg = {
    message_id = '123123123""',
    addr_a = {
        addr = "123456789"
    },
    addr_b = {
        addr = "123456888"
    },
    parts = 3
}

print(create_json_line(msg))