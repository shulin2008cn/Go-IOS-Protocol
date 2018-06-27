--- main bet lucky number for walleta
-- @gas_limit 1000000
-- @gas_price 0.0001
-- @param_cnt 0
-- @return_cnt 0
-- @publisher walleta
function main()
	tx = "DtA6uaLRHTXSCTCZ86wUsjaUuqvjxFCTEB723qNmuduU"
	-- tx = "main"
	a = "gvCQNmkuA6AwdddRMSUg6jr8W7swKWAnhEY3cAthj9bX"
	-- a = "walleta"
	ok, r = Call(tx, "QueryLastLuckyBlock")
	Log(string.format("last lucky block %s", tostring(ok)))
	Log(string.format("last lucky block r = %s", tostring(r)))
	Assert(ok)
	-- Assert(r == -1)

	nonce = 0
    for i=0,48 do
		ok, r = Call(tx, "Bet", a, i % 10, 1.1, nonce)
		nonce = nonce + 1
		Log(string.format("bet %s", tostring(ok)))
		Log(string.format("bet r = %s", tostring(r)))
		Assert(ok)
		Assert(r == 0)
    end
    ok, r = Call(tx, "QueryUserNumber")
	Assert(ok)
	-- Assert(r == 49)
    ok, r = Call(tx, "QueryTotalCoins")
	Assert(ok)
	Log(string.format("total coins = %f", r))
	-- Assert(r == 53.9)

    for i=0,49 do
		ok, r = Call(tx, "Bet", a, i % 10, 4.9, nonce)
		nonce = nonce + 1
		Log(string.format("bet %s", tostring(ok)))
		Log(string.format("bet r = %s", tostring(r)))
		Assert(ok)
		Assert(r == 0)
    end
    ok, r = Call(tx, "QueryUserNumber")
	Assert(ok)
	-- Assert(r == 99)
    ok, r = Call(tx, "QueryTotalCoins")
	Assert(ok)
	-- Assert(r == 298.9)
end--f
