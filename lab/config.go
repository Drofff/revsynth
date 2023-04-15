package main

import "drofff.com/revsynth/aco"

var experiments = []experimentSettings{
	{
		filename: "num_of_ants",
		config: aco.Config{
			NumOfAnts:       10,
			NumOfIterations: 15,
			Alpha:           2.0,
			Beta:            1.0,
			EvaporationRate: 0.3,
			DepositStrength: 100,

			LocalLoops:  4,
			SearchDepth: 6,
		},
		configModifier: func(config aco.Config) aco.Config {
			config.NumOfAnts += 10
			return config
		},
		modificationsCount: 9,
	},
	{
		filename: "num_of_iterations",
		config: aco.Config{
			NumOfAnts:       40,
			NumOfIterations: 5,
			Alpha:           2.0,
			Beta:            1.0,
			EvaporationRate: 0.3,
			DepositStrength: 100,

			LocalLoops:  4,
			SearchDepth: 6,
		},
		configModifier: func(config aco.Config) aco.Config {
			config.NumOfIterations += 5
			return config
		},
		modificationsCount: 5,
	},
	{
		filename: "evaporation_rate",
		config: aco.Config{
			NumOfAnts:       40,
			NumOfIterations: 15,
			Alpha:           2.0,
			Beta:            1.0,
			EvaporationRate: 0.1,
			DepositStrength: 100,

			LocalLoops:  4,
			SearchDepth: 6,
		},
		configModifier: func(config aco.Config) aco.Config {
			config.EvaporationRate += 0.2
			return config
		},
		modificationsCount: 4,
	},
	{
		filename: "local_loops",
		config: aco.Config{
			NumOfAnts:       40,
			NumOfIterations: 15,
			Alpha:           2.0,
			Beta:            1.0,
			EvaporationRate: 0.3,
			DepositStrength: 100,

			LocalLoops:  2,
			SearchDepth: 6,
		},
		configModifier: func(config aco.Config) aco.Config {
			config.LocalLoops += 4
			return config
		},
		modificationsCount: 9,
	},
	{
		filename: "search_depth",
		config: aco.Config{
			NumOfAnts:       40,
			NumOfIterations: 15,
			Alpha:           2.0,
			Beta:            1.0,
			EvaporationRate: 0.3,
			DepositStrength: 100,

			LocalLoops:  4,
			SearchDepth: 2,
		},
		configModifier: func(config aco.Config) aco.Config {
			config.SearchDepth += 3
			return config
		},
		modificationsCount: 5,
	},
}
