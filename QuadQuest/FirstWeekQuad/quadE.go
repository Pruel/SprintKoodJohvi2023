package main

import "fmt"

func QuadE(x, y int) {
	if x > 0 && y > 0 { // Проверяет положительный ли размер фигуры, если да то мы можем её нарисовать, если нет то ничего не рисуем.
		for i := 0; i < y; i++ { // Этот цикл поможет нам нарисовать каждую строку вертикали
			for j := 0; j < x; j++ { // Этот цикл поможет нам нарисовать каждую строку горизонтали
				if i == 0 && j == 0 { // Само построение фигуры осуществляется через if else, это точка старта рисунка.
					fmt.Print("A")
				} else if i == 0 && j == x-1 { // Верхний правый
					fmt.Print("C")
				} else if i == y-1 && j == 0 { // Нижний левый
					fmt.Print("C")
				} else if i == y-1 && j == x-1 { // Нижний правый
					fmt.Print("A")
				} else if i == 0 || i == y-1 || j == 0 || j == x-1 { // Стороны фигуры
					fmt.Print("B")
				} else { // Печать пробела в остальных позициях.
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
