from urllib import request,parse

url = 'http://jinjideyiyi.cn/mybooklist'

u = request.urlopen(url)
resp  = u.read()
print(resp)