#include <stdio.h>

int  main()
{
        int a,x,nilai;
        printf("Masukkan angka : ");
        scanf("%d", &a);
    if(a <= 1) {
	     printf("Masukkan angka mulai dari 2.");
	  }else{
	   if(a == 2){
	   printf("%d merupakan bilangan prima.", a);
	  }else{
        for(x=2; x < a; x++){
        	nilai = x;
        if(nilai % 2==1){
             printf("%d adalah bilangan ganjil\n", nilai);
          if (nilai == 3 || nilai % nilai ==0 && nilai % 3 !=0 ){//formula bilangan prima
	           printf("%d adalah bilangan prima\n", nilai);
            } 

          }
         else if (nilai % 2==0){
            printf("%d adalah bilangan genap\n", nilai);
	        if (nilai == 2 ){
	           printf("%d adalah bilangan prima \n", nilai);
            }
          }
         
      }
    }
  }
}	
