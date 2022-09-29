__author__ = "newbiecihuy"
__date__ = "$Mar 31, 2021 8:08:57 PM$"

angka = 100
for x in range(2, angka):
    # print(x)
        if x % 2 == 1:
            print ("%i adalah bilangan ganjil" % x)
            if x <= 10:
                if x == 3 or x % x == 0 and x % 3 != 0:
                    print ("%i adalah bilangan prima" % x)
            else :
             if x == 3 or x % x == 0 and x % 3 != 0 and x % 5 != 0 and x % 7 != 0:
                  print ("%i adalah bilangan prima" % x)
        else:
            print ("%i adalah bilangan genap" % x)	
            if x == 2:   	 
                print ("%i adalah bilangan prima" % x)
