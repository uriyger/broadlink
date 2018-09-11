package main

import (
	"fmt"
	"time"

	"github.com/uriyger/broadlink"
)

func main() {
	manager, err := broadlink.NewManager()
	if err != nil {
		panic(err)
	}

	fmt.Println("Discovering...")
	devs, err := manager.Discover(8 * time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Devices: %v\n", devs)

	for _, dev := range devs {
		sp3sdev := dev.(*broadlink.Sp3sDevice)
		err = sp3sdev.BaseDevice.Auth()
		if err != nil {
			panic(err)
		}

		fmt.Println("Start checking power ...")
		resp, err := sp3sdev.CheckPower()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Check power response %v\n", resp)

		en, err := sp3sdev.GetEnergy()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Energy: %2.1f\n", en)

		//resp, err := sp3sdev.BaseDevice.EnterLearning()
		//if err != nil {
		//  panic(err)
		//}
		//fmt.Printf( "enter learning mode:%v\n", resp )
		//
		//var data []byte
		//for {
		//  data, err = sp3sdev.BaseDevice.CheckData()
		//  if err == nil {
		//    break
		//  }
		//  time.Sleep( 200 * time.Millisecond )
		//}
		//fmt.Printf( "Learnt Data: %v\n", data )
		//
		//time.Sleep( 5000 * time.Millisecond )
		//err = sp3sdev.BaseDevice.SendData( data )
		//if err != nil {
		//  panic(err)
		//}
	}
}
