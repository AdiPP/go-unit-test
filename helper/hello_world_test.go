package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Adi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adi")
		}
	})

	b.Run("Putra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Putra")
		}
	})
}

func BenchmarkHelloWorldAdi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adi")
	}
}

func BenchmarkHelloWorldPutra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Putra")
	}
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Adi)",
			request: "Adi",
		},
		{
			name:    "HelloWorld(Putra)",
			request: "Putra",
		},
		{
			name:    "HelloWorld(Permana)",
			request: "Permana",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	m.Run() // eksekusi semua unit test

	// after
	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}

	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Result is not 'Hello Adi'")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Adi)",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "HelloWorld(Putra)",
			request:  "Putra",
			expected: "Hello Putra",
		},
		{
			name:     "HelloWorld(Permana)",
			request:  "Permana",
			expected: "Hello Permana",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Adi", func(t *testing.T) {
		result := HelloWorld("Adi")
		require.Equal(t, "Hello Adi", result, "Result is not 'Hello Adi'")
	})

	t.Run("Putra", func(t *testing.T) {
		result := HelloWorld("Putra")
		require.Equal(t, "Hello Putra", result, "Result is not 'Hello Putra'")
	})

	t.Run("Permana", func(t *testing.T) {
		result := HelloWorld("Permana")
		require.Equal(t, "Hello Permana", result, "Result is not 'Hello Permana'")
	})
}

func TestHelloWorldAdi(t *testing.T) {
	result := HelloWorld("Adi")

	// if result != "Hello Adi" {
	// panic("Result is not 'Hello Adi'")
	// t.Fail()
	// t.Error("Result is not 'Hello Adi'")
	// }

	// assert.Equal(t, "Hello Adi", result, "Result is not 'Hello Adi'")
	require.Equal(t, "Hello Adi", result, "Result is not 'Hello Adi'")

	fmt.Println("TestHelloWorldAdi Done")
}

func TestHelloWorldPutra(t *testing.T) {
	result := HelloWorld("Putra")

	// if result != "Hello Putra" {
	// panic("Result is not 'Hello Putra'")
	// t.FailNow()
	// t.Fatal("Result is not 'Hello Putra'")
	// }

	// assert.Equal(t, "Hello Putra", result, "Result is not 'Hello Putra'")
	require.Equal(t, "Hello Putra", result, "Result is not 'Hello Putra'")

	fmt.Println("TestHelloWorldPutra Done")
}

func TestHelloWorldPermana(t *testing.T) {
	result := HelloWorld("Permana")

	// if result != "Hello Putra" {
	// panic("Result is not 'Hello Putra'")
	// t.FailNow()
	// t.Fatal("Result is not 'Hello Putra'")
	// }

	// assert.Equal(t, "Hello Putra", result, "Result is not 'Hello Putra'")
	require.Equal(t, "Hello Permana", result, "Result is not 'Hello Permana'")

	fmt.Println("TestHelloWorldPermana Done")
}
