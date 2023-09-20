package particle_velocity

func particleVelocitySoln1(particles []int) int {
	count := 0 // Count of stable periods
	for i := 0; i < len(particles)-2; i++ {
		velocity := particles[i+1] - particles[i]
		for j := i + 2; j < len(particles); j++ {
			if particles[j]-particles[j-1] != velocity {
				break
			}
			count++
		}
	}
	return count
}

func particleVelocitySoln2(particles []int) int {
	totalPeriods := 0
	pLen := len(particles)
	for i := 0; i < pLen; i++ {
		count := 0
		for i+2 < pLen && particles[i+1]-particles[i] == particles[i+2]-particles[i+1] {
			count++
			totalPeriods += count
			i++
		}
	}
	if totalPeriods < 1000000000 {
		return totalPeriods
	}
	return -1
}
