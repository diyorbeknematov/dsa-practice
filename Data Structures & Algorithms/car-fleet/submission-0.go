type Car struct {
    position int
    speed    int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	for i := range position {
		cars[i] = Car{
			position: position[i],
			speed:    speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleetCount := 0
	lastTime := 0.0

	for i := 0; i < len(cars); i ++ {
		time := float64(target-cars[i].position) / float64(cars[i].speed)

		if time <= lastTime {
			continue
		}

		fleetCount++
		lastTime = time
	}

	return fleetCount
}

// _ _ _ _ _ _ _ _ _ _ _
//         4(2)
//   1(2) 
// 0(1)  
//               7(1)