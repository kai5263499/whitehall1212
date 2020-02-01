package whitehall1212

// This file generated with mapgen -in data/game_board.csv -out map_data.go; DO NOT EDIT.

import (
	"github.com/kai5263499/whitehall1212/types"
)

// InitializeMap builds the in-memory game board representation
func (m *Map) InitializeMap() {
	m.vertices = make(map[types.Vertex]types.Edges, 201)
	for pos := range m.vertices {
		m.vertices[pos] = make([]types.Edge, 0)
	}

	m.vertices[types.Vertex(1)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(8),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(9),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(58),
		},
	}

	m.vertices[types.Vertex(2)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(10),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(20),
		},
	}

	m.vertices[types.Vertex(3)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(4),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(11),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(12),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(23),
		},
	}

	m.vertices[types.Vertex(4)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(3),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(13),
		},
	}

	m.vertices[types.Vertex(5)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(16),
		},
	}

	m.vertices[types.Vertex(6)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(7),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(29),
		},
	}

	m.vertices[types.Vertex(7)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(6),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(17),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(42),
		},
	}

	m.vertices[types.Vertex(8)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(1),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(18),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(19),
		},
	}

	m.vertices[types.Vertex(9)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(1),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(19),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(20),
		},
	}

	m.vertices[types.Vertex(10)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(2),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(11),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(21),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(34),
		},
	}

	m.vertices[types.Vertex(11)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(3),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(10),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(22),
		},
	}

	m.vertices[types.Vertex(12)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(3),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(23),
		},
	}

	m.vertices[types.Vertex(13)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(4),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(14),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(14),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(24),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(89),
		},
	}

	m.vertices[types.Vertex(14)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(25),
		},
	}

	m.vertices[types.Vertex(15)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(5),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(14),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(14),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(16),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(26),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(28),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(41),
		},
	}

	m.vertices[types.Vertex(16)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(5),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(28),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(29),
		},
	}

	m.vertices[types.Vertex(17)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(7),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(30),
		},
	}

	m.vertices[types.Vertex(18)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(8),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(31),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(43),
		},
	}

	m.vertices[types.Vertex(19)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(8),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(9),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(32),
		},
	}

	m.vertices[types.Vertex(20)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(2),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(9),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(33),
		},
	}

	m.vertices[types.Vertex(21)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(10),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(33),
		},
	}

	m.vertices[types.Vertex(22)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(3),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(11),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(35),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(65),
		},
	}

	m.vertices[types.Vertex(23)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(3),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(12),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(37),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(67),
		},
	}

	m.vertices[types.Vertex(24)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(37),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(38),
		},
	}

	m.vertices[types.Vertex(25)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(14),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(38),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(39),
		},
	}

	m.vertices[types.Vertex(26)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(27),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(39),
		},
	}

	m.vertices[types.Vertex(27)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(26),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(28),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(40),
		},
	}

	m.vertices[types.Vertex(28)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(16),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(27),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(41),
		},
	}

	m.vertices[types.Vertex(29)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(6),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(16),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(17),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(42),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(42),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(55),
		},
	}

	m.vertices[types.Vertex(30)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(17),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(42),
		},
	}

	m.vertices[types.Vertex(31)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(18),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(43),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(44),
		},
	}

	m.vertices[types.Vertex(32)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(19),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(33),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(44),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(45),
		},
	}

	m.vertices[types.Vertex(33)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(20),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(21),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(32),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(46),
		},
	}

	m.vertices[types.Vertex(34)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(10),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(47),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(48),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(63),
		},
	}

	m.vertices[types.Vertex(35)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(36),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(48),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(65),
		},
	}

	m.vertices[types.Vertex(36)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(35),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(37),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(49),
		},
	}

	m.vertices[types.Vertex(37)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(24),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(36),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(50),
		},
	}

	m.vertices[types.Vertex(38)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(24),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(25),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(50),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(51),
		},
	}

	m.vertices[types.Vertex(39)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(25),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(26),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(51),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(52),
		},
	}

	m.vertices[types.Vertex(40)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(27),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(53),
		},
	}

	m.vertices[types.Vertex(41)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(15),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(28),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(40),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(54),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(87),
		},
	}

	m.vertices[types.Vertex(42)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(7),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(30),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(56),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(72),
		},
	}

	m.vertices[types.Vertex(43)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(18),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(31),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(57),
		},
	}

	m.vertices[types.Vertex(44)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(31),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(32),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
	}

	m.vertices[types.Vertex(45)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(32),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(59),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(60),
		},
	}

	m.vertices[types.Vertex(46)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(1),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(1),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(33),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(45),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(47),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(61),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(79),
		},
	}

	m.vertices[types.Vertex(47)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(62),
		},
	}

	m.vertices[types.Vertex(48)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(35),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(62),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(63),
		},
	}

	m.vertices[types.Vertex(49)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(36),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(50),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(66),
		},
	}

	m.vertices[types.Vertex(50)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(37),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(38),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(49),
		},
	}

	m.vertices[types.Vertex(51)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(38),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(39),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(68),
		},
	}

	m.vertices[types.Vertex(52)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(39),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(40),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(51),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(69),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(86),
		},
	}

	m.vertices[types.Vertex(53)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(40),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(54),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(69),
		},
	}

	m.vertices[types.Vertex(54)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(53),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(55),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(70),
		},
	}

	m.vertices[types.Vertex(55)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(29),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(54),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(71),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(89),
		},
	}

	m.vertices[types.Vertex(56)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(42),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(91),
		},
	}

	m.vertices[types.Vertex(57)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(43),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(73),
		},
	}

	m.vertices[types.Vertex(58)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(1),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(44),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(45),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(57),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(59),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(75),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(77),
		},
	}

	m.vertices[types.Vertex(59)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(45),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(75),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(76),
		},
	}

	m.vertices[types.Vertex(60)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(45),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(61),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(76),
		},
	}

	m.vertices[types.Vertex(61)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(60),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(62),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(76),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(78),
		},
	}

	m.vertices[types.Vertex(62)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(47),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(48),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(61),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(79),
		},
	}

	m.vertices[types.Vertex(63)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(34),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(48),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(64),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(80),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(100),
		},
	}

	m.vertices[types.Vertex(64)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(81),
		},
	}

	m.vertices[types.Vertex(65)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(22),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(35),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(64),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(66),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(82),
		},
	}

	m.vertices[types.Vertex(66)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(49),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(82),
		},
	}

	m.vertices[types.Vertex(67)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(23),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(51),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(66),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(68),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(84),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(102),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(111),
		},
	}

	m.vertices[types.Vertex(68)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(51),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(69),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(85),
		},
	}

	m.vertices[types.Vertex(69)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(53),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(68),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(86),
		},
	}

	m.vertices[types.Vertex(70)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(54),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(71),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(87),
		},
	}

	m.vertices[types.Vertex(71)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(55),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(70),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(89),
		},
	}

	m.vertices[types.Vertex(72)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(42),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(42),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(71),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(90),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(91),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(107),
		},
	}

	m.vertices[types.Vertex(73)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(57),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(92),
		},
	}

	m.vertices[types.Vertex(74)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(73),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(75),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(92),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(94),
		},
	}

	m.vertices[types.Vertex(75)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(59),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(94),
		},
	}

	m.vertices[types.Vertex(76)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(59),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(60),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(61),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(77),
		},
	}

	m.vertices[types.Vertex(77)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(58),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(76),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(94),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(95),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(96),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(124),
		},
	}

	m.vertices[types.Vertex(78)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(61),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(97),
		},
	}

	m.vertices[types.Vertex(79)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(46),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(62),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(93),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(98),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(111),
		},
	}

	m.vertices[types.Vertex(80)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(99),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(100),
		},
	}

	m.vertices[types.Vertex(81)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(64),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(100),
		},
	}

	m.vertices[types.Vertex(82)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(65),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(66),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(81),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(100),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(101),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(140),
		},
	}

	m.vertices[types.Vertex(83)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(101),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(102),
		},
	}

	m.vertices[types.Vertex(84)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(85),
		},
	}

	m.vertices[types.Vertex(85)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(68),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(84),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(103),
		},
	}

	m.vertices[types.Vertex(86)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(52),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(69),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(87),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(102),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(103),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(104),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(116),
		},
	}

	m.vertices[types.Vertex(87)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(41),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(70),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(86),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(88),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(105),
		},
	}

	m.vertices[types.Vertex(88)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(87),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(117),
		},
	}

	m.vertices[types.Vertex(89)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(13),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(55),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(71),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(88),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(140),
		},
	}

	m.vertices[types.Vertex(90)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(91),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(105),
		},
	}

	m.vertices[types.Vertex(91)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(56),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(90),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(107),
		},
	}

	m.vertices[types.Vertex(92)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(73),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(93),
		},
	}

	m.vertices[types.Vertex(93)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(92),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(94),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(94),
		},
	}

	m.vertices[types.Vertex(94)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(74),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(75),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(93),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(93),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(95),
		},
	}

	m.vertices[types.Vertex(95)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(94),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(122),
		},
	}

	m.vertices[types.Vertex(96)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(97),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(109),
		},
	}

	m.vertices[types.Vertex(97)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(78),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(96),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(98),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(109),
		},
	}

	m.vertices[types.Vertex(98)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(97),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(99),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(110),
		},
	}

	m.vertices[types.Vertex(99)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(80),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(98),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(110),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(112),
		},
	}

	m.vertices[types.Vertex(100)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(63),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(80),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(81),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(101),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(112),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(113),
		},
	}

	m.vertices[types.Vertex(101)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(83),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(100),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
	}

	m.vertices[types.Vertex(102)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(83),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(86),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(103),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(127),
		},
	}

	m.vertices[types.Vertex(103)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(85),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(86),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(102),
		},
	}

	m.vertices[types.Vertex(104)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(86),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(116),
		},
	}

	m.vertices[types.Vertex(105)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(87),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(90),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(91),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(106),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(107),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(108),
		},
	}

	m.vertices[types.Vertex(106)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(107),
		},
	}

	m.vertices[types.Vertex(107)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(72),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(91),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(106),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(119),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(161),
		},
	}

	m.vertices[types.Vertex(108)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(105),
		},
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(117),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(119),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(135),
		},
	}

	m.vertices[types.Vertex(109)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(96),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(97),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(110),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
	}

	m.vertices[types.Vertex(110)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(98),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(99),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(109),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(111),
		},
	}

	m.vertices[types.Vertex(111)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(67),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(79),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(100),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(110),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(112),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(163),
		},
	}

	m.vertices[types.Vertex(112)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(99),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(100),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(125),
		},
	}

	m.vertices[types.Vertex(113)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(100),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(125),
		},
	}

	m.vertices[types.Vertex(114)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(101),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(113),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(126),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(131),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(132),
		},
	}

	m.vertices[types.Vertex(115)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(102),
		},
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(126),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(157),
		},
	}

	m.vertices[types.Vertex(116)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(86),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(104),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(117),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(118),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(142),
		},
	}

	m.vertices[types.Vertex(117)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(88),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(129),
		},
	}

	m.vertices[types.Vertex(118)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(129),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(134),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
	}

	m.vertices[types.Vertex(119)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(107),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(136),
		},
	}

	m.vertices[types.Vertex(120)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(121),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(144),
		},
	}

	m.vertices[types.Vertex(121)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(120),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(122),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(145),
		},
	}

	m.vertices[types.Vertex(122)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(95),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(121),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(144),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(146),
		},
	}

	m.vertices[types.Vertex(123)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(122),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(122),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(137),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(144),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(148),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(149),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(165),
		},
	}

	m.vertices[types.Vertex(124)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(77),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(109),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(130),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(138),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(138),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(153),
		},
	}

	m.vertices[types.Vertex(125)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(112),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(113),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(131),
		},
	}

	m.vertices[types.Vertex(126)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
	}

	m.vertices[types.Vertex(127)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(102),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(126),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(134),
		},
	}

	m.vertices[types.Vertex(128)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(143),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(160),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(161),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(172),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(187),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(188),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(129)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(117),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(118),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(143),
		},
	}

	m.vertices[types.Vertex(130)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(131),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(139),
		},
	}

	m.vertices[types.Vertex(131)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(125),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(130),
		},
	}

	m.vertices[types.Vertex(132)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(114),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
	}

	m.vertices[types.Vertex(133)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(141),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(157),
		},
	}

	m.vertices[types.Vertex(134)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(118),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(127),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(141),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
	}

	m.vertices[types.Vertex(135)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(108),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(129),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(136),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(143),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(161),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(161),
		},
	}

	m.vertices[types.Vertex(136)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(119),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(162),
		},
	}

	m.vertices[types.Vertex(137)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(147),
		},
	}

	m.vertices[types.Vertex(138)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(150),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(152),
		},
	}

	m.vertices[types.Vertex(139)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(130),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(154),
		},
	}

	m.vertices[types.Vertex(140)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(82),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(89),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(126),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(132),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(139),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(156),
		},
	}

	m.vertices[types.Vertex(141)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(134),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(158),
		},
	}

	m.vertices[types.Vertex(142)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(116),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(118),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(129),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(134),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(141),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(143),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(158),
		},
	}

	m.vertices[types.Vertex(143)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(129),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(160),
		},
	}

	m.vertices[types.Vertex(144)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(120),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(122),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(145),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(163),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(177),
		},
	}

	m.vertices[types.Vertex(145)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(121),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(144),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(146),
		},
	}

	m.vertices[types.Vertex(146)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(122),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(145),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(147),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(163),
		},
	}

	m.vertices[types.Vertex(147)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(137),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(146),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(164),
		},
	}

	m.vertices[types.Vertex(148)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(149),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(164),
		},
	}

	m.vertices[types.Vertex(149)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(148),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(150),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(165),
		},
	}

	m.vertices[types.Vertex(150)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(138),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(149),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(151),
		},
	}

	m.vertices[types.Vertex(151)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(150),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(152),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(165),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(166),
		},
	}

	m.vertices[types.Vertex(152)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(138),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(151),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(153),
		},
	}

	m.vertices[types.Vertex(153)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(124),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(139),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(152),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(163),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(166),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(167),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(185),
		},
	}

	m.vertices[types.Vertex(154)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(139),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(155),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(156),
		},
	}

	m.vertices[types.Vertex(155)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(167),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(168),
		},
	}

	m.vertices[types.Vertex(156)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(140),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(154),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(155),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(169),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(184),
		},
	}

	m.vertices[types.Vertex(157)] = types.Edges{
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(115),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(133),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(158),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(170),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(194),
		},
	}

	m.vertices[types.Vertex(158)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(141),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(142),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(159),
		},
	}

	m.vertices[types.Vertex(159)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(158),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(170),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(172),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(186),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(198),
		},
	}

	m.vertices[types.Vertex(160)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(143),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(161),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(173),
		},
	}

	m.vertices[types.Vertex(161)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(107),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(135),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(160),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(174),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(162)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(136),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(175),
		},
	}

	m.vertices[types.Vertex(163)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(111),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(144),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(146),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(176),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(177),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(191),
		},
	}

	m.vertices[types.Vertex(164)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(147),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(148),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(178),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(179),
		},
	}

	m.vertices[types.Vertex(165)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(123),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(149),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(151),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(179),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(191),
		},
	}

	m.vertices[types.Vertex(166)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(151),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(181),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(183),
		},
	}

	m.vertices[types.Vertex(167)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(155),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(168),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(183),
		},
	}

	m.vertices[types.Vertex(168)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(155),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(167),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(184),
		},
	}

	m.vertices[types.Vertex(169)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(184),
		},
	}

	m.vertices[types.Vertex(170)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(159),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(185),
		},
	}

	m.vertices[types.Vertex(171)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(173),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(175),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(172)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(159),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(187),
		},
	}

	m.vertices[types.Vertex(173)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(160),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(171),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(174),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(188),
		},
	}

	m.vertices[types.Vertex(174)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(161),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(173),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(175),
		},
	}

	m.vertices[types.Vertex(175)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(162),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(171),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(174),
		},
	}

	m.vertices[types.Vertex(176)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(163),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(177),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(189),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(190),
		},
	}

	m.vertices[types.Vertex(177)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(144),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(163),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(176),
		},
	}

	m.vertices[types.Vertex(178)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(164),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(189),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(191),
		},
	}

	m.vertices[types.Vertex(179)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(164),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(165),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(191),
		},
	}

	m.vertices[types.Vertex(180)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(165),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(165),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(181),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(190),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(193),
		},
	}

	m.vertices[types.Vertex(181)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(166),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(182),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(193),
		},
	}

	m.vertices[types.Vertex(182)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(181),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(183),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(195),
		},
	}

	m.vertices[types.Vertex(183)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(166),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(167),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(182),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(196),
		},
	}

	m.vertices[types.Vertex(184)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(156),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(168),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(169),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(196),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(197),
		},
	}

	m.vertices[types.Vertex(185)] = types.Edges{
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Underground"),
			Destination:    types.Vertex(153),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(170),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(186),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(187),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(186)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(159),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(198),
		},
	}

	m.vertices[types.Vertex(187)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(172),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(188),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(198),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(188)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(173),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(187),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(189)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(176),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(178),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(190),
		},
	}

	m.vertices[types.Vertex(190)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(176),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(189),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(191),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(191),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(192),
		},
	}

	m.vertices[types.Vertex(191)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(163),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(165),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(178),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(179),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(190),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(190),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(192),
		},
	}

	m.vertices[types.Vertex(192)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(190),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(191),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(194),
		},
	}

	m.vertices[types.Vertex(193)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(180),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(181),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(194),
		},
	}

	m.vertices[types.Vertex(194)] = types.Edges{
		{
			Transportation: types.Transportation("Ship"),
			Destination:    types.Vertex(157),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(192),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(193),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(195),
		},
	}

	m.vertices[types.Vertex(195)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(182),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(194),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(197),
		},
	}

	m.vertices[types.Vertex(196)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(183),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(197),
		},
	}

	m.vertices[types.Vertex(197)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(184),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(195),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(196),
		},
	}

	m.vertices[types.Vertex(198)] = types.Edges{
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(159),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(186),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(187),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(199),
		},
	}

	m.vertices[types.Vertex(199)] = types.Edges{
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(128),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(161),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(171),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(185),
		},
		{
			Transportation: types.Transportation("Bus"),
			Destination:    types.Vertex(187),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(188),
		},
		{
			Transportation: types.Transportation("Taxi"),
			Destination:    types.Vertex(198),
		},
	}

}
