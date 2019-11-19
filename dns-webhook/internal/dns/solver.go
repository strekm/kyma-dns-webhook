package dns

import (
	"github.com/jetstack/cert-manager/pkg/acme/webhook"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

// Solver implements `github.com/jetstack/cert-manager/pkg/acme/webhook.Solver`
// In order to delegate DNS challenge to DNS client deployed in different namespace/k8s cluster.
type Solver struct {
}

// NewSolver returned initialied Solver
func NewSolver() webhook.Solver {
	return &Solver{}
}

//Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource.
func (s *Solver) Name() string {
	return "kyma-dns"
}

func (s *Solver) Present(cr *v1alpha1.ChallengeRequest) error {
	log.Info("DNS Challenge Present ")

	//TODO @piotrmsc add rest client here to communicate with DNS client in order to add TXT record
	return nil
}

func (s *Solver) CleanUp(cr *v1alpha1.ChallengeRequest) error {

	//TODO @piotrmsc add rest communication here to perform cleanup
	return nil
}

func (s *Solver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {

	return nil
}
