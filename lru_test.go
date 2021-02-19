package lru

import (
	"testing"
)

func TestLRUCache_Add(t *testing.T) {
	l := NewCache(100)
	l.Add("Python", 1)
	l.Add("Perl", 2)
	l.Add("Go", 3)
	l.Add("C++", 4)
	l.Add("Java", 5)

	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testcase1", args: args{key: "Java", value: 5}},
		{name: "testcase2", args: args{key: "Go", value: 3}},
		{name: "testcase3", args: args{key: "Python", value: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := l.Get(tt.args.key); got.(int) != tt.args.value.(int) {
				t.Errorf("Get = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestLRUCache_Add2(t *testing.T) {
	l := NewCache(-1)
	l.Add("Python", 1)
	l.Add("Perl", 2)
	l.Add("Go", 3)
	l.Add("C++", 4)
	l.Add("Java", 5)
	l.Add("Python", 10)

	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testcase1", args: args{key: "Java", value: 5}},
		{name: "testcase2", args: args{key: "Go", value: 3}},
		{name: "testcase3", args: args{key: "Python", value: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := l.Get(tt.args.key); got.(int) != tt.args.value.(int) {
				t.Errorf("Get = %v, want %v", got, tt.args.value)
			}
			if l.Len() != 5 {
				t.Errorf("Len() = %v, want 5", l.Len())
			}
		})
	}
}

func TestLRUCache_Remove(t *testing.T) {
	l := NewCache(5)
	l.Add("Python", 1)
	l.Add("Perl", 2)
	l.Add("Go", 3)
	l.Add("C++", 4)
	l.Add("Java", 5)
	l.Add("C", 6)

	l.Remove("Go")
	l.Remove("Perl")

	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testcase1", args: args{key: "Go", value: 3}},
		{name: "testcase2", args: args{key: "Perl", value: 2}},
		{name: "testcase3", args: args{key: "Python", value: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, ok := l.Get(tt.args.key); ok {
				t.Errorf("Get = %v, want nil", got)
			}
		})
	}
}

func TestLRUCache_Purge(t *testing.T) {
	l := NewCache(5)
	l.Add("Python", 1)
	l.Add("Perl", 2)
	l.Add("Go", 3)
	l.Add("C++", 4)
	l.Add("Java", 5)
	l.Add("C", 6)

	l.Purge()

	tests := []struct {
		name string
	}{
		{name: "testcase1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Len(); got != 0 {
				t.Errorf("Len = %v, want 0", got)
			}
		})
	}
}

func TestLRUCache_PurgeAdd(t *testing.T) {
	l := NewCache(5)
	l.Add("Python", 1)
	l.Add("Perl", 2)
	l.Add("Go", 3)
	l.Add("C++", 4)
	l.Add("Java", 5)
	l.Add("C", 6)

	l.Purge()

	l.Add("C++", 4)
	l.Add("Java", 5)
	l.Add("C", 6)

	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "testcase1", args: args{key: "C++", value: 4}},
		{name: "testcase2", args: args{key: "Java", value: 5}},
		{name: "testcase3", args: args{key: "C", value: 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := l.Get(tt.args.key); got.(int) != tt.args.value.(int) {
				t.Errorf("Get = %v, want nil", got)
			}
		})
	}
}
