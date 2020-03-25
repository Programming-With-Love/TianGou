import json
import random
import requests
import re
from DBUtils.PooledDB import PooledDB
from multiprocessing.pool import ThreadPool

tgi=0
proxys=[]
header = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'}
#获取代理
def get_proxy(i):
	global proxys
	for num in range(0,i):
		r=requests.get("http://api.p00q.cn/proxy/get")
		if(r.status_code==200):
			rjson = r.text
			rjsonp = json.loads(rjson)
			dl=rjsonp.get('ip')+":"+rjsonp.get("port")
			proxys.append(dl)
#获取页面html
def getHtml(url,p):
	if p:
		global proxys
		try:
			i=random.randint(0,len(proxys)-1)
			html = requests.get(url,proxies={'http':"http://"+proxys[i]},headers=header,timeout=5)
			if html!=None:
				return html.text
		except Exception as e:
			print(e)
			proxys.remove(proxys[i])
			if len(proxys)<5:
				get_proxy(10)
			return getHtml(url,True)
	else :
		html = requests.get(url,headers=header,timeout=5)
		return html.text
def main():
	print("开始")
	global proxys
	get_proxy(10)
	print(proxys)
	#创建线程池
	pool = ThreadPool(20)
	for num in range(1,20):
		url="http://m.weibo.cn/api/container/getIndex?containerid=100103type%3D61%26q%3D%E8%88%94%E7%8B%97%E6%97%A5%E8%AE%B0%26t%3D0&page_type=searchall&page="+str(num)
		try:
			pool.apply_async(run, args=(url,))
		except Exception as e:
			print(e)
	pool.close()
	pool.join()
	print("结束")
def run(url):
	#获取页面数据
	r=getHtml(url,True)
	data2 = json.loads(r)
	global tgi
	try:
		for i in range(0,len(data2['data']['cards'])):
			url = 'http://tiangou.p00q.cn/diary'
			d = {'time': '来自微博的日记', 'content':data2['data']['cards'][i]['mblog']["text"]}
			r = requests.post(url, data=d)
			tgi+=1
			print(str(tgi)+"--"+data2['data']['cards'][i]['mblog']["text"])
	except Exception as e:
		print(e)
if __name__ == "__main__":
	main()