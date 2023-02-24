package internal

import (
	pb "calculator/protos/gen/calculator/v1"
)

type counts = map[pb.ItemType]int

type calculationItemOutputs = []*pb.CalculationItemOutput

// Inspired by https://stackoverflow.com/a/42184212
func CountItems(arr []pb.ItemType) counts {
	dict := make(counts)
	for _, item := range arr {
		dict[item] = dict[item] + 1
	}
	return dict
}

func CastCounts(cnts counts) calculationItemOutputs {
	var outputs calculationItemOutputs
	for key, value := range cnts {
		var calculation_item_output_instance = new(pb.CalculationItemOutput)
		calculation_item_output_instance.Item = key
		calculation_item_output_instance.Count = int32(value)
		outputs = append(outputs, calculation_item_output_instance)
	}
	return outputs
}
