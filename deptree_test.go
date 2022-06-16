package main

//
//import "testing"
//
//func TestDepTree(t *testing.T) {
//	dt := &DepTree{}
//
//	dt.Add("serve", "buy:ingredients", "cook:ingredients", "present:ingredients")
//	deps := dt.Flatten("serve")
//
//
//	if !testEq( deps, "serve", "buy:ingredients", "cook:ingredients", "present:ingredients") {
//		t.Fatal(deps)
//	}
//
//}
//
//
//
//func testEq(a []string, b... string) bool {
//    if len(a) != len(b) {
//        return false
//    }
//    for i := range a {
//        if a[i] != b[i] {
//            return false
//        }
//    }
//    return true
//}
//
