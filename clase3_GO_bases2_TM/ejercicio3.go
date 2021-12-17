package main
import "fmt"

func main() {
	category:="C"
	workedTimeInMinutes:= 3000
	selectByCategory:= calcularSalario(category)

	if selectByCategory==nil {
		fmt.Println("Categoría no soportada")
		return
	}
	salarioBySelectedCategory:= selectByCategory(workedTimeInMinutes)
	fmt.Println("Total a cobrar por categoría ", category, ": ", salarioBySelectedCategory)


}

func sueldoC(time int) float64 {
	timeInHours:= float64(time/60)
	cCategorySalary:= float64(1000)
	totalSalary:=float64(timeInHours*cCategorySalary)
	return totalSalary
}

func sueldoB(time int) float64 {
	timeInHours:= float64(time/60)
	bCategorySalary:= float64(1500)
	totalSalary:=float64(timeInHours*bCategorySalary)
	totalSalary = (totalSalary + (totalSalary*20)/100)
	return totalSalary

}

func sueldoA(time int) float64 {
	timeInHours:= float64(time/60)
	aCategorySalary:=float64(3000)
	totalSalary:=float64(timeInHours*aCategorySalary)
	totalSalary = (totalSalary + (totalSalary*50)/100)
	return totalSalary
}

func calcularSalario(category string) func(int) float64 {
	switch category {
		case "C":
			return sueldoC
		case "B":
			return sueldoB
		case "A":
			return sueldoA
		default:
			return nil
	}

}