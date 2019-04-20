function testFun(tab)
    result = {}
    result["key"] = "test"
    result["key1"] = "val2"

    if(tab["user"]=="test")then
        result["title"]="good"
    end
    if(tab["os"]=="ios")then
        result["url"]="http://www.google.com"
    else
        result["url"]="http://www.baidu.com"
    end

    return result
end

function condition(value)
    return value == "123"
end