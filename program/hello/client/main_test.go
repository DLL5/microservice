package main

import (
	pb "awesomeProject/proto/hello"
	"strconv"
	"testing"
)

func TestSampleCallHello(t *testing.T) {
	name, name1 := "ldl", "WuDi"
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  *pb.HelloResponse
		want1 string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{name: name}, want1: "Hello " + name + "."},
		{name: "test2", args: args{name: name1}, want1: "Hello " + name1 + "."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := SampleCallHello(tt.args.name)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("SampleCallHello() got = %v, want %v", got, tt.want)
			//}
			if got1 != tt.want1 {
				t.Errorf("SampleCallHello() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// go test 说明： -bench 进行性能测试； -benchtime=5s 指定测试时间为5s； -benchtime=1000x 指定执行一次测试时待测代码运行的次数；
// -count 10 指定执行几次测试，减少一次测试的较大不确定性； -cpu 指定使用的cpu核心数
// go test -bench . -benchtime=5s | -benchtime=1000x -count=10 -cpu=1,4,6,8 -benchmem
func BenchmarkSampleCallHello(b *testing.B) {
	//b.ResetTimer()
	//b.StopTimer()
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SampleCallHello("LiuDongLiang" + strconv.Itoa(i))
		//fmt.Println(v)
	}
}

func BenchmarkConcurrentCallHello(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SampleCallHello("LiuDongLiang")
		}
	})
	//}
}
