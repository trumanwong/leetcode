package parkingSystem

type ParkingSystem struct {
	cars []int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	cars := []int{big, medium, small}
	parkeingSystem := ParkingSystem{
		cars: cars,
	}
	return parkeingSystem
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cars[carType - 1] == 0 {
		return false
	}
	this.cars[carType - 1]--
	return true
}