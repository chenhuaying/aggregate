aggregate
=========
## zset
the key of zset is uint32 like [redis](http://redis.io), but haven't support range feature

## zmap
the key of zmap is an interface
```
type Comparer interface {
    Less(x Comparer) bool
}
```
and the skiplist is not the same as zset, it using [container](https://github.com/chenhuaying/container) package 
every type implements the **Comparer** interface can be using as the key

## tool
the tool generate a template for the new type

`usage: tool *< type name >*`

It will generate aggregate type template with its' iterator naming *name.go*.
