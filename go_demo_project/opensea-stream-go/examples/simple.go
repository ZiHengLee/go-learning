package main

import (
	"fmt"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	"github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
)

func main() {
	client := opensea.NewStreamClient(types.MAINNET, "5168b3128b494bc589b6c742d47de729", phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	client.Connect()

	//client.OnItemListed("*", func(response interface{}) {
	//	var itemListedEvent entity.ItemListedEvent
	//	err := mapstructure.Decode(response, &itemListedEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemListedEvent)
	//})

	//client.OnItemSold("collection-slug", func(response interface{}) {
	//	var itemSoldEvent entity.ItemSoldEvent
	//	err := mapstructure.Decode(response, &itemSoldEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemSoldEvent)
	//})
	//
	//client.OnItemTransferred("collection-slug", func(response interface{}) {
	//	var itemTransferredEvent entity.ItemTransferredEvent
	//	err := mapstructure.Decode(response, &itemTransferredEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemTransferredEvent)
	//})
	//
	//client.OnItemCancelled("collection-slug", func(response interface{}) {
	//	var itemCancelledEvent entity.ItemCancelledEvent
	//	err := mapstructure.Decode(response, &itemCancelledEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemCancelledEvent)
	//})
	//
	//client.OnItemReceivedBid("collection-slug", func(response interface{}) {
	//	var itemReceivedBidEvent entity.ItemReceivedBidEvent
	//	err := mapstructure.Decode(response, &itemReceivedBidEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemReceivedBidEvent)
	//})
	//client.OnItemReceivedOffer("*", func(response interface{}) {
	//	var itemReceivedOfferEvent entity.ItemReceivedOfferEvent
	//	err := mapstructure.Decode(response, &itemReceivedOfferEvent)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//	}
	//	fmt.Printf("%+v\n", itemReceivedOfferEvent)
	//})
	//
	client.OnItemMetadataUpdated("*", func(response interface{}) {
		var itemMetadataUpdateEvent entity.ItemMetadataUpdateEvent
		err := mapstructure.Decode(response, &itemMetadataUpdateEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		//if len(itemMetadataUpdateEvent.Payload.Traits) > 0 {
		fmt.Printf("%+v\n", itemMetadataUpdateEvent)
		//}
	})

	select {}
}
