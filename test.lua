local function level4() error("error!") end
local function level3() level4() end
local function level2() level3() end
local function level1() level2() end
local ok, result = xpcall(level1, function(err)
  return debug.traceback("msg", 3)
end)
print('Result:')
print(result)
print('Expected:')
print(string.gsub([[msg
stack traceback:
	test.lua:1: in function 'level4'
	test.lua:2: in function 'level3'
	test.lua:3: in function 'level1'
	test.lua:4: in function <test.lua:4>
	[G]: in function 'xpcall'
	test.lua:5: in main chunk
	[G]: ?]], "\t", "\t"))

local result = os.date("!*t", os.time())
print("Result type:", type(result))
if result then
    for k, v in pairs(result) do
        print(k, v)
    end
else
    print("Result is nil!")
end
