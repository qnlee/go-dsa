package particle_velocity

import (
	"testing"
)

func TestParticleVelocitySoln1(t *testing.T) {
	particles := []int{-1, 1, 3, 3, 3, 2, 3, 2, 1, 0}
	expected := 5
	result := particleVelocitySoln1(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParticleVelocitySoln2(t *testing.T) {
	particles := []int{-1, 1, 3, 3, 3, 2, 3, 2, 1, 0}
	expected := 5
	result := particleVelocitySoln2(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParticleVelocitySoln1_EmptyInput(t *testing.T) {
	particles := []int{}
	expected := 0
	result := particleVelocitySoln1(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParticleVelocitySoln2_EmptyInput(t *testing.T) {
	particles := []int{}
	expected := 0
	result := particleVelocitySoln2(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParticleVelocitySoln1_LargeResult(t *testing.T) {
	particles := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	expected := 45
	result := particleVelocitySoln1(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestParticleVelocitySoln2_LargeResult(t *testing.T) {
	particles := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	expected := 45
	result := particleVelocitySoln2(particles)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func BenchmarkParticleVelocitySoln1(b *testing.B) {
	particles := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	for n := 0; n < b.N; n++ {
		particleVelocitySoln1(particles)
	}
}

func BenchmarkParticleVelocitySoln2(b *testing.B) {
	particles := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	for n := 0; n < b.N; n++ {
		particleVelocitySoln2(particles)
	}
}
