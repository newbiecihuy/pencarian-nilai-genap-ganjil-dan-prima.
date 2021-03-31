/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.util;

/**
 *
 * @author newbiecihuy
 */
public class Bilangan {

    public static void main(String args[]) {
        Integer a = null, x = null, nilai = null;
        a = 10;
        if (a <= 1) {
            System.out.println(" Masukkan angka mulai dari 2.");
        } else {
            if (a == 2) {
                System.out.println(a + " merupakan bilangan prima.");
            } else {
                for (x = 2; x < a; x++) {
                    nilai = x;
                    if (nilai % 2 == 1) {
                        System.out.println(nilai + " adalah bilangan ganjil");
                        if (nilai == 3 || nilai % nilai == 0 && nilai % 3 != 0) {//formula bilangan prima
                            System.out.println(nilai + " adalah bilangan prima");
                        }

                    } else{
                        System.out.println(nilai + " adalah bilangan genap");
                        if (nilai == 2) {
                            System.out.println(nilai + " adalah bilangan prima");
                        }
                    }

                }
            }
        }
    }
}
