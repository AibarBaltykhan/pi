

 
package piscine


func FindNextPrime(nb int)int{
	nextprime:=nb-1
	i:=nb+1
	for i>nb{
		nextprime++
		if IsPrime(nextprime){
			return nextprime
		}
		i++	
	}
	return nextprime
}





func IsPrime2(value int) bool {
	decision := true

	if value <= 1 {
		return false
	}
	for i := 2; i < value+1; i++ {
		if value%i == 0 {
			decision = false
		}
	}
	return decision
}
