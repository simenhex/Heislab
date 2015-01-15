from threading import Thread

i=0

def thread0():
	global i
	for x in range (1000000):
		i+=1
	
def thread1():
	global i
	for x in range(1000000):
		i-=1

def main():
	inc_i_thread = Thread(target=thread0, args = (), )
	inc_i_thread.start()

	dec_i_thread = Thread(target=thread1, args = (), )
	dec_i_thread.start()

	inc_i_thread.join()
	dec_i_thread.join()

	print(i)

main()
