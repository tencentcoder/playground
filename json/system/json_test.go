package system

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	b.Run("OneLevelFastJSON", func(b *testing.B) {
		data := []byte(`{"info":"{\"address\":{\"city\":\"Beijing\"}}"}`)
		one := OneLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &one, FastJSON)
		}
	})

	b.Run("TwoLevelFastJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":"{\"city\":\"Beijing\"}"}}`)
		two := TwoLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &two, FastJSON)
		}
	})
	b.Run("ThreeLevelFastJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":{"city":"Beijing"}}}`)
		three := ThreeLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &three, FastJSON)
		}
	})

	b.Run("OneLevelJSON", func(b *testing.B) {
		data := []byte(`{"info":"{\"address\":{\"city\":\"Beijing\"}}"}`)
		one := OneLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &one, JSON)
		}
	})

	b.Run("TwoLevelJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":"{\"city\":\"Beijing\"}"}}`)
		two := TwoLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &two, JSON)
		}
	})

	b.Run("ThreeLevelJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":{"city":"Beijing"}}}`)
		three := ThreeLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &three, JSON)
		}
	})

	b.Run("OneLevelFFJSON", func(b *testing.B) {
		data := []byte(`{"info":"{\"address\":{\"city\":\"Beijing\"}}"}`)
		one := OneLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &one, FFJSON)
		}
	})

	b.Run("TwoLevelFFJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":"{\"city\":\"Beijing\"}"}}`)
		two := TwoLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &two, FFJSON)
		}
	})

	b.Run("ThreeLevelFFJSON", func(b *testing.B) {
		data := []byte(`{"info":{"address":{"city":"Beijing"}}}`)
		three := ThreeLevel{}
		for i := 0; i < b.N; i++ {
			_ = Parse(data, &three, FFJSON)
		}
	})
}

func BenchmarkUnParse(b *testing.B) {
	b.Run("OneLevelFastJSON", func(b *testing.B) {
		one := OneLevel{Info: "{\"address\":{\"city\":\"Beijing\"}}"}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&one, FastJSON)
		}
	})

	b.Run("TwoLevelFastJSON", func(b *testing.B) {
		two := TwoLevel{Info: TwoInfo{Address: "{\"city\":\"Beijing\"}"}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&two, FastJSON)
		}
	})
	b.Run("ThreeLevelFastJSON", func(b *testing.B) {
		three := ThreeLevel{Info: ThreeInfo{ThreeAddress{City: "Beijing"}}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&three, FastJSON)
		}
	})
	b.Run("OneLevelJSON", func(b *testing.B) {
		one := OneLevel{Info: "{\"address\":{\"city\":\"Beijing\"}}"}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&one, JSON)
		}
	})
	b.Run("TwoLevelJSON", func(b *testing.B) {
		two := TwoLevel{Info: TwoInfo{Address: "{\"city\":\"Beijing\"}"}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&two, JSON)
		}
	})

	b.Run("ThreeLevelJSON", func(b *testing.B) {
		three := ThreeLevel{Info: ThreeInfo{ThreeAddress{City: "Beijing"}}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&three, JSON)
		}

	})
	b.Run("OneLevelFFJSON", func(b *testing.B) {
		one := OneLevel{Info: "{\"address\":{\"city\":\"Beijing\"}}"}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&one, FFJSON)
		}
	})
	b.Run("TwoLevelFFJSON", func(b *testing.B) {
		two := TwoLevel{Info: TwoInfo{Address: "{\"city\":\"Beijing\"}"}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&two, FFJSON)
		}
	})

	b.Run("ThreeLevelFFJSON", func(b *testing.B) {
		three := ThreeLevel{Info: ThreeInfo{ThreeAddress{City: "Beijing"}}}
		for i := 0; i < b.N; i++ {
			_, _ = UnParse(&three, FFJSON)
		}

	})
}
