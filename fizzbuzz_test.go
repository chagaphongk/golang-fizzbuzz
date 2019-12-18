package fizzbuzz_test

import (
	"testing"

	"github.com/chagaphongk/fizzbuzz"
)

// 1 =1
// 2=2
// 3=fizzbuzz
// 4=4
// 5 =buzz

func TestFizzbuzz(t *testing.T) {
	t.Run("1 should say 1", func(t *testing.T) {
		get := fizzbuzz.Say(1)
		want := "1"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
func TestFizzbuzz2Should2(t *testing.T) {
	t.Run("2 should say 2", func(t *testing.T) {
		get := fizzbuzz.Say(2)
		want := "2"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}

func TestFizzbuzz3ShouldBeFizz(t *testing.T) {
	t.Run("3 should say fizz", func(t *testing.T) {
		get := fizzbuzz.Say(3)
		want := "fizz"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
func TestFizzbuzz4ShouldBe4(t *testing.T) {
	t.Run("4 should say 4", func(t *testing.T) {
		get := fizzbuzz.Say(4)
		want := "4"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
func TestFizzbuzz5ShouldBeBuzz(t *testing.T) {
	t.Run("5 should say Buzz", func(t *testing.T) {
		get := fizzbuzz.Say(5)
		want := "buzz"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
func TestFizzbuzz6ShouldBeNotSupport(t *testing.T) {
	t.Run("6 should say Buzz", func(t *testing.T) {
		get := fizzbuzz.Say(6)
		want := "6 is not support"
		if want != get {
			t.Errorf("it should say %s but get %s", want, get)
		}
	})
}
