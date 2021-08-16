package redis

function SetTx(c *RedisClient, key string, value interface{}) error {
    p, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return c.Set(key, p)
}

function GetTx(c *RedisClient, key string, dest interface{}) error {
    p, err := c.Get(key)
    if err != nil {
        return err
    }
    return json.Unmarshal(p, dest)
}
