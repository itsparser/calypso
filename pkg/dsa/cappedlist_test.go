package dsa

//func TestCList_Dequeue(t *testing.T) {
//	type fields struct {
//		len  int
//		head int
//		tail int
//		q    []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields *CList
//		want   interface{}
//		want1  bool
//	}{
//		{
//			name:   "dequeue empty list",
//			fields: NewCList(2),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &CList{
//				len:  tt.fields.len,
//				head: tt.fields.head,
//				tail: tt.fields.tail,
//				q:    tt.fields.q,
//			}
//			got, got1 := p.Dequeue()
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func TestCList_Enqueue(t *testing.T) {
//	type fields struct {
//		len  int
//		head int
//		tail int
//		q    []interface{}
//	}
//	type args struct {
//		x interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &CList{
//				len:  tt.fields.len,
//				head: tt.fields.head,
//				tail: tt.fields.tail,
//				q:    tt.fields.q,
//			}
//			if got := p.Enqueue(tt.args.x); got != tt.want {
//				t.Errorf("Enqueue() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewCList(t *testing.T) {
//	type args struct {
//		n int
//	}
//	tests := []struct {
//		name string
//		args args
//		want *CList
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewCList(tt.args.n); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewCList() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
