package main

func main() {

	//testApplication()

	// c1, _ := newNSXTConfiguration("manager1", "sjors", false, true, "3.2.0.2")

	// c2, _ := newNSXTConfiguration("manager2", "admin", false, true, "3.2.0.1")

	//f1 := newConfigurationFactory(nsxt)

	testBackend()
	testApplication()

	// c1.setManager("henk")
	// b1 := newBackend("backend1", nsxt, c1)
	// b2 := newBackend("backend2", nsxt, c2)
	// c1.setManager("updated1")
	// c2.setManager("updated2")
	// cr1 := b1.getConfiguration()
	// cr2 := b2.getConfiguration()
	// //	cr3 := b2.getConfiguration().getConfiguration()
	// fmt.Printf("%v", cr1)
	// fmt.Printf("%v", cr2)
	// x := cr1.(*nsxtBackendConfiguration)
	// fmt.Println(x.getManager())
	// x.setManager("blabla")
	// fmt.Printf("%v", cr1)
	// fmt.Println(x.getManager())
	// fmt.Println(getManager(&cr1))
	// fmt.Printf("%v", cr3)
}
