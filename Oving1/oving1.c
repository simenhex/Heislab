#include <pthread.h>
#include <stdio.h>
 int i = 0 ;

void* thread0(){

	for(int counter=0; counter<1000000; counter++){
		i++;	
	}	
	return NULL;

};

void* thread1(){

	for(int counter=0; counter<1000000; counter++){
		i--;	
	}
return NULL;
};	

int main(){
	pthread_t inc_i_thread;
	pthread_t dec_i_thread;
	pthread_create(&inc_i_thread, NULL, thread0, NULL);
	pthread_create(&dec_i_thread, NULL, thread1, NULL);
	pthread_join(inc_i_thread,NULL);
	pthread_join(dec_i_thread,NULL);
	printf("%d",i);
	printf("\n");

	return 0;
};

