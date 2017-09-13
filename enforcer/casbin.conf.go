package enforcer

var CasbinConf = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act, eft

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) 

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

//说明：如果规则是只要一个权限是不允许的则不允许，则e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
