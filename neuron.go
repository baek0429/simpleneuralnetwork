package neural

/* This is go-lang implementation of neural network based on github.com/NOX73/go-neural
/* I am doing it for fun.
*/

/** model structure **/
type Neuron struct {
	OutSynapses          []*Synapse
	InSynapses           []*Synapse
	ActionvationFunction ActionvationFunction
	Out                  float64
}

// function type, use with closures.
type ActionvationFunction func(float64) float64

// receives value and returns weight multiplied value
// Insynapse resides in Neuron and receives from OutSynapse
type Synapse struct {
	Weight float64
	In     float64
	Out    float64
}

func (s *Synapse) Signal(value float64) {
	s.In = value
	s.Out = s.In * s.Weight
}

type Network struct {
	Enters []*Enter
	Layers []*Layer
	Out    []float64
}

// initiates enters layers out
func (n *Network) initLayers() {

}

type Layer struct {
	Neurons []*Neuron
}

type Enter struct {
	OutSynapses []*Synapse
	Input       float64
}

func (e *Enter) SynapseTo(nTo *Neuron, weight float64) {
	syn := &Synapse{Weight: weight}
	e.OutSynapses = append(e.OutSynapses, syn)
}
