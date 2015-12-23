package neural

/* This is go-lang implementation of neural network based on github.com/NOX73/go-neural
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

// enter point of input
type Enter struct {
	OutSynapses []*Synapse
	Input       float64
}

// enter has list of neurons in a layer enter -> neurons in a layer with 0 weight
func (e *Enter) ConnectTo(layer *Layer) {
	for _, n := range layer.Neurons {
		e.SynapseTo(n, 0)
	}
}

// connect to a single neuron, applied above
func (e *Enter) SynapseTo(nTo *Neuron, weight float64) {
	syn := &Synapse{Weight: weight}
	e.OutSynapses = append(e.OutSynapses, syn)
	nTo.InSynapses = append(nTo.InSynapses, syn)
}

// signal from output synapses to neurons with weight
func (e *Enter) Signal() {
	for _, s := range e.OutSynapses {
		s.Signal(e.Input)
	}
}

/* learning algorithm */

type Deltas [][]float64

type Sample struct {
	In    []float64
	Ideal []float64
}

func Learn(n *Network, in, ideal []float64, speed float64) {
	BackPropagation(n, in, ideal, speed)
}

func (n *Network) Calculate(in []float64) {

}

func BackPropagation(n *Network, in, ideal []float64, speed float64) {
	n.Calculate(in)

	deltas := make([][]float64, len(n.Layers)) // 1 by n(layers) matrix

	last := len(n.Layers) - 1 // layer index
	l := n.Layers[last]

	deltas[last] = make([]float64, len(l.Neurons)) // make detal[i] = neurons of ith layers
	for i, n := range l.Neurons {
		deltas[last][i] = n.Out * (1 - n.Out) * (ideal[i] - n.Out)
	}

	for i := last - 1; i >= 0; i-- {
		l := n.Layers[i]
		deltas[i] = make([]float64, len(l.Neurons))
		for j, n := range l.Neurons {

			var sum float64 = 0
			for k, s := range n.OutSynapses {
				sum += s.Weight * deltas[i+1][k]
			}

			deltas[i][j] = n.Out * (1 - n.Out) * sum
		}
	}
}
